/**
 * image_proc.h — 图像处理 DLL C 接口
 *
 * 所有函数均为数组版：传入 uint16 像素数据，返回 uint16 像素数据。
 * 不涉及文件 I/O，调用方负责读写文件。
 *
 * 使用流程：
 *   1. ip_init() 初始化
 *   2. 分配 input/output 缓冲区，调用 ip_xxx() 处理
 *   3. ip_finalize() 释放
 */
#ifndef IMAGE_PROC_H
#define IMAGE_PROC_H

#include <stdint.h>

#ifdef _WIN32
    #ifdef IMAGE_PROC_EXPORTS
        #define IP_API __declspec(dllexport)
    #else
        #define IP_API __declspec(dllimport)
    #endif
#else
    #define IP_API __attribute__((visibility("default")))
#endif

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    int code;            /* 0=成功, 非0=错误 */
    char message[512];   /* 错误信息 */
} IPResult;

/* ===================== 初始化 / 清理 ===================== */

IP_API IPResult ip_init(const char* python_home, const char* src_dir);
IP_API void ip_finalize(void);

/* ===================== 去模糊 ===================== */

IP_API IPResult ip_richardson_lucy(const uint16_t* input, uint16_t* output, int w, int h, double psf_sigma, int num_iter);
IP_API IPResult ip_wiener_deblur(const uint16_t* input, uint16_t* output, int w, int h, double psf_sigma, double balance);

/* ===================== 去噪 ===================== */

IP_API IPResult ip_adaptive_median_filter(const uint16_t* input, uint16_t* output, int w, int h, int window_size);
IP_API IPResult ip_weighted_median_filter(const uint16_t* input, uint16_t* output, int w, int h, int kernel_size);
IP_API IPResult ip_denoise_bilateral(const uint16_t* input, uint16_t* output, int w, int h, int d, double sigma_color, double sigma_space);
IP_API IPResult ip_nlm_denoise(const uint16_t* input, uint16_t* output, int w, int h, double h_factor, int patch_size, int patch_distance);
IP_API IPResult ip_bm3d_denoise(const uint16_t* input, uint16_t* output, int w, int h, double sigma_psd);
IP_API IPResult ip_tv_denoise(const uint16_t* input, uint16_t* output, int w, int h, double weight);
IP_API IPResult ip_wavelet_denoise(const uint16_t* input, uint16_t* output, int w, int h, double sigma, const char* method);
IP_API IPResult ip_filter(const uint16_t* input, uint16_t* out_gauss, uint16_t* out_median, uint16_t* out_mean, int w, int h, double sigma, int median_size, int mean_size);

/* ===================== 数学变换 ===================== */

IP_API IPResult ip_exp_transform(const uint16_t* input, uint16_t* output, int w, int h, double c);
IP_API IPResult ip_gamma_correction(const uint16_t* input, uint16_t* output, int w, int h, double gamma);
IP_API IPResult ip_log_transform(const uint16_t* input, uint16_t* output, int w, int h, double c);

/* ===================== 锐化 ===================== */

IP_API IPResult ip_imagej_sharpen(const uint16_t* input, uint16_t* output, int w, int h);
IP_API IPResult ip_sharpen(const uint16_t* input, uint16_t* output, int w, int h, double radius, double amount);

#ifdef __cplusplus
}
#endif

#endif /* IMAGE_PROC_H */
