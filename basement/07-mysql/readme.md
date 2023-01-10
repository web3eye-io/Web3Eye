# mysql

read password

```shell
kubectl get secret --namespace "default" mysql -o jsonpath="{.data.mysql-root-password}" | base64 -d
```
