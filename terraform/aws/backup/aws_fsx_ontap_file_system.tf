resource "aws_fsx_ontap_file_system" "pike" {
  storage_capacity    = 1024
  subnet_ids          = [data.aws_subnets.example.ids[0], data.aws_subnets.example.ids[2]]
  deployment_type     = "MULTI_AZ_1"
  throughput_capacity = 512
  preferred_subnet_id = data.aws_subnets.example.ids[0]
}
