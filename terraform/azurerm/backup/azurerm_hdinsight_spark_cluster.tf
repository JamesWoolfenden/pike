resource "azurerm_hdinsight_spark_cluster" "pike_gen" {
  name                = "example-hdicluster"
  resource_group_name = "pike"
  location            = "pike"
  cluster_version     = "3.6"
  tier                = "Standard"

  component_version {
    spark = "2.3"
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
      vm_size  = "Standard_A3"
      username = "acctestusrvm"
    }

    worker_node {
      vm_size               = "Standard_A3"
      username              = "acctestusrvm"
      target_instance_count = 3
    }

    zookeeper_node {
      vm_size  = "Medium"
      username = "acctestusrvm"
    }
  }
}
