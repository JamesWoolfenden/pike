data "aws_elastic_beanstalk_solution_stack" "example" {
  most_recent = true
  name_regex  = "^64bit Amazon Linux 2018.03 v2.9.15 running Python"
}
