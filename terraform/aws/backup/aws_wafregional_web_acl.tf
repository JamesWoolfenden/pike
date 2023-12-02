resource "aws_wafregional_web_acl" "pike" {
  name        = "tfWebACL"
  metric_name = "tfWebACL"

  default_action {
    type = "ALLOW"
  }

  rule {
    action {
      type = "BLOCK"
    }

    priority = 1
    rule_id  = aws_wafregional_rule.pike.id
    type     = "REGULAR"
  }
  tags = {
    pike = "permissions"
    #    another="tag"
  }
}
