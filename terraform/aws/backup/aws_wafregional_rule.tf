resource "aws_wafregional_rule" "pike" {
  name        = "tfWAFRule"
  metric_name = "tfWAFRule"

  predicate {
    data_id = aws_wafregional_ipset.pike.id
    negated = false
    type    = "IPMatch"
  }
  tags = {
    pike = "permission"
    #    another="tag"
  }
}
