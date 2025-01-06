# Istio Config

## Check VirtualService
```bash
kubectl get dr -A
kubectl get vs -A

SLEEP_POD=$(kubectl get pod -l app=sleep -o jsonpath={.items..metadata.name})

# Check the Endpoints
istioctl proxy-config endpoints $SLEEP_POD | grep 7777

# Test Traffic Routing 99% V2
TARGET_URL="hw-svc.istio-testapp.svc.cluster.local:7777"
while true; do kubectl exec "$SLEEP_POD" -c sleep -- curl -sS $TARGET_URL; done;
```

## Check Ingress GW
```bash
kubectl get gw -n shop
kubectl get vs -n shop

export INGRESS_NS=istio-gateways
export INGRESS_NAME=istio-ingressgateway
kubectl get svc "$INGRESS_NAME" -n "$INGRESS_NS"

NAME                   TYPE           CLUSTER-IP      EXTERNAL-IP    PORT(S)                                      AGE
istio-ingressgateway   LoadBalancer   10.103.68.167   172.18.0.152   15021:32691/TCP,80:32220/TCP,443:31012/TCP   5h44m

export INGRESS_HOST=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export INGRESS_PORT=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')

curl -s -I -HHost:gateway.cnk.com "http://$INGRESS_HOST:$INGRESS_PORT/frontend"

# add DNS for gateway
vi /etc/hosts
# 172.18.0.152 gateway.cnk.com

curl gateway.cnk.com:80/frontend
```

## Authzation Policy
```bash
# allow to call frontend endpoint directly
curl cnk.com:9990
curl cnk.com:9990/health
curl gateway.cnk.com:80/frontend

# all to call backend endpoint via frontend and gateway
curl cnk.com:9990/fetch-item?index=1
curl gateway.cnk.com:80/backend-inventory

# deny to call backend endpoint directly
curl backend.cnk.com:9991
curl backend.cnk.com:9991/health
curl backend.cnk.com:9991/items/1
```