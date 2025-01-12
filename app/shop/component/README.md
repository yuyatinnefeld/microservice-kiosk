## SHOP APPS

### Quick Debug
```bash
#---------------------
# Inventory Service
#---------------------

cd app/shop/component/inventory
export IMAGE_NAME="yuyatinnefeld/cnk-backend-inventory"
docker build -t ${IMAGE_NAME} .
docker run -it -p 9991:9991 -e ENV=DOCKER-DEV --net=host ${IMAGE_NAME}

# verify
curl http://localhost:9991
curl http://localhost:9991/health
curl http://localhost:9991/items/2

# push image
docker image tag ${IMAGE_NAME} ${IMAGE_NAME}:1.1.0
docker image push ${IMAGE_NAME}:1.1.0

#---------------------
# frontend Service
#---------------------

cd app/shop/component/frontend
export IMAGE_NAME="yuyatinnefeld/cnk-frontend"
docker build -t ${IMAGE_NAME} .
docker run -it -p 9990:9990 -e ENV=DOCKER-DEV --net=host ${IMAGE_NAME}

# verify
curl http://localhost:9990
curl http://localhost:9990/health
curl http://localhost:9990/fetch-item?index=1

# push image
docker image tag ${IMAGE_NAME} ${IMAGE_NAME}:1.1.0
docker image push ${IMAGE_NAME}:1.1.0

#---------------------
# ML Service
#---------------------

cd app/shop/component/ml
export IMAGE_NAME="yuyatinnefeld/cnk-ml"
docker build -t ${IMAGE_NAME} .
docker run -it -p 9992:9992 -e ENV=DOCKER-DEV --net=host ${IMAGE_NAME}

docker image tag ${IMAGE_NAME} ${IMAGE_NAME}:1.1.0
docker image push ${IMAGE_NAME}:1.1.0
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
# 172.18.0.150 gateway.cnk.com

curl gateway.cnk.com:80
```

### Authzation Policy
```bash
# allow access frontend endpoint directly
curl cnk.com:9990
curl cnk.com:9990/health
curl gateway.cnk.com:80
curl gateway.cnk.com:80/health

# allow access backend endpoint via frontend and gateway
curl cnk.com:9990/fetch-item?index=1
curl gateway.cnk.com:80/backend-inventory
curl gateway.cnk.com:80/fetch-item?index=1
curl gateway.cnk.com:80/fetch-item?index=2

# access denied
curl backend.cnk.com:9991
curl backend.cnk.com:9991/health
curl backend.cnk.com:9991/items/1
```