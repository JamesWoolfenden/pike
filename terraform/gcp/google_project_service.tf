#resource "google_project_service" "pike" {
#    project = "pike-361314"
#    service = "iam.googleapis.com"
#
#    timeouts {
#      create = "30m"
#      update = "40m"
#    }
#
#    disable_dependent_services = true
#  }

resource "google_storage_bucket" "bucket" {
  name     = "test-bucket-jgw-today"
  location = "US"
}
