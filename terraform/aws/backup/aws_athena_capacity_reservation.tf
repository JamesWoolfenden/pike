resource "aws_athena_capacity_reservation" "pike" {
  name        = "example-reservation"
  target_dpus = 24
  tags = {
    pike = "permissions"
  }
}
