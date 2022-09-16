#resource "aws_db_snapshot_copy" "pike" {
#    source_db_snapshot_identifier = aws_db_snapshot.rds.db_snapshot_arn
#    target_db_snapshot_identifier = "testsnapshot1234-copy"
#  copy_tags=true
#  destination_region=""
#  kms_key_id= ""
#          option_group_name=""
#  presigned_ulr=""
#  target_custom_availability_zone=""
#  tags={
#    pike="permissions"
#    delete="me"
#  }
#
#}
