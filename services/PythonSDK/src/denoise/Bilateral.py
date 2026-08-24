import cv2
import numpy as np

def denoise_bilateral_16bit(img_16bit: np.ndarray, d=5, sigma_color=0.08, sigma_space=80):
    """
    16bit(uint16)单通道图像双边滤波
    :param img_16bit: 输入 uint16 图像
    :param d: 滤波邻域直径
    :param sigma_color: float32域灰度sigma，0‑1范围，一般0.05‑0.15
    :param sigma_space: 空间sigma
    :return: 滤波后 uint16 图像
    """
    if img_16bit.dtype != np.uint16:
        raise ValueError("输入必须为uint16 16bit图像")

    # 16bit [0‑65535] 归一化 float32 [0‑1]
    img_float32 = img_16bit.astype(np.float32) / 65535.0

    # 在float32域执行双边滤波
    filtered_float = cv2.bilateralFilter(img_float32, d=d, sigmaColor=sigma_color, sigmaSpace=sigma_space)

    # 还原回16bit，裁剪防止越界
    filtered_uint16 = np.clip(filtered_float * 65535.0, 0, 65535).astype(np.uint16)
    return filtered_uint16



if __name__ == "__main__":
    # 读取16bit图像，cv2.IMREAD_UNCHANGED必须加，否则自动转为8bit
    img_raw = cv2.imread("../../imgs/al-20s-50kv1000ua.tif", cv2.IMREAD_UNCHANGED)
    print(f"输入图像 dtype:{img_raw.dtype}, shape:{img_raw.shape}")

    # 双边滤波处理
    img_bilateral = denoise_bilateral_16bit(img_raw, d=7, sigma_color=120, sigma_space=120)

    # 保存16bit结果
    cv2.imwrite("../../imgs/output_Bilateral.tif", img_bilateral)

