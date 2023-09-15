resource "google_compute_security_policy" "pike" {
  name = "pike"

  rule {
    action      = "deny(403)"
    priority    = "2147483647"
    description = "Default rule"

    match {
      versioned_expr = "SRC_IPS_V1"

      config {
        src_ip_ranges = ["*"]
      }
    }
  }

  rule {
    action      = "deny(403)"
    priority    = "2147483646"
    description = "log4j"

    match {
      expr {
        expression = "evaluatePreconfiguredExpr('cve-canary')"
      }
    }
  }
}
