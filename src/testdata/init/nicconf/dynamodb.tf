module "dynamodb_table" {
  source  = "terraform-aws-modules/dynamodb-table/aws"
  version = "v3.1.2"

  name      = random_pet.this.id
  hash_key  = "id"
  range_key = "name"

  attributes = [
    {
      name = "id"
      type = "S"
    },
    {
      name = "name"
      type = "S"
    },
  ]
}
