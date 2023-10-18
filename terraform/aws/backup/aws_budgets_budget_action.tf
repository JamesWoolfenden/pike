resource "aws_budgets_budget_action" "pike" {
  budget_name        = "budget-monthly"
  action_type        = "APPLY_IAM_POLICY"
  approval_model     = "AUTOMATIC"
  notification_type  = "ACTUAL"
  execution_role_arn = "arn:aws:iam::680235478471:role/apigatewaytoclouwatch"

  action_threshold {
    action_threshold_type  = "ABSOLUTE_VALUE"
    action_threshold_value = 100
  }

  definition {
    iam_action_definition {
      policy_arn = "arn:aws:iam::680235478471:policy/aaa"
      roles      = ["arn:aws:iam::680235478471:role/apigatewaytoclouwatch"]
    }
  }

  subscriber {
    address           = "example@example.example"
    subscription_type = "EMAIL"
  }
}

output "actions" {
  value = aws_budgets_budget_action.pike
}
