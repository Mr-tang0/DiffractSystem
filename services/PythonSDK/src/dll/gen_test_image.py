"""生成测试用的 uint16 tif 图像"""
import numpy as np
import cv2

W, H = 256, 256
img = np.zeros((H, W), dtype=np.uint16)

for y in range(H):
    for x in range(W):
        val = (x * 256 + y * 16) % 65536
        if (x * 7 + y * 13) % 100 == 0:
            val = 65535 if (x % 2) else 0
        img[y, x] = val

cv2.imwrite("test_input.tif", img)
print(f"生成测试图像: {W}x{H}, dtype={img.dtype}")
