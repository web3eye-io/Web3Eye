from fileinput import close
from typing import Tuple
from uuid import uuid4
import urllib3

typedic = {
    'image/png': "png",
    'image/jpeg': "jpg",
}

IPFS_HTTP_Gateway = "https://ipfs.io/"


def DownloadUrlImg(url) -> Tuple[str, bool]:
    if url.startswith("http"):
        return DownloadHttpImg(url=url)
    elif url.startswith("ipfs"):
        return DownloadIPFSImg(url=url)
    # just try to download with http
    return DownloadHttpImg(url=url)


def DownloadIPFSImg(url) -> Tuple[str, bool]:
    url = url.replace("//", "/")
    url = url.replace(":", "")
    url = f"{IPFS_HTTP_Gateway}/{url}"
    return DownloadHttpImg(url=url)


def DownloadHttpImg(url) -> Tuple[str, bool]:
    urllib3.disable_warnings()
    http = urllib3.PoolManager()

    try:
        req = http.urlopen("GET", url=url, preload_content=False)
    except Exception:
        return "", False

    meta = req.info()
    content_type = str(meta.getheaders("Content-Type")[0])
    accept_ranges = str(meta.getheaders("Accept-Ranges")[0])

    if not content_type in typedic.keys():
        return "", False
    if not accept_ranges == "bytes":
        return "", False

    file_path = f"./tmp/{str(uuid4())}.{typedic[content_type]}"
    file = open(file_path, 'wb')

    file_size_dl = 0
    block_sz = 4*1024
    while True:
        buffer = req.read(block_sz)
        if not buffer:
            break

        file_size_dl += len(buffer)
        file.write(buffer)
    file.close()
    http.clear()
    return file_path, True


url = "https://mirrors.aliyun.com/deepin-cd/20.1/deepin-desktop-community-1010-amd64.iso"
url = "https://ipfs.io/ipfs/QmdJk8kfwacmT4FPXEDTQax9bvSYxDy5XNr5rZWq46f3ip/Teddies_hidden.mp4"
url = "https://ipfs.io/ipfs/QmQqzMTavQgT4f4T5v6PWBp7XNKtoPmC9jvn12WPT3gkSE"
url = "https://ipfs.io/ipfs/QmddokWqSLYp1vUP4XNaYzAbdDWeDLA4uyapN9fsDrSRv2/3679.png"

ipfsurl = "ipfs://QmdJk8kfwacmT4FPXEDTQax9bvSYxDy5XNr5rZWq46f3ip/Teddies_hidden.mp4"
ipfsurl = "ipfs://QmddokWqSLYp1vUP4XNaYzAbdDWeDLA4uyapN9fsDrSRv2/3679.png"
ipfsurl = "ipfs://QmQqzMTavQgT4f4T5v6PWBp7XNKtoPmC9jvn12WPT3gkSE"
DownloadUrlImg(url=ipfsurl)
