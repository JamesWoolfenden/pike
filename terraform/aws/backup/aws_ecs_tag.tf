resource "aws_ecs_tag" "pike" {
  resource_arn = aws_batch_compute_environment.pike.ecs_cluster_arn
  key          = "Name"
  value        = "Hello World"
}

resource "aws_batch_compute_environment" "pike" {
  compute_environment_name_prefix = "pike"
  service_role                    = "arn:aws:iam::680235478471:role/aws-service-role/batch.amazonaws.com/AWSServiceRoleForBatch"
  type                            = "MANAGED"

  compute_resources {
    bid_percentage = 0
    desired_vcpus  = 0
    instance_role  = "arn:aws:iam::680235478471:instance-profile/ecsInstanceRole"
    instance_type = [
      "optimal",
    ]
    max_vcpus = 256
    min_vcpus = 0
    security_group_ids = [
      "sg-05b27cb61c9c46bd2",
    ]

    subnets = [
      "subnet-03fdfb13a135366a7",
    ]
    tags = {
      pike = "permissions"
    }
    type = "EC2"
  }


  tags = {
    pike = "permissions"
  }
}
