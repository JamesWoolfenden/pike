# locals {
#   project = "412922" # Google Cloud Platform Project ID
# }
#
# resource "google_service_account" "account" {
#   account_id = "gcf-sa"
#   display_name = "Test Service Account"
# }
#
# //allows the build to assign this user to cloud run
# resource "google_service_account_iam_member" "pike2" {
#   service_account_id = google_service_account.account.name //service for cloud run
#   role               = "roles/iam.serviceAccountUser"
#   member             = "serviceAccount:${google_service_account.account.email}" //allow who context of build
# }
#
# # resource "google_pubsub_topic" "topic" {
# #   name = "functions2-topic"
# # }
#
# resource "google_storage_bucket" "bucket" {
#   name     = "${local.project}-gcf-source-jgw"  # Every bucket name must be globally unique
#   location = "US"
#   uniform_bucket_level_access = true
# }
#
# resource "google_storage_bucket_object" "object" {
#   name   = "test.zip"
#   bucket = google_storage_bucket.bucket.name
#   source = "./test.zip"  # Add path to the zipped function source code
# }
#
# resource "google_cloudfunctions2_function" "function" {
#   name = "gcf-function"
#   location = "europe-west2"
#   description = "a new function"
#
#   build_config {
#     runtime = "go123"
#     entry_point = "hello_http"  # Set the entry point
#     environment_variables = {
#       BUILD_CONFIG_TEST = "build_test"
#     }
#     source {
#       storage_source {
#         bucket = google_storage_bucket.bucket.name
#         object = google_storage_bucket_object.object.name
#       }
#     }
#   }
#
#   service_config {
#     max_instance_count  = 3
#     min_instance_count = 1
#     available_memory    = "4Gi"
#     timeout_seconds     = 60
#     max_instance_request_concurrency = 80
#     available_cpu = "4"
#     environment_variables = {
#       SERVICE_CONFIG_TEST      = "config_test"
#       SERVICE_CONFIG_DIFF_TEST = google_service_account.account.email
#     }
#     ingress_settings = "ALLOW_INTERNAL_ONLY"
#     all_traffic_on_latest_revision = true
#  //   service_account_email = google_service_account.account.email
#   }
#
#   # event_trigger {
#   #   trigger_region = "us-central1"
#   #   event_type = "google.cloud.pubsub.topic.v1.messagePublished"
#   #   pubsub_topic = google_pubsub_topic.topic.id
#   #   retry_policy = "RETRY_POLICY_RETRY"
#   # }
#
#   depends_on = [
#     google_service_account_iam_member.pike2
#   ]
# }
