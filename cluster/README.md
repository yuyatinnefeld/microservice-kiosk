# Cluster Setup

## Setup Proxy
To configure your proxy settings, run the following commands:
```bash
export http_proxy=""
export https_proxy="$http_proxy"
export ftp_proxy="$http_proxy"
export no_proxy=localhost,127.0.0.1
export NO_PROXY=$no_proxy
unset http_proxy
unset https_proxy
```

## Install 2 Clusters
You can create two Kubernetes clusters using Kind with the following commands:

```bash
kind create cluster --name=my-cluster --config=cluster/kind-config-my-cluster.yaml 
kind get clusters  # List the created clusters
```

## Create Testing Environment in the Cluster
To set up the testing environment, follow these steps:

```bash
# Set the Kind cluster as the current contex
kubectl config use-context kind-my-cluster

# Check current cluster
kubectl config view --minify -o jsonpath='{..namespace}'
kubectl cluster-info --context kind-my-cluster # https://127.0.0.1:6443

# Apply the Namespace Configuration:
kubectl apply -f cluster/namespace.yaml

# Switch to the Service Mesh Testing Namespace:
kubectl config set-context --current --namespace=istio-testapp
```
