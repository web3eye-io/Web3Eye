from boto3.session import Session
import boto3
from singleton import singleton
#Client初始化
access_key = "2IVIqQk0n42MbzMCuigp"
secret_key = "b9lXn6Xg7WkBrXS5BJqgnSsJe7RAFFXkPI1ylsRK"
url = "http://web3eye-minio:9000" 
bucket="fil-alaws-on"
session = Session(access_key, secret_key)
s3_client = session.client('s3', endpoint_url=url)

key="bbxxx"
file_path="/root/code/Web3Eye/README.md"



@singleton
class OSS(object):
    s3_client
    def __init__(self) -> None:
        session = Session(access_key, secret_key)
        self.s3_client = session.client('s3', endpoint_url=url)
    
    def put_object(self,file_path:str)->bool:
        try:
            self.s3_client.put_object(
                Bucket=bucket, 
                Key=key, 
                Body=open(file_path, 'rb').read())
            return True
        except:
            return False
