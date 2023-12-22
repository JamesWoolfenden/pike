resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "connect:DescribeContactFlow",
          "connect:CreateContactFlow",
          "connect:DeleteContactFlow",
          "connect:UntagResource",
          "connect:TagResource",

          "connect:AssociateBot",
          "connect:DisassociateBot",

          "connect:CreateContactFlowModule",
          "connect:DeleteContactFlowModule",
          "connect:DescribeContactFlowModule",
          "connect:UntagResource",
          "connect:TagResource",

          "connect:CreateHoursOfOperation",
          "connect:DeleteHoursOfOperation",
          "connect:UpdateHoursOfOperation",
          "connect:DescribeHoursOfOperation",
          "connect:UntagResource",
          "connect:TagResource",

          //aws_connect_instance
          "connect:DescribeInstanceAttribute",
          "connect:CreateInstance",
          "connect:DeleteInstance",
          "connect:DescribeInstance",
          //"ds:CreateAlias",
          "ds:CheckAlias",
          "iam:CreateServiceLinkedRole",
          "ds:DescribeDirectories",

          "connect:DescribeInstanceStorageConfig",
          "connect:UpdateInstanceStorageConfig",
          "connect:DisassociateInstanceStorageConfig",
          "connect:AssociateInstanceStorageConfig",

          "connect:AssociateLambdaFunction",
          "connect:DisassociateLambdaFunction",
          "connect:ListLambdaFunctions",

          //aws_connect_phone_number
          "connect:SearchAvailablePhoneNumbers",
          "connect:DescribePhoneNumber",
          "connect:UpdatePhoneNumber",
          "connect:ClaimPhoneNumber",

          "connect:DescribeQueue",
          "connect:CreateQueue",
          "connect:DeleteQueue",
          "connect:ListSecurityProfilePermissions",

          "connect:DescribeQuickConnect",
          "connect:CreateQuickConnect",
          "connect:DeleteQuickConnect",
          "connect:ListQueueQuickConnects",
          "connect:UntagResource",
          "connect:TagResource",

          "connect:DescribeRoutingProfile",
          "connect:CreateRoutingProfile",
          "connect:DeleteRoutingProfile",

          "connect:DescribeSecurityProfile",
          "connect:CreateSecurityProfile",
          "connect:UpdateSecurityProfile",
          "connect:DeleteSecurityProfile",
          "connect:ListSecurityProfilePermissions",
          "connect:UntagResource",
          "connect:TagResource",

          "connect:DescribeUser",
          "connect:CreateUser",
          "connect:DeleteUser",

          "connect:DescribeUserHierarchyGroup",
          "connect:CreateUserHierarchyGroup",
          "connect:DeleteUserHierarchyGroup",
          "connect:UntagResource",
          "connect:TagResource",

          "connect:DescribeUserHierarchyStructure",
          "connect:UpdateUserHierarchyStructure",

          "connect:DescribeVocabulary",
          "connect:CreateVocabulary",
          "connect:DeleteVocabulary",
          "connect:UntagResource",
          "connect:TagResource",


          "firehose:CreateDeliveryStream",
          "firehose:DeleteDeliveryStream",
          "firehose:DescribeDeliveryStream",
          "firehose:ListTagsForDeliveryStream",
          "firehose:TagDeliveryStream",
          "firehose:UntagDeliveryStream",
          "iam:PassRole"
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
