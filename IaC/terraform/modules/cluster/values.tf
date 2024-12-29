variable "cluster_name" {
  description = "Name of the cluster"
  type        = string
}

variable "kind_cluster_name" {
  description = "Name of the Kind cluster"
  type        = string
}

variable "node_image" {
  description = "Image to use for the nodes"
  type        = string
}

variable "pod_subnet" {
  description = "Subnet for Pods"
  type        = string
}

variable "service_subnet" {
  description = "Subnet for Services"
  type        = string
}

variable "api_server_address" {
  description = "Address for the API server"
  type        = string
}

variable "api_server_port" {
  description = "Port for the API server"
  type        = number
}

variable "nodes" {
  description = "List of nodes for the cluster"
  type = list(object({
    role = string
  }))
}

variable "namespaces" {
  description = "Map of namespaces to be created with annotations"
  type        = map(string)
  default = {
    "test"         = "test"
  }
}