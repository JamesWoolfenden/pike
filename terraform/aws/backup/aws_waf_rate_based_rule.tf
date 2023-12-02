resource "aws_waf_rate_based_rule" "pike" {
  depends_on  = [aws_waf_ipset.pike]
  name        = "tfWAFRule"
  metric_name = "tfWAFRule"

  rate_key   = "IP"
  rate_limit = 100

  predicates {
    data_id = aws_waf_ipset.pike.id
    negated = false
    type    = "IPMatch"
  }
}
