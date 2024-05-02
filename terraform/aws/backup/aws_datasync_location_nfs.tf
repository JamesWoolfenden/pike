resource "aws_datasync_location_nfs" "source_pike" {
  server_hostname = "nfs.example.com"
  subdirectory    = "/exported/path"

  on_prem_config {
    agent_arns = [aws_datasync_agent.pike.arn]
  }

  tags = {
    pike = "permissions"
  }
}

resource "aws_datasync_location_nfs" "destination_pike" {
  server_hostname = "nfs.example.com"
  subdirectory    = "/exported/path"

  on_prem_config {
    agent_arns = [aws_datasync_agent.pike.arn]
  }

  tags = {
    pike = "permissions"
  }
}
