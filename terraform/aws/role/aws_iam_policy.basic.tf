resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "imagebuilder:CreateInfrastructureConfiguration",
          "imagebuilder:GetInfrastructureConfiguration",
          "imagebuilder:DeleteInfrastructureConfiguration",
          "imagebuilder:UpdateInfrastructureConfiguration",
          "iam:PassRole",
          "sns:Publish",

          "ec2:DescribeAccountAttributes",
          "imagebuilder:CreateComponent",
          "imagebuilder:GetComponent",
          "imagebuilder:DeleteComponent",

          #          "iam:CreateServiceLinkedRole",
          #
          "kms:GenerateDataKey",
          "kms:GenerateDataKeyWithoutPlaintext",
          "kms:Encrypt",
          "kms:Decrypt",


          "imagebuilder:CreateImageRecipe",
          "imagebuilder:GetImageRecipe",
          "imagebuilder:DeleteImageRecipe",
          "ec2:DescribeImages",
          "imagebuilder:Getimage",
          "imagebuilder:GetComponent",
          "iam:CreateServiceLinkedRole",

          "imagebuilder:CreateImagePipeline",
          "imagebuilder:GetImagePipeline",
          "imagebuilder:DeleteImagePipeline",
          "imagebuilder:UpdateImagePipeline",
          "imagebuilder:GetImageRecipe",
          "iam:CreateServiceLinkedRole",
          "imagebuilder:GetContainerRecipe",


          "imagebuilder:CreateDistributionConfiguration",
          "imagebuilder:DeleteDistributionConfiguration",
          "imagebuilder:GetDistributionConfiguration",
          "iam:CreateServiceLinkedRole",
          #          "imagebuilder:ListDistributionConfigurations",
          "imagebuilder:UpdateDistributionConfiguration",

          "imagebuilder:TagResource",
          "imagebuilder:UntagResource",
          "imagebuilder:CreateImage",
          "imagebuilder:GetImage",
          "imagebuilder:DeleteImage",


        ]
        "Resource" : "*"
      }
    ]
  })
  tags = {
    pike      = "permissions"
    createdby = "JamesWoolfenden"
  }
}

resource "aws_iam_role_policy_attachment" "basic" {
  role       = aws_iam_role.basic.name
  policy_arn = aws_iam_policy.basic.arn
}

resource "aws_iam_user_policy_attachment" "basic" {
  # checkov:skip=CKV_AWS_40: By design
  user       = "basic"
  policy_arn = aws_iam_policy.basic.arn
}
