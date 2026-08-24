import numpy as np
from skimage.restoration import denoise_tv_chambolle

def tv_denoise_16bit(img_raw, weight=0.1):
    """
    16bit(uint16)图像 TV‑Chambolle全变分去噪
    :param img_raw: 输入 uint16 图像 (numpy ndarray / cv2 Mat)
    :param weight: TV正则权重，0.05~0.2；越大平滑越强，易出现阶梯效应
    :return: 去噪后 uint16 图像数组 (numpy ndarray)
    """
    assert img_raw.dtype == np.uint16, "输入必须为uint16 16bit图像"

    img_float = img_raw.astype(np.float32) / 65535.0

    img_denoised = denoise_tv_chambolle(
        img_float,
        weight=weight,
        channel_axis=None
    )

    img_denoised = np.clip(img_denoised, 0.0, 1.0)
    img_out = (img_denoised * 65535).astype(np.uint16)

    return img_out


if __name__ == "__main__":
    from skimage import io
    img = io.imread("../../imgs/al-20s-50kv1000ua.tif")
    res = tv_denoise_16bit(img, weight=0.001)
    print(f"输出图像shape:{res.shape}, dtype:{res.dtype}")