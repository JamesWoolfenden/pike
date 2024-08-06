resource "aws_chimesdkvoice_voice_profile_domain" "pike" {
  provider = aws.central
  name     = "ExampleVoiceProfileDomain"
  server_side_encryption_configuration {
    kms_key_arn = aws_kms_key.example.arn
  }
  description = "My Voice Profile Domain"
  tags = {
    key1 = "value1"
  }
}

resource "aws_kms_key" "example" {

  enable_key_rotation                = true
  is_enabled                         = true
  key_usage                          = "ENCRYPT_DECRYPT"
  description                        = "myfavouritekey"
  bypass_policy_lockout_safety_check = false
  deletion_window_in_days            = 7
  multi_region                       = false
  policy                             = <<POLICY
{
    "Version": "2012-10-17",
    "Id": "key-default-1",
    "Statement": [
        {
            "Sid": "Enable IAM User Permissions",
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::680235478471:root"
            },
            "Action": "kms:*",
            "Resource": "*"
        }
    ]
}
POLICY
}
