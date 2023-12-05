resource "aws_vpc_ipam_scope" "example" {
  ipam_id     = aws_vpc_ipam.example.id
  description = "Another Scope"
  tags = {
    pike    = "permissions"
    another = "tag"
  }
}
