
resource "aws_vpc_endpoint_policy" "example" {
  vpc_endpoint_id = aws_vpc_endpoint.example.id
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "AllowAll",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : "*"
        },
        "Action" : [
          "dynamodb:*"
        ],
        "Resource" : "*"
      }
    ]
  })
}
