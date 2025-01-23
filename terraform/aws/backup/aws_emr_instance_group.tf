resource "aws_emr_instance_group" "pike" {
  cluster_id     = aws_emr_cluster.tf-test-cluster.id
  instance_count = 1
  instance_type  = "m5.xlarge"
  name           = "my little instance group"
}
