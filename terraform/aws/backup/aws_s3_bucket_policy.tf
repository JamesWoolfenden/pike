resource "aws_s3_bucket_policy" "example" {
  bucket = "testbucketforlbjgw"
  policy = <<POLICY
  {
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::652711504416:root"
            },
            "Action": "s3:PutObject",
            "Resource": "arn:aws:s3:::testbucketforlbjgw/prefix/AWSLogs/680235478471/*"
        }
    ]
}
POLICY
}
