#data aws_organizations_organization "example" {}
data "aws_route_tables" "example" {}

output "routes" {
  value = data.aws_route_tables.example
}
#data aws_s3_bucket_object "example" {}
#data aws_ssoadmin_instances "example" {}
#data aws_workspaces_bundle "example" {}
