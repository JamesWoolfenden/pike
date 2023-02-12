data "aws_cloudcontrolapi_resource" "pike" {
  identifier = "pike"
  type_name  = "AWS::ECS::Cluster"
}
