resource "aws_neptune_global_cluster" "pike" {
  global_cluster_identifier = "global-test"
  engine                    = "neptune"
  engine_version            = "1.2.0.0"
}
