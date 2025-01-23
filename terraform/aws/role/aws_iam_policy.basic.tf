resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [
          # aws_appmesh_virtual_gateway
          "appmesh:DescribeVirtualGateway",
          "appmesh:CreateVirtualGateway",
          "appmesh:DeleteVirtualGateway",
          "appmesh:UpdateVirtualGateway",
          "appmesh:TagResource",
          "appmesh:UntagResource",

          # aws_appmesh_mesh
          "appmesh:DescribeMesh",
          "appmesh:CreateMesh",
          "appmesh:DeleteMesh",
          "appmesh:UpdateMesh",
          "appmesh:ListTagsForResource",

          # aws_appmesh_route
          "appmesh:DescribeRoute",
          "appmesh:CreateRoute",
          "appmesh:DeleteRoute",
          "appmesh:UpdateRoute",

          # aws_appmesh_gateway_route
          "appmesh:DescribeGatewayRoute",
          "appmesh:CreateGatewayRoute",
          "appmesh:DeleteGatewayRoute",
          "appmesh:UpdateGatewayRoute",

          # aws_appmesh_virtual_node
          "appmesh:DescribeVirtualNode",
          "appmesh:CreateVirtualNode",
          "appmesh:DeleteVirtualNode",
          "appstream:DescribeUsers",
          "appstream:CreateUser",
          "appstream:DeleteUser",
          "appstream:DescribeStacks",
          "appstream:CreateStack",
          "appstream:DeleteStack",
          "appstream:UpdateStack",
          "appstream:DescribeFleets",
          "appstream:CreateFleet",
          "appstream:DeleteFleet",
          "appstream:UpdateFleet",
          "appstream:ListTagsForResource",
          "appstream:CreateDirectoryConfig",
          "appstream:DeleteDirectoryConfig",
          "appstream:DeleteDirectoryConfig",
          "appstream:DescribeDirectoryConfigs",
          "appstream:UpdateDirectoryConfig",
          "appstream:DescribeUserStackAssociations",
          "appstream:BatchAssociateUserStack",
          "appstream:BatchDisassociateUserStack",
        ],
        "Resource" : [
          "*"
        ]
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
