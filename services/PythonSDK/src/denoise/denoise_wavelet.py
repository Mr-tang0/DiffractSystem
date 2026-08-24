import numpy as np
from skimage.restoration import denoise_wavelet

def wavelet_denoise_16bit(img_raw, sigma=None, method="BayesShrink"):
    """
    16bit(uint16)单通道图像 小波阈值去噪
    :param img_raw: 输入 uint16 图像 (numpy ndarray / cv2 Mat)
    :param sigma: 噪声标准差(归一化0‑1范围)；sigma=None自动估计噪声
    :param method: 阈值策略 "BayesShrink"(推荐) / "VisuShrink"
    :return: denoised uint16 图像数组 (numpy ndarray)
    """
    assert img_raw.dtype == np.uint16, "输入必须为uint16 16bit图像"

    img_float = img_raw.astype(np.float32) / 65535.0

    img_denoised = denoise_wavelet(
        img_float,
        sigma=sigma,
        method=method,
        wavelet="db4",
        mode="soft",
        channel_axis=None,
        rescale_sigma=True
    )

    img_denoised = np.clip(img_denoised, 0.0, 1.0)
    img_out = (img_denoised * 65535).astype(np.uint16)

    return img_out


if __name__ == "__main__":
    from skimage import io
    img = io.imread("../../imgs/al-20s-50kv1000ua.tif")
    res = wavelet_denoise_16bit(img, sigma=0.0003, method="BayesShrink")
    print(f"shape:{res.shape}, dtype:{res.dtype}")