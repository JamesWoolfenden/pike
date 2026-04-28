resource "azurerm_kubernetes_cluster_deployment_safeguard" "pike_gen" {
  kubernetes_cluster_id        = "pike"
  level                        = "Enforce"
  excluded_namespaces          = ["my-app-namespace", "legacy-app"]
  pod_security_standards_level = "Restricted"
}
