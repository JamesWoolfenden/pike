resource "aws_kms_replica_external_key" "pike" {
  description             = "Multi-Region replica key"
  deletion_window_in_days = 7
  primary_key_arn         = aws_kms_external_key.pike.arn

  key_material_base64 = "â¡¨Võb¹_^V<95>Ob<94>\\¬µxZî<82>1àíÎ¯^MpªQ^Bd{"
}