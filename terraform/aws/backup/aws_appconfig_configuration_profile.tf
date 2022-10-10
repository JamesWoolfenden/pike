resource "aws_appconfig_configuration_profile" "pike" {
  application_id = aws_appconfig_application.pike.id
  description    = "Example Configuration Profile"
  name           = "pike"
  location_uri   = "hosted"

  validator {
    content = "arn:aws:lambda:eu-west-2:680235478471:function:message-processor"
    type    = "LAMBDA"
  }

  tags = {
    pike = "permissions"
  }
}

resource "aws_appconfig_application" "pike" {
  name        = "pike-tf"
  description = "Pike is permissions"

  tags = {
    pike   = "Permissions"
    delete = "me"
  }
}
