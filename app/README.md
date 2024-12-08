# Target Application Setup

Navigate to the Application directory:

```bash
cd app
```

## Access the hw1 Application
Define the namespace and get the Fortio pod name:
```bash
NS=istio-testapp
FORTIO_POD="$(kubectl get pod -n $NS -l app=fortio -o jsonpath='{.items[0].metadata.name}')"
TARGET_URL=hw-svc.$NS.svc.cluster.local:7777
```

Run Load Test with Fortio:
```bash
kubectl exec -it $FORTIO_POD --container fortio -- /usr/bin/fortio load -k -c 1 -qps 20 -t 0.5m -loglevel Warning $TARGET_URL
```

Forward Ports to Access the Service:
```bash
kubectl port-forward svc/hw-svc 7777:7777
```