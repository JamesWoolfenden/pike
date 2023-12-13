resource "aws_waf_web_acl" "pike" {
  depends_on = [
    aws_waf_ipset.pike,
    aws_waf_rule.pike,
  ]
  name        = "tfWebACL"
  metric_name = "tfWebACL"

  default_action {
    type = "ALLOW"
  }

  rules {
    action {
      type = "BLOCK"
    }

    priority = 1
    rule_id  = aws_waf_rule.pike.id
    type     = "REGULAR"
  }
}
