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

docker image tag ${IMAGE_NAME} ${IMAGE_NAME}:1.0.1
docker image push ${IMAGE_NAME}:1.0.1
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
curl frontend.testapp.com:9990
curl frontend.testapp.com:9990/fetch-item?index=1
curl backend-inventory.testapp.com:9991/items/2
```