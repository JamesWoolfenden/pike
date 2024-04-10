data "aws_devopsguru_notification_channel" "pike" {
  id = "channel-1234"
}

output "aws_devopsguru_notification_channel" {
  value = data.aws_devopsguru_notification_channel.pike
}
