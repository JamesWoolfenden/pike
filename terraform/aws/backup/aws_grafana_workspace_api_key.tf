resource "aws_grafana_workspace_api_key" "pike" {
  workspace_id    = "g-1234567890"
  key_role        = "ADMIN"
  seconds_to_live = 5
  key_name        = "pike"
}
