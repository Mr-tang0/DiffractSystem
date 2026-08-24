import numpy as np

def gamma_correction_16bit(img_raw, gamma):
    """
    16bit(uint16)图像 Gamma变换（幂律变换）
    I_out = I_in^gamma ，输入归一化到 [0,1]
    :param img_raw: 输入 uint16 图像 (numpy ndarray / cv2 Mat)
    :param gamma: γ值
        gamma < 1：图像变亮，拉伸暗区；
        gamma > 1：图像变暗，压缩暗区，增强亮区对比度；
        gamma = 1：图像不变
    :return: 变换后 uint16 图像数组 (numpy ndarray)
    """
    assert img_raw.dtype == np.uint16, "输入必须为uint16 16bit图像"

    # 归一化 [0, 1]
    img_float = img_raw.astype(np.float32) / 65535.0

    # Gamma幂律变换
    img_gamma = np.power(img_float, gamma)

    # 截断，防止浮点计算数值越界
    img_gamma = np.clip(img_gamma, 0.0, 1.0)

    # 映射回16bit
    img_out = (img_gamma * 65535).astype(np.uint16)

    return img_out


if __name__ == "__main__":
    from skimage import io
    img = io.imread("../../imgs/al-20s-50kv1000ua.tif")
    res = gamma_correction_16bit(img, gamma=0.1)
    print(f"shape:{res.shape}, dtype:{res.dtype}")