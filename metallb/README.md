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
kubectl get svc # External IP = 172.18.0.150
curl 172.18.0.150:9991

vi /etc/hosts # add testapp.com
#172.18.0.150 backend-inventory.testapp.com
#172.18.0.151 frontend.testapp.com

curl backend-inventory.testapp.com:9991
curl frontend.testapp.com:9990
```