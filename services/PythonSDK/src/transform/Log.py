import numpy as np

def log_transform_16bit(img_raw, c=1.0):
    """
    16bit(uint16)图像 对数变换
    公式：S = c * log(1 + r)
    r 归一化到 [0,1]
    作用：拉伸暗灰度区域，压缩高亮度区域，适合大动态范围X射线/中子图像
    :param img_raw: 输入 uint16 图像 (numpy ndarray / cv2 Mat)
    :param c: 比例系数
    :return: 变换后 uint16 图像数组 (numpy ndarray)
    """
    assert img_raw.dtype == np.uint16, "输入图像必须为uint16 16bit"

    # 归一化到 [0, 1]
    img_float = img_raw.astype(np.float32) / 65535.0

    # log(1+x) 避免log(0)
    img_log = c * np.log1p(img_float)

    # log变换后值域不再0‑1，执行min‑max归一化映射回0~1
    img_log = (img_log - np.min(img_log)) / (np.max(img_log) - np.min(img_log) + 1e-8)
    img_log = np.clip(img_log, 0.0, 1.0)

    # 转回16bit
    img_out = (img_log * 65535).astype(np.uint16)

    return img_out


if __name__ == "__main__":
    from skimage import io
    img = io.imread("../../imgs/al-20s-50kv1000ua.tif")
    res = log_transform_16bit(img, c=1.0)
    print(f"shape:{res.shape}, dtype:{res.dtype}")