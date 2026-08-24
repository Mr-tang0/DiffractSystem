"""
图像处理函数统一接口模块
=========================
将 deblur、denoise、transform、sharpen 各子目录下的图像处理函数统一导入到本文件。

所有函数均接收 numpy ndarray (cv2 Mat) 作为输入，返回 numpy ndarray 作为输出，
不涉及文件 I/O。

使用方式：
    import cv2
    from image_processing import gamma_correction_16bit, sharpen_16bit
    img = cv2.imread("input.tif", cv2.IMREAD_UNCHANGED)
    result = gamma_correction_16bit(img, gamma=0.5)
    cv2.imwrite("output.tif", result)

所有函数均面向 uint16 16bit 图像（X射线/中子成像等）。
"""

import os
import sys

_SRC_DIR = os.path.dirname(os.path.abspath(__file__))
if _SRC_DIR not in sys.path:
    sys.path.insert(0, _SRC_DIR)

# =====================================================================
#  去模糊 (deblur)
# =====================================================================

from deblur.Richardson_Lucy import richardson_lucy_16bit
from deblur.wiener import wiener_deblur_16bit

# =====================================================================
#  去噪 (denoise)
# =====================================================================

from denoise.AMF import adaptiveMedianFilter, weightedMedianFilter
from denoise.Bilateral import denoise_bilateral_16bit
from denoise.denoise_tv import tv_denoise_16bit
from denoise.denoise_wavelet import wavelet_denoise_16bit
from denoise.Filters import filter_16bit
from denoise.NLM import nlm_denoise_16bit_skimage

try:
    from denoise.BM3D import image_bm3d_denoise_16bit
except ImportError as _e:
    _bm3d_err = str(_e)
    def image_bm3d_denoise_16bit(*args, **kwargs):
        raise ImportError(
            f"函数 'image_bm3d_denoise_16bit' 不可用：未安装 bm3d 库\n"
            f"原始错误：{_bm3d_err}\n"
            f"请执行 pip install bm3d 后重试。"
        )

# =====================================================================
#  数学变换 (transform)
# =====================================================================

from transform.Exp import exp_transform_16bit
from transform.Gamma import gamma_correction_16bit
from transform.Log import log_transform_16bit

# =====================================================================
#  锐化 (sharpen)
# =====================================================================

from sharpen.Laplace import imagej_sharpen_16bit
from sharpen.USM import sharpen_16bit


# =====================================================================
#  __all__ 导出清单
# =====================================================================

__all__ = [
    # 去模糊
    "richardson_lucy_16bit",
    "wiener_deblur_16bit",
    # 去噪
    "adaptiveMedianFilter",
    "weightedMedianFilter",
    "denoise_bilateral_16bit",
    "image_bm3d_denoise_16bit",
    "tv_denoise_16bit",
    "wavelet_denoise_16bit",
    "filter_16bit",
    "nlm_denoise_16bit_skimage",
    # 数学变换
    "exp_transform_16bit",
    "gamma_correction_16bit",
    "log_transform_16bit",
    # 锐化
    "imagej_sharpen_16bit",
    "sharpen_16bit",
]
