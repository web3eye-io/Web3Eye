# cyber tracer

### 功能计划
- [x] 拆分成web3eye独立的message
- [ ] 支持k8s
- [ ] 支持多副本
- [x] 合并成一个git项目
- [ ] 服务发现和配置管理
  
#### CICD
- [ ] Docker Compose
- [ ] K8s
- [ ] Jenkins
Golang
- [x] Verify
    - [x] lint
    - [x] spell
    - [x] build
- [ ] Test
- [x] Dockerfile

Python
- [ ] Verify
    - [ ] lint
    - [ ] spell
    - [ ] build
- [ ] Test
- [x] Dockerfile


#### NFT-META
存储nft信息
- [x] 创建NFT信息
- [x] 以图搜索nft信息
- [x] 与milvus交互（创建collection及向量创建、删除、查询、搜索等）
- [ ] milvus交互优化，使用官方的代码生成器生成交互代码
- [x] 记录扫描过的高度及状态
- [ ] 解析错误的过程记录
- [ ] 管理kafka的topic(创建、删除)

#### BLOCK-ETL
从链上索引nft相关信息
- [x] 自动索引NFT信息(转账信息)
- [x] 索引合约创建信息
- [ ] 索引NFT-mint信息
- [x] 补齐未进行交易的NFT信息(主要补齐资源地址)
- [ ] 支持以chain-contract-tokenid生成一个identify帮助快速查询，主要方便关系表的使用
- [x] 热查询（经常会用到的固定信息）缓存进redis
    - [x] tokentransfer中用到的identifier
    - [x] 是否已存在contract
    - [ ] 区块高度对应时间戳

#### NFT-MEDIA
解析nft资产带的媒体资源
- [ ] 自动解析资源地址

#### IMAGE-CONVERT
- [x] 图片向量化
- [x] http图片向量化
- [x] ipfs图片向量化（采用http方式解析）
- [ ] ipfs图片向量化
- [ ] 链上图片数据向量化
- [ ] 架构更换或打包优化（现在的docker镜像太大）

### CyberTracer 目录结构
服务一（golang）:
- [ ] 代码
- [ ] Dockerfile、K8s_yaml
- [ ] Script
  - [ ] Build
  - [ ] Test
  
服务二（golang）  
服务三（python）  
服务四（typescript）  

公共代码: 
- [ ] network
- [ ] math
- [ ] ...

hack:
- [ ] Verify (Lint、Spell...)
- [ ] ProjectManager(ToolsInstall\ Message\abi-gen\Deps ...)

Makefile:  
- [ ] Verify
- [ ] Test
- [ ] Build
- [ ] Run
- [ ] ProjectManager

DockerCompose    
Jenkinsfile  

## 第三方组件搭建方案
### mysql
### milvus
### redis