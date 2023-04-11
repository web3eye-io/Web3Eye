# 问题

## 描述

多server对于多client长连接均匀分配问题

在aws和idc之间需要使用流连接，来保证aws上的服务可以使用idc的数据。通过idc中的服务主动连接aws上的服务，可以减少aws的运营负担，也可以保证idc的安全。

但是使用长连接无法做到负载均衡，很容易导致个别pod负载过高，比如server的两个pod启动时间有差时，快的一个就会接受较多连接，而且长连接会一直保持。

## 解决办法

每次启动服务（server和client）都生成uuid，server将自己的uuid上报（可能是redis中，这个上报带有过期时间，在过期时间内不再次刷新，会被过期删除），这样其他server也能知道总pod数量。client发起连接时就能知道server的个数，可重复建立多个连接，如果UUID相同则中断，重试多次让client和server尽量均匀的保持连接。
