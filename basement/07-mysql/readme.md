# mysql

read password

```shell
kubectl get secret -n kube-system mysql -o jsonpath="{.data.mysql-root-password}" | base64 -d
```
