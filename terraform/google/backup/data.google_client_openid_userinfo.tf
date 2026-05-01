data "google_client_openid_userinfo" "me" {
}

output "google_client_openid_userinfo" {
  value = data.google_client_openid_userinfo.me.email
}
