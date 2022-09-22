resource "aws_directory_service_directory" "pike" {
  name     = "corp.pike.com"
  password = "PikeI$Permission2"
  vpc_settings {
    vpc_id     = "vpc-032c4e15f24d3628e"
    subnet_ids = ["subnet-04338b6369d8288a5", "subnet-0d99e0f4e4f1ff84f"]
  }

  tags = {
    pike   = "permissions"
    delete = "me"
  }
}
