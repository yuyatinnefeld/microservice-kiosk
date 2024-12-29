# Cluster Setup

## Setup Kind-Cluster with Kind Config Manifest
```bash
cd IaC/kind
kind create cluster --name=my-cluster --config=kind-config-my-cluster.yaml
kubectl apply -f cluster/namespace.yaml
```

## Setup Kind-Cluster via Terraform
```bash
cd IaC/terraform
export TF_LOG=DEBUG
terraform init
terraform plan -var-file=dev.tfvars
terraform approve -var-file=dev.tfvars

# Switch to the Service Mesh Testing Namespace:
kubectl config set-context --current --namespace=istio-testapp
```