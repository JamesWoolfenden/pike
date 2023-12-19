resource "aws_vpclattice_listener_rule" "pike" {
  name                = "example"
  listener_identifier = aws_vpclattice_listener.pike.listener_id
  service_identifier  = aws_vpclattice_service.pike.id
  priority            = 10
  match {
    http_match {
      path_match {
        case_sensitive = false
        match {
          exact = "/example-path"
        }
      }
    }
  }
  action {
    fixed_response {
      status_code = 404
    }
  }

}
