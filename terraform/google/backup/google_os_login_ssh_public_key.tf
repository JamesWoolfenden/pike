# data "google_client_openid_userinfo" "me" {
# }

resource "google_os_login_ssh_public_key" "cache" {
  user = "pike-service@pike-477416.iam.gserviceaccount.com"
  key  = file("~/.ssh//id_rsa.pub")
}
