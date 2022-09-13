resource "aws_redshift_authentication_profile" "pike" {
  authentication_profile_name = "example"
  authentication_profile_content = jsonencode(
    {
      AllowDBUserOverride = "2"
      Client_ID           = "ExampleClientID"
      App_ID              = "example"
    }
  )
}
