from pkg.model.resnet50 import Resnet50

# Load the operator ahead of time
Resnet50().resnet50_extract_feat(img_path="./tmp/milvus.jpg")
