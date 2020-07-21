FROM centos:latest
LABEL maintainers="xuezhiyou"
LABEL description="Demo CSI Driver"

RUN yum install -y util-linux e2fsprogs && yum clean all
COPY bin/csi-demo-driver /csi-demo-driver
ENTRYPOINT ["/csi-demo-driver"]