resource "aws_iam_role" "terraform_pike" {
  name_prefix = "terraform_pike_"
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


resource "aws_iam_role_policy_attachment" "pike-attach" {
  role       = aws_iam_role.terraform_pike.name
  policy_arn = aws_iam_policy.terraform_pike.arn
}

output "arn" {
  value = aws_iam_role.terraform_pike.arn
}
