resource "aws_fsx_windows_file_system" "pike" {
  storage_capacity    = 64
  subnet_ids          = [data.aws_subnets.example.ids[0]]
  deployment_type     = "SINGLE_AZ_1"
  throughput_capacity = 64
}
