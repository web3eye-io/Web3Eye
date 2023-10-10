# 测试集群信息

主机规划表

| IP           | hostname | 硬件配置                           | 角色                  | 位置 | 系统       |
|--------------|----------|--------------------------------|-----------------------|------|------------|
| 172.16.29.49 |          | CPU:4核  内存：4G  磁盘：50G         | gateway               | IDC  | ubuntu20.4 |
| 172.16.29.51 | idcnode1 | CPU:24核  内存：16G  磁盘：200G      | k8s-master            | IDC  | ubuntu20.4 |
| 172.16.29.52 | idcnode2 | CPU:24核  内存：32G  磁盘：100G      | k8s-worker            | IDC  | ubuntu20.4 |
| 172.16.29.53 | idcnode3 | CPU:24核  内存：32G  磁盘：100G      | k8s-worker            | IDC  | ubuntu20.4 |
| 172.16.29.54 | idcnode4 | CPU:24核  内存：16G  磁盘：100G+400G | k8s-worker            | IDC  | ubuntu20.4 |
| 172.23.10.87 | awsnode1 | CPU:1核  内存：16G  磁盘：100G       | k8s-master,k8s-worker | AWS  | ubuntu20.4 |

# 存储

测试数据：1110 个高度（18061000附近，目前总高度18318388）

MySql和Milvus都是3个节点的部署结构，以下对数据量最多的节点进行描述，其他两个节点数据观察下来数据量在主节点的70%（数据量大的场景未验证）。

MySql用量：
- 磁盘总占用：1.4G
- 现存表占用：113M
- bin-log：～1G
- logfile+ib*：160M

索引完所有数据不考虑bin-log，总占用量大概2TB

Milvus用量：
- etcd： 4.6G（其他两个节点未超过1G）
- minio： 309M
- pulsar-bookie-journal：1.6G
- pulsar-bookie-ledgers：2.5G

milvus的使用上需要注意向量维度大小，需要考虑是否需要缩小维度

# 接口速度

接口测试单位 并发×轮循

| API                  | 1*1      | 1*10       | 10*1     | 10*10    | 20*1     | 50*1      | 100*1   | 100*10 |
|----------------------|----------|------------|----------|----------|----------|-----------|---------|--------|
| SearchFile           | 3236 ms | 1860.2 ms | 7344 ms | 6991 ms | 4040 ms | 44893 ms | x       | x      |
| GetTransfers         | 283 ms  | 221.3 ms  | 432 ms  | 673 ms  | 650 ms  | 788 ms   | 709 ms | x      |
| GetContractAndTokens | 341 ms  | 415.2 ms  | 508 ms  | 576 ms  | 549 ms  | 511 ms   | 663 ms | x      |


# 解析速度

Block-ETL部署情况为2 Job * 5 Goroutine

解析每一个块的Transfer和Order大于耗时**1.1s**

Transform部署情况为6 Job

解析速度取决于网络带宽，大约**0.5～0.7张/s**
