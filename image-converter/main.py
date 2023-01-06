from uuid import uuid4
from pkg.model.resnet50 import Resnet50
from pkg.queue_task.queue_deal import QueueDealImageURL2Vector
from pkg.utils import imggetter
from pkg.utils import imgcheck
import _thread
from lib.bottle import route, run, request
import os


@route('/img2vector/file', method='POST')
def img2ventor():
    upload = request.files.get('upload')
    name, ext = os.path.splitext(upload.filename)
    if ext not in ('.png', '.jpg', '.jpeg'):
        return vectorResp(msg='File extension not allowed.')

    image_path = "./img/"+str(uuid4())+ext
    # appends upload.filename automatically
    upload.save(image_path)

    imgcheck.CheckImg(path=image_path)
    vector = Resnet50().resnet50_extract_feat(img_path=image_path)

    os.remove(path=image_path)
    return vectorResp(vector=vector, success=True)


@route('/img2vector/url', method='POST')
def img2ventor():
    url = request.forms.get("url")
    image_path, ok = imggetter.DownloadUrlImg(url)
    if not ok:
        return vectorResp(msg="url format cannot parse")

    imgcheck.CheckImg(path=image_path)
    vector = Resnet50().resnet50_extract_feat(img_path=image_path)

    os.remove(path=image_path)
    return vectorResp(vector=vector, success=True)


@route('/')
def upload_page():
    page = """
    <form action="/img2vector/file" method="post" enctype="multipart/form-data">
  Category:      <input type="text" name="category" />
  Select a file: <input type="file" name="upload" />
  <input type="submit" value="Start upload" />
</form>
    """
    return page


def vectorResp(vector=[], msg="success", success=False):
    return dict({
        "vector": vector,
        "msg": msg,
        "success": success
    })


if __name__ == '__main__':
    _thread.start_new_thread(QueueDealImageURL2Vector, ())
    run(host='0.0.0.0', port=8080, debug=True)
