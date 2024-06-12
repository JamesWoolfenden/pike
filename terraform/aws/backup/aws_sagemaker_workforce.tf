resource "aws_sagemaker_workforce" "pike" {
  workforce_name = "pike"
  oidc_config {
    authorization_endpoint = "https://example.com"
    client_id              = "example"
    client_secret          = "example"
    issuer                 = "https://example.com"
    jwks_uri               = "https://example.com"
    logout_endpoint        = "https://example.com"
    token_endpoint         = "https://example.com"
    user_info_endpoint     = "https://example.com"
  }
}
