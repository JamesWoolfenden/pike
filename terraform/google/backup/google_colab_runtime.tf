resource "google_colab_runtime" "runtime" {
  name     = "colab-runtime"
  location = "us-central1"

  notebook_runtime_template_ref {
    notebook_runtime_template = google_colab_runtime_template.my_template.id
  }

  display_name = "Runtime basic"
  runtime_user = "james.woolfenden@gmail.com"

  depends_on = [
    google_colab_runtime_template.my_template,
  ]
}
