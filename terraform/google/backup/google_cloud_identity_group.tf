# resource "google_cloud_identity_group" "pike" {
#   display_name         = "my-identity-group"
#   initial_group_config = "WITH_INITIAL_OWNER"
#
#   parent = "customers/A01b123xz"
#
#   group_key {
#     id = "james.woolfenden@gmail.com"
#   }
#
#   labels = {
#     "cloudidentity.googleapis.com/groups.discussion_forum" = ""
#   }
# }
