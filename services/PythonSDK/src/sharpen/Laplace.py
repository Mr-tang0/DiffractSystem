import numpy as np
from scipy.ndimage import convolve

def imagej_sharpen_16bit(img_raw):
    """
    复刻 ImageJ Process->Sharpen 3×3卷积锐化
    :param img_raw: 输入 uint16 图像 (numpy ndarray / cv2 Mat)
    :return: uint16 锐化图像 (numpy ndarray)
    """
    # ImageJ Sharpen 固定卷积核
    kernel = np.array([
        [ 0, -1,  0],
        [-1,  5, -1],
        [ 0, -1,  0]
    ], dtype=np.float32)

    assert img_raw.dtype == np.uint16, "输入必须uint16 16bit"

    img_float = img_raw.astype(np.float32) / 65535.0

    # 卷积；mode="reflect" 和ImageJ边界处理一致
    img_sharp = convolve(img_float, kernel, mode="reflect")

    # 必须clip，卷积会出现超界，16bit极易溢出
    img_sharp = np.clip(img_sharp, 0.0, 1.0)
    img_out = (img_sharp * 65535).astype(np.uint16)

    return img_out


if __name__ == "__main__":
    from skimage import io
    img = io.imread("../../imgs/al-20s-50kv1000ua.tif")
    res = imagej_sharpen_16bit(img)
    print(f"shape:{res.shape}, dtype:{res.dtype}")