# Istio Debug Test Application

## HW APP
```bash
kubectl port-forward svc/hw-svc 7777:7777
```

## HTTPBIN
```bash
kubectl port-forward svc/httpbin 7775:7775
curl -v localhost:7775/status/200
```

## FORTIO
Define the namespace and get the Fortio pod name:
```bash
TARGET_URL=hw-svc.istio-testapp.svc.cluster.local:7777
FORTIO_POD="$(kubectl get pod -l app=fortio -o jsonpath='{.items[0].metadata.name}')"
kubectl exec -it $FORTIO_POD --container fortio -- /usr/bin/fortio load -k -c 1 -qps 10 -t 0.5m -loglevel Warning $TARGET_URL
```

## SLEEP
```bash
SOURCE_POD=$(kubectl get pod -l app=sleep -o jsonpath={.items..metadata.name})
TARGET_URL=hw-svc.istio-testapp.svc.cluster.local:7777
kubectl exec "$SOURCE_POD" -c sleep -- curl -sSI $TARGET_URL | grep  "HTTP/"

for i in {1..20}; do kubectl exec "$SOURCE_POD" -c sleep -- curl -sSI $TARGET_URL | grep  "HTTP/"; done

# v1 1%, v2 99%
while true; do kubectl exec "$SOURCE_POD" -c sleep -- curl $TARGET_URL 2>/dev/null | grep -E "APP: .* "; done
```

## SHOP APPS

### Quick Debug
```bash
# Inventory Service
cd app/shop/component/inventory
export IMAGE_NAME="yuyatinnefeld/cnk-backend-inventory"
docker build -t ${IMAGE_NAME} .
docker run -it -p 9991:9991 -e ENV=DOCKER-DEV --net=host ${IMAGE_NAME}

docker image push $IMAGE_NAME

# verify
curl http://localhost:9991
curl http://localhost:9991/health
curl http://localhost:9991/items/2

docker image tag ${IMAGE_NAME} ${IMAGE_NAME}:1.0.1
docker image push ${IMAGE_NAME}:1.0.1

# frontend Service
cd app/shop/component/frontend
export IMAGE_NAME="yuyatinnefeld/cnk-frontend"
docker build -t ${IMAGE_NAME} .
docker run -it -p 9990:9990 -e ENV=DOCKER-DEV --net=host ${IMAGE_NAME}

docker image tag ${IMAGE_NAME} ${IMAGE_NAME}:1.0.2
docker image push ${IMAGE_NAME}:1.0.2
# verify
curl http://localhost:9990
curl http://localhost:9990/health
curl http://localhost:9990/fetch-item?index=1
```

### Image Build Pipeline
- Github Action

### Deployment
- ArgoCD (incl. MetalLB)

```bash
curl cnk.com:9990
curl cnk.com:9990/fetch-item?index=1 # Allow the access from frontend with authz policy
curl backend.cnk.com:9991/items/1 # Deny the direct access with authz policy
curl backend.cnk.com:9991/health # Deny the direct access with authz policy
```

## Istio Config

### Check VirtualService
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

### Check Ingress GW
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

### Authzation Policy
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