# ============ #
# Load library #
# ============ #
import string
import sys
import torch
import cv2 as cv
from torch import nn
import json


def train(s):
    # ============== #
    # Argument Check #
    # ============== #

    # if len(sys.argv) != 3:
    #   print("The format shoud be: \npython cat_dog_classifier.py ./model/CNN_model_weights.pth ./image/Cat_Sample_01.jpg")
    #  exit()
   # if type(sys.argv[1]) != str:
    #    print(
    #       "The second argument should be string, i.e., ./PATH/TO/YOUR_TRAINED_MODEL.pth")
    #  exit()
    # if type(sys.argv[2]) != str:
    #   print("The third argument should be string, i.e., ./PATH/TO/YOU_NAME_IT.jpg")
    #  exit()

    # print("Image Path: ", sys.argv[1])
    # print("Model Path: ", sys.argv[2])
    model_path = ".\\SSG\\Golang_AI\\Model\\CNN_model_weights.pth"
    image_path = s

    # ================ #
    # Define CNN Class #
    # ================ #

    class CNN_v1(nn.Module):

        # I add one more param here, i.e., img_size, for changing CNN structure auto
        def __init__(self, img_size):
            super(CNN_v1, self).__init__()

            self.img_size = img_size  # assume (B, C=3, H=256, W=256)

            self.cspec = [3, 64, 128, 256, 512, 1024,
                          512, 256]  # cspec stands for conv spec
            # fspec stands for fully connected layer spec
            self.fspec = [128, 64, 1]

            self.repeat_conv = nn.Sequential(

                # 換換不同的寫法 v1
                nn.Conv2d(
                    in_channels=self.cspec[0],
                    out_channels=self.cspec[1],
                    # this could be tuple, i.e., (3,3), or just integer i.e., 3.
                    kernel_size=(3, 3),
                    stride=2,  # based on the calculator mentioned above, this setting will make spatial size half
                    padding=1
                ),
                nn.ReLU(),                   # (B, C=  64, H=128, W=128)
                nn.BatchNorm2d(self.cspec[1]),

                # 換換不同的寫法 v2
                nn.Conv2d(in_channels=self.cspec[1], out_channels=self.cspec[2], kernel_size=(
                    3, 3), stride=2, padding=1),
                nn.ReLU(),                   # (B, C= 128, H= 64, W= 64)
                nn.BatchNorm2d(self.cspec[2]),

                nn.Conv2d(in_channels=self.cspec[2], out_channels=self.cspec[3], kernel_size=(
                    3, 3), stride=2, padding=1),
                nn.ReLU(),                   # (B, C= 256, H= 32, W= 32)
                nn.BatchNorm2d(self.cspec[3]),

                nn.Conv2d(in_channels=self.cspec[3], out_channels=self.cspec[4], kernel_size=(
                    3, 3), stride=2, padding=1),
                nn.ReLU(),                   # (B, C= 512, H= 16, W= 16)
                nn.BatchNorm2d(self.cspec[4]),

                nn.Conv2d(in_channels=self.cspec[4], out_channels=self.cspec[5], kernel_size=(
                    3, 3), stride=2, padding=1),
                nn.ReLU(),                   # (B, C=1024, H=  8, W=  8)
                nn.BatchNorm2d(self.cspec[5]),

                nn.Conv2d(in_channels=self.cspec[5], out_channels=self.cspec[6], kernel_size=(
                    3, 3), stride=2, padding=1),
                nn.ReLU(),                   # (B, C= 512, H=  4, W=  4)
                nn.BatchNorm2d(self.cspec[6]),

                nn.Conv2d(in_channels=self.cspec[6], out_channels=self.cspec[7], kernel_size=(
                    3, 3), stride=2, padding=1),
                nn.ReLU(),                   # (B, C= 256, H=  2, W=  2)
                nn.BatchNorm2d(self.cspec[7]),
            )

            self.flatten = nn.Flatten()

            # 需要優化 #
            C = 256
            H = 1
            W = H  # assume square
            self.repeat_dense = nn.Sequential(
                nn.Linear(in_features=C*H*W, out_features=self.fspec[0]),
                nn.ReLU(),
                nn.Linear(in_features=self.fspec[0],
                          out_features=self.fspec[1]),
                nn.ReLU(),
                nn.Linear(in_features=self.fspec[1],
                          out_features=self.fspec[2]),
            )

        def forward(self, img):
            feature_map = self.repeat_conv(img)
            features = self.flatten(feature_map)
            logits = self.repeat_dense(features)
            return logits

        def __preprocess__(self, img_path):
            image = cv.imread(img_path)
            image = cv.resize(image, (self.img_size, self.img_size))
            image = image / 255.

            if None:
                image = self.x_transform(image)

            # channel first
            image = image.reshape(3, self.img_size, self.img_size)
            return image

    # ======================== #
    # Initialize CNN structure #
    # ======================== #
    device = "cuda" if torch.cuda.is_available() else "cpu"

    CNN_model = CNN_v1(img_size=128).to(device)

    try:
        CNN_model.load_state_dict(torch.load(
            model_path, map_location=torch.device(device)))
        test_img_numpy = CNN_model.__preprocess__(image_path)
        test_img_tensor = torch.tensor(test_img_numpy).to(device)
        test_img_tensor = torch.unsqueeze(test_img_tensor, 0).float()
        CNN_model.eval()
    except Exception as e:
        print(e)

    # ================ #
    # Preprocess input #
    # ================ #

    # =================== #
    # Predict input image #
    # =================== #
    try:
        with torch.no_grad():
            logit = CNN_model(test_img_tensor)
            logit = logit.cpu().numpy()[0][0]
            if logit > 0:

                result = {"result": "dog", "logit": logit}

            else:
                result = {"result": "cat", "logit": logit}

    except Exception as e:
        print(e)

    print(str(result))
    try:
        a = str(result)
    except Exception as e:
        print(e)
    return(a)


def Train(s):

    #sys.argv[1] = "./Golang_AI/model/CNN_model_weights.pth"
    #sys.argv[2] = "..\image\20220406192317WIN_20220216_08_47_44_Pro.jpg"
    a = train(s)
    print("result======"+a+"\n")
    return train(s)
