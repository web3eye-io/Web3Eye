# 使用grpc-stream做grpc代理

将微服务集群分成两个部分，IDC和AWS两个部分。当AWS想使用IDC时，不想让IDC有其他人连接不暴露接口，让IDC主动与AWS建立长连接，AWS就可以向IDC发送请求。

类似于WebSocket的使用场景，只是这里基于grpc-stream实现。

## 对比

相比直接使用client调用server，中间加了cloud-proxy和gateway，效率会降低但是带来的是AWS和IDC中服务之间的无缝衔接。

![grpc-stream-proxy](https://github.com/web3eye-io/Web3Eye/tree/master/doc/picture/grpc-stream-proxy.png)

在client调用server时和原来的方式基本一致，只需要在调用server接口之前指定使用Proxy提供的连接方法即可。

指定使用Proxy提供的连接方法后，会将Server地址、grpc方法、请求体发送到proxy，proxy再使用流发送到gateway，由gateway调用IDC的服务。

两种方式调用代码差别不大：

```Golang
// 直接调用
contract.GetContracts(context.TODO(), nil, 0, 10)

// 指定使用proxy调用
contract.UseCloudProxyCC()
contract.GetContracts(context.TODO(), nil, 0, 10)
```

要让client中间使用proxy和gateway做传话人，最重要的就是让client想要发出的请求放到gateway发出，然后再原路返回到client。

grpc的底层其实是HTTP/2协议，所以发送请求就是将方法和数据封装成HTTP/2数据包的过程。最后发出去的HTTP/2数据包中，请求体是以字节数组的方式存放的，所以大概的思路就是让proxy和gateway来回都直接传输请求体的字节数组，不编码也不解码。gateway调用server时直接用编码后的请求体，收到response也直接将字节数组方式的响应返回到client。

主要的工作：

- 重写Client用到的Invoke，当Client调用接口时可以获取Server地址、方法名、编码好的请求体，将以上信息转发到proxy，proxy处理好后再返回，解码好后返回结果。
- proxy到gateway之间的流管理，建立grpc-stream池，当AWS中有请求过来时选择其中一条stream发送请求。
- gateway使用自定义的codec，保证发送请求和接收请求都不做解析，将解析交给client调用的接口。这一步是做grpc代理的关键，grpc库自己就支持自定义的codec让编解码和发送请求分离。
