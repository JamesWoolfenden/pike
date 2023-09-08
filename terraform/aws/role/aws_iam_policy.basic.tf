resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_ssmincidents_replication_set
          "ssm-incidents:ListReplicationSets",
          //aws_resourcegroupstaggingapi_resources
          "tag:GetResources",
          //aws_ram_resource_share
          "ram:GetResourceShares",
          //aws_mskconnect_worker_configuration
          "kafkaconnect:ListWorkerConfigurations",
          //aws_mskconnect_custom_plugin
          "kafkaconnect:ListCustomPlugins",
          //aws_mskconnect_connector
          "kafkaconnect:ListConnectors",
          //aws_msk_vpc_connection
          "kafka:DescribeVpcConnection",
          //aws_serverlessapplicationrepository_application
          "serverlessrepo:GetApplication",
          //aws_signer_signing_profile
          "Signer:GetSigningProfile",
          //aws_signer_signing_job
          "Signer:DescribeSigningJob"

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
