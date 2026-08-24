package PythonSDK

/*
#cgo CFLAGS: -I${SRCDIR}/src/dll
#include <stdlib.h>
#include <string.h>
#include <windows.h>
#include "image_proc.h"

// ===== DLL 动态加载 =====

static HMODULE g_dll = NULL;

// 函数指针类型定义（全部为数组版：uint16 输入 → uint16 输出）
typedef IPResult (*fn_ip_init)(const char*, const char*);
typedef void    (*fn_ip_finalize)(void);
typedef IPResult (*fn_ip_richardson_lucy)(const uint16_t*, uint16_t*, int, int, double, int);
typedef IPResult (*fn_ip_wiener_deblur)(const uint16_t*, uint16_t*, int, int, double, double);
typedef IPResult (*fn_ip_adaptive_median_filter)(const uint16_t*, uint16_t*, int, int, int);
typedef IPResult (*fn_ip_weighted_median_filter)(const uint16_t*, uint16_t*, int, int, int);
typedef IPResult (*fn_ip_denoise_bilateral)(const uint16_t*, uint16_t*, int, int, int, double, double);
typedef IPResult (*fn_ip_nlm_denoise)(const uint16_t*, uint16_t*, int, int, double, int, int);
typedef IPResult (*fn_ip_bm3d_denoise)(const uint16_t*, uint16_t*, int, int, double);
typedef IPResult (*fn_ip_tv_denoise)(const uint16_t*, uint16_t*, int, int, double);
typedef IPResult (*fn_ip_wavelet_denoise)(const uint16_t*, uint16_t*, int, int, double, const char*);
typedef IPResult (*fn_ip_filter)(const uint16_t*, uint16_t*, uint16_t*, uint16_t*, int, int, double, int, int);
typedef IPResult (*fn_ip_exp_transform)(const uint16_t*, uint16_t*, int, int, double);
typedef IPResult (*fn_ip_gamma_correction)(const uint16_t*, uint16_t*, int, int, double);
typedef IPResult (*fn_ip_log_transform)(const uint16_t*, uint16_t*, int, int, double);
typedef IPResult (*fn_ip_imagej_sharpen)(const uint16_t*, uint16_t*, int, int);
typedef IPResult (*fn_ip_sharpen)(const uint16_t*, uint16_t*, int, int, double, double);

// 函数指针全局变量
static fn_ip_init                g_ip_init = NULL;
static fn_ip_finalize            g_ip_finalize = NULL;
static fn_ip_richardson_lucy     g_ip_richardson_lucy = NULL;
static fn_ip_wiener_deblur       g_ip_wiener_deblur = NULL;
static fn_ip_adaptive_median_filter g_ip_adaptive_median_filter = NULL;
static fn_ip_weighted_median_filter g_ip_weighted_median_filter = NULL;
static fn_ip_denoise_bilateral  g_ip_denoise_bilateral = NULL;
static fn_ip_nlm_denoise         g_ip_nlm_denoise = NULL;
static fn_ip_bm3d_denoise        g_ip_bm3d_denoise = NULL;
static fn_ip_tv_denoise          g_ip_tv_denoise = NULL;
static fn_ip_wavelet_denoise     g_ip_wavelet_denoise = NULL;
static fn_ip_filter              g_ip_filter = NULL;
static fn_ip_exp_transform       g_ip_exp_transform = NULL;
static fn_ip_gamma_correction    g_ip_gamma_correction = NULL;
static fn_ip_log_transform       g_ip_log_transform = NULL;
static fn_ip_imagej_sharpen      g_ip_imagej_sharpen = NULL;
static fn_ip_sharpen             g_ip_sharpen = NULL;

static void ip_set_dll_dir(const char* dir) {
    SetDllDirectoryA(dir);
}

static int ip_load_dll(const char* path) {
    g_dll = LoadLibraryA(path);
    if (!g_dll) return 0;

    g_ip_init                = (fn_ip_init)GetProcAddress(g_dll, "ip_init");
    g_ip_finalize            = (fn_ip_finalize)GetProcAddress(g_dll, "ip_finalize");
    g_ip_richardson_lucy     = (fn_ip_richardson_lucy)GetProcAddress(g_dll, "ip_richardson_lucy");
    g_ip_wiener_deblur       = (fn_ip_wiener_deblur)GetProcAddress(g_dll, "ip_wiener_deblur");
    g_ip_adaptive_median_filter = (fn_ip_adaptive_median_filter)GetProcAddress(g_dll, "ip_adaptive_median_filter");
    g_ip_weighted_median_filter = (fn_ip_weighted_median_filter)GetProcAddress(g_dll, "ip_weighted_median_filter");
    g_ip_denoise_bilateral    = (fn_ip_denoise_bilateral)GetProcAddress(g_dll, "ip_denoise_bilateral");
    g_ip_nlm_denoise          = (fn_ip_nlm_denoise)GetProcAddress(g_dll, "ip_nlm_denoise");
    g_ip_bm3d_denoise         = (fn_ip_bm3d_denoise)GetProcAddress(g_dll, "ip_bm3d_denoise");
    g_ip_tv_denoise           = (fn_ip_tv_denoise)GetProcAddress(g_dll, "ip_tv_denoise");
    g_ip_wavelet_denoise      = (fn_ip_wavelet_denoise)GetProcAddress(g_dll, "ip_wavelet_denoise");
    g_ip_filter               = (fn_ip_filter)GetProcAddress(g_dll, "ip_filter");
    g_ip_exp_transform        = (fn_ip_exp_transform)GetProcAddress(g_dll, "ip_exp_transform");
    g_ip_gamma_correction     = (fn_ip_gamma_correction)GetProcAddress(g_dll, "ip_gamma_correction");
    g_ip_log_transform        = (fn_ip_log_transform)GetProcAddress(g_dll, "ip_log_transform");
    g_ip_imagej_sharpen       = (fn_ip_imagej_sharpen)GetProcAddress(g_dll, "ip_imagej_sharpen");
    g_ip_sharpen              = (fn_ip_sharpen)GetProcAddress(g_dll, "ip_sharpen");

    if (!g_ip_init || !g_ip_finalize) {
        FreeLibrary(g_dll);
        g_dll = NULL;
        return 0;
    }
    return 1;
}

static IPResult ip_not_loaded(void) {
    IPResult r;
    r.code = -1;
    strcpy(r.message, "DLL not loaded, call InitPython first");
    return r;
}

// 包装宏
#define WRAP_IP(name, params, args) \
    static IPResult w_##name params { \
        if (!g_##name) return ip_not_loaded(); \
        return g_##name args; \
    }
#define WRAP_VOID(name, params, args) \
    static void w_##name params { \
        if (g_##name) g_##name args; \
    }

WRAP_IP(ip_init, (const char* a, const char* b), (a, b))
WRAP_VOID(ip_finalize, (void), ())
WRAP_IP(ip_richardson_lucy, (const uint16_t* a, uint16_t* b, int c, int d, double e, int f), (a, b, c, d, e, f))
WRAP_IP(ip_wiener_deblur, (const uint16_t* a, uint16_t* b, int c, int d, double e, double f), (a, b, c, d, e, f))
WRAP_IP(ip_adaptive_median_filter, (const uint16_t* a, uint16_t* b, int c, int d, int e), (a, b, c, d, e))
WRAP_IP(ip_weighted_median_filter, (const uint16_t* a, uint16_t* b, int c, int d, int e), (a, b, c, d, e))
WRAP_IP(ip_denoise_bilateral, (const uint16_t* a, uint16_t* b, int c, int d, int e, double f, double g), (a, b, c, d, e, f, g))
WRAP_IP(ip_nlm_denoise, (const uint16_t* a, uint16_t* b, int c, int d, double e, int f, int g), (a, b, c, d, e, f, g))
WRAP_IP(ip_bm3d_denoise, (const uint16_t* a, uint16_t* b, int c, int d, double e), (a, b, c, d, e))
WRAP_IP(ip_tv_denoise, (const uint16_t* a, uint16_t* b, int c, int d, double e), (a, b, c, d, e))
WRAP_IP(ip_wavelet_denoise, (const uint16_t* a, uint16_t* b, int c, int d, double e, const char* f), (a, b, c, d, e, f))
WRAP_IP(ip_filter, (const uint16_t* a, uint16_t* b, uint16_t* c, uint16_t* d, int e, int f, double g, int h, int i), (a, b, c, d, e, f, g, h, i))
WRAP_IP(ip_exp_transform, (const uint16_t* a, uint16_t* b, int c, int d, double e), (a, b, c, d, e))
WRAP_IP(ip_gamma_correction, (const uint16_t* a, uint16_t* b, int c, int d, double e), (a, b, c, d, e))
WRAP_IP(ip_log_transform, (const uint16_t* a, uint16_t* b, int c, int d, double e), (a, b, c, d, e))
WRAP_IP(ip_imagej_sharpen, (const uint16_t* a, uint16_t* b, int c, int d), (a, b, c, d))
WRAP_IP(ip_sharpen, (const uint16_t* a, uint16_t* b, int c, int d, double e, double f), (a, b, c, d, e, f))

#undef WRAP_IP
#undef WRAP_VOID
*/
import "C"

import (
	"fmt"
	"unsafe"
)

const (
	DefaultVenvDir = "d:/NIMTE/Code/go/Diffract/services/PythonSDK/.venv"
	DefaultSrcDir  = "d:/NIMTE/Code/go/Diffract/services/PythonSDK/src"
	DefaultDllDir  = "d:/NIMTE/Code/go/Diffract/services/PythonSDK/src/dll/build"
)

func ipResultToError(r C.IPResult) error {
	if r.code == 0 {
		return nil
	}
	return fmt.Errorf("%s", C.GoString(&r.message[0]))
}

func cString(s string) (*C.char, func()) {
	cs := C.CString(s)
	return cs, func() { C.free(unsafe.Pointer(cs)) }
}

// ptr 返回 []uint16 底层指针（slice 为空时返回 NULL）
func ptr(s []uint16) *C.uint16_t {
	if len(s) == 0 {
		return nil
	}
	return (*C.uint16_t)(unsafe.Pointer(&s[0]))
}

// ─── 初始化 / 清理 ───

func InitPython() error {
	return InitPythonWith(DefaultVenvDir, DefaultSrcDir, DefaultDllDir)
}

func InitPythonWith(venvDir, srcDir, dllDir string) error {
	cdllDir, freedll := cString(dllDir)
	defer freedll()
	C.ip_set_dll_dir(cdllDir)

	dllPath := dllDir + "/libimage_proc.dll"
	cPath, freePath := cString(dllPath)
	defer freePath()
	if C.ip_load_dll(cPath) == 0 {
		return fmt.Errorf("无法加载 DLL: %s", dllPath)
	}

	cVenv, freeVenv := cString(venvDir)
	defer freeVenv()
	cSrc, freeSrc := cString(srcDir)
	defer freeSrc()
	r := C.w_ip_init(cVenv, cSrc)
	return ipResultToError(r)
}

func ClosePython() {
	C.w_ip_finalize()
}

// ─── 去模糊 ───

func RichardsonLucyDeblur(input []uint16, width, height int, psfSigma float64, numIter int) ([]uint16, error) {
	output := make([]uint16, width*height)
	r := C.w_ip_richardson_lucy(ptr(input), ptr(output), C.int(width), C.int(height), C.double(psfSigma), C.int(numIter))
	if err := ipResultToError(r); err != nil {
		return nil, err
	}
	return output, nil
}

func WienerDeblur(input []uint16, width, height int, psfSigma, balance float64) ([]uint16, error) {
	output := make([]uint16, width*height)
	r := C.w_ip_wiener_deblur(ptr(input), ptr(output), C.int(width), C.int(height), C.double(psfSigma), C.double(balance))
	if err := ipResultToError(r); err != nil {
		return nil, err
	}
	return output, nil
}

// ─── 去噪 ───

func AdaptiveMedianFilter(input []uint16, width, height, windowSize int) ([]uint16, error) {
	output := make([]uint16, width*height)
	r := C.w_ip_adaptive_median_filter(ptr(input), ptr(output), C.int(width), C.int(height), C.int(windowSize))
	if err := ipResultToError(r); err != nil {
		return nil, err
	}
	return output, nil
}

func WeightedMedianFilter(input []uint16, width, height, kernelSize int) ([]uint16, error) {
	output := make([]uint16, width*height)
	r := C.w_ip_weighted_median_filter(ptr(input), ptr(output), C.int(width), C.int(height), C.int(kernelSize))
	if err := ipResultToError(r); err != nil {
		return nil, err
	}
	return output, nil
}

func BilateralDenoise(input []uint16, width, height, d int, sigmaColor, sigmaSpace float64) ([]uint16, error) {
	output := make([]uint16, width*height)
	r := C.w_ip_denoise_bilateral(ptr(input), ptr(output), C.int(width), C.int(height), C.int(d), C.double(sigmaColor), C.double(sigmaSpace))
	if err := ipResultToError(r); err != nil {
		return nil, err
	}
	return output, nil
}

func NLMDenoise(input []uint16, width, height int, hFactor float64, patchSize, patchDistance int) ([]uint16, error) {
	output := make([]uint16, width*height)
	r := C.w_ip_nlm_denoise(ptr(input), ptr(output), C.int(width), C.int(height), C.double(hFactor), C.int(patchSize), C.int(patchDistance))
	if err := ipResultToError(r); err != nil {
		return nil, err
	}
	return output, nil
}

func BM3DDenoise(input []uint16, width, height int, sigmaPsd float64) ([]uint16, error) {
	output := make([]uint16, width*height)
	r := C.w_ip_bm3d_denoise(ptr(input), ptr(output), C.int(width), C.int(height), C.double(sigmaPsd))
	if err := ipResultToError(r); err != nil {
		return nil, err
	}
	return output, nil
}

func TVDenoise(input []uint16, width, height int, weight float64) ([]uint16, error) {
	output := make([]uint16, width*height)
	r := C.w_ip_tv_denoise(ptr(input), ptr(output), C.int(width), C.int(height), C.double(weight))
	if err := ipResultToError(r); err != nil {
		return nil, err
	}
	return output, nil
}

func WaveletDenoise(input []uint16, width, height int, sigma float64, method string) ([]uint16, error) {
	output := make([]uint16, width*height)
	cMethod, freeMethod := cString(method)
	defer freeMethod()
	r := C.w_ip_wavelet_denoise(ptr(input), ptr(output), C.int(width), C.int(height), C.double(sigma), cMethod)
	if err := ipResultToError(r); err != nil {
		return nil, err
	}
	return output, nil
}

// Filter16bit 返回高斯、中值、均值三种滤波结果
func Filter16bit(input []uint16, width, height int, sigma float64, medianSize, meanSize int) (gauss, median, mean []uint16, err error) {
	gauss = make([]uint16, width*height)
	median = make([]uint16, width*height)
	mean = make([]uint16, width*height)
	r := C.w_ip_filter(ptr(input), ptr(gauss), ptr(median), ptr(mean),
		C.int(width), C.int(height), C.double(sigma), C.int(medianSize), C.int(meanSize))
	err = ipResultToError(r)
	return
}

// ─── 数学变换 ───

func ExpTransform(input []uint16, width, height int, c float64) ([]uint16, error) {
	output := make([]uint16, width*height)
	r := C.w_ip_exp_transform(ptr(input), ptr(output), C.int(width), C.int(height), C.double(c))
	if err := ipResultToError(r); err != nil {
		return nil, err
	}
	return output, nil
}

func GammaCorrection(input []uint16, width, height int, gamma float64) ([]uint16, error) {
	output := make([]uint16, width*height)
	r := C.w_ip_gamma_correction(ptr(input), ptr(output), C.int(width), C.int(height), C.double(gamma))
	if err := ipResultToError(r); err != nil {
		return nil, err
	}
	return output, nil
}

func LogTransform(input []uint16, width, height int, c float64) ([]uint16, error) {
	output := make([]uint16, width*height)
	r := C.w_ip_log_transform(ptr(input), ptr(output), C.int(width), C.int(height), C.double(c))
	if err := ipResultToError(r); err != nil {
		return nil, err
	}
	return output, nil
}

// ─── 锐化 ───

func ImageJSharpen(input []uint16, width, height int) ([]uint16, error) {
	output := make([]uint16, width*height)
	r := C.w_ip_imagej_sharpen(ptr(input), ptr(output), C.int(width), C.int(height))
	if err := ipResultToError(r); err != nil {
		return nil, err
	}
	return output, nil
}

func Sharpen(input []uint16, width, height int, radius, amount float64) ([]uint16, error) {
	output := make([]uint16, width*height)
	r := C.w_ip_sharpen(ptr(input), ptr(output), C.int(width), C.int(height), C.double(radius), C.double(amount))
	if err := ipResultToError(r); err != nil {
		return nil, err
	}
	return output, nil
}

// ─── 测试 ───

func PythonTest() {
	if err := InitPython(); err != nil {
		fmt.Println("初始化失败:", err)
		return
	}
	defer ClosePython()

	fmt.Println("Python 初始化成功")

	W, H := 128, 128
	input := make([]uint16, W*H)
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			val := uint16(x*512 + y*4)
			if (x*7+y*13)%100 == 0 {
				if x%2 == 0 {
					val = 0
				} else {
					val = 65535
				}
			}
			input[y*W+x] = val
		}
	}

	diff := func(a, b []uint16) int {
		d := 0
		for i := range a {
			if a[i] != b[i] {
				d++
			}
		}
		return d
	}

	// 去模糊
	fmt.Println("[1] Richardson-Lucy...")
	if out, err := RichardsonLucyDeblur(input, W, H, 1.5, 5); err != nil {
		fmt.Println("  FAIL:", err)
	} else {
		fmt.Printf("  OK, %d 像素变化\n", diff(input, out))
	}

	fmt.Println("[2] Wiener...")
	if out, err := WienerDeblur(input, W, H, 1.5, 0.1); err != nil {
		fmt.Println("  FAIL:", err)
	} else {
		fmt.Printf("  OK, %d 像素变化\n", diff(input, out))
	}

	// 去噪
	fmt.Println("[3] 自适应中值滤波...")
	if out, err := AdaptiveMedianFilter(input, W, H, 3); err != nil {
		fmt.Println("  FAIL:", err)
	} else {
		fmt.Printf("  OK, %d 像素变化\n", diff(input, out))
	}

	fmt.Println("[4] 双边滤波...")
	if out, err := BilateralDenoise(input, W, H, 5, 0.08, 80); err != nil {
		fmt.Println("  FAIL:", err)
	} else {
		fmt.Printf("  OK, %d 像素变化\n", diff(input, out))
	}

	fmt.Println("[5] TV 去噪...")
	if out, err := TVDenoise(input, W, H, 0.05); err != nil {
		fmt.Println("  FAIL:", err)
	} else {
		fmt.Printf("  OK, %d 像素变化\n", diff(input, out))
	}

	fmt.Println("[6] 小波去噪...")
	if out, err := WaveletDenoise(input, W, H, 0.001, "BayesShrink"); err != nil {
		fmt.Println("  FAIL:", err)
	} else {
		fmt.Printf("  OK, %d 像素变化\n", diff(input, out))
	}

	fmt.Println("[7] 滤波(3路)...")
	if g, m, me, err := Filter16bit(input, W, H, 2, 3, 3); err != nil {
		fmt.Println("  FAIL:", err)
	} else {
		fmt.Printf("  OK, gauss=%d median=%d mean=%d\n", diff(input, g), diff(input, m), diff(input, me))
	}

	// 变换
	fmt.Println("[8] Gamma 校正...")
	if out, err := GammaCorrection(input, W, H, 0.5); err != nil {
		fmt.Println("  FAIL:", err)
	} else {
		fmt.Printf("  OK, %d 像素变化\n", diff(input, out))
	}

	fmt.Println("[9] 对数变换...")
	if out, err := LogTransform(input, W, H, 1.0); err != nil {
		fmt.Println("  FAIL:", err)
	} else {
		fmt.Printf("  OK, %d 像素变化\n", diff(input, out))
	}

	fmt.Println("[10] 指数变换...")
	if out, err := ExpTransform(input, W, H, 1.0); err != nil {
		fmt.Println("  FAIL:", err)
	} else {
		fmt.Printf("  OK, %d 像素变化\n", diff(input, out))
	}

	// 锐化
	fmt.Println("[11] ImageJ 锐化...")
	if out, err := ImageJSharpen(input, W, H); err != nil {
		fmt.Println("  FAIL:", err)
	} else {
		fmt.Printf("  OK, %d 像素变化\n", diff(input, out))
	}

	fmt.Println("[12] USM 锐化...")
	if out, err := Sharpen(input, W, H, 2.0, 1.5); err != nil {
		fmt.Println("  FAIL:", err)
	} else {
		fmt.Printf("  OK, %d 像素变化\n", diff(input, out))
	}

	fmt.Println("=== 测试完成 ===")
}
