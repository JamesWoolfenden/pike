resource "aws_iam_user_ssh_key" "pike" {
  encoding   = "SSH"
  public_key = file("testdata.pub")
  username   = "basic"
}
