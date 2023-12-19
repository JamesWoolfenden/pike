resource "aws_vpclattice_listener" "pike" {
  name               = "pike"
  protocol           = "HTTPS"
  service_identifier = aws_vpclattice_service.pike.id
  default_action {
    fixed_response {
      status_code = 404
    }
  }
  tags = {
    pike = "permission"
  }
}
