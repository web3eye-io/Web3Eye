# mysql

read accesskey and secretkey

```shell
kubectl get secret -n kube-system redis-cluster -o jsonpath="{.data.redis-password}" | base64 -d && echo ""
```