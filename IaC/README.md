# Cluster Setup

## Setup Kind-Cluster via Terraform
```bash
cd IaC
export TF_LOG=DEBUG
terraform init
terraform plan -var-file=dev.tfvars”
terraform approve -var-file=dev.tfvars”

# Switch to the Service Mesh Testing Namespace:
kubectl config set-context --current --namespace=istio-testapp
```