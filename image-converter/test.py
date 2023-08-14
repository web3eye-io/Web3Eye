from pkg.model.resnet50 import Resnet50
import torch.nn.functional as F
import torch


# print(len(Resnet50().resnet50_extract_feat(img_path="/home/coast/cybertracer/Web3Eye/image-converter/img/milvus.jpg")))
r1=Resnet50().resnet50_extract_feat(img_path="/home/coast/cybertracer/Web3Eye/image-converter/img/22.png")
r2=Resnet50().resnet50_extract_feat(img_path="/home/coast/cybertracer/Web3Eye/image-converter/img/22_1.png")
# print(len(Resnet50().resnet50_extract_feat(img_path="/home/coast/cybertracer/Web3Eye/image-converter/img/ce521f1c-463d-44e4-9c0d-424eb3e1095f.png")))

t1=torch.FloatTensor(r1)
t2=torch.FloatTensor(r2)
distance = F.pairwise_distance(t1, t2, p=2)
print(distance)
