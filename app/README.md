# Target Application Setup

Navigate to the Application directory:

```bash
cd app
```

## Install Required Docker Images

Before installing the target applications, run the following commands to work around the image pull error by kind. This ensures that the necessary Docker images are available:

```bash
# First pull the image in your local system using docker pull xxx and then use below command to load that image to the kind cluster
docker pull hashicorp/http-echo
docker pull fortio/fortio
docker pull sverrirab/sleep


# Kind uses containerd instead of docker as runtime, that's why docker is not installed on the nodes.
kind load docker-image hashicorp/http-echo --name north-cluster
kind load docker-image fortio/fortio --name north-cluster
kind load docker-image sverrirab/sleep --name north-cluster
```

## Install the Target Applications
To deploy the target applications, apply the following configurations:
```bash
k apply -f app/$CLUSTER_NAME/fortio.yaml
k apply -f app/$CLUSTER_NAME/hw.yaml
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