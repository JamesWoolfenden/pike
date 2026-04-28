resource "azurerm_oracle_cloud_vm_cluster" "pike_gen" {
  name                            = "example-cloud-vm-cluster"
  resource_group_name             = "pike"
  location                        = "pike"
  gi_version                      = "23.0.0.0"
  virtual_network_id              = "pike"
  license_model                   = "BringYourOwnLicense"
  db_servers                      = [for obj in "pike" : "pike"]
  ssh_public_keys                 = [file("~/.ssh/id_rsa.pub")]
  display_name                    = "example-cloud-vm-cluster"
  cloud_exadata_infrastructure_id = "pike"
  cpu_core_count                  = 2
  hostname                        = "hostname"
  subnet_id                       = "pike"
  system_version                  = "23.1.19.0.0.241015"
  file_system_configuration {
    mount_point = "/var"
    size_in_gb  = 32
  }
}
