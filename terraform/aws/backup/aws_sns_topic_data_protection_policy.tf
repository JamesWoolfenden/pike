resource "aws_sns_topic_data_protection_policy" "pike" {
  arn = "arn:aws:sns:eu-west-2:680235478471:pike"
  policy = jsonencode(
    {
      "Description" = "Example data protection policy"
      "Name"        = "__example_data_protection_policy"
      "Statement" = [
        {
          "DataDirection" = "Inbound"
          "DataIdentifier" = [
            "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
          ]
          "Operation" = {
            "Deny" = {}
          }
          "Principal" = [
            "*",
          ]
          "Sid" = "__deny_statement_11ba9d96"
        },
      ]
      "Version" = "2021-06-01"
    }
  )
}

output "policy" {
  value = aws_sns_topic_data_protection_policy.pike
}
