resource "aws_storagegateway_cache" "pike" {
  disk_id     = data.aws_storagegateway_local_disk.test.id
  gateway_arn = aws_storagegateway_gateway.pike.arn
}

data "aws_storagegateway_local_disk" "test" {
  disk_node   = aws_volume_attachment.test.device_name
  gateway_arn = aws_storagegateway_gateway.pike.arn
}

resource "aws_volume_attachment" "test" {
  device_name = "/dev/xvdb"
  volume_id   = aws_ebs_volume.test.id
  instance_id = aws_instance.test.id
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

resource "aws_instance" "test" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = "t3.micro"

  tags = {
    Name = "HelloWorld"
  }
}


resource "aws_ebs_volume" "test" {
  availability_zone = "eu-west-2a"
  size              = 50
  encrypted         = true
  type              = "gp3"
  #  tags = {
  #    Name = "HelloWorld"
  #    pike="permissions"
  #  }
}
