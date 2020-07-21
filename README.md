# csi-demo-plugin

1. 当前demo目录下，按需修改 Dockerfile，之后build出镜像
docker build -t xxxxxxx:5000/online/test/csi-demo-driver:v1.0 .
2. cd deploy/kubernetes，按需修改namespace，csi-demo-driver，rbac文件，注意namespace必须保持一致，然后应用上
```
kubectl apply -f xxx.yaml
kubectl get pods --namespace=csi-dev-demo  #注意修改的namespace
```
output like:
<pre>
I0110 03:15:13.125126       1 driver.go:27] Driver: kvm.csi.dianduidian.com version: 1.0.0
I0110 03:15:13.125350       1 mount_linux.go:160] Detected OS without systemd
I0110 03:15:13.125388       1 driver.go:54] protocol: unix,addr: /csi/csi.sock
I0110 03:15:13.259096       1 driver.go:75] GRPC call: /csi.v1.Identity/GetPluginInfo
I0110 03:15:13.259119       1 driver.go:76] GRPC request: {}
I0110 03:15:13.260501       1 indentityserver.go:18] Using default GetPluginInfo
I0110 03:15:13.260507       1 driver.go:81] GRPC response: {"name":"kvm.csi.dianduidian.com","vendor_version":"1.0.0"}
I0110 03:15:13.401214       1 driver.go:75] GRPC call: /csi.v1.Identity/GetPluginInfo
I0110 03:15:13.401232       1 driver.go:76] GRPC request: {}
I0110 03:15:13.401744       1 indentityserver.go:18] Using default GetPluginInfo
</pre>

3. 按需修改 deploy/example/xxx.yaml, kubectl apply,注意里面的namespace，driver名必须保持一致

