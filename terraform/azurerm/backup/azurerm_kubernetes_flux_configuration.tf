resource "azurerm_kubernetes_flux_configuration" "pike_gen" {
  name       = "example-fc"
  cluster_id = "pike"
  namespace  = "flux"

  git_repository {
    url             = "https://github.com/Azure/arc-k8s-demo"
    reference_type  = "branch"
    reference_value = "main"
  }

  kustomizations {
    name = "kustomization-1"

    post_build {
      substitute = {
        example_var = "substitute_with_this"
      }
      substitute_from {
        kind = "ConfigMap"
        name = "example-configmap"
      }
    }
  }

  depends_on = [
    azurerm_kubernetes_cluster_extension.example
  ]
}
