resource "google_cloud_tasks_queue" "pike" {
  name     = "cloud-tasks-queue-test"
  location = "us-central1"
}
