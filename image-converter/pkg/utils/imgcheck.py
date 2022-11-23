import cv2
from skimage import io


def CheckImg(path):
    if path.endswith("png"):
        image = io.imread(path)
        image = cv2.cvtColor(image, cv2.COLOR_RGBA2BGRA)
        cv2.imencode('.png', image)[1].tofile(path)
