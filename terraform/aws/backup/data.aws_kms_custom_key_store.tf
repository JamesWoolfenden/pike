data "aws_kms_custom_key_store" "pike" {
  custom_key_store_name = "my_cloudhsm"
}
