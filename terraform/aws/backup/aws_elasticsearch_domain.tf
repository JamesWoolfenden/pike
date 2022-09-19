resource "aws_elasticsearch_domain" "pike" {
  domain_name           = "pike"
  elasticsearch_version = "7.10"

  ebs_options {
    ebs_enabled = true
    volume_size = "10"
  }
  cluster_config {
    instance_type = "t2.small.elasticsearch"
  }

  tags = {
    pike   = "permission"
    delete = "me"
  }
}
