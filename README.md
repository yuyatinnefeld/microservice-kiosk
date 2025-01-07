# Microservice Kiosk 🍬
Welcome to the Microservice Kiosk project! This repository is a hands-on exploration of modern cloud-native technologies, including ArgoCD, Istio as a service mesh, HashiCorp Vault, Terraform, Google Cloud Platform (GCP), and more. The goal is to deepen understanding and practical experience with these tools in a microservices architecture. In a typical microservices architecture, each service is often stored in its own repository. However, for the purpose of this project, I have consolidated all important source code into one repo to simplify management.

## Deployment Architecture
![Screenshot](/images/cnk-architecture.png)

## Application Architecture
![Screenshot](/images/cnk-application-diagram.png)

## Table of Contents
- [Cluster Provisioning: Kind-Cluster + Terraform](IaC/README.md)
- [Continuous Deployment: ArgoCD](argocd/README.md)
- [Software: Shop-Apps + Istio-Test-Apps](app/README.md)
- [Cloud Native Testing: Testkube](testkube/README.md)
- [Secret Manament: Vault + External Secret Operator](vault/README.md)
- [Service Mesh: Istio](istio/README.md)
- [Localtest: External IP](metallb/README.md)
- Observability:  Prometheus + Grafana + Kiali
