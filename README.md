# Microservice Kiosk 🍬
Welcome to the Microservice Kiosk project! This repository is a hands-on exploration of modern cloud-native technologies, including ArgoCD, Istio as a service mesh, HashiCorp Vault, Terraform, Google Cloud Platform (GCP), and more. The goal is to deepen understanding and practical experience with these tools in a microservices architecture. In a typical microservices architecture, each service is often stored in its own repository. However, for the purpose of this project, I have consolidated all important source code into one repo to simplify management.

## Deployment Architecture
![Screenshot](/images/cnk-deployment.png)

- Provision the kind-cluster using Terraform.
- Build images with GitHub Actions and push to DockerHub.
- Manage secrets with Vault and ExternalSecretOperator.
- Implement service mesh using Istio.
- Conduct application and Istio testing via Testkube.
- Deploy and manage the cluster using ArgoCD.


## Application Architecture
![Screenshot](/images/cnk-application.png)

Key components include:
- Istio Gateway (GW) for routing traffic, connected to VirtualService (VS) and DestinationRule (DR) for traffic management.
- Frontend service, interacting with users and external APIs.
- Backend services like item, account, inventory, and ML for handling business logic and data processing.
- Istiod in the istio-system manages service mesh configurations.
- Self-signed Certs for secure communication is pushed into the Vault via and External Secret Operator.

## Table of Contents
- [Cluster Provisioning: Kind-Cluster + Terraform](IaC/README.md)
- [Continuous Deployment: ArgoCD](argocd/README.md)
- [Software: Shop-Apps + Istio-Test-Apps](app/README.md)
- [Cloud Native Testing: Testkube](testkube/README.md)
- [Secret Manament: Vault + External Secret Operator](vault/README.md)
- [Service Mesh: Istio](istio/README.md)
- [Localtest: External IP](metallb/README.md)
- Observability:  Prometheus + Grafana + Kiali
