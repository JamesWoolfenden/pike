resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "iam:GetSAMLProvider",
          //aws_sfn_state_machine
          "states:ListStateMachines",
          "states:DescribeStateMachine",
          "states:ListStateMachineVersions",
          //aws_kinesis_stream
          "kinesis:DescribeStreamSummary",
          //aws_kinesis_firehose_delivery_stream
          "firehose:DescribeDeliveryStream",
          //aws_iam_session_context
          "iam:GetRole",
          //aws_iam_principal_policy_simulation
          "iam:SimulatePrincipalPolicy",
          //aws_iam_access_keys
          "iam:ListAccessKeys"
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
