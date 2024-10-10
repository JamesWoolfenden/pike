resource "aws_quicksight_ingestion" "pike" {
  data_set_id    = aws_quicksight_data_set.pike.data_set_id
  ingestion_id   = "example-id"
  ingestion_type = "FULL_REFRESH"
}
