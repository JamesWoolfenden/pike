# resource "google_cloud_identity_group_membership" "pike" {
#
#     //group    = google_cloud_identity_group.pike.id
#     group      = "pike"
#
#     preferred_member_key {
#       id = "cloud_identity_user@example.com"
#     }
#
#     roles {
#       name = "MEMBER"
#     }
#
#     roles {
#       name = "MANAGER"
#     }
#   }
