resource "aws_wafv2_regex_pattern_set" "example" {
  name        = "example"
  description = "Example regex pattern set"
  scope       = "REGIONAL"

  regular_expression {
    regex_string = "one"
  }

  regular_expression {
    regex_string = "two"
  }

  tags = {
    Tag1 = "Value1"
    Tag2 = "Value2"
  }
}
