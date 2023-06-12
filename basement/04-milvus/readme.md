# mysql

read accesskey and secretkey

```shell
kubectl get secret --namespace "default" milvus-minio -o jsonpath="{.data.accesskey}" | base64 -d && echo ""
kubectl get secret --namespace "default" milvus-minio -o jsonpath="{.data.secretkey}" | base64 -d && echo ""
```
