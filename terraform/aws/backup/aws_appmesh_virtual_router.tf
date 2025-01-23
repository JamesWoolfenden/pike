resource "aws_appmesh_virtual_router" "pike" {
  name      = "serviceB"
  mesh_name = aws_appmesh_mesh.pike.id

  spec {
    listener {
      port_mapping {
        port     = 8080
        protocol = "http"
      }
    }
  }
  tags = {
    pike = "permimssion"
  }
}
