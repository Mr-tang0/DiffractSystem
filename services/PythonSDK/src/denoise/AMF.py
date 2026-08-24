import numpy as np
import cv2 as cv


def adaptiveMedianFilter(image, windowSize):
    # padding
    top = bottom = left = right = windowSize
    padImg = cv.copyMakeBorder(image, top, bottom, left, right, cv.BORDER_REPLICATE, value=0)

    row = padImg.shape[0]
    col = padImg.shape[1]

    sMax = 7
    a = sMax // 2
    filter_output = np.zeros(padImg.shape)
    for y in range(a, row - 1):
        for x in range(a, col - 1):
            filter_output[y, x] = stageA(padImg, y, x, windowSize, sMax)

    return filter_output[a:-a, a:-a].astype(np.uint16)


def stageA(img, y, x, windowSize, sMax):
    img_part = img[y - (windowSize // 2):y + (windowSize // 2) + 1, x - (windowSize // 2):x + (windowSize // 2) + 1]

    zmin = np.min(img_part)
    zmed = np.median(img_part)
    zmax = np.max(img_part)

    A1 = zmed - zmin
    A2 = zmed - zmax

    if A1 > 0 and A2 < 0:  # go to level B
        return stageB(img_part, zmin, zmed, zmax)
    else:
        windowSize = windowSize + 2  # increase window size (must be odd so add 2)
        if windowSize <= sMax:  # window size is lower than Smax, repeat level A
            return stageA(img, y, x, windowSize, sMax)
        else:  # return zmed
            return zmed


def stageB(img, zmin, zmed, zmax):
    h, w = img.shape
    zxy = img[h // 2, w // 2]

    B1 = zxy - zmin
    B2 = zxy - zmax

    if B1 > 0 and B2 < 0:  # return Zxy
        return zxy
    else:
        return zmed  # return zmed


# From HW3
# center weighted median filter
def weightedMedianFilter(image, kernelSize):
    h, w = image.shape
    # padding
    top = bottom = int((kernelSize - 1) / 2)  # rows
    left = right = int((kernelSize - 1) / 2)  # cols
    padImg = cv.copyMakeBorder(image, top, bottom, left, right, cv.BORDER_REPLICATE)

    imageW, imageH = image.shape  # get image dimensions
    filter_out = np.zeros((imageW, imageH))  # intialize output image
    # convolution
    for i in range(w):
        for j in range(h):
            temp = padImg[i:i + kernelSize, j:j + kernelSize]
            centerValue = temp[kernelSize // 2, kernelSize // 2]
            flattenedImg = temp.flatten()
            flattenedImg = np.append(flattenedImg, centerValue)
            flattenedImg = np.append(flattenedImg, centerValue)
            median = np.median(flattenedImg)
            filter_out[i, j] = median
    filter_out = filter_out.astype(np.uint16)
    return filter_out


if __name__ == "__main__":
    img_filename = '../../imgs/al-20s-50kv1000ua.tif'
    image = cv.imread(img_filename, -1)

    output_AMF = adaptiveMedianFilter(image, 3)
    output_WMF = weightedMedianFilter(image, 3)

    cv.imwrite('../../imgs/output_AMF.tif', output_AMF)
    cv.imwrite('../../imgs/output_WMF.tif', output_WMF)
