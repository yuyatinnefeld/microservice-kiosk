# Istion

## Check Istio Config
```bash
# Verify
istioctl verify-install

# activate namespace sidecar injection
kubectl label namespace shop istio-injection=enabled
kubectl label namespace istio-testapp istio-injection=enabled

# check if namespaces are injected
kubectl get ns shop --show-labels
kubectl get namespace -L istio-injection

# testapp as default ns
kubectl config set-context --current --namespace=istio-testapp

## Verify Sidecar Injection
TARGET_POD="$(kubectl get pod -l app=httpbin -n istio-testapp -o jsonpath='{.items[0].metadata.name}')"
kubectl logs $TARGET_POD -c istio-proxy
ISTIOD_POD="$(kubectl get pod -n istio-system -l app=istiod -o jsonpath='{.items[0].metadata.name}')"
kubectl logs $ISTIOD_POD -n istio-system
```

## Check Monitoring Tools

```bash
# Prometheus
kubectl port-forward svc/prometheus -n istio-monitor 9090

# Kiali
kubectl port-forward svc/kiali -n istio-monitor 20001

# Grafana
kubectl port-forward svc/grafana -n istio-monitor 3000
```
