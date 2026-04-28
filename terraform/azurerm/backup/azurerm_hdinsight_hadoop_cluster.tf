resource "azurerm_hdinsight_hadoop_cluster" "pike_gen" {
  name                = "example-hdicluster"
  resource_group_name = "pike"
  location            = "pike"
  cluster_version     = "3.6"
  tier                = "Standard"

  component_version {
    hadoop = "2.7"
  }

  gateway {
    username = "acctestusrgw"
  }

  storage_account {
    storage_container_id = "pike"
    storage_account_key  = "pike"
    is_default           = true
  }

  roles {
    head_node {
      vm_size  = "Standard_D3_V2"
      username = "acctestusrvm"
    }

    worker_node {
      vm_size               = "Standard_D4_V2"
      username              = "acctestusrvm"
      target_instance_count = 3
    }

    zookeeper_node {
      vm_size  = "Standard_D3_V2"
      username = "acctestusrvm"
    }
  }
}
