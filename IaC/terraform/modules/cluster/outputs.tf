output "cluster_name" {
  description = "The name of the Kind cluster"
  value       = kind_cluster.my_cluster.name
}
