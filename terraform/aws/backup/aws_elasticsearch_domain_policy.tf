resource "aws_elasticsearch_domain_policy" "pike" {
  domain_name = aws_elasticsearch_domain.pike.domain_name

  access_policies = <<POLICIES
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "es:*",
            "Principal": "*",
            "Effect": "Allow",
            "Condition": {
                "IpAddress": {"aws:SourceIp": "127.0.0.1/32"}
            },
            "Resource": "${aws_elasticsearch_domain.pike.arn}/*"
        }
    ]
}
POLICIES
}
