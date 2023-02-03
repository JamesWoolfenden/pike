data "aws_dx_router_configuration" "pike" {
  virtual_interface_id   = "dxvif-abcde123"
  router_type_identifier = "CiscoSystemsInc-2900SeriesRouters-IOS124"
}
