## Enable External Cluster IP with Metallb

### Inspect the CIDR block
```bash
kubectl get nodes -o json | jq .items[].status.addresses[].address # "172.18.0.2" "my-cluster-control-plane"
docker inspect my-cluster-control-plane | jq -r ".[0].NetworkSettings.Networks.kind.IPAddress"
```

### Deploy Metallab speaker and controller
- ArgoCD

## Test External IP and Configure Local DNS
```bash
kubectl get svc -n shop
# External IP = 172.18.0.150
curl 172.18.0.151:9990


# Setup DNS
vi /etc/hosts 
# 172.18.0.150 gateway.cnk.com
# 172.18.0.151 cnk.com
# 172.18.0.152 backend.cnk.com
# 172.18.0.153 testapp.cnk.com

curl cnk.com:9990
curl cnk.com:9990/health
curl cnk.com:9990/fetch-item?index=1

# RBAC: access denied -> Istio AuthZ Policy
curl backend.cnk.com:9991
curl backend.cnk.com:9991/health
curl backend.cnk.com:9991/items/1

curl testapp.cnk.com:7777 # hw1 and hw2
```