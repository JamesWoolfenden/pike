resource "aws_datasync_location_object_storage" "pike" {
  agent_arns      = [aws_datasync_agent.pike.arn]
  server_hostname = "example"
  bucket_name     = "example"

  tags = {
    pike = "permissions"
  }
}
