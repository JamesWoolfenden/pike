resource "aws_ssmcontacts_contact" "pike" {
  alias        = "jameswoolfenden"
  type         = "PERSONAL"
  display_name = "mr james woolfenden"
  tags = {
    pike    = "permissions"
    another = "permissions"
  }
  depends_on = [aws_ssmincidents_replication_set.pike]
}
