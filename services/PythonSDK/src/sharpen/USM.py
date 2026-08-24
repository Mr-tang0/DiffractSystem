import numpy as np
from skimage.filters import unsharp_mask

def sharpen_16bit(img_raw, radius=1.0, amount=1.2):
    """
    16bit(uint16)单通道图像 USM非锐化掩模锐化
    :param img_raw: 输入 uint16 图像 (numpy ndarray / cv2 Mat)
    :param radius: 高斯模糊半径，控制锐化空间尺度，1.0~2.5
    :param amount: 锐化强度；0关闭；1.0~2.0常用；数值越大锐化越强，噪声同步放大
    :return: 锐化后 uint16 图像数组 (numpy ndarray)
    """
    assert img_raw.dtype == np.uint16, "输入必须为uint16 16bit图像"

    # 归一化0‑1 float
    img_float = img_raw.astype(np.float32) / 65535.0

    # USM非锐化掩模锐化
    img_sharp = unsharp_mask(
        img_float,
        radius=radius,
        amount=amount,
        channel_axis=None
    )

    # 截断，防止数值溢出，转回16bit
    img_sharp = np.clip(img_sharp, 0.0, 1.0)
    img_out = (img_sharp * 65535).astype(np.uint16)

    return img_out


if __name__ == "__main__":
    from skimage import io
    img = io.imread("../../imgs/al-20s-50kv1000ua.tif")
    res = sharpen_16bit(img, radius=4, amount=2)
    print(f"shape:{res.shape}, dtype:{res.dtype}")