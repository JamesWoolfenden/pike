resource "aws_inspector2_organization_configuration" "pike" {
  auto_enable {
    ec2 = false
    ecr = false
  }
}
