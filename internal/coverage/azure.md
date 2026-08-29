# Azure Resource Status

| Terraform  | Coverage % | Resources | Total Resources |
|------------|------------|-----------|-----------------|
| Resources  | 96.47      |  1065       |  1104            |
| Datasource | 98.48      |   389       |   395             |

## Deprecated

30 resources and 2 datasources are flagged as deprecated in provider schema v5.3.0. Users pinned to an older provider major may already be affected when they upgrade.

### Deprecated Resources

| Resource | Note |
|---|---|
| azurerm_spring_cloud_accelerator | — |
| azurerm_spring_cloud_active_deployment | — |
| azurerm_spring_cloud_api_portal | — |
| azurerm_spring_cloud_api_portal_custom_domain | — |
| azurerm_spring_cloud_app | — |
| azurerm_spring_cloud_app_cosmosdb_association | — |
| azurerm_spring_cloud_app_dynamics_application_performance_monitoring | — |
| azurerm_spring_cloud_app_mysql_association | — |
| azurerm_spring_cloud_app_redis_association | — |
| azurerm_spring_cloud_application_insights_application_performance_monitoring | — |
| azurerm_spring_cloud_application_live_view | — |
| azurerm_spring_cloud_build_deployment | — |
| azurerm_spring_cloud_build_pack_binding | — |
| azurerm_spring_cloud_builder | — |
| azurerm_spring_cloud_certificate | — |
| azurerm_spring_cloud_configuration_service | — |
| azurerm_spring_cloud_container_deployment | — |
| azurerm_spring_cloud_custom_domain | — |
| azurerm_spring_cloud_customized_accelerator | — |
| azurerm_spring_cloud_dev_tool_portal | — |
| azurerm_spring_cloud_dynatrace_application_performance_monitoring | — |
| azurerm_spring_cloud_elastic_application_performance_monitoring | — |
| azurerm_spring_cloud_gateway | — |
| azurerm_spring_cloud_gateway_custom_domain | — |
| azurerm_spring_cloud_gateway_route_config | — |
| azurerm_spring_cloud_java_deployment | — |
| azurerm_spring_cloud_new_relic_application_performance_monitoring | — |
| azurerm_spring_cloud_service | — |
| azurerm_spring_cloud_storage | — |
| azurerm_virtual_machine_scale_set | — |

### Deprecated Data Sources

| Data Source | Note |
|---|---|
| azurerm_spring_cloud_app | — |
| azurerm_spring_cloud_service | — |

```shell
./resource.ps1 azurerm_app_service_certificate_binding
./resource.ps1 azurerm_cdn_frontdoor_batch_rule_set
./resource.ps1 azurerm_cognitive_account_connection_account_key
./resource.ps1 azurerm_cognitive_account_connection_account_managed_identity
./resource.ps1 azurerm_cognitive_account_connection_api_key
./resource.ps1 azurerm_cognitive_account_connection_custom_keys
./resource.ps1 azurerm_cognitive_account_connection_entra_id
./resource.ps1 azurerm_communication_service_email_domain_association
./resource.ps1 azurerm_container_registry_task_schedule_run_now
./resource.ps1 azurerm_cosmosdb_postgresql_cluster
./resource.ps1 azurerm_cosmosdb_postgresql_coordinator_configuration
./resource.ps1 azurerm_cosmosdb_postgresql_firewall_rule
./resource.ps1 azurerm_cosmosdb_postgresql_node_configuration
./resource.ps1 azurerm_cosmosdb_postgresql_role
./resource.ps1 azurerm_data_protection_resource_guard
./resource.ps1 azurerm_dynatrace_monitor
./resource.ps1 azurerm_key_vault_managed_hardware_security_module_key
./resource.ps1 azurerm_key_vault_managed_hardware_security_module_key_rotation_policy
./resource.ps1 azurerm_key_vault_managed_hardware_security_module_role_assignment
./resource.ps1 azurerm_key_vault_managed_hardware_security_module_role_definition
./resource.ps1 azurerm_kubernetes_automatic_cluster
./resource.ps1 azurerm_log_analytics_workspace_table
./resource.ps1 azurerm_log_analytics_workspace_table_microsoft
./resource.ps1 azurerm_management_group_subscription_association
./resource.ps1 azurerm_netapp_volume_bucket
./resource.ps1 azurerm_netapp_volume_bucket_with_server
./resource.ps1 azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association
./resource.ps1 azurerm_palo_alto_local_rulestack_outbound_untrust_certificate_association
./resource.ps1 azurerm_pim_active_role_assignment
./resource.ps1 azurerm_pim_eligible_role_assignment
./resource.ps1 azurerm_playwright_workspace
./resource.ps1 azurerm_private_endpoint_application_security_group_association
./resource.ps1 azurerm_resource_management_private_link
./resource.ps1 azurerm_resource_management_private_link_association
./resource.ps1 azurerm_resource_provider_feature_registration
./resource.ps1 azurerm_resource_provider_registration
./resource.ps1 azurerm_storage_account_table_properties
./resource.ps1 azurerm_virtual_desktop_scaling_plan_host_pool_association
./resource.ps1 azurerm_virtual_machine_gallery_application_assignment
./resource.ps1 azurerm_cdn_frontdoor_batch_rule_set -type data
./resource.ps1 azurerm_cdn_frontdoor_security_policy -type data
./resource.ps1 azurerm_kubernetes_automatic_cluster -type data
./resource.ps1 azurerm_netapp_volume_bucket -type data
./resource.ps1 azurerm_netapp_volume_bucket_with_server -type data
./resource.ps1 azurerm_playwright_workspace -type data
```
