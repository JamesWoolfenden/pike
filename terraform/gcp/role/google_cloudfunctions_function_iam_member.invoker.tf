# IAM entry for all users to invoke the function
resource "google_cloudfunctions_function_iam_member" "invoker" {
  project        = google_cloudfunctions_function.pikey.project
  region         = google_cloudfunctions_function.pikey.region
  cloud_function = google_cloudfunctions_function.pikey.name

  role   = "roles/cloudfunctions.invoker"
  member = var.invoker
}


variable "invoker" {
  description = "Set who can invoke the lambda"
  default     = "allUsers"
}
