resource "azurerm_mssql_job_step" "pike_gen" {
  name                = "example-job-step"
  job_id              = "pike"
  job_credential_id   = "pike"
  job_target_group_id = "pike"

  job_step_index = 1
  sql_script     = <<EOT
IF NOT EXISTS (SELECT * FROM sys.objects WHERE [name] = N'Pets')
  CREATE TABLE Pets (
    Animal NVARCHAR(50),
    Name NVARCHAR(50),
  );
EOT
}
