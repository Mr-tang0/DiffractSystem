import numpy as np
from skimage.filters import gaussian
from skimage.filters import median
from skimage.morphology import square
from scipy.ndimage import uniform_filter

def filter_16bit(img_raw, sigma=1.2, median_size=3, mean_size=3):
    """
    16bit(uint16)单通道图像：高斯滤波、中值滤波、均值滤波
    :param img_raw: 输入 uint16 图像 (numpy ndarray / cv2 Mat)
    :param sigma: 高斯核标准差
    :param median_size: 中值滤波窗口大小(奇数)
    :param mean_size: 均值滤波窗口大小(奇数)
    :return: gauss_out, median_out, mean_out  三个 uint16 图像数组
    """
    assert img_raw.dtype == np.uint16, "输入必须为uint16 16bit图像"

    img_float = img_raw.astype(np.float32) / 65535.0

    img_gauss = gaussian(img_float, sigma=sigma, channel_axis=None)
    img_med = median(img_float, footprint=square(median_size))
    img_mean = uniform_filter(img_float, size=mean_size)

    def to_uint16(arr):
        arr = np.clip(arr, 0.0, 1.0)
        return (arr * 65535).astype(np.uint16)

    gauss_out = to_uint16(img_gauss)
    median_out = to_uint16(img_med)
    mean_out = to_uint16(img_mean)

    return gauss_out, median_out, mean_out


if __name__ == "__main__":
    from skimage import io
    img = io.imread("../../imgs/al-20s-50kv1000ua.tif")
    g, m, me = filter_16bit(img, sigma=2, median_size=3, mean_size=3)
    print(f"gauss shape:{g.shape}, dtype:{g.dtype}")