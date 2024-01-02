resource "aws_verifiedaccess_instance_trust_provider_attachment" "pike" {
  verifiedaccess_instance_id       = aws_verifiedaccess_instance.pike.id
  verifiedaccess_trust_provider_id = aws_verifiedaccess_trust_provider.pike.id
}
