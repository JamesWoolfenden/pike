resource "aws_appmesh_gateway_route" "pike" {
  name                 = "example-gateway-route"
  mesh_name            = aws_appmesh_mesh.pike.name
  virtual_gateway_name = aws_appmesh_virtual_gateway.pike.name

  spec {
    http_route {
      action {
        target {
          virtual_service {
            virtual_service_name = aws_appmesh_virtual_service.pike.name
          }
        }
      }

      match {
        prefix = "/"
      }
    }
  }

  tags = {
    Environment = "test"
  }
}
