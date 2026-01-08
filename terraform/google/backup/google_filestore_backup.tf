resource "google_filestore_backup" "pike" {
  name              = "fs-bkup"
  location          = "us-central1"
  description       = "This is a filestore backup for the test instance"
  source_instance   = google_filestore_instance.pike.id
  source_file_share = "share1"

  labels = {
    "files" : "label1",
    "other-label" : "label2"
  }
}
