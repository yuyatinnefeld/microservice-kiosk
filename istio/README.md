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
export INGRESS_NS=istio-gateways
export INGRESS_NAME=istio-ingressgateway
kubectl get svc "$INGRESS_NAME" -n "$INGRESS_NS"

NAME                   TYPE           CLUSTER-IP      EXTERNAL-IP    PORT(S)                                      AGE
istio-ingressgateway   LoadBalancer   10.103.68.167   172.18.0.152   15021:32691/TCP,80:32220/TCP,443:31012/TCP   5h44m

export INGRESS_HOST=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export INGRESS_PORT=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')

curl -s -I -HHost:gateway.cnk.com "http://$INGRESS_HOST:$INGRESS_PORT/backend"

# add DNS for gateway
vi /etc/hosts
# 172.18.0.152 gateway.cnk.com

curl gateway.cnk.com:80/backend
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
