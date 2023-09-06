data "aws_appmesh_route" "pike" {
  mesh_name           = "pike"
  name                = "pike"
  virtual_router_name = "pike"
}
