resource "google_data_catalog_tag" "pike" {
  parent   = google_data_catalog_entry.pike.id
  template = google_data_catalog_tag_template.pike.id

  fields {
    field_name   = "source"
    string_value = "my-string2"
  }
}
