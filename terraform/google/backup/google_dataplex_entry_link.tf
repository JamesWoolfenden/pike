terraform {
  required_providers {
    random = {
      source  = "hashicorp/random"
      version = "3.7.2"
    }
  }
}
resource "google_dataplex_entry_group" "entry-group-basic" {
  location       = "us-central1"
  entry_group_id = "tf-test-entry-group${random_id.suffix.id}"
  project        = "1111111111111"
}

resource "google_dataplex_entry" "source" {
  location       = "us-central1"
  entry_group_id = google_dataplex_entry_group.entry-group-basic.entry_group_id
  entry_id       = "tf-test-source-entry${random_id.suffix.id}"
  entry_type     = google_dataplex_entry_type.entry-type-basic.name
  project        = "1111111111111"
}

resource "google_dataplex_entry_type" "entry-type-basic" {
  entry_type_id = "tf-test-entry-type${random_id.suffix.id}"
  location      = "us-central1"
  project       = "1111111111111"
}

resource "google_dataplex_entry" "target" {
  location       = "us-central1"
  entry_group_id = google_dataplex_entry_group.entry-group-basic.entry_group_id
  entry_id       = "tf-test-target-entry${random_id.suffix.id}"
  entry_type     = google_dataplex_entry_type.entry-type-basic.name
  project        = "1111111111111"
}

resource "google_dataplex_entry_link" "basic_entry_link" {
  project        = "1111111111111"
  location       = "us-central1"
  entry_group_id = google_dataplex_entry_group.entry-group-basic.entry_group_id
  entry_link_id  = "tf-test-entry-link${random_id.suffix.id}"

  entry_link_type = "projects/655216118709/locations/global/entryLinkTypes/related"

  entry_references {
    name = google_dataplex_entry.source.name
  }

  entry_references {
    name = google_dataplex_entry.target.name
  }
}

resource "random_id" "suffix" {
  byte_length = 6
}
