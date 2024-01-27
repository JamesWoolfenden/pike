resource "google_access_context_manager_access_levels" "access-levels" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"
  access_levels {
    name  = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/accessLevels/chromeos_no_lock"
    title = "chromeos_no_lock"
    basic {
      conditions {
        device_policy {
          require_screen_lock = true
          os_constraints {
            os_type = "DESKTOP_CHROME_OS"
          }
        }
        regions = [
          "CH",
          "IT",
          "US",
        ]
      }
    }
  }

  access_levels {
    name  = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/accessLevels/mac_no_lock"
    title = "mac_no_lock"
    basic {
      conditions {
        device_policy {
          require_screen_lock = true
          os_constraints {
            os_type = "DESKTOP_MAC"
          }
        }
        regions = [
          "CH",
          "IT",
          "US",
        ]
      }
    }
  }
}
