# Google Resource Status

| Terraform  | Coverage % | Resources | Total Resources |
|------------|------------|-----------|-----------------|
| Resources  | 89.44      |  1093       |  1222            |
| Datasource | 100.00      |   421       |   421             |

## Deprecated

28 resources and 7 datasources are flagged as deprecated in the latest provider schema. Users pinned to an older provider major may already be affected when they upgrade.

### Deprecated Resources

| Resource | Note |
|---|---|
| google_container_registry | — |
| google_data_catalog_entry | — |
| google_data_catalog_entry_group | — |
| google_data_catalog_entry_group_iam_binding | — |
| google_data_catalog_entry_group_iam_member | — |
| google_data_catalog_entry_group_iam_policy | — |
| google_data_catalog_tag | — |
| google_data_catalog_tag_template | — |
| google_data_catalog_tag_template_iam_binding | — |
| google_data_catalog_tag_template_iam_member | — |
| google_data_catalog_tag_template_iam_policy | — |
| google_iap_brand | — |
| google_iap_client | — |
| google_ml_engine_model | — |
| google_network_services_service_binding | — |
| google_notebooks_environment | — |
| google_notebooks_instance | — |
| google_notebooks_instance_iam_binding | — |
| google_notebooks_instance_iam_member | — |
| google_notebooks_instance_iam_policy | — |
| google_notebooks_runtime | — |
| google_notebooks_runtime_iam_binding | — |
| google_notebooks_runtime_iam_member | — |
| google_notebooks_runtime_iam_policy | — |
| google_pubsub_lite_reservation | — |
| google_pubsub_lite_subscription | — |
| google_pubsub_lite_topic | — |
| google_scc_v2_organization_scc_big_query_exports | — |

### Deprecated Data Sources

| Data Source | Note |
|---|---|
| google_container_registry_image | — |
| google_container_registry_repository | — |
| google_data_catalog_entry_group_iam_policy | — |
| google_data_catalog_tag_template_iam_policy | — |
| google_kms_secret_ciphertext | — |
| google_notebooks_instance_iam_policy | — |
| google_notebooks_runtime_iam_policy | — |

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
./resource.ps1 google_apigee_space
./resource.ps1 google_biglake_iceberg_table
./resource.ps1 google_biglake_iceberg_table_iam_binding
./resource.ps1 google_biglake_iceberg_table_iam_member
./resource.ps1 google_biglake_iceberg_table_iam_policy
./resource.ps1 google_bigquery_reservation_group
./resource.ps1 google_chronicle_data_table
./resource.ps1 google_chronicle_data_table_row
./resource.ps1 google_compute_firewall_policy_with_rules
./resource.ps1 google_compute_network_firewall_policy_with_rules
./resource.ps1 google_compute_preview_feature
./resource.ps1 google_compute_region_composite_health_check
./resource.ps1 google_compute_region_health_aggregation_policy
./resource.ps1 google_compute_region_health_source
./resource.ps1 google_compute_region_network_firewall_policy
./resource.ps1 google_compute_region_network_firewall_policy_association
./resource.ps1 google_compute_region_network_firewall_policy_rule
./resource.ps1 google_compute_region_network_firewall_policy_with_rules
./resource.ps1 google_compute_region_ssl_policy
./resource.ps1 google_compute_resize_request
./resource.ps1 google_contact_center_insights_assessment_rule
./resource.ps1 google_contact_center_insights_auto_labeling_rule
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
./resource.ps1 google_dialogflow_environment
./resource.ps1 google_discovery_engine_serving_config
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
./resource.ps1 google_kms_project_autokey_config
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
./resource.ps1 google_network_services_multicast_consumer_association
./resource.ps1 google_network_services_multicast_domain
./resource.ps1 google_network_services_multicast_domain_activation
./resource.ps1 google_network_services_multicast_domain_group
./resource.ps1 google_network_services_multicast_group_consumer_activation
./resource.ps1 google_network_services_multicast_group_producer_activation
./resource.ps1 google_network_services_multicast_group_range
./resource.ps1 google_network_services_multicast_group_range_activation
./resource.ps1 google_network_services_multicast_producer_association
./resource.ps1 google_scc_v2_organization_source
./resource.ps1 google_service_networking_peered_dns_domain
./resource.ps1 google_service_networking_vpc_service_controls
./resource.ps1 google_sql_provision_script
./resource.ps1 google_workload_identity_service_agent
```
