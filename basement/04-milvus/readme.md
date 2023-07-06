# mysql

read accesskey and secretkey

```shell
kubectl get secret -n kube-system milvus-minio -o jsonpath="{.data.accesskey}" | base64 -d && echo ""
kubectl get secret -n kube-system milvus-minio -o jsonpath="{.data.secretkey}" | base64 -d && echo ""
```
