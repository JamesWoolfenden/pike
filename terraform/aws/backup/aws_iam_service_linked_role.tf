resource "aws_iam_service_linked_role" "elasticbeanstalk" {
  aws_service_name = "autoscaling.amazonaws.com"
  custom_suffix    = "jgw1"
  tags = {
    custom = "tag"
  }
}
