resource "aws_lightsail_static_ip_attachment" "pike" {
  instance_name  = aws_lightsail_instance.pike.name
  static_ip_name = aws_lightsail_static_ip.pike.name
}
