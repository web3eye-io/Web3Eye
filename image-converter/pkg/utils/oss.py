from boto3.session import Session
import boto3
#Client初始化
access_key = "2IVIqQk0n42MbzMCuigp"
secret_key = "b9lXn6Xg7WkBrXS5BJqgnSsJe7RAFFXkPI1ylsRK"
url = "http://web3eye-minio:9000" 
session = Session(access_key, secret_key)
s3_client = session.client('s3', endpoint_url=url)
bucket="fil-alaws-on"
key="bbxxx"
file_path="/root/code/Web3Eye/README.md"
resp = s3_client.put_object(Bucket=bucket, Key=key, Body=open(file_path, 'rb').read())