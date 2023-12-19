resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "s3:CreateBucket",
          "s3:DeleteBucket",
          "s3:GetAccelerateConfiguration",
          "s3:GetBucketAcl",
          "s3:GetBucketCORS",
          "s3:GetBucketLogging",
          "s3:GetBucketObjectLockConfiguration",
          "s3:GetBucketPolicy",
          "s3:GetBucketRequestPayment",
          "s3:GetBucketTagging",
          "s3:GetBucketVersioning",
          "s3:GetBucketWebsite",
          "s3:GetEncryptionConfiguration",
          "s3:GetLifecycleConfiguration",
          "s3:GetObject",
          "s3:GetObjectAcl",
          "s3:GetReplicationConfiguration",
          "s3:ListBucket",
          "elasticloadbalancing:AddTags",
          "elasticloadbalancing:AttachLoadBalancerToSubnets",
          "elasticloadbalancing:CreateLoadBalancer",
          "elasticloadbalancing:CreateLoadBalancerListeners",
          "elasticloadbalancing:DeleteLoadBalancer",
          "elasticloadbalancing:DescribeLoadBalancerAttributes",
          "elasticloadbalancing:DescribeLoadBalancers",
          "elasticloadbalancing:DescribeTags",
          "elasticloadbalancing:ModifyLoadBalancerAttributes",
          "elasticloadbalancing:RemoveTags",
          "elasticloadbalancing:SetSecurityGroups",

          //aws_vpclattice_access_log_subscription
          "vpc-lattice:GetAccessLogSubscription",
          "vpc-lattice:CreateAccessLogSubscription",
          "vpc-lattice:DeleteAccessLogSubscription",
          "vpc-lattice:UpdateAccessLogSubscription",

          //aws_vpclattice_auth_policy
          "vpc-lattice:PutAuthPolicy",
          "vpc-lattice:DeleteAuthPolicy",
          "vpc-lattice:GetAuthPolicy",

          //aws_vpclattice_listener
          "vpc-lattice:GetListener",
          "vpc-lattice:CreateListener",
          "vpc-lattice:DeleteListener",
          "vpc-lattice:UpdateListener",
          "vpc-lattice:UntagResource",
          "vpc-lattice:TagResource",

          //aws_vpclattice_listener_rule
          "vpc-lattice:GetRule",
          "vpc-lattice:CreateRule",
          "vpc-lattice:DeleteRule",
          "vpc-lattice:UpdateRule",


          //aws_vpclattice_resource_policy
          "vpc-lattice:GetResourcePolicy",
          "vpc-lattice:DeleteResourcePolicy",
          "vpc-lattice:PutResourcePolicy",

          //aws_vpclattice_service
          "vpc-lattice:GetService",
          "vpc-lattice:CreateService",
          "vpc-lattice:DeleteService",
          "vpc-lattice:UpdateService",
          "iam:CreateServiceLinkedRole",
          "vpc-lattice:ListTagsForResource",
          "vpc-lattice:UntagResource",
          "vpc-lattice:TagResource",

          //aws_vpclattice_service_network
          "vpc-lattice:GetServiceNetwork",
          "vpc-lattice:CreateServiceNetwork",
          "vpc-lattice:DeleteServiceNetwork",
          "vpc-lattice:UpdateServiceNetwork",
          "iam:CreateServiceLinkedRole",
          "vpc-lattice:ListTagsForResource",
          "vpc-lattice:UntagResource",
          "vpc-lattice:TagResource",

          //aws_vpclattice_service_network_service_association
          "vpc-lattice:GetServiceNetworkServiceAssociation",
          "vpc-lattice:CreateServiceNetworkServiceAssociation",
          "vpc-lattice:DeleteServiceNetworkServiceAssociation",

          //aws_vpclattice_service_network_vpc_association
          "vpc-lattice:GetServiceNetworkVpcAssociation",
          "vpc-lattice:CreateServiceNetworkVpcAssociation",
          "vpc-lattice:DeleteServiceNetworkVpcAssociation",
          "vpc-lattice:UpdateServiceNetworkVpcAssociation",
          "ec2:DescribeSecurityGroups",

          //aws_vpclattice_target_group
          "vpc-lattice:GetTargetGroup",
          "vpc-lattice:CreateTargetGroup",
          "vpc-lattice:DeleteTargetGroup",
          "vpc-lattice:UpdateTargetGroup",
          "ec2:DescribeVpcs",
          "vpc-lattice:ListTagsForResource",
          "vpc-lattice:UntagResource",
          "vpc-lattice:TagResource",


          //aws_vpclattice_target_group_attachment

          //aws_codebuild_webhook
          "codebuild:CreateWebhook",
          "codebuild:DeleteWebhook",
          "codebuild:UpdateWebhook",

          //aws_location_tracker_association
          "location:AssociateTrackerConsumer",

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
