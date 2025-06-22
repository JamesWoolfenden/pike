resource "google_project_iam_custom_role" "pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //google_data_catalog_tag_template
    "datacatalog.tagTemplates.create",
    "datacatalog.tagTemplates.getIamPolicy",
    "datacatalog.tagTemplates.get",

    //"datacatalog.tagTemplates.setIamPolicy",

    //google_data_catalog_entry_group
    "datacatalog.entryGroups.create",
    "datacatalog.entryGroups.get",
    "datacatalog.entryGroups.update",
    "datacatalog.entryGroups.delete",


    //google_data_catalog_tag
    "datacatalog.tagTemplates.get",
    "datacatalog.tagTemplates.getTag",
    "datacatalog.tagTemplates.use",
    "datacatalog.tagTemplates.delete",
    "datacatalog.entries.updateTag",


    //google_data_catalog_entry
    "datacatalog.entries.create",
    "datacatalog.entries.get",
    "datacatalog.entries.update",
    "datacatalog.entries.delete",

    //google_bigquery_datapolicy_data_policy
    "datacatalog.tagTemplates.setIamPolicy",
    "datacatalog.tagTemplates.getIamPolicy",

    //google_data_catalog_taxonomy
    "datacatalog.taxonomies.create",
    "datacatalog.taxonomies.get",
    "datacatalog.taxonomies.update",
    "datacatalog.taxonomies.delete"
  ]
}
