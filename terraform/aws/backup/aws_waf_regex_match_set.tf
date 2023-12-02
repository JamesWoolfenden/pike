resource "aws_waf_regex_match_set" "pike" {
  name = "example"

  regex_match_tuple {
    field_to_match {
      data = "User-Agent"
      type = "HEADER"
    }

    regex_pattern_set_id = aws_waf_regex_pattern_set.pike.id
    text_transformation  = "NONE"
  }
}
