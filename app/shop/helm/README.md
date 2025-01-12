

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
