resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //data.aws_connect_bot_association.pike
          "connect:ListBots",
          // data.aws_connect_contact_flow.pike
          "connect:ListContactFlows",
          //data.aws_connect_contact_flow_module.pike
          "connect:ListContactFlowModules",
          // data.aws_connect_hours_of_operation
          "connect:ListHoursOfOperations",
          //data.aws_connect_routing_profile.pike
          "connect:ListRoutingProfiles",
          //data.aws_connect_instance
          "connect:DescribeInstance",
          //data.aws_connect_instance_storage_config.pike
          "connect:DescribeInstanceStorageConfig",
          // data.aws_connect_lambda_function_association.pike
          "connect:ListLambdaFunctions",
          // data.aws_connect_prompt.pike
          "connect:ListPrompts",
          //data.aws_connect_queue.
          "connect:ListQueues",
          //data.aws_connect_quick_connect.pike
          "connect:ListQuickConnects",
          //data.aws_connect_routing_profile.pike
          "connect:ListRoutingProfiles",
          //data.aws_connect_security_profile.pike
          "connect:ListSecurityProfiles",
          //data.aws_connect_user_hierarchy_structure.pike
          "connect:DescribeUserHierarchyStructure",
          //data.aws_connect_user_hierarchy_group.pike
          "connect:ListUserHierarchyGroups"
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
