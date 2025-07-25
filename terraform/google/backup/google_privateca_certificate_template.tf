resource "google_privateca_certificate_template" "default" {
  name        = "my-template2"
  location    = "europe-west2"
  description = "A sample certificate template"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }

  maximum_lifetime = "86400s"

  passthrough_extensions {
    additional_extensions {
      object_id_path = [1, 6]
    }
    known_extensions = ["EXTENDED_KEY_USAGE"]
  }

  predefined_values {
    additional_extensions {
      object_id {
        object_id_path = [1, 6]
      }
      value    = "c3RyaW5nCg=="
      critical = true
    }
    aia_ocsp_servers = ["string"]
    ca_options {
      is_ca                  = false
      max_issuer_path_length = 6
    }
    key_usage {
      base_key_usage {
        cert_sign          = false
        content_commitment = true
        crl_sign           = false
        data_encipherment  = true
        decipher_only      = true
        digital_signature  = true
        encipher_only      = true
        key_agreement      = true
        key_encipherment   = true
      }
      extended_key_usage {
        client_auth      = true
        code_signing     = true
        email_protection = true
        ocsp_signing     = true
        server_auth      = true
        time_stamping    = true
      }
      unknown_extended_key_usages {
        object_id_path = [1, 6]
      }
    }
    policy_ids {
      object_id_path = [1, 6]
    }
  }

  labels = {
    label-one = "value-one"
  }
}
