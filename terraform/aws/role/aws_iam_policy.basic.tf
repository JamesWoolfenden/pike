resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_transfer_connector
          "transfer:DescribeConnector",
          //aws_ec2_transit_gateway_peering_attachments
          "ec2:DescribeTransitGatewayPeeringAttachments",
          //aws_appstream_image
          "appstream:DescribeImages",
          //aws_cloudfront_origin_access_control
          "cloudfront:GetOriginAccessControl",
          //aws_cognito_user_pool
          "cognito-idp:DescribeUserPool",
          //aws_timestreamwrite_table, aws_timestreamwrite_database
          "timestream:DescribeEndpoints",
          //aws_timestreamwrite_table
          "timestream:DescribeTable",
          //aws_timestreamwrite_database
          "timestream:DescribeDatabase"
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
