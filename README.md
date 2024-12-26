# Microservice Kiosk 🍬
Welcome to the Microservice Kiosk project! This repository is a hands-on exploration of modern cloud-native technologies, including ArgoCD, Istio as a service mesh, HashiCorp Vault, Terraform, Google Cloud Platform (GCP), and more. The goal is to deepen understanding and practical experience with these tools in a microservices architecture. In a typical microservices architecture, each service is often stored in its own repository. However, for the purpose of this project, I have consolidated all important source code into one repo to simplify management.

## Table of Contents
- [Cluster Setup](cluster/README.md)
- [ArgoCD Deployment](argocd/README.md)
- [Application](app/README.md)
- [External IP Setup](metallb/README.md)
- [Cloud Native Testing](testkube/README.md)

## System Architecture
![Screenshot](/images/cnk-architecture.png)

### Namespace
- istio-system
- istio-testapp
- istio-monitor
- shop
- testkube
- argocd
- vault

### CI/CD
- CI: Github Action
- CD: argoCD