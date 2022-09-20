resource "aws_media_convert_queue" "pike" {
  name        = "pike"
  description = "Pike is for permissioms"
  tags = {
    pike   = "permissions"
    delete = "me"
  }
}
