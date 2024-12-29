terraform {
  required_providers {
    kind = {
      source  = "tehcyx/kind"
      version = "~> 0.6.0"
    }

    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.11.0"
    }
  }
}

## CLUSTER CONFIG

resource "kind_cluster" "my_cluster" {
  name       = var.cluster_name
  node_image = var.node_image

  kind_config {
    kind        = "Cluster"
    api_version = "kind.x-k8s.io/v1alpha4"

    networking {
      pod_subnet         = var.pod_subnet
      service_subnet     = var.service_subnet
      api_server_address = var.api_server_address
      api_server_port    = var.api_server_port
    }

    dynamic "node" {
      for_each = var.nodes
      content {
        role = node.value.role
      }
    }
  }
}

## NAMESPACES

provider "kubernetes" {
  config_path    = "~/.kube/config"
  config_context = var.kind_cluster_name
}

resource "kubernetes_namespace" "namespace" {
  for_each = var.namespaces

  metadata {
    name = each.key

    annotations = {
      provider = each.value
    }

    labels = {
      cluster_name = var.cluster_name
    }
  }

  depends_on = [
    kind_cluster.my_cluster  # Ensure the cluster is created before the namespaces
  ]
}