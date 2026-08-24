import numpy as np
import bm3d


def image_bm3d_denoise_16bit(img_raw, sigma_psd=0.04):
    """
    16位单通道图像 BM3D去噪
    适配X射线、中子成像 uint16 图像
    :param img_raw: 输入 uint16 图像 (numpy ndarray / cv2 Mat)
    :param sigma_psd: 噪声强度 [0~1]，低剂量图像常用0.02~0.08
    :return: 去噪后 uint16 图像 (numpy ndarray)
    """
    # 如果是多通道，转为灰度单通道
    if len(img_raw.shape) == 3:
        import cv2
        img_raw = cv2.cvtColor(img_raw, cv2.COLOR_BGR2GRAY)

    if img_raw.dtype != np.uint16:
        raise TypeError("输入图像必须为uint16 16位图像")

    img_float = img_raw.astype(np.float32) / 65535.0

    denoised_float = bm3d.bm3d(
        img_float,
        sigma_psd=sigma_psd,
        stage_arg=bm3d.BM3DStages.ALL_STAGES
    )

    denoised_float = np.clip(denoised_float, 0.0, 1.0)
    denoised_16 = (denoised_float * 65535).astype(np.uint16)

    return denoised_16


if __name__ == "__main__":
    from skimage import io
    img = io.imread("../../imgs/al-20s-50kv1000ua.tif")
    result = image_bm3d_denoise_16bit(img, sigma_psd=0.0004)