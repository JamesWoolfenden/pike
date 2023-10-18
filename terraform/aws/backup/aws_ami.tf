resource "aws_ami" "pike" {
  name = "pike"

  virtualization_type = "hvm"
  root_device_name    = "/dev/xvda"
  imds_support        = "v2.0"
  ebs_block_device {
    device_name = "/dev/xvda"
    snapshot_id = "snap-0a55469d8aa4dcdb9"
    volume_size = 8
  }

  ephemeral_block_device {
    device_name  = "/dev/xvdb"
    virtual_name = "ephemeral1"
  }

  tags = {
    pike = "permissions"
  }
}

output "ami" {
  value = aws_ami.pike
}
