resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_workspaces_image
          "workspaces:DescribeWorkspaceImages",
          //aws_workspaces_directory
          "workspaces:DescribeWorkspaceDirectories",
          //aws_iot_endpoint
          "iot:DescribeEndpoint",
          //aws_ssoadmin_instances
          "sso:ListInstances",
          //aws_kinesis_stream_consumer
          "kinesis:ListStreamConsumers",
          //aws_identitystore_user
          "identitystore:GetUserId",
          //aws_identitystore_group
          "identitystore:GetGroupId"
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
