resource "google_workflows_workflow" "pike" {
  name        = "pike-test-workflow"
  region      = "us-central1"
  description = "Test workflow for Pike permission discovery"

  source_contents = <<-EOF
    - step1:
        call: http.get
        args:
          url: https://www.google.com
        result: response
    - step2:
        return: $${response.code}
  EOF

  labels = {
    environment = "test"
    managed_by  = "pike"
  }
}
