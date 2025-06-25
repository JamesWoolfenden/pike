resource "google_project_iam_custom_role" "pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "dataplex.glossaries.create",
    "dataplex.glossaries.get",
    "dataplex.glossaries.update",
    "dataplex.glossaries.delete",
    "dataplex.operations.get",

    "dataplex.glossaries.getIamPolicy",
    "dataplex.glossaries.setIamPolicy",


    "dataplex.glossaryCategories.create",
    "dataplex.glossaryCategories.get",
    "dataplex.glossaryCategories.update",
    "dataplex.glossaryCategories.delete",

    "dataplex.glossaryTerms.create",
    "dataplex.glossaryTerms.get",
    "dataplex.glossaryTerms.update",
    "dataplex.glossaryTerms.delete",

  ]
}
