resource "azurerm_redhat_openshift_cluster" "pike_gen" {
  name                = "examplearo"
  location            = "pike"
  resource_group_name = "pike"

  cluster_profile {
    domain  = "aro-example.com"
    version = "4.13.23"
  }

  network_profile {
    pod_cidr     = "10.128.0.0/14"
    service_cidr = "172.30.0.0/16"
  }

  main_profile {
    vm_size   = "Standard_D8s_v3"
    subnet_id = "pike"
  }

  api_server_profile {
    visibility = "Public"
  }

  ingress_profile {
    visibility = "Public"
  }

  worker_profile {
    vm_size      = "Standard_D4s_v3"
    disk_size_gb = 128
    node_count   = 3
    subnet_id    = "pike"
  }

  service_principal {
    client_id = "pike"
  }

  depends_on = [
    "azurerm_role_assignment.role_network1",
    "azurerm_role_assignment.role_network2",
  ]
}
