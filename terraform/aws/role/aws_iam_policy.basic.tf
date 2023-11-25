resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_opensearchserverless_lifecycle_policy
          "aoss:BatchGetLifecyclePolicy",
          //aws_apigatewayv2_vpc_link
          "apigateway:GET",
          //aws_athena_named_query
          "athena:ListNamedQueries",
          //aws_iot_registration_code
          "iot:GetRegistrationCode",

          "bedrock:ListFoundationModels",
          "bedrock:GetFoundationModel",
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
