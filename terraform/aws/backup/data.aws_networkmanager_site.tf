data "aws_networkmanager_site" "pike" {
  global_network_id = "var.global_network_id"
  site_id           = "var.site_id"
}
