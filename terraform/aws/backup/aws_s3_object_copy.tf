resource "aws_s3_object_copy" "pike" {
  bucket = "pike-680235478471"
  key    = "terraform.tfstate"
  source = "680235478471-terraform-state/pike-aws/terraform.tfstate"
  override_provider {
    default_tags {
      tags = {}
    }
  }
}
