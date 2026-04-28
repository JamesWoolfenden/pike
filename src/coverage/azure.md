# Azure Resource Status

| Terraform  | Coverage % | Resources | Total Resources |
|------------|------------|-----------|-----------------|
| Resources  | 96.55      |  1093       |  1132            |
| Datasource | 100.00      |   395       |   395             |

## Deprecated

38 resources and 6 datasources are flagged as deprecated in the latest provider schema. Users pinned to an older provider major may already be affected when they upgrade.

### Deprecated Resources

| Resource | Note |
|---|---|
| azurerm_app_service | — |
| azurerm_app_service_active_slot | — |
| azurerm_app_service_hybrid_connection | — |
| azurerm_app_service_plan | — |
| azurerm_app_service_slot | — |
| azurerm_app_service_source_control_token | — |
| azurerm_automation_software_update_configuration | — |
| azurerm_batch_certificate | — |
| azurerm_data_protection_backup_instance_postgresql | — |
| azurerm_data_protection_backup_policy_postgresql | — |
| azurerm_databricks_workspace_customer_managed_key | — |
| azurerm_extended_custom_location | — |
| azurerm_function_app | — |
| azurerm_function_app_slot | — |
| azurerm_hpc_cache | — |
| azurerm_hpc_cache_access_policy | — |
| azurerm_hpc_cache_blob_nfs_target | — |
| azurerm_hpc_cache_blob_target | — |
| azurerm_hpc_cache_nfs_target | — |
| azurerm_maps_creator | — |
| azurerm_network_packet_capture | — |
| azurerm_orbital_contact | — |
| azurerm_orbital_contact_profile | — |
| azurerm_orbital_spacecraft | — |
| azurerm_postgresql_active_directory_administrator | — |
| azurerm_postgresql_configuration | — |
| azurerm_postgresql_database | — |
| azurerm_postgresql_firewall_rule | — |
| azurerm_postgresql_server | — |
| azurerm_postgresql_server_key | — |
| azurerm_postgresql_virtual_network_rule | — |
| azurerm_redis_enterprise_cluster | — |
| azurerm_redis_enterprise_database | — |
| azurerm_restore_point_collection | — |
| azurerm_security_center_auto_provisioning | — |
| azurerm_static_site | — |
| azurerm_static_site_custom_domain | — |
| azurerm_virtual_machine_scale_set | — |

### Deprecated Data Sources

| Data Source | Note |
|---|---|
| azurerm_app_service | — |
| azurerm_app_service_plan | — |
| azurerm_batch_certificate | — |
| azurerm_function_app | — |
| azurerm_postgresql_server | — |
| azurerm_redis_enterprise_database | — |

```shell
./resource.ps1 azurerm_app_service_active_slot
./resource.ps1 azurerm_app_service_certificate_binding
./resource.ps1 azurerm_app_service_source_control_token
./resource.ps1 azurerm_communication_service_email_domain_association
./resource.ps1 azurerm_container_registry_task_schedule_run_now
./resource.ps1 azurerm_cosmosdb_postgresql_cluster
./resource.ps1 azurerm_cosmosdb_postgresql_coordinator_configuration
./resource.ps1 azurerm_cosmosdb_postgresql_firewall_rule
./resource.ps1 azurerm_cosmosdb_postgresql_node_configuration
./resource.ps1 azurerm_cosmosdb_postgresql_role
./resource.ps1 azurerm_data_protection_resource_guard
./resource.ps1 azurerm_databricks_workspace_customer_managed_key
./resource.ps1 azurerm_dynatrace_monitor
./resource.ps1 azurerm_extended_custom_location
./resource.ps1 azurerm_key_vault_managed_hardware_security_module_key
./resource.ps1 azurerm_key_vault_managed_hardware_security_module_key_rotation_policy
./resource.ps1 azurerm_key_vault_managed_hardware_security_module_role_assignment
./resource.ps1 azurerm_key_vault_managed_hardware_security_module_role_definition
./resource.ps1 azurerm_log_analytics_workspace_table
./resource.ps1 azurerm_management_group_subscription_association
./resource.ps1 azurerm_network_packet_capture
./resource.ps1 azurerm_orbital_contact
./resource.ps1 azurerm_orbital_contact_profile
./resource.ps1 azurerm_orbital_spacecraft
./resource.ps1 azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association
./resource.ps1 azurerm_palo_alto_local_rulestack_outbound_untrust_certificate_association
./resource.ps1 azurerm_pim_active_role_assignment
./resource.ps1 azurerm_pim_eligible_role_assignment
./resource.ps1 azurerm_private_endpoint_application_security_group_association
./resource.ps1 azurerm_resource_management_private_link
./resource.ps1 azurerm_resource_management_private_link_association
./resource.ps1 azurerm_resource_provider_feature_registration
./resource.ps1 azurerm_resource_provider_registration
./resource.ps1 azurerm_restore_point_collection
./resource.ps1 azurerm_security_center_auto_provisioning
./resource.ps1 azurerm_virtual_desktop_scaling_plan_host_pool_association
./resource.ps1 azurerm_virtual_machine_gallery_application_assignment
./resource.ps1 azurerm_voice_services_communications_gateway
./resource.ps1 azurerm_voice_services_communications_gateway_test_line
```
