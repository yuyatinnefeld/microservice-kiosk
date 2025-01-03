# Cluster Setup

## Setup Kind-Cluster with Kind Config Manifest
```bash
cd IaC/kind
kind create cluster --name=my-cluster --config=kind-config-my-cluster.yaml
kubectl apply -f namespace.yaml
```

## Setup Kind-Cluster via Terraform
```bash
cd IaC/terraform
export TF_LOG=DEBUG
terraform init
terraform plan -var-file=dev.tfvars
terraform apply -var-file=dev.tfvars
```

## Setup Default Cluster and Namespace
```bash
kubectl config use-context kind-my-cluster
kubectl config set-context --current --namespace=shop
```