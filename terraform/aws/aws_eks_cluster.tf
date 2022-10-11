resource "aws_eks_cluster" "pike" {
  name     = "pike2"
  role_arn = "arn:aws:iam::680235478471:role/eks_cluster"
  vpc_config {
    subnet_ids = ["subnet-09ff91b5b0adb1fd4", "subnet-05e87623a2a5c41f0", "subnet-043bb893867355740"]
  }

  #  encryption_config {
  #    resources = ["secrets"]
  #    provider {
  #      key_arn="arn:aws:kms:eu-west-2:680235478471:key/03a3077b-1b63-4f42-98a1-20ea7a2fabba"
  #    }
  #  }
  enabled_cluster_log_types = ["api", "audit", "authenticator", "controllerManager", "scheduler"]

  tags = {
    pike = "permissions"
  }
}
