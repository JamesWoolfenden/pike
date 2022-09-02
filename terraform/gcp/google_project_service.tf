#resource "google_project_service" "pike" {
#    project = "pike-361314"
#    service = "artifactregistry.googleapis.com"
#
#    timeouts {
#      create = "30m"
#      update = "40m"
#    }
#
#    disable_dependent_services = true
#  }
#
#output "project_service" {
#  value=google_project_service.pike
#}
