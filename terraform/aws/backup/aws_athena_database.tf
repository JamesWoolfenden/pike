resource "aws_athena_database" "pike" {
  name   = "pike2"
  bucket = "athenaresults-mp42nr"
  acl_configuration {
    s3_acl_option = "BUCKET_OWNER_FULL_CONTROL"
  }
  comment = "dave"
  encryption_configuration {
    encryption_option = "SSE_KMS"
    kms_key           = "8d2c4754-28b5-4590-8627-e4ac0792980e"
  }
  expected_bucket_owner = "680235478471"
  force_destroy         = true
  properties = {
    pike = "permissions"
  }

}
