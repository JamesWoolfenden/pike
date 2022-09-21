resource "aws_imagebuilder_component" "pike" {
  data = yamlencode({
    phases = [{
      name = "build"
      steps = [{
        action = "ExecuteBash"
        inputs = {
          commands = ["echo 'hello world'"]
        }
        name      = "example"
        onFailure = "Continue"
      }]
    }]
    schemaVersion = 1.0
  })
  name        = "pike-hello"
  platform    = "Linux"
  version     = "1.0.0"
  kms_key_id  = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
  description = "Pike"
  tags = {
    pike   = "permissions"
    delete = "me"
  }
}
