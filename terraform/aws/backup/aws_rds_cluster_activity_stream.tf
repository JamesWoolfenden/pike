resource "aws_rds_cluster_activity_stream" "pike" {
  resource_arn                        = "arn:aws:rds:eu-west-2:680235478471:cluster:aurora-cluster-pike"
  mode                                = "async"
  kms_key_id                          = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
  engine_native_audit_fields_included = true
}
