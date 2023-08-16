# model description

the resnet50_v2.onnx from torch`model

```Python
import torch
import torchvision.models as models
import torchvision.transforms as transforms
from PIL import Image
import torch.nn as nn
import torch.nn.functional as F
import torch

# should only cpu
device = torch.device('cuda:0' if torch.cuda.is_available() else 'cpu')

# 加载 ResNet50 模型并设置为评估模式
resnet50 = models.resnet50(weights=models.ResNet50_Weights.DEFAULT)
model = resnet50.eval().to(device)
resnet50.fc = nn.Linear(2048, 2048)
torch.nn.init.eye(resnet50.fc.weight)  # 将二维tensor初始化为单位矩阵

for param in resnet50.parameters():
    param.requires_grad = False

x = torch.randn(1, 3, 256, 256).to(device)
output = model(x)

with torch.no_grad():
    torch.onnx.export(
        model,                       # 要转换的模型
        x,                           # 模型的任意一组输入
        'test_resnet50.onnx',    # 导出的 ONNX 文件名
        opset_version=11,            # ONNX 算子集版本
        input_names=['input'],       # 输入 Tensor 的名称（自己起名字）
        output_names=['output']      # 输出 Tensor 的名称（自己起名字）
    )
```

model info

```Python
class ResNet50_Weights(WeightsEnum):
    IMAGENET1K_V1 = Weights(
        url="https://download.pytorch.org/models/resnet50-0676ba61.pth",
        transforms=partial(ImageClassification, crop_size=224),
        meta={
            **_COMMON_META,
            "num_params": 25557032,
            "recipe": "https://github.com/pytorch/vision/tree/main/references/classification#resnet",
            "_metrics": {
                "ImageNet-1K": {
                    "acc@1": 76.130,
                    "acc@5": 92.862,
                }
            },
            "_ops": 4.089,
            "_file_size": 97.781,
            "_docs": """These weights reproduce closely the results of the paper using a simple training recipe.""",
        },
    )
    IMAGENET1K_V2 = Weights(
        url="https://download.pytorch.org/models/resnet50-11ad3fa6.pth",
        transforms=partial(ImageClassification, crop_size=224, resize_size=232),
        meta={
            **_COMMON_META,
            "num_params": 25557032,
            "recipe": "https://github.com/pytorch/vision/issues/3995#issuecomment-1013906621",
            "_metrics": {
                "ImageNet-1K": {
                    "acc@1": 80.858,
                    "acc@5": 95.434,
                }
            },
            "_ops": 4.089,
            "_file_size": 97.79,
            "_docs": """
                These weights improve upon the results of the original paper by using TorchVision's `new training recipe
                <https://pytorch.org/blog/how-to-train-state-of-the-art-models-using-torchvision-latest-primitives/>`_.
            """,
        },
    )
    DEFAULT = IMAGENET1K_V2
```