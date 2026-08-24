import numpy as np
from skimage.restoration import wiener

def create_gaussian_psf(sigma):
    """生成高斯点扩散函数PSF小核"""
    size = int(2 * np.ceil(3 * sigma) + 1)
    ax = np.arange(-size // 2 + 1., size // 2 + 1.)
    xx, yy = np.meshgrid(ax, ax)
    kernel = np.exp(-(xx**2 + yy**2) / (2.0 * sigma**2))
    kernel = kernel / np.sum(kernel)
    return kernel


def wiener_deblur_16bit(img_raw, psf_sigma=1.2, balance=0.02):
    """
    16bit(uint16)图像 维纳滤波反卷积去模糊
    :param img_raw: 输入 uint16 图像 (numpy ndarray / cv2 Mat)
    :param psf_sigma: 高斯PSF的sigma，模糊越大数值越大 0.8‑2.0
    :param balance: 噪声平衡系数；0.01~0.05；越大对噪声抑制越强，细节恢复越弱
    :return: 去模糊后 uint16 图像 (numpy ndarray)
    """
    assert img_raw.dtype == np.uint16, "输入必须为uint16 16bit图像"

    img_float = img_raw.astype(np.float32) / 65535.0

    psf = create_gaussian_psf(psf_sigma)

    img_deblur = wiener(img_float, psf=psf, balance=balance)

    img_deblur = np.clip(img_deblur, 0.0, 1.0)
    img_out = (img_deblur * 65535).astype(np.uint16)

    return img_out


if __name__ == "__main__":
    from skimage import io
    img = io.imread("../../imgs/al-20s-50kv1000ua.tif")
    res = wiener_deblur_16bit(img, psf_sigma=1.5, balance=0.1)
    print(f"shape:{res.shape}, dtype:{res.dtype}")