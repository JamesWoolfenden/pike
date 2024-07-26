resource "aws_workspaces_ip_group" "pike" {
  name        = "Contractors"
  description = "Contractors IP access control group"
  rules {
    source      = "150.24.14.0/24"
    description = "NY"
  }
  rules {
    source      = "125.191.14.85/32"
    description = "LA"
  }
  rules {
    source      = "44.98.100.0/24"
    description = "STL"
  }
}
