resource "azurerm_workloads_sap_single_node_virtual_instance" "pike_gen" {
  name                        = "X05"
  resource_group_name         = "pike"
  location                    = "pike"
  environment                 = "NonProd"
  sap_product                 = "S4HANA"
  managed_resource_group_name = "managedTestRG"
  app_location                = "pike"
  sap_fqdn                    = "sap.bpaas.com"

  single_server_configuration {
    app_resource_group_name = "pike"
    subnet_id               = "pike"
    database_type           = "HANA"
    secondary_ip_enabled    = true

    virtual_machine_configuration {
      virtual_machine_size = "Standard_E32ds_v4"

      image {
        offer     = "RHEL-SAP-HA"
        publisher = "RedHat"
        sku       = "82sapha-gen2"
        version   = "latest"
      }

      os_profile {
        admin_username  = "testAdmin"
        ssh_private_key = "pike"
        ssh_public_key  = "pike"
      }
    }

    disk_volume_configuration {
      volume_name     = "hana/data"
      number_of_disks = 3
      size_in_gb      = 128
      sku_name        = "Premium_LRS"
    }

    disk_volume_configuration {
      volume_name     = "hana/log"
      number_of_disks = 3
      size_in_gb      = 128
      sku_name        = "Premium_LRS"
    }

    disk_volume_configuration {
      volume_name     = "hana/shared"
      number_of_disks = 1
      size_in_gb      = 256
      sku_name        = "Premium_LRS"
    }

    disk_volume_configuration {
      volume_name     = "usr/sap"
      number_of_disks = 1
      size_in_gb      = 128
      sku_name        = "Premium_LRS"
    }

    disk_volume_configuration {
      volume_name     = "backup"
      number_of_disks = 2
      size_in_gb      = 256
      sku_name        = "StandardSSD_LRS"
    }

    disk_volume_configuration {
      volume_name     = "os"
      number_of_disks = 1
      size_in_gb      = 64
      sku_name        = "StandardSSD_LRS"
    }

    virtual_machine_resource_names {
      host_name               = "apphostName0"
      os_disk_name            = "app0osdisk"
      virtual_machine_name    = "appvm0"
      network_interface_names = ["appnic0"]

      data_disk {
        volume_name = "default"
        names       = ["app0disk0"]
      }
    }
  }

  identity {
    type = "UserAssigned"

    identity_ids = [
      azurerm_user_assigned_identity.example.id,
    ]
  }

  depends_on = [
    azurerm_role_assignment.example
  ]
}
