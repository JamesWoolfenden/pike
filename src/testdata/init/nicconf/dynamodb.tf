module "dynamodb_table" {
  source = "git::https://github.com/terraform-aws-modules/terraform-aws-dynamodb-table.git?ref=b120f36167a93ec3bf454884db808df903babb57"

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
