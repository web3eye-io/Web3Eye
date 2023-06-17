# mysql

read accesskey and secretkey

```shell
kubectl get secret --namespace "default" redis-cluster -o jsonpath="{.data.redis-password}" | base64 -d && echo ""
```