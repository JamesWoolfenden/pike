resource "aws_inspector2_enabler" "pike" {
  account_ids    = ["680235478471"]
  resource_types = ["EC2"]
}
