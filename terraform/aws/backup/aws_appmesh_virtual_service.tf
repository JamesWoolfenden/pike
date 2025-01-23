resource "aws_appmesh_virtual_service" "pike" {
  name      = "servicea.simpleapp.local"
  mesh_name = aws_appmesh_mesh.pike.id

  spec {
    provider {
      virtual_node {
        virtual_node_name = aws_appmesh_virtual_node.serviceb1.name
      }
    }
  }
}
