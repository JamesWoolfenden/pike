resource "aws_route53_zone" "primary" {
  name    = "myexample.com"
  comment = "Pike example"
  vpc {
    vpc_id = "vpc-0103cd49182c5cca5"
  }

  tags = {
    pike   = "Permission"
    remove = "this"
  }

}
