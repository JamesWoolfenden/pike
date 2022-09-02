resource "aws_codeartifact_domain_permissions_policy" "example" {
  domain          = aws_codeartifact_domain.pike.domain
  policy_document = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "codeartifact:CreateRepository",
            "Effect": "Allow",
            "Principal": "*",
            "Resource": "${aws_codeartifact_domain.pike.arn}"
        }
    ]
}
EOF
}
