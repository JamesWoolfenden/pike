resource "aws_ce_cost_category" "pike" {
  name         = "NAME"
  rule_version = "CostCategoryExpression.v1"
  rule {
    value = "production"
    type  = "REGULAR"
    rule {
      dimension {
        key           = "LINKED_ACCOUNT_NAME"
        values        = ["-prod"]
        match_options = ["ENDS_WITH"]
      }
    }
  }
  rule {
    value = "staging"
    type  = "REGULAR"
    rule {
      dimension {
        key           = "LINKED_ACCOUNT_NAME"
        values        = ["-stg"]
        match_options = ["ENDS_WITH"]
      }
    }
  }
  rule {
    value = "testing"
    type  = "REGULAR"
    rule {
      dimension {
        key           = "LINKED_ACCOUNT_NAME"
        values        = ["-dev"]
        match_options = ["ENDS_WITH"]
      }
    }
  }
}
