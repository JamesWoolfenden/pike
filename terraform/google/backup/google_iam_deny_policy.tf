resource "google_project" "project" {
  project_id      = "my-project"
  name            = "my-project"
  org_id          = "123456789"
  billing_account = "000000-0000000-0000000-000000"
  deletion_policy = "DELETE"
}

resource "google_iam_deny_policy" "example" {
  parent       = urlencode("cloudresourcemanager.googleapis.com/projects/${google_project.project.project_id}")
  name         = "my-deny-policy"
  display_name = "A deny rule"
  rules {
    description = "First rule"
    deny_rule {
      denied_principals = ["principalSet://goog/public:all"]
      denial_condition {
        title      = "Some expr"
        expression = "!resource.matchTag('12345678/env', 'test')"
      }
      denied_permissions = ["cloudresourcemanager.googleapis.com/projects.update"]
    }
  }
  rules {
    description = "Second rule"
    deny_rule {
      denied_principals = ["principalSet://goog/public:all"]
      denial_condition {
        title      = "Some expr"
        expression = "!resource.matchTag('12345678/env', 'test')"
      }
      denied_permissions   = ["cloudresourcemanager.googleapis.com/projects.update"]
      exception_principals = ["principal://iam.googleapis.com/projects/-/serviceAccounts/${google_service_account.test-account.email}"]
    }
  }
}

resource "google_service_account" "test-account" {
  account_id   = "svc-acc"
  display_name = "Test Service Account"
  project      = google_project.project.project_id
}
