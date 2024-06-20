resource "aws_dms_s3_endpoint" "pike" {
  endpoint_id   = "donnedtipi"
  endpoint_type = "target"
  ssl_mode      = "none"

  add_column_name                             = true
  add_trailing_padding_character              = false
  bucket_folder                               = "folder"
  bucket_name                                 = "bucket_name"
  canned_acl_for_objects                      = "private"
  cdc_inserts_and_updates                     = true
  cdc_inserts_only                            = false
  cdc_max_batch_interval                      = 100
  cdc_min_file_size                           = 16
  cdc_path                                    = "cdc/path"
  compression_type                            = "GZIP"
  csv_delimiter                               = ";"
  csv_no_sup_value                            = "x"
  csv_null_value                              = "?"
  csv_row_delimiter                           = "\\r\\n"
  data_format                                 = "parquet"
  data_page_size                              = 1100000
  date_partition_delimiter                    = "UNDERSCORE"
  date_partition_enabled                      = true
  date_partition_sequence                     = "yyyymmddhh"
  date_partition_timezone                     = "Asia/Seoul"
  dict_page_size_limit                        = 1000000
  enable_statistics                           = false
  encoding_type                               = "plain"
  encryption_mode                             = "SSE_S3"
  expected_bucket_owner                       = data.aws_caller_identity.current.account_id
  ignore_header_rows                          = 1
  include_op_for_full_load                    = true
  max_file_size                               = 1000000
  parquet_timestamp_in_millisecond            = true
  parquet_version                             = "parquet-2-0"
  preserve_transactions                       = false
  rfc_4180                                    = false
  row_group_length                            = 11000
  server_side_encryption_kms_key_id           = data.aws_kms_key.example.arn
  service_access_role_arn                     = aws_iam_role.example.arn
  timestamp_column_name                       = "tx_commit_time"
  use_csv_no_sup_value                        = false
  use_task_start_time_for_full_load_timestamp = true
  #   glue_catalog_generation                     = true

  depends_on = [aws_iam_role_policy.example]
}

data "aws_caller_identity" "current" {}

resource "aws_iam_role" "example" {
  assume_role_policy = data.aws_iam_policy_document.dms_assume_role.json
}

resource "aws_iam_role_policy" "example" {
  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ec2:CreateNetworkInterface",
                "ec2:DescribeAvailabilityZones",
                "ec2:DescribeInternetGateways",
                "ec2:DescribeSecurityGroups",
                "ec2:DescribeSubnets",
                "ec2:DescribeVpcs",
                "ec2:DeleteNetworkInterface",
                "ec2:ModifyNetworkInterfaceAttribute"
            ],
            "Resource": "*"
        }
    ]
}
EOF
  role   = aws_iam_role.example.id
}

data "aws_iam_policy_document" "dms_assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["dms.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}


data "aws_kms_key" "example" {
  key_id = "6018184d-f061-446a-9e0c-81aa3c6ba33c"
}
