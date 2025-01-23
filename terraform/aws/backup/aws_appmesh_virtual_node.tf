resource "aws_appmesh_virtual_node" "serviceb1" {
  name      = "serviceBv1"
  mesh_name = aws_appmesh_mesh.pike.id

  spec {
    backend {
      virtual_service {
        virtual_service_name = "servicea.simpleapp.local"
      }
    }

    listener {
      port_mapping {
        port     = 8080
        protocol = "http"
      }
    }

    service_discovery {
      dns {
        hostname = "serviceb.simpleapp.local"
      }
    }
  }
}

resource "aws_appmesh_virtual_node" "serviceb2" {
  name      = "serviceBv2"
  mesh_name = aws_appmesh_mesh.pike.id

  spec {
    backend {
      virtual_service {
        virtual_service_name = "servicea.simpleapp.local"
      }
    }

    listener {
      port_mapping {
        port     = 8080
        protocol = "http"
      }
    }

    service_discovery {
      dns {
        hostname = "serviceb2.simpleapp.local"
      }
    }
  }
}
