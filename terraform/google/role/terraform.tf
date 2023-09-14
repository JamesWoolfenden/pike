terraform {
  backend "s3" {
    profile        = "personal"
    encrypt        = true
    bucket         = "680235478471-terraform-state"
    key            = "pike-iam-gcp/terraform.tfstate"
    dynamodb_table = "dynamodb-state-lock"
    region         = "eu-west-2"
  }
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.82.0"
    }
  }
}
