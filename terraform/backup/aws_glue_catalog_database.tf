resource "aws_glue_catalog_database" "aws_glue_catalog_database" {
  name = "mycatalogdatabase"
  create_table_default_permission {
    permissions = ["SELECT"]

    principal {
      data_lake_principal_identifier = "IAM_ALLOWED_PRINCIPALS"
    }
  }
}
