resource "aws_ecr_repository" "example" {
  name                 = "pike"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
  encryption_configuration {
    encryption_type = "KMS"
    kms_key         = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
  }
  tags = {
    pike = "permissions"
    #    another="tag"
  }
}
