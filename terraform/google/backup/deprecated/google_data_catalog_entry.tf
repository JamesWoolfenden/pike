resource "google_data_catalog_entry" "pike" {
  entry_group = google_data_catalog_entry_group.pike.id
  entry_id    = "my_entry"

  user_specified_type   = "my_custom_type"
  user_specified_system = "SomethingExternal"
}
