resource "azurerm_data_factory_linked_service_azure_databricks" "pike_gen" {
  name            = "ADBLinkedServiceViaMSI"
  data_factory_id = "pike"
  description     = "ADB Linked Service via MSI"
  adb_domain      = "https://${azurerm_databricks_workspace.example.workspace_url}"

  msi_workspace_id = "pike"

  new_cluster_config {
    node_type             = "Standard_NC12"
    cluster_version       = "5.5.x-gpu-scala2.11"
    min_number_of_workers = 1
    max_number_of_workers = 5
    driver_node_type      = "Standard_NC12"
    log_destination       = "dbfs:/logs"

    custom_tags = {
      custom_tag1 = "sct_value_1"
      custom_tag2 = "sct_value_2"
    }

    spark_config = {
      config1 = "value1"
      config2 = "value2"
    }

    spark_environment_variables = {
      envVar1 = "value1"
      envVar2 = "value2"
    }

    init_scripts = ["init.sh", "init2.sh"]
  }
}
