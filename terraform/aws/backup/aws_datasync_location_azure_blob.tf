resource "aws_datasync_location_azure_blob" "pike" {
  agent_arns          = [aws_datasync_agent.pike.arn]
  authentication_type = "SAS"
  container_url       = "https://myaccount.blob.core.windows.net/mycontainer"

  sas_configuration {
    token = "sp=r&st=2023-12-20T14:54:52Z&se=2023-12-20T22:54:52Z&spr=https&sv=2021-06-08&sr=c&sig=aBBKDWQvyuVcTPH9EBp%2FXTI9E%2F%2Fmq171%2BZU178wcwqU%3D"
  }

  tags = {
    pike = "permissions"
  }
}
