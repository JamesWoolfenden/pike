resource "aws_lb" "test" {
  name               = "test-lb-tf"
  internal           = false
  load_balancer_type = "application"
  security_groups    = ["sg-05b27cb61c9c46bd2"]
  subnets            = ["subnet-03fdfb13a135366a7","subnet-08d97e381dbc80d40"]

  enable_deletion_protection = false

  access_logs {
    bucket  = "testbucketforlbjgw"
    prefix  = "prefix"
    enabled = true
  }

  tags = {
    Environment = "production"
    another="tag"
  }
}

