resource "aws_ses_configuration_set" "pike" {
  name = "some-configuration-set-test"

  delivery_options {
    tls_policy = "Require"
  }
  reputation_metrics_enabled = false
  sending_enabled            = true
  #   tracking_options {
  #     custom_redirect_domain = "pike-iac.com"
  #   }
}
