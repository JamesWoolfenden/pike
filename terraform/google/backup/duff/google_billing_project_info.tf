# resource "google_billing_project_info" "pike" {
#   # deletion_policy = "DELETE"
#   billing_account = data.google_billing_account.pike.id
#
#
#   lifecycle {
#     ignore_changes = [billing_account]
#   }
# }
