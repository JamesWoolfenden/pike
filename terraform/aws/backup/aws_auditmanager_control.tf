resource "aws_auditmanager_control" "pike" {
  name = "example"

  control_mapping_sources {
    source_name          = "example"
    source_set_up_option = "Procedural_Controls_Mapping"
    source_type          = "MANUAL"
  }
  tags = {
    pike = "permission"
  }
}
