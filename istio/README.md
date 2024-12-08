# Istion

## Check
```bash
# Verify
istioctl verify-install

## Verify Sidecar Injection
TARGET_POD="$(kubectl get pod -l app=httpbin -n istio-testapp -o jsonpath='{.items[0].metadata.name}')"
kubectl logs $TARGET_POD -c istio-proxy
ISTIOD_POD="$(kubectl get pod -n istio-system -l app=istiod -o jsonpath='{.items[0].metadata.name}')"
kubectl logs $ISTIOD_POD -n istio-system
```

## Deploy Gateway
```bash
kubectl apply -f dr.yaml
kubectl get dr -A
kubectl apply -f vs-1.yaml
kubectl get vs -A

SLEEP_POD=$(kubectl get pod -l app=sleep -o jsonpath={.items..metadata.name})

# Check the Endpoints
istioctl proxy-config endpoints $SLEEP_POD | grep 7777

# Test Traffic Routing
TARGET_URL="hw-svc.istio-testapp.svc.cluster.local:7777"
while true; do kubectl exec "$SLEEP_POD" -c sleep -- curl -sS $TARGET_URL; done;
```