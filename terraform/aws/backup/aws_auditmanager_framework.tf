resource "aws_auditmanager_framework" "pike" {
  name = "example"

  control_sets {
    name = "example"
    controls {
      id = aws_auditmanager_control.pike.id
    }
  }
}
