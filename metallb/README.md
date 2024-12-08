## Enable External Cluster IP with Metallb

### Inspect the CIDR block
```bash
kubectl get nodes -o json | jq .items[].status.addresses[].address # "172.18.0.2" "my-cluster-control-plane"
docker inspect my-cluster-control-plane | jq -r ".[0].NetworkSettings.Networks.kind.IPAddress"
```

### Install Metallab
```bash
helm repo add metallb https://metallb.github.io/metallb
helm install metallb metallb/metallb
kubectl get pod -w # wait until metallb-speaker and metallb-controller are installed
kubectl apply -f metallb/metallb-config-istio-testapp.yaml
```

## Test External IP and Configure Local DNS
```bash
kubectl get svc # External IP = 172.18.0.100
curl 172.18.0.100:7777

vi /etc/hosts # add testapp.com
#172.18.0.100 shop.testapp.com
#172.18.0.200 shop-south.testapp.com

curl shop.testapp.com:7777
```