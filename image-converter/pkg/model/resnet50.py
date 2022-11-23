import towhee


class Resnet50:
    """
    Say something about the ExampleCalass...

    Args:
        args_0 (`type`):
        ...
    """

    # TODO: research how to adjust model params
    def resnet50_extract_feat(self, img_path):
        vector = towhee.glob(img_path) \
            .image_decode() \
            .image_embedding.timm(model_name='resnet50') \
            .tensor_normalize() \
            .to_list()

        # convert numpy array to float list
        # convert for json dump
        vector = vector[0].astype(float)
        vector = vector.tolist()
        return vector
