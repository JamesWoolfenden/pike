resource "google_bigtable_table" "table" {
  name          = "tf-table"
  instance_name = google_bigtable_instance.instance.name

  column_family {
    family = "name"
  }
}

resource "google_bigtable_gc_policy" "policy" {
  instance_name   = google_bigtable_instance.instance.name
  table           = google_bigtable_table.table.name
  column_family   = "name"
  deletion_policy = "ABANDON"


  gc_rules = <<EOF
  {
    "rules": [
      {
        "max_age": "168h"
      }
    ]
  }
  EOF
}
