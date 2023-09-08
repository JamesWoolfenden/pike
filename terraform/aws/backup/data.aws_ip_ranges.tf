data "aws_ip_ranges" "pike" {
  services = []
}


output "ranges" {
  value = data.aws_ip_ranges.pike
}
