

resource "aws_dynamodb_global_table" "pike" {

  provider = aws.alt

  name = "pike"

  replica {
    region_name = "eu-west-1"
  }

  replica {
    region_name = "eu-west-2"
  }
}
