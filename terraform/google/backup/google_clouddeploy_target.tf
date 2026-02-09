resource "google_clouddeploy_target" "pike" {
  location          = "us-west1"
  name              = "target"
  deploy_parameters = {}
  description       = "multi-target description"

  execution_configs {
    usages            = ["RENDER", "DEPLOY"]
    execution_timeout = "3600s"
  }

  multi_target {
    target_ids = ["1", "2"]
  }
  project          = "pike-477416"
  require_approval = false

  annotations = {
    my_first_annotation = "example-annotation-1"

    my_second_annotation = "example-annotation-2"
  }

  labels = {
    my_first_label = "example-label-1"

    my_second_label = "example-label-2"
  }
  provider = google-beta
}
