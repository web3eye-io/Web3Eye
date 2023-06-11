import urllib3
import json
from pkg.utils import config

def report_file_to_gen_car(id:str,file_s3_key)-> bool:
    try:
        http = urllib3.PoolManager()
        data = json.dumps({'ID':id,"FileName":file_s3_key}).encode()
        print(config.gen_car_ip)
        http.request(
            method="POST",
            url=f"http://{config.gen_car_ip}:{config.gen_car_http_port}/v1/report/file",
            body=data
            )
        return True
    except:
        return False    