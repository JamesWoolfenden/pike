data "aws_prefix_list" "pike" {
  prefix_list_id = aws_vpc_endpoint.private_s3.prefix_list_id
}
