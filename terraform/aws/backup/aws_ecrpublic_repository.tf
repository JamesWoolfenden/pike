provider "aws" {
  alias  = "us_east_1"
  region = "us-east-1"
}

resource "aws_ecrpublic_repository" "pike" {
  provider        = aws.us_east_1
  repository_name = "pike"

  catalog_data {
    about_text        = "About Text"
    architectures     = ["ARM"]
    description       = "Description"
    logo_image_blob   = filebase64("Frank_Pike.png")
    operating_systems = ["Linux"]
    usage_text        = "Usage Text"
  }

  tags = {
    pike = "permissions"
  }
}
