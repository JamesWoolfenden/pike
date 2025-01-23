resource "aws_appmesh_virtual_gateway" "pike" {
  name      = "example-virtual-gateway"
  mesh_name = aws_appmesh_mesh.pike.name

  spec {
    listener {
      port_mapping {
        port     = 8080
        protocol = "http"
      }
    }
  }

  tags = {
    Environment = "test"
  }
}
