data "aws_ssmincidents_response_plan" "pike" {
  arn = aws_ssmincidents_response_plan.example.arn
}

#resource "aws_ssmincidents_response_plan" "example" {
#  name = "name"
#
#  incident_template {
#    title  = "title"
#    impact = "3"
#  }
#
#  tags = {
#    key = "value"
#  }
#
#}
