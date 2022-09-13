resource "aws_redshift_cluster" "pike" {
  cluster_identifier = "redshift-cluster-1"
  database_name      = "dev"
  //default_iam_role_arn = "arn:aws:iam::680235478471:role/aws-service-role/redshift.amazonaws.com/AWSServiceRoleForRedshift"
  master_username = "awsuser"
  master_password = "Mustbe8characters"
  node_type       = "dc2.large"
  cluster_type    = "single-node"

  vpc_security_group_ids    = ["sg-05b27cb61c9c46bd2"]
  cluster_subnet_group_name = "pike"
  //availability_zone = "eu-west-2b"
  availability_zone_relocation_enabled = false
  preferred_maintenance_window         = "thu:17:00-thu:17:30"
  cluster_parameter_group_name         = "default.redshift-1.0"
  automated_snapshot_retention_period  = 1
  port                                 = 5439
  #  //  cluster_version = ""
  allow_version_upgrade     = true
  apply_immediately         = true
  aqua_configuration_status = "auto"
  number_of_nodes           = 1
  publicly_accessible       = false
  //encrypted = true
  enhanced_vpc_routing = true
  //kms_key_id = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
  elastic_ip                  = ""
  skip_final_snapshot         = true
  snapshot_cluster_identifier = ""
  //owner_account = ""
  iam_roles = ["arn:aws:iam::680235478471:role/aws-service-role/redshift.amazonaws.com/AWSServiceRoleForRedshift"]
  logging {
    log_destination_type = "cloudwatch"
    enable               = true
    log_exports = [
    "connectionlog"]
  }
  maintenance_track_name           = "current"
  manual_snapshot_retention_period = -1
  tags = {
    pike   = "permissions"
    delete = "me"
  }
}
