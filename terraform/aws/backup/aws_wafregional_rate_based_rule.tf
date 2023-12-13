resource "aws_wafregional_rate_based_rule" "pike" {
  depends_on  = [aws_wafregional_ipset.pike]
  name        = "tfWAFRule"
  metric_name = "tfWAFRule"

  rate_key   = "IP"
  rate_limit = 100

  predicate {
    data_id = aws_wafregional_ipset.pike.id
    negated = false
    type    = "IPMatch"
  }

  tags = {
    pike = "permission"
  }
}
