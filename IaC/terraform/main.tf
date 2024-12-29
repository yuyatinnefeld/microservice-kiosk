# TODO: Activate Github Action for CD
# TODO: Use GCP Storage for Statefile Managment 

module "kind_cluster" {
  source             = "./modules/cluster"
  cluster_name       = var.cluster_name
  kind_cluster_name  = var.kind_cluster_name
  node_image         = "kindest/node:v1.27.1"
  pod_subnet         = "10.244.0.0/16"
  service_subnet     = "10.96.0.0/12"
  api_server_address = "127.0.0.1"
  api_server_port    = 6443
  nodes = [
    { role = "control-plane" },
    { role = "worker" }
  ]
  
  namespaces = {
    "argocd"         = "argocd"
    "istio-gateways" = "istio-gateways"
    "istio-monitor"  = "istio-monitor"
    "istio-system"   = "istio-system"
    "istio-testapp"  = "istio-testapp"
    "shop"           = "shop"
    "testkube"       = "testkube"
    "vault"          = "vault"
  }
}
