resource "aws_waf_rule" "pike" {
  depends_on  = [aws_waf_ipset.pike]
  name        = "tfWAFRule"
  metric_name = "tfWAFRule"

  predicates {
    data_id = aws_waf_ipset.pike.id
    negated = false
    type    = "IPMatch"
  }
}
