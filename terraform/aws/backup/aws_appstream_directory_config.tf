resource "aws_appstream_directory_config" "pike" {
  directory_name                          = "NAME OF DIRECTORY"
  organizational_unit_distinguished_names = ["DISTINGUISHED NAME"]

  service_account_credentials {
    account_name     = "NAME OF ACCOUNT"
    account_password = "PASSWORD OF ACCOUNT"
  }
}
