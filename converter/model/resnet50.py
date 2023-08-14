import onnxruntime as ort
from PIL import Image
import torchvision.transforms as transforms
import torch.nn.functional as F
import torch
import os
import aspose.words as aw


class Resnet50(object):
    _instance = None
    transform = None
    ort_sess = None
    def __new__(cls, *args, **kw):
        if cls._instance is None:
            cls.transform=transforms.Compose([
                transforms.Resize(256),
                transforms.CenterCrop(256),
                transforms.ToTensor(),
                # the result is normalizeed, so the image input don`t normalize
                # transforms.Normalize(mean=[0.485, 0.456, 0.406], std=[0.229, 0.224, 0.225])
            ])
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
        if img_path.endswith("svg"):
            pngPath=img_path+".png"
            svg2png(img_path,pngPath)
            img = Image.open(pngPath).convert('RGB')
            os.remove(pngPath)
        else:
            img = Image.open(img_path).convert('RGB')
        img_tensor =self.transform(img).unsqueeze(0) # 增加一个 batch 维度
        outputs = self.ort_sess.run(None,{"input":img_tensor.numpy()})
        features=torch.FloatTensor(outputs[0])
        features = F.normalize(features,dim=1)
        print(features)
        return features.tolist()[0]



def svg2png(filePath,pngPath):
	doc = aw.Document()
	builder = aw.DocumentBuilder(doc)
	shape = builder.insert_image(filePath)

	pageSetup = builder.page_setup
	pageSetup.page_width = shape.width
	pageSetup.page_height = shape.height
	pageSetup.top_margin = 0
	pageSetup.left_margin = 0
	pageSetup.bottom_margin = 0
	pageSetup.right_margin = 0

	doc.save(pngPath)

# resnet50=Resnet50()

# r1= resnet50.toVector("/home/coast/cybertracer/Web3Eye/image-converter-v2/img/22.png")
# r2= resnet50.toVector("/home/coast/cybertracer/Web3Eye/image-converter-v2/img/22_1.png")


# distance1 = F.pairwise_distance(r1, r2, p=2)
# print(distance1)