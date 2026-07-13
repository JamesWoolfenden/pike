terraform {
  backend "s3" {
    profile      = "personal"
    encrypt      = true
    bucket       = "680235478471-terraform-state"
    key          = "pike-iam/terraform.tfstate"
    use_lockfile = true
    region       = "eu-west-2"
  }
}
