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
          "lightsail:CreateInstances",
          "lightsail:TagResource",
          "lightsail:UntagResource",
          "lightsail:GetInstance",

          "lightsail:DeleteInstance",
          "lightsail:ReleaseStaticIp",

          "lightsail:CreateKeyPair",
          "lightsail:GetKeyPair",
          "lightsail:DeleteKeyPair",
          "lightsail:GetOperation",

          "lightsail:PutInstancePublicPorts",
          "lightsail:GetInstancePortStates",
          "lightsail:CloseInstancePublicPorts",

          "lightsail:GetStaticIp",
          "lightsail:AllocateStaticIp",


          "lightsail:AttachStaticIp",
          "lightsail:DetachStaticIp",
          "lightsail:ReleaseStaticIp"
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
