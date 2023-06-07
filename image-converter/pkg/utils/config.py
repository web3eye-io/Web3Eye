import rtoml
import pathlib
import os
from singleton import singleton

    
@singleton
class Config(object):
    config=dict()

    def _analyze_config(self,data,preffix:str):
        if isinstance(data, dict):
            for k, v in data.items():
                self._analyze_config(v,preffix+"_"+k)
        else:
            preffix=str.lstrip(preffix,"_")
            preffix=str.strip(preffix," ")
            preffix=str.replace(preffix,"-","_")
            self.config[preffix]=data

    def __init__(self):
        path=pathlib.Path("config.toml")
        configStr=rtoml.load(toml=path)
        self._analyze_config(configStr,"")

        env_dist = os.environ
        for i in self.config:
            try:
                env=env_dist[i]
                self.config[str.strip(i," ")]=str.strip(env," ")
            except:
                continue

minio_address=Config().config["minio_address"]
minio_access_key=Config().config["minio_access_key"]
minio_secret_key=Config().config["minio_secret_key"]
minio_token_image_bucket=Config().config["minio_token_image_bucket"]

# use demo
# Config().config["mysql_ip"]