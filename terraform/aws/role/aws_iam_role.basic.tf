resource "aws_iam_role" "basic" {
  name = "pike-test-role"
  assume_role_policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Principal" : { "AWS" : "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root" },
          "Action" : "sts:AssumeRole",
        }
      ]
    }
  )
}


data "aws_caller_identity" "current" {}
