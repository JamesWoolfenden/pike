resource "google_monitoring_snooze" "pike" {
  display_name = "pike test snooze"

  criteria {
    policies = ["projects/pike-477416/alertPolicies/example"]
  }

  interval {
    start_time = "2026-10-19T13:30:00Z"
    end_time   = "2026-10-20T00:00:00Z"
  }
}