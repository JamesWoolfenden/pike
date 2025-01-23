resource "aws_appmesh_mesh" "pike" {
  name = "simpleapp"

  spec {
    egress_filter {
      type = "ALLOW_ALL"
    }
  }

  tags = {
    pike = "permimssion"
  }
}
