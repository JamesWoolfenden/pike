
resource "aws_cloudwatch_log_delivery_source" "example" {
  name         = "example"
  log_type     = "APPLICATION_LOGS"
  resource_arn = aws_instance.pike.arn
}


data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_instance" "pike" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = "t3.micro"
}
