# ArgoCD
```bash
# deploy argocd server
helm repo add argo https://argoproj.github.io/argo-helm
helm repo update
helm install argocd argo/argo-cd --namespace argocd
kubectl get all -n argocd

# username: admin, password:
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d

# open argocd ui
kubectl port-forward svc/argocd-server -n argocd 8080:443

# deploy service mesh components
kubectl apply -f argocd/applicationset-istio-system.yaml
kubectl apply -f argocd/applicationset-istio-gateways.yaml

# deploy service mesh test components
kubectl apply -f argocd/applicationset-istio-testapp.yaml

# deploy service mesh monitor components
kubectl apply -f argocd/applicationset-istio-monitor.yaml

# deploy testkube components
kubectl apply -f argocd/applicationset-testkube-init.yaml
kubectl apply -f argocd/applicationset-testkube-tests.yaml

# vault components
kubectl apply -f argocd/applicationset-vault.yaml
```