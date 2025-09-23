data "aws_media_convert_queue" "pike" {
  id = "pike"
}

output "aws_media_convert_queue" {
  value = data.aws_media_convert_queue.pike
}
