#!/usr/bin/env python

import json
import os
import uuid
import urllib3
import logging
from pkg.model.resnet50 import Resnet50
from pkg.utils import imggetter
from pkg.utils import imgcheck
from pkg.utils import oss
from pkg.utils import config


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
                    vectorInfo.url = info["TokenURI"] 
                    vectorInfo.id = info["ID"]

                    if vectorInfo.url != "":
                        image_path, ok = imggetter.DownloadUrlImg(vectorInfo.url)
                        if not ok:
                            vectorInfo.msg = "url format cannot parse,url: " + vectorInfo.url
                            logging.warning(vectorInfo.msg)
                            vectorInfo.success = False
                            continue

                        imgcheck.CheckImg(path=image_path)
                        vectorInfo.vector = Resnet50().resnet50_extract_feat(img_path=image_path)

                        os.remove(path=image_path)
                        vectorInfo.success = True
                except Exception as e:
                    vectorInfo.msg=e
                finally:
                    update_token_vstate(vectorInfo.id,vectorInfo.success,vectorInfo.msg)

        except Exception as e:
            logging.error("parse nft-token image url failed,",e)
       

def report_file_to_gen_car(id:str,file_s3_key)-> bool:
    try:
        http = urllib3.PoolManager()
        data = json.dumps({'ID':id,"S3Key":file_s3_key}).encode()
        logging.info(config.gen_car_ip)
        http.request(
            method="POST",
            url=f"http://{config.gen_car_ip}:{config.gen_car_http_port}/v1/report/file",
            body=data
            )
        return True
    except:
        return False

def get_waiting_tokens(limit:int)-> list:
    try:
        http = urllib3.PoolManager()
        data = json.dumps({"Conds":{"VectorState":{"Op":"eq","Value":"Waiting"}},"Limit":limit}).encode()
        config.nft_meta_ip="127.0.0.1"
        resp=http.request(
            method="POST",
            url=f"http://{config.nft_meta_ip}:{config.nft_meta_http_port}/v1/get/tokens",
            body=data
            ).data
        return json.loads(resp)["Infos"]
    except Exception as e:
        logging.error(e)
        return []
    
def update_token_vstate(ID:str,vstate:bool,msg:str)-> any:
    vector_state="Success"
    if not vstate:
        vector_state="Failed"
    try:
        http = urllib3.PoolManager()
        data = json.dumps({"Info":{"ID":ID,"VectorState":vector_state,"Remark":msg}}).encode()
        config.nft_meta_ip="127.0.0.1"
        resp=http.request(
            method="POST",
            url=f"http://{config.nft_meta_ip}:{config.nft_meta_http_port}/v1/update/token",
            body=data
            ).data
        return json.loads(resp)["Info"]
    except Exception as e:
        logging.error(e)
        return 