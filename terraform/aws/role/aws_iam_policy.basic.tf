resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //dependson
          "ec2:CreateSecurityGroup",
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DeleteSecurityGroup",

          //aws_vpc_ipam
          "ec2:CreateIpam",
          "ec2:ModifyIpam",
          "iam:CreateServiceLinkedRole",
          "ec2:DescribeIpams",
          "ec2:DeleteIpam",
          "ec2:CreateTags",
          "ec2:DeleteTags",
          "ec2:DescribeTags",

          //aws_vpc_ipam_pool
          "ec2:CreateIpamPool",
          "ec2:DescribeIpamPools",
          "ec2:DeleteIpamPool",
          "ec2:CreateTags",
          "ec2:DeleteTags",
          "ec2:ModifyIpamPool",

          //aws_vpc_ipam_pool_cidr
          "ec2:ProvisionIpamPoolCidr",
          "ec2:GetIpamPoolCidrs",

          //aws_vpc_ipam_preview_next_cidr
          "ec2:AllocateIpamPoolCidr",

          //aws_vpc_ipam_pool_cidr_allocation
          "ec2:AllocateIpamPoolCidr",
          "ec2:GetIpamPoolAllocations",
          "ec2:ReleaseIpamPoolAllocation",
          "ec2:DeprovisionIpamPoolCidr",
        ],
        "Resource" : "*",
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
