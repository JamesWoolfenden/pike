resource "google_bigtable_instance" "instance" {
  name = "tf-instance"

  cluster {
    cluster_id   = "tf-instance-cluster"
    zone         = "us-central1-b"
    num_nodes    = 3
    storage_type = "HDD"
  }
  deletion_protection = false

  #
  # lifecycle {
  #   prevent_destroy = true
  # }
}

resource "google_bigtable_table" "table" {
  name          = "tf-table"
  instance_name = google_bigtable_instance.instance.name
  split_keys    = ["a", "b", "c"]

  # lifecycle {
  #   prevent_destroy = true
  # }

  column_family {
    family = "family-first"
  }

  column_family {
    family = "family-second"
  }

  change_stream_retention = "24h0m0s"
}

resource "google_bigtable_authorized_view" "authorized_view" {
  name          = "tf-authorized-view"
  instance_name = google_bigtable_instance.instance.name
  table_name    = google_bigtable_table.table.name

  # lifecycle {
  #   prevent_destroy = true
  # }

  subset_view {
    row_prefixes = [base64encode("prefix#")]

    family_subsets {
      family_name = "family-first"
      qualifiers  = [base64encode("qualifier"), base64encode("qualifier-second")]
    }

    family_subsets {
      family_name        = "family-second"
      qualifier_prefixes = [""]
    }
  }
}
