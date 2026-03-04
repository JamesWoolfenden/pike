# Google Resource Status

| Terraform  | Coverage % | Resources | Total Resources |
|------------|------------|-----------|-----------------|
| Resources  | 88.18      |  1044       |  1184            |
| Datasource | 96.81      |   395       |   408             |

```shell
./resource.ps1 google_access_context_manager_egress_policy
./resource.ps1 google_access_context_manager_ingress_policy
./resource.ps1 google_apigee_endpoint_attachment
./resource.ps1 google_apigee_env_references
./resource.ps1 google_apigee_flowhook
./resource.ps1 google_apigee_keystores_aliases_key_cert_file
./resource.ps1 google_apigee_keystores_aliases_pkcs12
./resource.ps1 google_apigee_sharedflow
./resource.ps1 google_apigee_sharedflow_deployment
./resource.ps1 google_compute_firewall_policy_with_rules
./resource.ps1 google_compute_network_firewall_policy_with_rules
./resource.ps1 google_compute_preview_feature
./resource.ps1 google_compute_region_network_firewall_policy
./resource.ps1 google_compute_region_network_firewall_policy_association
./resource.ps1 google_compute_region_network_firewall_policy_rule
./resource.ps1 google_compute_region_network_firewall_policy_with_rules
./resource.ps1 google_compute_region_ssl_policy
./resource.ps1 google_compute_resize_request
./resource.ps1 google_container_analysis_occurrence
./resource.ps1 google_container_attached_cluster
./resource.ps1 google_container_aws_cluster
./resource.ps1 google_container_aws_node_pool
./resource.ps1 google_container_azure_client
./resource.ps1 google_container_azure_cluster
./resource.ps1 google_container_azure_node_pool
./resource.ps1 google_dataproc_gdc_application_environment
./resource.ps1 google_dataproc_gdc_service_instance
./resource.ps1 google_dataproc_gdc_spark_application
./resource.ps1 google_dataproc_metastore_federation
./resource.ps1 google_dataproc_metastore_service
./resource.ps1 google_dialogflow_cx_test_case
./resource.ps1 google_dialogflow_cx_tool_version
./resource.ps1 google_dialogflow_generator
./resource.ps1 google_dialogflow_version
./resource.ps1 google_discovery_engine_serving_config
./resource.ps1 google_document_ai_processor_default_version
./resource.ps1 google_document_ai_warehouse_document_schema
./resource.ps1 google_document_ai_warehouse_location
./resource.ps1 google_firebase_app_check_app_attest_config
./resource.ps1 google_firebase_app_check_debug_token
./resource.ps1 google_firebase_app_check_device_check_config
./resource.ps1 google_firebase_app_check_play_integrity_config
./resource.ps1 google_firebase_app_check_recaptcha_enterprise_config
./resource.ps1 google_firebase_app_check_recaptcha_v3_config
./resource.ps1 google_firebase_app_check_service_config
./resource.ps1 google_firebase_app_hosting_backend
./resource.ps1 google_firebase_app_hosting_build
./resource.ps1 google_firebase_app_hosting_default_domain
./resource.ps1 google_firebase_app_hosting_domain
./resource.ps1 google_firebase_app_hosting_traffic
./resource.ps1 google_firebase_data_connect_service
./resource.ps1 google_gke_hub_scope_rbac_role_binding
./resource.ps1 google_identity_platform_config
./resource.ps1 google_identity_platform_default_supported_idp_config
./resource.ps1 google_identity_platform_inbound_saml_config
./resource.ps1 google_identity_platform_oauth_idp_config
./resource.ps1 google_identity_platform_tenant
./resource.ps1 google_identity_platform_tenant_default_supported_idp_config
./resource.ps1 google_identity_platform_tenant_inbound_saml_config
./resource.ps1 google_identity_platform_tenant_oauth_idp_config
./resource.ps1 google_integration_connectors_connection
./resource.ps1 google_integration_connectors_endpoint_attachment
./resource.ps1 google_integration_connectors_managed_zone
./resource.ps1 google_integrations_auth_config
./resource.ps1 google_integrations_client
./resource.ps1 google_netapp_active_directory
./resource.ps1 google_netapp_backup
./resource.ps1 google_netapp_host_group
./resource.ps1 google_netapp_storage_pool
./resource.ps1 google_netapp_volume
./resource.ps1 google_netapp_volume_quota_rule
./resource.ps1 google_netapp_volume_replication
./resource.ps1 google_netapp_volume_snapshot
./resource.ps1 google_network_connectivity_destination
./resource.ps1 google_network_connectivity_group
./resource.ps1 google_network_connectivity_internal_range
./resource.ps1 google_network_connectivity_multicloud_data_transfer_config
./resource.ps1 google_network_connectivity_policy_based_route
./resource.ps1 google_network_connectivity_regional_endpoint
./resource.ps1 google_network_connectivity_service_connection_policy
./resource.ps1 google_network_connectivity_spoke
./resource.ps1 google_network_management_connectivity_test
./resource.ps1 google_network_management_vpc_flow_logs_config
./resource.ps1 google_network_security_authz_policy
./resource.ps1 google_network_security_backend_authentication_config
./resource.ps1 google_network_security_dns_threat_detector
./resource.ps1 google_network_security_firewall_endpoint
./resource.ps1 google_network_security_firewall_endpoint_association
./resource.ps1 google_network_security_gateway_security_policy
./resource.ps1 google_network_security_gateway_security_policy_rule
./resource.ps1 google_network_security_intercept_deployment
./resource.ps1 google_network_security_intercept_deployment_group
./resource.ps1 google_network_security_intercept_endpoint_group
./resource.ps1 google_network_security_intercept_endpoint_group_association
./resource.ps1 google_network_security_mirroring_deployment
./resource.ps1 google_network_security_mirroring_deployment_group
./resource.ps1 google_network_security_mirroring_endpoint
./resource.ps1 google_network_security_mirroring_endpoint_group
./resource.ps1 google_network_security_mirroring_endpoint_group_association
./resource.ps1 google_network_security_security_profile
./resource.ps1 google_network_security_security_profile_group
./resource.ps1 google_network_security_server_tls_policy
./resource.ps1 google_network_security_tls_inspection_policy
./resource.ps1 google_network_security_url_lists
./resource.ps1 google_network_services_authz_extension
./resource.ps1 google_network_services_edge_cache_keyset
./resource.ps1 google_network_services_edge_cache_origin
./resource.ps1 google_network_services_edge_cache_service
./resource.ps1 google_network_services_endpoint_policy
./resource.ps1 google_network_services_gateway
./resource.ps1 google_network_services_grpc_route
./resource.ps1 google_network_services_http_route
./resource.ps1 google_network_services_lb_edge_extension
./resource.ps1 google_network_services_lb_route_extension
./resource.ps1 google_network_services_lb_traffic_extension
./resource.ps1 google_network_services_mesh
./resource.ps1 google_network_services_multicast_consumer_association
./resource.ps1 google_network_services_multicast_domain
./resource.ps1 google_network_services_multicast_domain_activation
./resource.ps1 google_network_services_multicast_domain_group
./resource.ps1 google_network_services_multicast_group_consumer_activation
./resource.ps1 google_network_services_multicast_group_producer_activation
./resource.ps1 google_network_services_multicast_group_range
./resource.ps1 google_network_services_multicast_group_range_activation
./resource.ps1 google_network_services_multicast_producer_association
./resource.ps1 google_network_services_service_binding
./resource.ps1 google_network_services_tcp_route
./resource.ps1 google_network_services_tls_route
./resource.ps1 google_network_services_wasm_plugin
./resource.ps1 google_oracle_database_exadb_vm_cluster
./resource.ps1 google_resource_manager_capability
./resource.ps1 google_scc_v2_organization_source
./resource.ps1 google_secure_source_manager_hook
./resource.ps1 google_securityposture_posture_deployment
./resource.ps1 google_service_networking_peered_dns_domain
./resource.ps1 google_service_networking_vpc_service_controls
./resource.ps1 google_storage_anywhere_cache
./resource.ps1 google_storage_batch_operations_job
./resource.ps1 google_storage_transfer_agent_pool
./resource.ps1 google_storage_transfer_job
./resource.ps1 google_artifact_registry_package -type data
./resource.ps1 google_artifact_registry_tag -type data
./resource.ps1 google_backup_dr_backup_plan -type data
./resource.ps1 google_backup_dr_management_server -type data
./resource.ps1 google_cloud_asset_search_all_resources -type data
./resource.ps1 google_compute_region_ssl_policy -type data
./resource.ps1 google_firestore_document -type data
./resource.ps1 google_iam_testable_permissions -type data
./resource.ps1 google_logging_sink -type data
./resource.ps1 google_network_management_connectivity_test_run -type data
./resource.ps1 google_network_management_connectivity_tests -type data
./resource.ps1 google_storage_managed_folder_iam_policy -type data
./resource.ps1 google_vertex_ai_index -type data
```
