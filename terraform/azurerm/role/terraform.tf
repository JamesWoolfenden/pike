
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "4.71.0"
    }
    azuread = {
      source  = "hashicorp/azuread"
      version = "3.8.0"
    }
  }
  backend "s3" {
    profile      = "personal"
    encrypt      = true
    bucket       = "680235478471-terraform-state"
    key          = "pike-iam-azure/terraform.tfstate"
    use_lockfile = true
    region       = "eu-west-2"
  }
}
