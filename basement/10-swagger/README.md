# swagger-ui

## 拷贝json文件到swagger-ui工作目录
* 获取swagger-ui-pod

  `kubectl get pod -A | grep swagger | awk '{print $2}'`

* 拷贝json文件到swagger-ui工作目录

  `kubectl cp xxx.json kube-system/$swagger-ui-pod:/usr/share/nginx/html/data`

## swagger-ui访问页面（添加本地域名解析）
http://swagger-ui.internal-devops.development.npool.top

`页面查看：data/xxx.json`
