resource "google_dialogflow_cx_webhook" "pike" {
  parent       = google_dialogflow_cx_agent.pike.id
  display_name = "MyFlow"
  generic_web_service {
    allowed_ca_certs = ["BQA="]
    uri              = "https://example.com"
    request_headers  = { "example-key" : "example-value" }
    webhook_type     = "STANDARD"
    oauth_config {
      client_id                        = "example-client-id"
      secret_version_for_client_secret = "projects/example-proj/secrets/example-secret/versions/example-version"
      token_endpoint                   = "https://example.com"
      scopes                           = ["example-scope"]
    }
    service_agent_auth                   = "NONE"
    secret_version_for_username_password = "projects/example-proj/secrets/example-secret/versions/example-version"
    secret_versions_for_request_headers {
      key            = "example-key-1"
      secret_version = "projects/example-proj/secrets/example-secret/versions/example-version"
    }
    secret_versions_for_request_headers {
      key            = "example-key-2"
      secret_version = "projects/example-proj/secrets/example-secret/versions/example-version-2"
    }
  }
}
