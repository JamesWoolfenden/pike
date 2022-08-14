resource "aws_spot_instance_request" "a" {
  //associate_public_ip_address = false
  ami                         = "ami-078a289ddf4b09ae0"
  instance_type               = "t3a.nano"
  subnet_id                   = "subnet-03fdfb13a135366a7"
  associate_public_ip_address = true
  spot_price                  = "0.001"
  # disable_api_stop = false
  # disable_api_termination = false
  #ebs_optimized=true
  #iam_instance_profile =
  key_name   = "test"
  monitoring = false
  tags = {
    "createdby" = "james"
  }
}
