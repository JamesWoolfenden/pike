
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "3.26.0"
    }
    azuread = {
      source  = "hashicorp/azuread"
      version = "2.29.0"
    }
  }
  backend "s3" {
    profile        = "personal"
    encrypt        = true
    bucket         = "680235478471-terraform-state"
    key            = "pike-iam-azure/terraform.tfstate"
    dynamodb_table = "dynamodb-state-lock"
    region         = "eu-west-2"
  }
}
