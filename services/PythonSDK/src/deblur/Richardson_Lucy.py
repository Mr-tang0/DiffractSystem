import numpy as np
from skimage.restoration import richardson_lucy

def create_gaussian_psf(sigma):
    size = int(2 * np.ceil(3 * sigma) + 1)
    ax = np.arange(-size // 2 + 1., size // 2 + 1.)
    xx, yy = np.meshgrid(ax, ax)
    kernel = np.exp(-(xx**2 + yy**2) / (2.0 * sigma**2))
    kernel = kernel / np.sum(kernel)
    return kernel


def richardson_lucy_16bit(img_raw, psf_sigma=1.2, num_iter=30):
    assert img_raw.dtype == np.uint16, "输入必须为uint16 16bit图像"

    img_float = img_raw.astype(np.float32) / 65535.0

    psf = create_gaussian_psf(psf_sigma)

    img_deblur = richardson_lucy(
        img_float,
        psf=psf,
        num_iter=num_iter,
        clip=False
    )

    img_deblur = np.clip(img_deblur, 0.0, 1.0)
    img_out = (img_deblur * 65535).astype(np.uint16)

    return img_out


if __name__ == "__main__":
    from skimage import io
    img = io.imread("../../imgs/al-20s-50kv1000ua.tif")
    res = richardson_lucy_16bit(img, psf_sigma=1.5, num_iter=5)
    print(f"shape:{res.shape}, dtype:{res.dtype}")