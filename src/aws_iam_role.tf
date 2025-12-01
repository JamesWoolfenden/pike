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

resource "aws_iam_policy" "version" {
  name_prefix = "terraform_pike_"

  description = "permissions to create and delete a policy"

  policy = jsonencode({
    "Sid" : "ManagePolicyVersions",
    "Effect" : "Allow",
    "Action" : [
      "iam:CreatePolicyVersion",
      "iam:DeletePolicyVersion"
    ],
    "Resource" : [
      aws_iam_policy.terraform_pike.arn
    ]
  })

}



data "aws_caller_identity" "current" {}


resource "aws_iam_role_policy_attachment" "pike-attach" {
  role       = aws_iam_role.terraform_pike.name
  policy_arn = aws_iam_policy.terraform_pike.arn
}

resource "aws_iam_role_policy_attachment" "pike-attach-versions" {
  role       = aws_iam_role.terraform_pike.name
  policy_arn = aws_iam_policy.version.arn
}



output "arn" {
  value = aws_iam_role.terraform_pike.arn
}
