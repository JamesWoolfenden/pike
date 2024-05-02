resource "aws_datasync_location_hdfs" "pike" {
  agent_arns          = [aws_datasync_agent.pike.arn]
  authentication_type = "SIMPLE"
  simple_user         = "example"

  name_node {
    hostname = aws_instance.example.private_dns
    port     = 80
  }
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

resource "aws_instance" "example" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = "t3.micro"

  tags = {
    Name = "HelloWorld"
  }
}
