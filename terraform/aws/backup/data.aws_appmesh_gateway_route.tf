data "aws_appmesh_gateway_route" "pike" {
  mesh_name            = "pike"
  name                 = "pike"
  virtual_gateway_name = "pike"
}
