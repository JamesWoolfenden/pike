resource "azurerm_vmware_netapp_volume_attachment" "pike_gen" {
  name              = "example-vmwareattachment"
  netapp_volume_id  = "pike"
  vmware_cluster_id = "pike"

  depends_on = ["pike"]
}
