import cv2
import numpy as np
from skimage.restoration import denoise_nl_means, estimate_sigma

def nlm_denoise_16bit_skimage(img_16bit: np.ndarray, h_factor=0.08,
                              patch_size=7, patch_distance=21):
    """
    NLM非局部均值去噪，原生支持uint16 16bit图像
    :param img_16bit: uint16 H×W
    :param h_factor: 滤波系数，0‑1；噪声越大调大；射线图像0.04‑0.12
    :param patch_size: 块大小，奇数
    :param patch_distance: 搜索窗口
    :return: denoised_16 uint16
    """
    if len(img_16bit.shape) == 3:
        img_16bit = cv2.cvtColor(img_16bit, cv2.COLOR_BGR2GRAY)

    # 归一化到[0,1]浮点
    img_float = img_16bit.astype(np.float32) / 65535.0

    denoised_float = denoise_nl_means(
        img_float,
        h=h_factor,
        patch_size=patch_size,
        patch_distance=patch_distance,
        fast_mode=True
    )

    denoised_float = np.clip(denoised_float, 0.0, 1.0)
    denoised_16 = (denoised_float * 65535).astype(np.uint16)
    return denoised_16


if __name__ == "__main__":
    input_tif = r"../../imgs/al-20s-50kv1000ua.tif"
    output_tif = r"../../imgs/output_NLM.tif"

    img_raw = cv2.imread(input_tif, cv2.IMREAD_UNCHANGED)
    if img_raw is None:
        raise FileNotFoundError("读取失败")

    result = nlm_denoise_16bit_skimage(img_raw, h_factor=0.0001)
    cv2.imwrite(output_tif, result)
    print(f"输出 {output_tif}, dtype={result.dtype}")