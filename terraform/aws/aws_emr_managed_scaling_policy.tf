resource "aws_emr_managed_scaling_policy" "pike" {
  cluster_id = aws_emr_cluster.pike.id
  compute_limits {
    unit_type                       = "Instances"
    minimum_capacity_units          = 2
    maximum_capacity_units          = 10
    maximum_ondemand_capacity_units = 2
    maximum_core_capacity_units     = 10
  }
}

resource "aws_emr_cluster" "pike" {
  name          = "pike"
  release_label = "emr-5.30.0"
  service_role  = ""
}
