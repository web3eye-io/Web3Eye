# cyber tracer
目前在NFT的世界中很多关于区块链的信息索取方法是复杂且难以上手，让用户很难获取信息；再者目前的各种区块链项目数据是割裂的，获取或整理信息就变得更难了。
CyberTracer是一个聚合历史NFT交易记录的搜素引擎；提供NFT资产的多链聚合搜索。

## quick start

## 架构
中间件：  
**MySql** 存储任务信息、NFT数据关系  
**Kafka** 主要用于任务分配  
**Redis** 用于缓存从MySql查到的热信息(计算机中的局部原理，减轻MySql的压力)  
**Milvus** 用于存储向量数据，以及提供向量搜索  

微服务模块：  
**NFT-Meta** 维护区块转储任务,存储NFT交易、NFT资产、NFT对应Contract等信息  
**Block-ETL** 负责与区块链节点交互，获取NFT的transfer日志，分析对应的Token信息以及Cantract信息  
**Image-Converter** 将图片转换为向量  

![架构](doc/picture/archi.jpg)

在主要的三个微服务模块中，NFT-Meta负责提供搜索、信息存储查询、任务分发等功能，其他两个模块更多的是获取并处理任务；其中Image-Converter不只是处理从Kafka获取由NFT-Meta发送的任务，还提供HTTP服务支持直接请求获得向量，主要用来为以图搜图服务；而Block-ETL不对外提供接口，只接收任务和提交任务。

### 模块设计：
#### Image-Converter:  
目前主要提供Jpg、Jepg、Png等常规图片格式的转向量操作，其他图像资源比如GIF、Base64等目前并不支持。  
服务启动后有两个线程，一个负责提供HTTP接口方式的转向量方式，提供同步的转向量方式，支持URL和文件两种方式；还有一个负责从Kafka获取转向量任务，转换后放入Kafka中NFT-Meta获取后存入Milvus和数据库。  

#### Block-ETL:  
目前仅支持Ethereum，所以以下描述都基于ETH背景；并且目前也只支持标准的ERC721和ERC1155至于其他玩法的NFT后续提供支持。

从区块链全节点（存有全部区块数据，下面称为钱包节点）中的log中**拉取transfer信息，解析出NFT交易、Token、Contract信息**；由于存在Swap合约、非标准NFT合约所以部分Token信息无法解析出资产信息（比如图片描述和图片地址）。

从NFT-Meta获取的任务粒度为区块高度，从一个区块高度中获取所有的transfer-log，记录每一笔transfer；再从transfer信息查找Token信息，因为多笔transfer可能会对应到同一个Token，所以会先向数据库查询Token是否存在，不存在会向钱包节点请求TokenURI同时也会检查对应的Contract是否存在。在此处查询Token和Contract是否存在时，其实会先检查Redis里是否有记录，没有记录再去数据库查询，当查询到后会在Redis中建立一条记录。

![数据处理关系](doc/picture/transfer-token-contract.jpg)

在解析一个区块的transfer日志时，大部分信息都可以从钱包节点获取。但是从TokenURI中所带的信息需要从互联网或者IPFS上获取，或者直接在区块链上存储Base64、SVG等，解析TokenURI的工作目前还属于这个模块，后续考虑独立成一个单独的模块。因为这样的解析工作费时费力，同时尽量保持Block-ETL只与钱包节点交互、只做链上数据的转存工作。

#### NFT-Meta:
分配到其他两个模块的任务都由NFT-Meta发出、存储其他两个模块处理过的数据并对外提供搜索功能。目前这个模块可能有些臃肿，比如关于搜索的功能可以独立出去、其他模块产生的数据直接与数据库交互这类问题后续会继续思考，但是前期搜索由于功能较少先放到这个模块中。

NFT-Meta主要维护四个表：  
1 **Transfers**  NFT交易记录  
2 **Tokens**  NFT-资产信息  
3 **Contracts**  NFT-合约信息  
4 **SyncTasks**  同步任务

NFT-Meta提供GRPC和HTTP两种协议的API接口，GRPC主要提供给对内的微服务模块，HTTP对外提供服务；向量数据主要存在Milvus中，关系型数据主要存在MySql中。Milvus与MySql中的数据依靠Milvus提供的ID关联。  
Milvus中结构为：

{  
    ID: 13125  
    Vector: [0.234,2.923,...]  
}

MySql中会关联Milvus中的ID字段  
如：  

{  
...  
    ID: 29aa144d-beb0-4d25-b7bb-95587fe06ba4  
    VectorID: 13125  
    VectorState: Success  
...  
}  

### 主要流程

如何做到搜索NFT资产？已有的搜索方式都是通过合约和TokenID，CyberTracer采用的是相似度搜索。NFT资产的形式多种多样，多是以非结构化数据展示，比如图像、音频、视频等转换成向量支持搜索；目前仅支持以图搜图，后续跟进其他形式的NFT资产。

有了非结构化数据的搜索，接下来就是聚合NFT数据，从钱包节点上获取transfer日志（NFT交易的日志）分析出Token和Contract信息。与NFT相关的三个数据中，transfer和Contract信息可以较简单的获取到，而Token的解析稍微复杂些。

#### 以图搜图

以图搜图就是特征向量求距离的过程，从一堆已有的特征向量中，求出与给定特征向量的距离，取出最小距离排名的前N个即可。

已有的特征向量的来源是NFT的图片数据转换而来，放到向量数据库中的。

用于搜索的图，也需要转成向量的过程，以便于与向量数据库中的数据做距离运算。

![图转向量（示例图片不代表真实的转向量结果，仅作为概念展示）](doc/picture/image-to-vector.jpg)

在CyberTracer中搜索一张图，大致经历四个阶段。

![以图搜图步骤](doc/picture/pictrue-search.jpg)

1. 用户带文件请求（图片文件）发送到NFT-Meta  
2. NFT-Meta将请求转发给Image-Converter转换成向量  
3. 拿到向量的NFT-Meta去Milvus中查找相似的向量，并且返回向量ID  
4. NFT-Meta拿到相似向量的ID，去MySql中将ID对应的Token信息查询出来  
最后返回到用户即可  

#### 获取Token信息

Token的主要字段如下：  
{  
"Contract"  
"TokenID"  
"TokenType"  
"URI"  
"ImageURL"  
"VectorID"  
}  

其中Contract、TokenID、URI都是可以直接从钱包节点获取的，而TokenType、ImageURL可以通过分析URI得到。而VectorID则需要在其他信息都获得之后，在NFT-Meta中插入一条Token记录时，将转向量的任务放到队列中，等待Image-Converter消费，转换完成后放入结果队列等待NFT-Meta更新VectorID字段。

#### 任务分发
目前用到Kafka的地方有两个，一是给Image-Converter放转向量任务，二是给Block-ETL放需要解析的区块高度。

图中就是需要大量转向量对时间要求不高时的转向量任务处理过程，因为转向量时消耗网络带宽和计算资源所以采用异步的方式提高稳定性。

![IC转换向量任务分配](doc/picture/to-vector-task.jpg)

但是在搜索时就直接请求Image-Converter提供的HTTP转向量方式，提高响应速度。

Block-ETL主要负责分析每一个区块高度出来的数据，并放入NFT-Meta中。

Blcok-ETL获取的任务（待同步的区块高度号），有两个过程：

1.管理员通过请求NFT-Meta建立同步任务，任务包含开始区块、结束区块、当前区块  
2.Blocke-ETL会定期检查(向NFT-Meta请求)是否有需要同步的Topic，有则监听并消费  

第一个过程很简单就是往数据库加入一条任务记录。第二个要求Block-ETL主动查询待同步的任务，主动触发NFT-Meta往Kafka中放数据，虽然是NFT-Meta放数据，但是消费的主动权交到Block-ETL，同时让NFT-Meta更加无状态化。

![IC转换向量任务分配](doc/picture/block-etl-task.jpg)

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

# 参考
