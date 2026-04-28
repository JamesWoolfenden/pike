resource "azurerm_role_management_policy" "pike_gen" {
  scope              = "pike"
  role_definition_id = "pike"

  active_assignment_rules {
    expire_after = "P365D"
  }

  eligible_assignment_rules {
    expiration_required = false
  }

  activation_rules {
    maximum_duration = "PT1H"
    require_approval = true
    approval_stage {
      primary_approver {
        object_id = "pike"
        type      = "Group"
      }
    }
  }

  notification_rules {
    eligible_assignments {
      approver_notifications {
        notification_level    = "Critical"
        default_recipients    = false
        additional_recipients = ["someone@example.com"]
      }
    }
    eligible_activations {
      assignee_notifications {
        notification_level    = "All"
        default_recipients    = true
        additional_recipients = ["someone.else@example.com"]
      }
    }
  }
}
