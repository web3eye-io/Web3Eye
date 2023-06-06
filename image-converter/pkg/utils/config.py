import rtoml
import pathlib
import os

def singleton(cls):
    _instance = {}

    def inner():
        if cls not in _instance:
            _instance[cls] = cls()
        return _instance[cls]
    return inner
    
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
                print(i)
                self.config[i]=env
            except:
                continue

# use demo
# Config().config["mysql_ip"]