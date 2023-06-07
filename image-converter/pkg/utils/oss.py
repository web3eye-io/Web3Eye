from boto3.session import Session
import config
from singleton import singleton
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
        except:
            return False

# use demo
# OSS().put_object(file_path="/root/code/Web3Eye/image-converter/pkg/utils/config.toml",key="config.toml",bucket="fil-alaws-on")
