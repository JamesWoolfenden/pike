data "aws_kms_secrets" "pike" {
  secret {
    name    = "master_password"
    payload = "AQECAHgaPa0J8WadplGCqqVAr4HNvDaFSQ+NaiwIBhmm6qDSFwAAAGIwYAYJKoZIhvcNAQcGoFMwUQIBADBMBgkqhkiG9w0BBwEwHgYJYIZIAWUDBAEuMBEEDI+LoLdvYv8l41OhAAIBEIAfx49FFJCLeYrkfMfAw6XlnxP23MmDBdqP8dPp28OoAQ=="

    context = {
      foo = "bar"
    }
  }
}
