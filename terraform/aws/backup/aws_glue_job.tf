resource "aws_glue_job" "pike" {
  name        = "example"
  role_arn    = "arn:aws:iam::680235478471:role/PikeGlueRole"
  description = "For Pike"

  command {
    script_location = "s3://testbucketforlbjgw/hello.py"
  }

  timeout = 2880
  tags = {
    pike = "permissions"
    //  createdby = "JamesWoolfenden"
  }
}
