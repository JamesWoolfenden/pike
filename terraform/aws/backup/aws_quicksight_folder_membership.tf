resource "aws_quicksight_folder_membership" "pike" {
  folder_id   = aws_quicksight_folder.pike.folder_id
  member_type = "DATASET"
  member_id   = aws_quicksight_data_set.pike.data_set_id
}
