resource "aws_datasync_location_smb" "pike" {
  server_hostname = "smb.example.com"
  subdirectory    = "/exported/path"

  user     = "Guest"
  password = "ANotGreatPassword"

  agent_arns = [aws_datasync_agent.pike.arn]

  tags = {
    pike = "permissions"
  }
}
