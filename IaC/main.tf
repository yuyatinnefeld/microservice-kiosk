module "kind_cluster" {
  source             = "./modules/cluster"
  cluster_name       = var.cluster_name
  node_image         = "kindest/node:v1.27.1"
  pod_subnet         = "10.244.0.0/16"
  service_subnet     = "10.96.0.0/12"
  api_server_address = "127.0.0.1"
  api_server_port    = 6443
  nodes = [
    { role = "control-plane" },
    { role = "worker" }
  ]
}