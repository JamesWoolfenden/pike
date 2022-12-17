data "aws_ebs_volume" "pike" {
  most_recent = true

  filter {
    name   = "volume-type"
    values = ["gp2"]
  }

}
