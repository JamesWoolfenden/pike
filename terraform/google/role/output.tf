output "service_account" {
  value = google_service_account.pike
}

output "iam_binding" {
  value = google_project_iam_binding.pike
}

output "custom_role" {
  value = google_project_iam_custom_role.terraform_pike
}

output "state_bucket" {
  value = google_storage_bucket.default
}

output "services" {
  value = google_project_service.project
}
