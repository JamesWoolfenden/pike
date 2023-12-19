resource "aws_signer_signing_profile_permission" "pike" {
  profile_name = aws_signer_signing_profile.pike.name
  action       = "signer:StartSigningJob"
  principal    = "680235478471"
}
