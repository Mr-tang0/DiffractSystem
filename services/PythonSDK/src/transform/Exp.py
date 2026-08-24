import numpy as np
import cv2

def exp_transform_16bit(img_raw, c=1.0):
    """
    16bit(uint16)图像 指数(反对数)变换
    公式：S = c * (exp(r) - 1)
    r 归一化 [0,1]
    作用：压缩暗灰度区域，拉伸高亮区域，增强亮区对比度；暗部细节会被抑制
    :param img_raw: 输入 uint16 图像 (numpy ndarray / cv2 Mat)
    :param c: 比例系数
    :return: 变换后 uint16 图像数组 (numpy ndarray)
    """
    assert img_raw.dtype == np.uint16, "输入图像必须为uint16 16bit"

    # 归一化到 [0, 1]
    img_float = img_raw.astype(np.float32) / 65535.0

    # exp变换
    img_exp = c * (np.exp(img_float) - 1.0)

    # exp变换后值域不再0‑1，min‑max归一化映射回0~1
    img_exp = (img_exp - np.min(img_exp)) / (np.max(img_exp) - np.min(img_exp) + 1e-8)
    img_exp = np.clip(img_exp, 0.0, 1.0)

    # 转回16bit
    img_out = (img_exp * 65535).astype(np.uint16)

    return img_out


if __name__ == "__main__":
    from skimage import io
    img = cv2.imread("../../imgs/al-20s-50kv1000ua.tif")
    res = exp_transform_16bit(img, c=1.0)
    print(f"shape:{res.shape}, dtype:{res.dtype}")