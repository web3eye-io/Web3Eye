from boto3.session import Session
from pkg.utils import config
from pkg.utils.singleton import singleton
import logging

#Client初始化


@singleton
class OSS(object):
    s3_client=None

    def __init__(self) -> None:
        session = Session(config.minio_access_key, config.minio_secret_key)
        self.s3_client = session.client('s3', endpoint_url=config.minio_address)
    
    def put_object(self,file_path:str,key:str,bucket:str)->bool:
        try:
            self.s3_client.put_object(
                Bucket=bucket, 
                Key=key, 
                Body=open(file_path, 'rb').read())
            
            return True
        except Exception as e:
            logging.error(e)
            return False
    
    def put_object_retries(self,file_path:str,key:str,bucket:str,retries:int)->bool:
        for i in range(retries):
            print(file_path, key,bucket)
            ok=self.put_object(file_path=file_path,key=key,bucket=bucket)
            if ok :
                return True
        return False

# use demo
# OSS().put_object(file_path="/root/code/Web3Eye/image-converter/pkg/utils/config.toml",key="config.toml",bucket="fil-alaws-on")
