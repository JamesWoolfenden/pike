resource "google_cloud_security_compliance_framework" "example" {
  organization = "123456789"
  location     = "global"
  framework_id = "example-framework"

  display_name = "Terraform Framework Name"
  description  = "An Terraform description for the framework"

  cloud_control_details {
    name              = "organizations/123456789/locations/global/cloudControls/builtin-assess-resource-availability"
    major_revision_id = "1"

    parameters {
      name = "location"
      parameter_value {
        string_value = "us-central1"
      }
    }
  }

  cloud_control_details {
    name              = "organizations/123456789/locations/global/cloudControls/builtin-cmek-key-in-use-for-bigquery-table"
    major_revision_id = "1"

    parameters {
      name = "location"
      parameter_value {
        string_list_value {
          values = ["us-central1", "us-west1"]
        }
      }
    }
  }

  cloud_control_details {
    name              = "organizations/123456789/locations/global/cloudControls/builtin-enable-automatic-backups-cloud-sql"
    major_revision_id = "1"

    parameters {
      name = "location"
      parameter_value {
        bool_value = true
      }
    }
  }

  cloud_control_details {
    name              = "organizations/123456789/locations/global/cloudControls/builtin-require-cmek-on-bigquery-datasets"
    major_revision_id = "1"

    parameters {
      name = "location"
      parameter_value {
        number_value = 1
      }
    }
  }


}
