data "aws_caller_identity" "current" {}

data "aws_regions" "example" {}

resource "aws_ecr_replication_configuration" "pike" {
  replication_configuration {
    rule {
      destination {
        region      = tolist(data.aws_regions.example.names)[0]
        registry_id = data.aws_caller_identity.current.account_id
      }
    }
  }
}
