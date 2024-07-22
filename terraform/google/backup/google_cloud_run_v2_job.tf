#module "backups" {
#  source="github.com/JamesWoolfenden/terraform-gcp-bigtable-backups"
#  image    = "us-central1-docker.pkg.dev/pike-gcp/pike/backup:latest"
#  job_name = "backup"
#  schedule = "0 10 * * 1"
#  labels = {
#    pike = "permissions"
#    more = "labels"
#  }
#  region     = "us-central1"
#  project_id = "pike-gcp"
#}

resource "google_cloud_run_v2_job" "pike" {
  location = ""
  name     = ""
}
