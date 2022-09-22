resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "ec2:DescribeAccountAttributes",
          "glacier:CreateVault",
          "glacier:AddTagsToVault",
          "glacier:RemoveTagsFromVault",
          "glacier:ListTagsForVault",
          "glacier:DescribeVault",
          "glacier:GetVaultAccessPolicy",
          "glacier:DeleteVault",
          "glacier:SetVaultAccessPolicy",
          "glacier:SetVaultNotifications",


          "glacier:InitiateVaultLock",
          "glacier:GetVaultLock",
          "glacier:AbortVaultLock",
          "glacier:CompleteVaultLock"

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
