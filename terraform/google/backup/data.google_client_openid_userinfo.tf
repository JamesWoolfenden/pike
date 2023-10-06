data "google_client_openid_userinfo" "me" {
}

output "my-email" {
  value = data.google_client_openid_userinfo.me.email
}
