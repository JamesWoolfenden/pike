resource "aws_iam_user" "example" {
  name = "example"
  tags = {
    tag-key = "tag-value"
  }
}
