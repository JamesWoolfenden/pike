resource "aws_vpclattice_target_group_attachment" "pike" {
  target_group_identifier = aws_vpclattice_target_group.pike.id

  target {
    id   = aws_lb.example.arn
    port = 80
  }
}

resource "aws_lb" "example" {
  name                       = "test2-lb-tf"
  internal                   = false
  load_balancer_type         = "application"
  security_groups            = ["sg-05b27cb61c9c46bd2"]
  subnets                    = ["subnet-03fdfb13a135366a7", "subnet-08d97e381dbc80d40"]
  enable_http2               = true
  enable_deletion_protection = false

  access_logs {
    bucket  = aws_s3_bucket.log.bucket
    prefix  = "prefix"
    enabled = true
  }

  tags = {
    Environment = "production"
    //another="tag"
  }
  drop_invalid_header_fields = true
}

resource "aws_s3_bucket" "log" {
  bucket = "anyoldbucketofrubbish"
}
