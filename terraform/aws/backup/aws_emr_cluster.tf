resource "aws_emr_cluster" "pike" {
  name          = "pike"
  release_label = "emr-4.6.0"
  applications  = ["Spark"]

  ec2_attributes {
    #    subnet_id                         = aws_subnet.main.id
    #    emr_managed_master_security_group = aws_security_group.allow_access.id
    #    emr_managed_slave_security_group  = aws_security_group.allow_access.id
    instance_profile = "arn:aws:iam::680235478471:instance-profile/emr_profile"
  }

  master_instance_group {
    //    instance_type = "m5.xlarge"
    instance_type = "m4.large"
  }

  core_instance_group {
    instance_count = 1
    //instance_type  = "m5.xlarge"
    instance_type = "m4.large"
  }

  tags = {
    role     = "rolename"
    dns_zone = "env_zone"
    env      = "env"
    pike     = "permissions"
  }

  bootstrap_action {
    path = "s3://elasticmapreduce/bootstrap-actions/run-if"
    name = "runif"
    args = ["instance.isMaster=true", "echo running on master node"]
  }
  configurations_json = <<EOF
  [
    {
      "Classification": "hadoop-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
          }
        }
      ],
      "Properties": {}
    },
    {
      "Classification": "spark-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.9.0"
          }
        }
      ],
      "Properties": {}
    }
  ]
EOF
  service_role        = "arn:aws:iam::680235478471:role/emr_service"

}
