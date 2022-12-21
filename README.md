# cyber tracer
目前在NFT的世界中很多关于区块链的信息索取方法是复杂且难以上手，让用户很难获取信息；再者目前的各种区块链项目数据是割裂的，获取或整理信息就变得更难了。
CyberTracer是一个聚合历史NFT交易记录的搜素引擎；提供NFT资产的多链聚合搜索。

## 架构
中间件：
MySql 存储任务信息、NFT数据关系
Kafka 主要用于任务分配
Redis 用于缓存从MySql查到的热信息(计算机中的局部原理，减轻MySql的压力)
Milvus 用于存储向量数据，以及提供向量搜索

微服务模块：
NFT-Meta 维护区块转储任务,存储NFT交易、NFT资产、NFT对应Contract等信息
Block-ETL 负责与区块链节点交互，获取NFT的transfer日志，分析对应的Token信息以及Cantract信息
Image-Converter 将图片转换为向量

![架构](doc/picture/archi.jpg)

在主要的三个微服务模块中，NFT-Meta负责提供搜索、信息存储查询、任务分发等功能，其他两个模块更多的是获取并处理任务；其中Image-Converter不只是处理从Kafka获取由NFT-Meta发送的任务，还提供HTTP服务支持直接请求获得向量，主要用来为以图搜图服务；而Block-ETL不对外提供接口，只接收任务和提交任务。

### 模块设计：
Image-Converter:
目前主要提供Jpg、Jepg、Png等常规图片格式的转向量操作，其他图像资源比如GIF、Base64等目前并不支持。
服务启动后有两个线程，一个负责提供HTTP接口方式的转向量方式，提供同步的转向量方式，支持URL和文件两种方式；还有一个负责从Kafka获取转向量任务，转换后放入Kafka中NFT-Meta获取后存入Milvus和数据库。

Block-ETL:
目前仅支持Ethereum，所以以下描述都基于ETH背景；并且目前也只支持标准的ERC721和ERC1155至于其他玩法的NFT后续提供支持。
从区块链全节点（存有全部区块数据，下面称为钱包节点）中的log中拉取transfer信息，解析出NFT交易、Token、Contract信息；由于存在Swap合约、非标准NFT合约所以部分Token信息无法解析出资产信息（比如图片描述和图片地址）。
从NFT-Meta获取的任务粒度为区块高度，从一个区块高度中获取所有的transfer-log，记录每一笔transfer；再从transfer信息查找Token信息，因为多笔transfer可能会对应到同一个Token，所以会先向数据库查询Token是否存在，不存在会向钱包节点请求TokenURI同时也会检查对应的Contract是否存在。在此处查询Token和Contract是否存在时，其实会先检查Redis里是否有记录，没有记录再去数据库查询，当查询到后会在Redis中建立一条记录。

![数据处理关系](doc/picture/transfer-token-contract.jpg)

在解析一个区块的transfer日志时，大部分信息都可以从钱包节点获取。但是从TokenURI中所带的信息需要从互联网或者IPFS上获取，或者直接在区块链上存储Base64、SVG等，解析TokenURI的工作目前还属于这个模块，后续考虑独立成一个单独的模块。因为这样的解析工作费时费力，同时尽量保持Block-ETL只与钱包节点交互、只做链上数据的转存工作。

NFT-Meta:

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


# 版本迭代计划

[0.1.0](doc/feature/0.1.0.md)

[100.0.0](doc/feature/100.0.0.md)