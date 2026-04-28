resource "azurerm_workloads_sap_three_tier_virtual_instance" "pike_gen" {
  name                        = "X05"
  resource_group_name         = "pike"
  location                    = "pike"
  environment                 = "NonProd"
  sap_product                 = "S4HANA"
  managed_resource_group_name = "exampleManagedRG"
  app_location                = "pike"
  sap_fqdn                    = "sap.bpaas.com"

  three_tier_configuration {
    app_resource_group_name = "pike"
    secondary_ip_enabled    = true

    application_server_configuration {
      instance_count = 1
      subnet_id      = "pike"

      virtual_machine_configuration {
        virtual_machine_size = "Standard_D16ds_v4"

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
    }

    central_server_configuration {
      instance_count = 1
      subnet_id      = "pike"

      virtual_machine_configuration {
        virtual_machine_size = "Standard_D16ds_v4"

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
    }

    database_server_configuration {
      instance_count = 1
      subnet_id      = "pike"
      database_type  = "HANA"

      virtual_machine_configuration {
        virtual_machine_size = "Standard_E16ds_v4"

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
    }

    resource_names {
      application_server {
        availability_set_name = "appAvSet"

        virtual_machine {
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

      central_server {
        availability_set_name = "csAvSet"

        load_balancer {
          name                            = "ascslb"
          backend_pool_names              = ["ascsBackendPool"]
          frontend_ip_configuration_names = ["ascsip0"]
          health_probe_names              = ["ascsHealthProbe"]
        }

        virtual_machine {
          host_name               = "ascshostName"
          os_disk_name            = "ascsosdisk"
          virtual_machine_name    = "ascsvm"
          network_interface_names = ["ascsnic"]

          data_disk {
            volume_name = "default"
            names       = ["ascsdisk"]
          }
        }
      }

      database_server {
        availability_set_name = "dbAvSet"

        load_balancer {
          name                            = "dblb"
          backend_pool_names              = ["dbBackendPool"]
          frontend_ip_configuration_names = ["dbip"]
          health_probe_names              = ["dbHealthProbe"]
        }

        virtual_machine {
          host_name               = "dbprhost"
          os_disk_name            = "dbprosdisk"
          virtual_machine_name    = "dbvmpr"
          network_interface_names = ["dbprnic"]

          data_disk {
            volume_name = "hanaData"
            names       = ["hanadatapr0", "hanadatapr1"]
          }

          data_disk {
            volume_name = "hanaLog"
            names       = ["hanalogpr0", "hanalogpr1", "hanalogpr2"]
          }

          data_disk {
            volume_name = "usrSap"
            names       = ["usrsappr0"]
          }

          data_disk {
            volume_name = "hanaShared"
            names       = ["hanasharedpr0", "hanasharedpr1"]
          }
        }
      }

      shared_storage {
        account_name          = "sharedexamplesa"
        private_endpoint_name = "examplePE"
      }
    }

    transport_create_and_mount {
      resource_group_id    = "pike"
      storage_account_name = "exampletranssa"
    }
  }

  identity {
    type = "UserAssigned"

    identity_ids = [
      azurerm_user_assigned_identity.example.id,
    ]
  }

  tags = {
    Env = "Test"
  }

  depends_on = [
    azurerm_role_assignment.example
  ]
}
