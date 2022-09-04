# IAM entry for all users to invoke the function
#resource "google_cloudfunctions_function_iam_member" "invoker" {
#  project        = google_cloudfunctions_function.lambda.project
#  region         = google_cloudfunctions_function.lambda.region
#  cloud_function = google_cloudfunctions_function.lambda.name
#
#  role   = "roles/cloudfunctions.invoker"
#  member = var.invoker
#}
#
#
#variable "invoker" {
#  description = "Set who can invoke the lambda"
#  default="allUsers"
#}
