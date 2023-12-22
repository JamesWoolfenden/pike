resource "aws_elb_attachment" "pike" {
  elb      = aws_elb.pike.id
  instance = aws_instance.pike.id
}


resource "aws_instance" "pike" {
  ami           = "ami-078a289ddf4b09ae0"
  instance_type = "t2.micro"
  subnet_id     = "subnet-03fdfb13a135366a7"
}

resource "aws_elb" "pike" {
  subnets = ["subnet-03fdfb13a135366a7", "subnet-05a6a6de2f4989d22", "subnet-08d97e381dbc80d40"]
  listener {
    instance_port     = 80
    instance_protocol = "http"
    lb_port           = 80
    lb_protocol       = "http"
  }
}
