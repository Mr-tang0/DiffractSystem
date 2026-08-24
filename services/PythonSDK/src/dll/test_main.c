/**
 * test_main.c — 全数组版 DLL 功能测试
 *
 * 所有测试用 uint16 像素数据直接传入，不涉及文件 I/O。
 */
#include "image_proc.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define VENV_DIR "d:/NIMTE/Code/go/Diffract/services/PythonSDK/.venv"
#define SRC_DIR  "d:/NIMTE/Code/go/Diffract/services/PythonSDK/src"

static int count_diff(const uint16_t* a, const uint16_t* b, int n) {
    int diff = 0;
    for (int i = 0; i < n; i++)
        if (a[i] != b[i]) diff++;
    return diff;
}

int main(void) {
    printf("=== 图像处理 DLL 全数组版测试 ===\n\n");

    /* ---- 初始化 ---- */
    printf("[1] ip_init...\n");
    IPResult r = ip_init(VENV_DIR, SRC_DIR);
    if (r.code != 0) { printf("  FAIL: %s\n", r.message); return 1; }
    printf("  OK\n");

    /* ---- 准备测试图像 ---- */
    int W = 128, H = 128;
    int N = W * H;
    uint16_t* input  = (uint16_t*)malloc((size_t)N * sizeof(uint16_t));
    uint16_t* output = (uint16_t*)malloc((size_t)N * sizeof(uint16_t));
    uint16_t* out2   = (uint16_t*)malloc((size_t)N * sizeof(uint16_t));
    uint16_t* out3   = (uint16_t*)malloc((size_t)N * sizeof(uint16_t));

    if (!input || !output || !out2 || !out3) {
        printf("  FAIL: 内存分配失败\n");
        goto cleanup;
    }

    /* 生成梯度 + 噪声 */
    for (int y = 0; y < H; y++)
        for (int x = 0; x < W; x++) {
            uint16_t val = (uint16_t)(x * 512 + y * 4);
            if ((x * 7 + y * 13) % 100 == 0) val = (x % 2) ? 0 : 65535;
            input[y * W + x] = val;
        }

    /* ---- 去模糊 ---- */
    printf("\n[2] ip_richardson_lucy...\n");
    r = ip_richardson_lucy(input, output, W, H, 1.5, 5);
    if (r.code) printf("  FAIL: %s\n", r.message);
    else        printf("  OK, %d 像素变化\n", count_diff(input, output, N));

    printf("\n[3] ip_wiener_deblur...\n");
    r = ip_wiener_deblur(input, output, W, H, 1.5, 0.1);
    if (r.code) printf("  FAIL: %s\n", r.message);
    else        printf("  OK, %d 像素变化\n", count_diff(input, output, N));

    /* ---- 去噪 ---- */
    printf("\n[4] ip_adaptive_median_filter...\n");
    r = ip_adaptive_median_filter(input, output, W, H, 3);
    if (r.code) printf("  FAIL: %s\n", r.message);
    else        printf("  OK, %d 像素变化\n", count_diff(input, output, N));

    printf("\n[5] ip_weighted_median_filter...\n");
    r = ip_weighted_median_filter(input, output, W, H, 3);
    if (r.code) printf("  FAIL: %s\n", r.message);
    else        printf("  OK, %d 像素变化\n", count_diff(input, output, N));

    printf("\n[6] ip_denoise_bilateral...\n");
    r = ip_denoise_bilateral(input, output, W, H, 5, 0.08, 80);
    if (r.code) printf("  FAIL: %s\n", r.message);
    else        printf("  OK, %d 像素变化\n", count_diff(input, output, N));

    printf("\n[7] ip_nlm_denoise...\n");
    r = ip_nlm_denoise(input, output, W, H, 0.08, 7, 21);
    if (r.code) printf("  FAIL: %s\n", r.message);
    else        printf("  OK, %d 像素变化\n", count_diff(input, output, N));

    printf("\n[8] ip_tv_denoise...\n");
    r = ip_tv_denoise(input, output, W, H, 0.05);
    if (r.code) printf("  FAIL: %s\n", r.message);
    else        printf("  OK, %d 像素变化\n", count_diff(input, output, N));

    printf("\n[9] ip_wavelet_denoise...\n");
    r = ip_wavelet_denoise(input, output, W, H, 0.001, "BayesShrink");
    if (r.code) printf("  FAIL: %s\n", r.message);
    else        printf("  OK, %d 像素变化\n", count_diff(input, output, N));

    printf("\n[10] ip_filter (返回3路输出)...\n");
    r = ip_filter(input, output, out2, out3, W, H, 2, 3, 3);
    if (r.code) printf("  FAIL: %s\n", r.message);
    else        printf("  OK, gauss=%d diff, median=%d diff, mean=%d diff\n",
                       count_diff(input, output, N), count_diff(input, out2, N), count_diff(input, out3, N));

    /* ---- 数学变换 ---- */
    printf("\n[11] ip_gamma_correction...\n");
    r = ip_gamma_correction(input, output, W, H, 0.5);
    if (r.code) printf("  FAIL: %s\n", r.message);
    else        printf("  OK, %d 像素变化\n", count_diff(input, output, N));

    printf("\n[12] ip_log_transform...\n");
    r = ip_log_transform(input, output, W, H, 1.0);
    if (r.code) printf("  FAIL: %s\n", r.message);
    else        printf("  OK, %d 像素变化\n", count_diff(input, output, N));

    printf("\n[13] ip_exp_transform...\n");
    r = ip_exp_transform(input, output, W, H, 1.0);
    if (r.code) printf("  FAIL: %s\n", r.message);
    else        printf("  OK, %d 像素变化\n", count_diff(input, output, N));

    /* ---- 锐化 ---- */
    printf("\n[14] ip_imagej_sharpen...\n");
    r = ip_imagej_sharpen(input, output, W, H);
    if (r.code) printf("  FAIL: %s\n", r.message);
    else        printf("  OK, %d 像素变化\n", count_diff(input, output, N));

    printf("\n[15] ip_sharpen (USM)...\n");
    r = ip_sharpen(input, output, W, H, 2.0, 1.5);
    if (r.code) printf("  FAIL: %s\n", r.message);
    else        printf("  OK, %d 像素变化\n", count_diff(input, output, N));

cleanup:
    free(input);
    free(output);
    free(out2);
    free(out3);

    printf("\n[16] ip_finalize...\n");
    ip_finalize();
    printf("  OK\n");

    printf("\n=== 测试完成 ===\n");
    return 0;
}
