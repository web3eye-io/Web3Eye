#!/usr/bin/env python

import json
import os
import uuid
import urllib3
import logging
from pkg.model.resnet50 import Resnet50
from pkg.utils import imggetter
from pkg.utils import imgcheck
import time
from pkg.utils import config

http = urllib3.PoolManager()


def UUIDCheck(id) -> bool:
    try:
        uuid.UUID(id)
        return True
    except ValueError:
        return False


class VectorInfo(object):
    def __init__(self, id="", url="", vector=[], msg="", success=False) -> None:
        self.id = id
        self.url = url
        self.vector = vector
        self.msg = msg
        self.success = success


class VectorInfoEncoder(json.JSONEncoder):
    def default(self, o):
        if isinstance(o, VectorInfo):
            return {
                'id': o.id,
                'url': o.url,
                'vector': o.vector,
                'msg': o.msg,
                'success': o.success,
            }
        return json.JSONEncoder.default(o)


def QueueDealImageURL2Vector():
    while True:
        try:
            # get tasks
            infos=get_waiting_tokens(10)
            for info in infos:
                try:
                    # init info
                    vectorInfo = VectorInfo(success=False)
                    vectorInfo.url = info["ImageURL"] 
                    vectorInfo.id = info["ID"]

                    if vectorInfo.url == "":
                        raise Exception("have no url".format(vectorInfo.url))
                        
                    image_path, ok = imggetter.DownloadUrlImg(vectorInfo.url)
                    if not ok:
                        vectorInfo.msg = "url format cannot parse,url: " + vectorInfo.url
                        logging.warning(vectorInfo.msg)
                        vectorInfo.success = False
                        raise Exception("download image from {} failed".format(vectorInfo.url))

                    imgcheck.CheckImg(path=image_path)
                    vectorInfo.vector = Resnet50().resnet50_extract_feat(img_path=image_path)

                    os.remove(path=image_path)
                    vectorInfo.success = True
                except Exception as e:
                    vectorInfo.msg=e
                finally:
                    update_token_vstate(vectorInfo.id,vectorInfo.success,vectorInfo.vector,vectorInfo.msg)

        except Exception as e:
            logging.error("parse nft-token image url failed,",e)
       
        time.sleep(5)
        
def report_file_to_gen_car(id:str,file_s3_key)-> bool:
    try:
        data = json.dumps({'ID':id,"S3Key":file_s3_key}).encode()
        http.request(
            method="POST",
            url=f"http://{config.gen_car_domain}:{config.gen_car_http_port}/v1/report/file",
            body=data
            )
        return True
    except:
        return False

def get_waiting_tokens(limit:int)-> list:
    try:
        data = json.dumps({"Conds":{"VectorState":{"Op":"eq","Value":"Waiting"}},"Limit":limit}).encode()
        resp=http.request(
            method="POST",
            url=f"http://{config.nft_meta_domain}:{config.nft_meta_http_port}/v1/get/tokens",
            body=data
            )
        
        return json.loads(resp.data)["Infos"]
    except Exception as e:
        logging.error(e)
        return []
    
def update_token_vstate(id:str,vstate:bool,vector:[],msg:str)-> any:
    vector_state="Success"
    if not vstate:
        vector=[]
        vector_state="Failed"
    try:
        data = json.dumps({"ID":id,"Vector":vector,"VectorState":vector_state,"Remark":f"{msg}"}).encode()
        resp=http.request(
            method="POST",
            url=f"http://{config.nft_meta_domain}:{config.nft_meta_http_port}/v1/update/image/vector",
            body=data
            )
        print(json.loads(resp.data))
        return json.loads(resp.data)["Info"]
    except Exception as e:
        logging.error(e)
        return 