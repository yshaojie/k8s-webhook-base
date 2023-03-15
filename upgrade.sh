tag=$(date '+%m%d%H%M%S')
image_name=base-webhook
#清除以前的镜像
crictl images | grep docker.io/library/${image_name} | awk '{print $3}' |  xargs crictl rmi
make docker-build IMG=${image_name}:${tag}
kind load docker-image ${image_name}:${tag}
sleep 3
make deploy IMG=${image_name}:${tag}