resource "aws_wafregional_rule_group" "pike" {
  name        = "example"
  metric_name = "example"

  activated_rule {
    action {
      type = "COUNT"
    }

    priority = 50
    rule_id  = aws_wafregional_rule.pike.id
  }
  tags = {
    pike = "permissions"
    #    another="tag"
  }
}
