/**
 * image_proc.cpp — 图像处理 DLL C++ 实现
 *
 * 通过嵌入 Python 解释器，将 image_processing.py 中的函数
 * 导出为 C 可调用接口。所有函数均为数组版 (uint16 输入 → uint16 输出)。
 */
#define IMAGE_PROC_EXPORTS
#include "image_proc.h"

#include <Python.h>
#include <string.h>
#include <string>
#include <fstream>

#ifdef _WIN32
    #ifndef NOMINMAX
    #define NOMINMAX
    #endif
    #include <windows.h>
#endif

// =====================================================================
//  静态状态
// =====================================================================

static PyObject* g_module = nullptr;
static bool g_initialized = false;

// =====================================================================
//  辅助函数
// =====================================================================

static IPResult make_ok() {
    IPResult r;
    r.code = 0;
    r.message[0] = '\0';
    return r;
}

static IPResult make_error(const char* context) {
    IPResult r;
    r.code = 1;
    r.message[0] = '\0';

    PyObject *type, *value, *traceback;
    PyErr_Fetch(&type, &value, &traceback);
    PyErr_NormalizeException(&type, &value, &traceback);

    if (value) {
        PyObject* str = PyObject_Str(value);
        if (str) {
            const char* msg = PyUnicode_AsUTF8(str);
            if (msg) {
                snprintf(r.message, sizeof(r.message), "%s: %s", context, msg);
            }
            Py_DECREF(str);
        }
    }
    if (r.message[0] == '\0') {
        snprintf(r.message, sizeof(r.message), "%s: unknown error", context);
    }

    Py_XDECREF(type);
    Py_XDECREF(value);
    Py_XDECREF(traceback);
    return r;
}

static PyObject* call_python(const char* func_name, PyObject* args, IPResult* out_result) {
    PyObject* func = PyObject_GetAttrString(g_module, func_name);
    if (!func) {
        *out_result = make_error(func_name);
        Py_XDECREF(args);
        return nullptr;
    }

    PyObject* result = PyObject_CallObject(func, args);
    Py_DECREF(func);
    Py_DECREF(args);

    if (!result) {
        *out_result = make_error(func_name);
        return nullptr;
    }

    *out_result = make_ok();
    return result;
}

static PyObject* create_ndarray(const uint16_t* data, int width, int height) {
    PyObject* numpy = PyImport_ImportModule("numpy");
    if (!numpy) return nullptr;

    PyObject* uint16_dtype = PyObject_GetAttrString(numpy, "uint16");
    if (!uint16_dtype) { Py_DECREF(numpy); return nullptr; }

    Py_ssize_t data_size = (Py_ssize_t)width * height * sizeof(uint16_t);
    PyObject* buffer = PyBytes_FromStringAndSize(reinterpret_cast<const char*>(data), data_size);
    if (!buffer) { Py_DECREF(uint16_dtype); Py_DECREF(numpy); return nullptr; }

    PyObject* array = PyObject_CallMethod(numpy, "frombuffer", "OO", buffer, uint16_dtype);
    Py_DECREF(uint16_dtype);

    if (!array) { Py_DECREF(buffer); Py_DECREF(numpy); return nullptr; }

    PyObject* shape = Py_BuildValue("(ii)", height, width);
    PyObject* reshaped = PyObject_CallMethod(array, "reshape", "O", shape);
    Py_DECREF(shape);
    Py_DECREF(array);

    if (!reshaped) { Py_DECREF(buffer); Py_DECREF(numpy); return nullptr; }

    Py_DECREF(buffer);
    Py_DECREF(numpy);

    PyObject* contiguous = PyObject_CallMethod(reshaped, "copy", nullptr);
    Py_DECREF(reshaped);

    return contiguous;
}

static int extract_ndarray(PyObject* result, uint16_t* output, int max_count) {
    PyObject* bytes = PyObject_CallMethod(result, "tobytes", nullptr);
    if (!bytes) return -1;

    char* data;
    Py_ssize_t size;
    if (PyBytes_AsStringAndSize(bytes, &data, &size) < 0) {
        Py_DECREF(bytes);
        return -1;
    }

    int count = (int)(size / sizeof(uint16_t));
    if (count > max_count) count = max_count;
    memcpy(output, data, (size_t)count * sizeof(uint16_t));

    Py_DECREF(bytes);
    return count;
}

/**
 * 通用辅助：单输入 → 单输出。
 * args 已包含 ndarray 作为第一个参数 + 额外参数。
 */
static IPResult call_single_io(const char* func_name,
                                const uint16_t* input, uint16_t* output,
                                int w, int h,
                                PyObject* args) {
    IPResult r;
    PyObject* ret = call_python(func_name, args, &r);
    if (ret) {
        extract_ndarray(ret, output, w * h);
        Py_DECREF(ret);
    }
    return r;
}

// =====================================================================
//  初始化 / 清理
// =====================================================================

static std::string read_pyvenv_home(const char* venv_path) {
    std::string cfg_path = std::string(venv_path) + "/pyvenv.cfg";
    std::ifstream file(cfg_path);
    if (!file.is_open()) return "";

    std::string line;
    while (std::getline(file, line)) {
        if (line.size() >= 3 && (unsigned char)line[0] == 0xEF && (unsigned char)line[1] == 0xBB && (unsigned char)line[2] == 0xBF)
            line = line.substr(3);
        if (line.find("home") == 0) {
            size_t eq = line.find('=');
            if (eq != std::string::npos) {
                std::string val = line.substr(eq + 1);
                size_t s = val.find_first_not_of(" \t");
                size_t e = val.find_last_not_of(" \t\r\n");
                if (s != std::string::npos && e != std::string::npos)
                    return val.substr(s, e - s + 1);
            }
        }
    }
    return "";
}

static void set_env(const char* key, const char* value) {
#ifdef _WIN32
    SetEnvironmentVariableA(key, value);
#else
    setenv(key, value, 1);
#endif
}

IP_API IPResult ip_init(const char* python_home, const char* src_dir) {
    if (g_initialized) {
        IPResult r;
        r.code = 0;
        snprintf(r.message, sizeof(r.message), "already initialized");
        return r;
    }

    std::string base_home = read_pyvenv_home(python_home);
    if (!base_home.empty()) {
        set_env("PYTHONHOME", base_home.c_str());
    }

    if (!Py_IsInitialized()) {
        PyConfig config;
        PyConfig_InitPythonConfig(&config);
        config.site_import = 0;

        PyStatus status = Py_InitializeFromConfig(&config);
        PyConfig_Clear(&config);

        if (PyStatus_Exception(status)) {
            IPResult r;
            r.code = 1;
            snprintf(r.message, sizeof(r.message), "Py_Initialize failed: %s",
                     status.err_msg ? status.err_msg : "unknown");
            return r;
        }
    }

    std::string site_packages = std::string(python_home) + "\\Lib\\site-packages";
    std::string dlls_dir = base_home + "\\DLLs";
    std::string lib_bin_dir = base_home + "\\Library\\bin";
    std::string setup_code =
        "import sys, os\n"
        #ifdef _WIN32
        "os.add_dll_directory(r'" + base_home + "')\n"
        "os.add_dll_directory(r'" + dlls_dir + "')\n"
        "if os.path.isdir(r'" + lib_bin_dir + "'):\n"
        "    os.add_dll_directory(r'" + lib_bin_dir + "')\n"
        "os.add_dll_directory(r'" + site_packages + "')\n"
        "for _d in os.listdir(r'" + site_packages + "'):\n"
        "    _p = os.path.join(r'" + site_packages + "', _d)\n"
        "    if os.path.isdir(_p):\n"
        "        try:\n"
        "            if any(_f.endswith('.dll') for _f in os.listdir(_p)):\n"
        "                os.add_dll_directory(_p)\n"
        "        except OSError:\n"
        "            pass\n"
        #endif
        "sys.path.insert(0, r'" + site_packages + "')\n"
        "sys.path.insert(0, r'" + std::string(src_dir) + "')\n";
    PyRun_SimpleString(setup_code.c_str());

    g_module = PyImport_ImportModule("image_processing");
    if (!g_module) {
        IPResult r = make_error("ip_init: import image_processing");
        PyEval_SaveThread();
        return r;
    }

    g_initialized = true;
    PyEval_SaveThread();

    return make_ok();
}

IP_API void ip_finalize(void) {
    if (!g_initialized) return;

    PyGILState_STATE gil = PyGILState_Ensure();

    Py_XDECREF(g_module);
    g_module = nullptr;

    if (Py_IsInitialized()) {
        Py_Finalize();
    }

    g_initialized = false;
}

// =====================================================================
//  去模糊
// =====================================================================

IP_API IPResult ip_richardson_lucy(const uint16_t* input, uint16_t* output, int w, int h, double psf_sigma, int num_iter) {
    PyGILState_STATE gil = PyGILState_Ensure();
    PyObject* arr = create_ndarray(input, w, h);
    if (!arr) { IPResult r = make_error("ip_richardson_lucy"); PyGILState_Release(gil); return r; }
    PyObject* args = Py_BuildValue("(Odi)", arr, psf_sigma, num_iter);
    Py_DECREF(arr);
    IPResult r = call_single_io("richardson_lucy_16bit", input, output, w, h, args);
    PyGILState_Release(gil);
    return r;
}

IP_API IPResult ip_wiener_deblur(const uint16_t* input, uint16_t* output, int w, int h, double psf_sigma, double balance) {
    PyGILState_STATE gil = PyGILState_Ensure();
    PyObject* arr = create_ndarray(input, w, h);
    if (!arr) { IPResult r = make_error("ip_wiener_deblur"); PyGILState_Release(gil); return r; }
    PyObject* args = Py_BuildValue("(Odd)", arr, psf_sigma, balance);
    Py_DECREF(arr);
    IPResult r = call_single_io("wiener_deblur_16bit", input, output, w, h, args);
    PyGILState_Release(gil);
    return r;
}

// =====================================================================
//  去噪
// =====================================================================

IP_API IPResult ip_adaptive_median_filter(const uint16_t* input, uint16_t* output, int w, int h, int window_size) {
    PyGILState_STATE gil = PyGILState_Ensure();
    PyObject* arr = create_ndarray(input, w, h);
    if (!arr) { IPResult r = make_error("ip_adaptive_median_filter"); PyGILState_Release(gil); return r; }
    PyObject* args = Py_BuildValue("(Oi)", arr, window_size);
    Py_DECREF(arr);
    IPResult r = call_single_io("adaptiveMedianFilter", input, output, w, h, args);
    PyGILState_Release(gil);
    return r;
}

IP_API IPResult ip_weighted_median_filter(const uint16_t* input, uint16_t* output, int w, int h, int kernel_size) {
    PyGILState_STATE gil = PyGILState_Ensure();
    PyObject* arr = create_ndarray(input, w, h);
    if (!arr) { IPResult r = make_error("ip_weighted_median_filter"); PyGILState_Release(gil); return r; }
    PyObject* args = Py_BuildValue("(Oi)", arr, kernel_size);
    Py_DECREF(arr);
    IPResult r = call_single_io("weightedMedianFilter", input, output, w, h, args);
    PyGILState_Release(gil);
    return r;
}

IP_API IPResult ip_denoise_bilateral(const uint16_t* input, uint16_t* output, int w, int h,
                                      int d, double sigma_color, double sigma_space) {
    PyGILState_STATE gil = PyGILState_Ensure();
    PyObject* arr = create_ndarray(input, w, h);
    if (!arr) { IPResult r = make_error("ip_denoise_bilateral"); PyGILState_Release(gil); return r; }
    PyObject* args = Py_BuildValue("(Oidd)", arr, d, sigma_color, sigma_space);
    Py_DECREF(arr);
    IPResult r = call_single_io("denoise_bilateral_16bit", input, output, w, h, args);
    PyGILState_Release(gil);
    return r;
}

IP_API IPResult ip_nlm_denoise(const uint16_t* input, uint16_t* output, int w, int h,
                                double h_factor, int patch_size, int patch_distance) {
    PyGILState_STATE gil = PyGILState_Ensure();
    PyObject* arr = create_ndarray(input, w, h);
    if (!arr) { IPResult r = make_error("ip_nlm_denoise"); PyGILState_Release(gil); return r; }
    PyObject* args = Py_BuildValue("(Odii)", arr, h_factor, patch_size, patch_distance);
    Py_DECREF(arr);
    IPResult r = call_single_io("nlm_denoise_16bit_skimage", input, output, w, h, args);
    PyGILState_Release(gil);
    return r;
}

IP_API IPResult ip_bm3d_denoise(const uint16_t* input, uint16_t* output, int w, int h, double sigma_psd) {
    PyGILState_STATE gil = PyGILState_Ensure();
    PyObject* arr = create_ndarray(input, w, h);
    if (!arr) { IPResult r = make_error("ip_bm3d_denoise"); PyGILState_Release(gil); return r; }
    PyObject* args = Py_BuildValue("(Od)", arr, sigma_psd);
    Py_DECREF(arr);
    IPResult r = call_single_io("image_bm3d_denoise_16bit", input, output, w, h, args);
    PyGILState_Release(gil);
    return r;
}

IP_API IPResult ip_tv_denoise(const uint16_t* input, uint16_t* output, int w, int h, double weight) {
    PyGILState_STATE gil = PyGILState_Ensure();
    PyObject* arr = create_ndarray(input, w, h);
    if (!arr) { IPResult r = make_error("ip_tv_denoise"); PyGILState_Release(gil); return r; }
    PyObject* args = Py_BuildValue("(Od)", arr, weight);
    Py_DECREF(arr);
    IPResult r = call_single_io("tv_denoise_16bit", input, output, w, h, args);
    PyGILState_Release(gil);
    return r;
}

IP_API IPResult ip_wavelet_denoise(const uint16_t* input, uint16_t* output, int w, int h,
                                    double sigma, const char* method) {
    PyGILState_STATE gil = PyGILState_Ensure();
    PyObject* arr = create_ndarray(input, w, h);
    if (!arr) { IPResult r = make_error("ip_wavelet_denoise"); PyGILState_Release(gil); return r; }

    PyObject* args;
    if (sigma < 0) {
        args = Py_BuildValue("(OOs)", arr, Py_None, method ? method : "BayesShrink");
    } else {
        args = Py_BuildValue("(Ods)", arr, sigma, method ? method : "BayesShrink");
    }
    Py_DECREF(arr);
    IPResult r = call_single_io("wavelet_denoise_16bit", input, output, w, h, args);
    PyGILState_Release(gil);
    return r;
}

IP_API IPResult ip_filter(const uint16_t* input, uint16_t* out_gauss, uint16_t* out_median, uint16_t* out_mean,
                          int w, int h, double sigma, int median_size, int mean_size) {
    PyGILState_STATE gil = PyGILState_Ensure();
    PyObject* arr = create_ndarray(input, w, h);
    if (!arr) { IPResult r = make_error("ip_filter"); PyGILState_Release(gil); return r; }

    PyObject* args = Py_BuildValue("(Odii)", arr, sigma, median_size, mean_size);
    Py_DECREF(arr);

    IPResult r;
    PyObject* ret = call_python("filter_16bit", args, &r);
    if (ret) {
        if (PyTuple_Check(ret) && PyTuple_Size(ret) >= 3) {
            extract_ndarray(PyTuple_GetItem(ret, 0), out_gauss, w * h);
            extract_ndarray(PyTuple_GetItem(ret, 1), out_median, w * h);
            extract_ndarray(PyTuple_GetItem(ret, 2), out_mean, w * h);
        }
        Py_DECREF(ret);
    }
    PyGILState_Release(gil);
    return r;
}

// =====================================================================
//  数学变换
// =====================================================================

IP_API IPResult ip_exp_transform(const uint16_t* input, uint16_t* output, int w, int h, double c) {
    PyGILState_STATE gil = PyGILState_Ensure();
    PyObject* arr = create_ndarray(input, w, h);
    if (!arr) { IPResult r = make_error("ip_exp_transform"); PyGILState_Release(gil); return r; }
    PyObject* args = Py_BuildValue("(Od)", arr, c);
    Py_DECREF(arr);
    IPResult r = call_single_io("exp_transform_16bit", input, output, w, h, args);
    PyGILState_Release(gil);
    return r;
}

IP_API IPResult ip_gamma_correction(const uint16_t* input, uint16_t* output, int w, int h, double gamma) {
    PyGILState_STATE gil = PyGILState_Ensure();
    PyObject* arr = create_ndarray(input, w, h);
    if (!arr) { IPResult r = make_error("ip_gamma_correction"); PyGILState_Release(gil); return r; }
    PyObject* args = Py_BuildValue("(Od)", arr, gamma);
    Py_DECREF(arr);
    IPResult r = call_single_io("gamma_correction_16bit", input, output, w, h, args);
    PyGILState_Release(gil);
    return r;
}

IP_API IPResult ip_log_transform(const uint16_t* input, uint16_t* output, int w, int h, double c) {
    PyGILState_STATE gil = PyGILState_Ensure();
    PyObject* arr = create_ndarray(input, w, h);
    if (!arr) { IPResult r = make_error("ip_log_transform"); PyGILState_Release(gil); return r; }
    PyObject* args = Py_BuildValue("(Od)", arr, c);
    Py_DECREF(arr);
    IPResult r = call_single_io("log_transform_16bit", input, output, w, h, args);
    PyGILState_Release(gil);
    return r;
}

// =====================================================================
//  锐化
// =====================================================================

IP_API IPResult ip_imagej_sharpen(const uint16_t* input, uint16_t* output, int w, int h) {
    PyGILState_STATE gil = PyGILState_Ensure();
    PyObject* arr = create_ndarray(input, w, h);
    if (!arr) { IPResult r = make_error("ip_imagej_sharpen"); PyGILState_Release(gil); return r; }
    PyObject* args = Py_BuildValue("(O)", arr);
    Py_DECREF(arr);
    IPResult r = call_single_io("imagej_sharpen_16bit", input, output, w, h, args);
    PyGILState_Release(gil);
    return r;
}

IP_API IPResult ip_sharpen(const uint16_t* input, uint16_t* output, int w, int h, double radius, double amount) {
    PyGILState_STATE gil = PyGILState_Ensure();
    PyObject* arr = create_ndarray(input, w, h);
    if (!arr) { IPResult r = make_error("ip_sharpen"); PyGILState_Release(gil); return r; }
    PyObject* args = Py_BuildValue("(Odd)", arr, radius, amount);
    Py_DECREF(arr);
    IPResult r = call_single_io("sharpen_16bit", input, output, w, h, args);
    PyGILState_Release(gil);
    return r;
}
