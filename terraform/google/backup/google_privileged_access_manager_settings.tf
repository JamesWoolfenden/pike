resource "google_privileged_access_manager_settings" "pike" {
  provider = google-beta
  location = "global"
  parent   = "projects/pike-477416"
  service_account_approver_settings {
    enabled = false
  }
  email_notification_settings {
    custom_notification_behavior {
      requester_notifications {
        entitlement_assigned      = "DISABLED"
        grant_activated           = "DISABLED"
        grant_denied              = "ENABLED"
        grant_expired             = "DISABLED"
        grant_ended               = "DISABLED"
        grant_revoked             = "DISABLED"
        grant_externally_modified = "DISABLED"
        grant_activation_failed   = "DISABLED"
      }
    }
  }

}
