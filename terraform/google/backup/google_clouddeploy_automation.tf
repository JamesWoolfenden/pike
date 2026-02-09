resource "google_clouddeploy_automation" "pike" {
  name              = "cd-automation"
  project           = google_clouddeploy_delivery_pipeline.pike.project
  location          = google_clouddeploy_delivery_pipeline.pike.location
  delivery_pipeline = google_clouddeploy_delivery_pipeline.pike.name
  service_account   = "test-964@pike-477416.iam.gserviceaccount.com"
  selector {
    targets {
      id = "*"
    }
  }
  rules {
    promote_release_rule {
      id = "promote-release"
    }
  }
  rules {
    advance_rollout_rule {
      id = "advance-rollout"
    }
  }
  rules {
    repair_rollout_rule {
      id = "repair-rollout"
      repair_phases {
        retry {
          attempts = "1"
        }
      }
      repair_phases {
        rollback {}
      }
    }
  }
  rules {
    timed_promote_release_rule {
      id        = "timed-promote-release"
      schedule  = "0 9 * * 1"
      time_zone = "America/New_York"
    }
  }
}
