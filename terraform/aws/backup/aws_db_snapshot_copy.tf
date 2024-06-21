resource "aws_db_snapshot_copy" "pike" {
  source_db_snapshot_identifier = aws_db_snapshot.pike.db_snapshot_arn
  target_db_snapshot_identifier = "testsnapshot1234-copy"
}
