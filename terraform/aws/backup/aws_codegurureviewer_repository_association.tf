#resource "aws_kms_key" "example" {
#  provider = aws.eire
#}



resource "aws_codegurureviewer_repository_association" "example" {
  provider = aws.eire
  repository {
    codecommit {
      name = "exampleA"
    }
  }
  #  kms_key_details {
  #    encryption_option = "CUSTOMER_MANAGED_CMK"
  #    kms_key_id        = aws_kms_key.example.key_id
  #  }

  tags = {
    pike    = "permissions"
    another = "one"
  }
}
