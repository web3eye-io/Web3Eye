# cyber tracer

[0.1.0](feature/0.1.0.md)
<a href="feature/0.1.0.md">0.1.0</a>

[100.0.0](feature/100.0.0.md)
<a href="feature/100.0.0.md">100.0.0</a>

## 配置
所有配置都在config/config.toml中，如果想修改有两种途径：
1.修改config/config.toml重新编译打包成Docker镜像
2.通过设置环境变量即可，在k8s中可设置configMap

config.toml -> environment 转换规则
例：
```toml
path="/uu/ii"
port=50515
project-name="cyber-tracer"

[mysql]
host="mysql"
port=3306
max-connect=100

log-dir="/var/log"
```

```shell
path=/uu/ii
port=50515
project_name=cyber-tracer

mysql_host=mysql
mysql_port=3306
mysql_max_connect=100

log_dir=/var/log
```