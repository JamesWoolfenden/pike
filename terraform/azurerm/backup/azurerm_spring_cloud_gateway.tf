resource "azurerm_spring_cloud_gateway" "pike_gen" {
  name                    = "default"
  spring_cloud_service_id = "pike"

  https_only                    = false
  public_network_access_enabled = true
  instance_count                = 2

  api_metadata {
    description       = "example description"
    documentation_url = "https://www.example.com/docs"
    server_url        = "https://wwww.example.com"
    title             = "example title"
    version           = "1.0"
  }

  cors {
    credentials_allowed = false
    allowed_headers     = ["*"]
    allowed_methods     = ["PUT"]
    allowed_origins     = ["example.com"]
    exposed_headers     = ["x-example-header"]
    max_age_seconds     = 86400
  }

  quota {
    cpu    = "1"
    memory = "2Gi"
  }

  sso {
    client_id  = "example id"
    issuer_uri = "https://www.test.com/issueToken"
    scope      = ["read"]
  }

  local_response_cache_per_instance {
    size         = "100MB"
    time_to_live = "30s"
  }
}
