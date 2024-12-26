cluster_name       = "my-cluster"
node_image         = "kindest/node:v1.28.0"
pod_subnet         = "10.250.0.0/16"
service_subnet     = "10.100.0.0/12"
api_server_address = "192.168.0.1"
api_server_port    = 6444
nodes = [
  { role = "control-plane" },
  { role = "worker" },
  { role = "worker" }
]