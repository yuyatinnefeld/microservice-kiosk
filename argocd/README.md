# ArgoCD
```bash
helm repo add argo https://argoproj.github.io/argo-helm
helm repo update
kubectl create namespace argocd
helm install argocd argo/argo-cd --namespace argocd
kubectl get all -n argocd

# username: admin, password:
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d

# open argocd ui
kubectl port-forward svc/argocd-server -n argocd 8080:443

# deploy applicationsets
kubectl apply -f argocd/applicationset.yaml
```