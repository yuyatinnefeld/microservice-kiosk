# Istio Debug Test Application

## HW APP
```bash
kubectl port-forward svc/hw-svc 7777:7777
```

## HTTPBIN
```bash
kubectl port-forward svc/hw-svc 7777:7777
```

## FORTIO
Define the namespace and get the Fortio pod name:
```bash
NS=istio-testapp
FORTIO_POD="$(kubectl get pod -n $NS -l app=fortio -o jsonpath='{.items[0].metadata.name}')"
TARGET_URL=hw-svc.$NS.svc.cluster.local:7777
kubectl exec -it $FORTIO_POD --container fortio -- /usr/bin/fortio load -k -c 1 -qps 10 -t 0.5m -loglevel Warning $TARGET_URL
```
