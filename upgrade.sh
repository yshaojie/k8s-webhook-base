tag=$(date '+%m%d%H%M%S')
make docker-build IMG=base-webhook:${tag}
kind load docker-image base-webhook:${tag}
sleep 3
make deploy IMG=base-webhook:${tag}