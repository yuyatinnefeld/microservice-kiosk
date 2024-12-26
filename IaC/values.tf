variable "cluster_name" {
  description = "The name of the Kubernetes cluster"
  type        = string
}

variable "node_image" {
  description = "The Docker image for the nodes in the Kind cluster"
  type        = string
  default     = "kindest/node:v1.27.1"
}

variable "pod_subnet" {
  description = "The pod subnet for the Kind cluster"
  type        = string
  default     = "10.244.0.0/16"
}

variable "service_subnet" {
  description = "The service subnet for the Kind cluster"
  type        = string
  default     = "10.96.0.0/12"
}

variable "api_server_address" {
  description = "The address of the API server for the Kind cluster"
  type        = string
  default     = "127.0.0.1"
}

variable "api_server_port" {
  description = "The port of the API server for the Kind cluster"
  type        = number
  default     = 6443
}

variable "nodes" {
  description = "The list of nodes for the Kind cluster"
  type = list(object({
    role = string
  }))
  default = [
    { role = "control-plane" },
    { role = "worker" }
  ]
}
