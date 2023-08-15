import onnxruntime as ort
from PIL import Image
import os
import numpy as np
from numpy import ndarray

class Resnet50(object):
    _instance = None
    transform = None
    ort_sess = None
    def __new__(cls, *args, **kw):
        if cls._instance is None:
            
            current_work_dir = os.path.dirname(__file__)  # 当前文件所在的目录
            onnx_file=os.path.join(current_work_dir,'resnet50_v2.onnx')
            if not os.path.exists(onnx_file):
                print("onnx file not exist")
            cls.ort_sess=ort.InferenceSession(onnx_file)
            cls._instance = object.__new__(cls, *args, **kw)
        return cls._instance
    def __init__(self):
        pass
    
    def toVector(self,img_path)->any:
        img = Image.open(img_path).convert('RGB')

        img_tensor=preprocess(img)[np.newaxis,:]
        print(img_tensor)
        outputs = self.ort_sess.run(None,{"input":img_tensor})
        
        # 范数归一化
        features = outputs[0][0]/np.linalg.norm(outputs[0][0])
        print(features)

        return features.tolist()


def preprocess(image: Image.Image) -> ndarray:
    w,h=256,256
    if image.width>image.height:
        w=h/image.height*image.width
    elif image.width<image.height:
        h=w/image.width*image.height
    resized_image = image.resize((int(w),int(h)))
    

    x0,y0,x1,y1=0,0,w,h
    if w>h:
        x0=(w-h)/2
        x1=w-x0
    else:
        y0=(h-w)/2
        y1=h-y0
    
    resized_image = resized_image.crop((int(x0),int(y0),int(x1),int(y1)))

    resized_image.save("./rrr.png")
    resized_image_ndarray = np.array(resized_image)
    transposed_image_ndarray = resized_image_ndarray.transpose((2, 0, 1))
    transposed_image_ndarrayfloat32 = transposed_image_ndarray.astype(np.float32)
    transposed_image_ndarrayfloat32 /= 255.0
    # mean = np.array([0.485, 0.456, 0.406]).reshape((3, 1, 1))
    # std = np.array([0.229, 0.224, 0.225]).reshape((3, 1, 1))
    # normalized_image_ndarray = (transposed_image_ndarrayfloat32 - mean) / std
    # normalized_image_ndarrayfloat32 = normalized_image_ndarray.astype(
    #     np.float32)
    return transposed_image_ndarrayfloat32


