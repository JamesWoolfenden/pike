
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    // google_dataplex_aspect_type


    // google_dataplex_entry


    // google_dataplex_entry_group
    // google_dataplex_entry_type

    // google_dataplex_lake
    // google_dataplex_task

    // google_dataplex_zone
  ]
}
