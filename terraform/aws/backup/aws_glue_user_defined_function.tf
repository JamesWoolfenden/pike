resource "aws_glue_user_defined_function" "pike" {
  name = "pike"

  catalog_id    = "680235478471"
  database_name = "pike"
  class_name    = "class"
  owner_name    = "owner"
  owner_type    = "GROUP"

  resource_uris {
    resource_type = "ARCHIVE"
    uri           = "uri"
  }

}
