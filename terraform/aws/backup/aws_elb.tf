resource "aws_elb" "example" {
  #  availability_zones = ["eu-west-2a","eu-west-2b"]
  subnets = ["subnet-03fdfb13a135366a7", "subnet-05a6a6de2f4989d22", "subnet-08d97e381dbc80d40"]
  listener {
    instance_port     = 80
    instance_protocol = "http"
    lb_port           = 80
    lb_protocol       = "http"
  }
  #  tags = {
  #    pike="permissions"
  #  }
}
