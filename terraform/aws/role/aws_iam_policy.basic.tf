resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [
          //aws_appconfig_application
          "appconfig:ListApplications",

          //aws_odb_cloud_autonomous_vm_clusters
          "odb:ListCloudAutonomousVmClusters",

          //aws_odb_cloud_autonomous_vm_cluster
          "odb:GetCloudAutonomousVmCluster",

          //aws_odb_cloud_exadata_infrastructures
          "odb:ListCloudExadataInfrastructures",

          //aws_odb_cloud_exadata_infrastructure
          "odb:GetCloudExadataInfrastructure",

          //aws_odb_cloud_vm_clusters
          "odb:ListCloudVmClusters",

          //aws_odb_cloud_vm_cluster
          "odb:GetCloudVmCluster",

          //aws_odb_db_nodes
          "odb:ListDbNodes",

          //aws_odb_db_node
          "odb:GetDbNode",

          //aws_odb_db_servers
          "odb:ListDbServers",

          //aws_odb_db_server
          "odb:GetDbServer",

          //aws_odb_db_system_shapes
          "odb:ListDbSystemShapes",

          //aws_odb_gi_versions
          "odb:ListGiVersions",

          //aws_odb_network_peering_connection
          "odb:GetOdbPeeringConnection",

          //aws_odb_network_peering_connections
          "odb:ListOdbPeeringConnections",

          //aws_odb_networks
          "odb:ListOdbNetworks",

          //aws_odb_network
          "odb:GetOdbNetwork",

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
