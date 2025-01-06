## Enable External Cluster IP with Metallb

### Inspect the CIDR block
```bash
kubectl get nodes -o json | jq .items[].status.addresses[].address # "172.18.0.2" "my-cluster-control-plane"
docker inspect my-cluster-control-plane | jq -r ".[0].NetworkSettings.Networks.kind.IPAddress"
```

### Deploy Metallab speaker and controller
- ArgoCD

```bash
kubectl get pod -n shop
kubectl apply -f metallb/metallb-config-shop.yaml
```

## Test External IP and Configure Local DNS
```bash
kubectl get svc -n shop
# External IP = 172.18.0.150
curl 172.18.0.150:9991

# Setup DNS
vi /etc/hosts 
# 172.18.0.150 cnk.com
# 172.18.0.151 backend.cnk.com
# 172.18.0.152 gateway.cnk.com
# 172.18.0.153 testapp.cnk.com

curl cnk.com:9990
curl backend.cnk.com:9991
curl backend.cnk.com:9991/items/1

curl testapp.cnk.com:7777 # hw1 and hw2
curl 172.18.0.153:7777 # hw1 and hw2
```