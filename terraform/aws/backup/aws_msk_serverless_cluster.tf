resource "aws_msk_serverless_cluster" "pike" {
  cluster_name = "pike"

  vpc_config {
    subnet_ids         = ["subnet-0d99e0f4e4f1ff84f", "subnet-02fbf2ee10897034a", "subnet-04338b6369d8288a5"]
    security_group_ids = ["sg-0403ba8bc0ba7876b"]
  }

  client_authentication {
    sasl {
      iam {
        enabled = true
      }
    }
  }
}
