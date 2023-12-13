resource "aws_waf_sql_injection_match_set" "pike" {
  name = "tf-sql_injection_match_set"

  sql_injection_match_tuples {
    text_transformation = "URL_DECODE"

    field_to_match {
      type = "QUERY_STRING"
    }
  }
}
