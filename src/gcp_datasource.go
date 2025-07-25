package pike

// GetGCPDataPermissions gets permissions required for datasources.
func GetGCPDataPermissions(result ResourceV2) ([]string, error) {
	temp := GCPDataLookup(result.Name)

	var (
		Permissions []string
		err         error
	)

	if temp != nil {
		Permissions, err = GetPermissionMap(temp.([]byte), result.Attributes, result.Name)
	} else {
		return nil, &notImplementedDatasourceError{result.Name}
	}

	return Permissions, err
}

func GCPDataLookup(result string) interface{} {
	TFLookup := map[string]interface{}{
		"google_access_approval_folder_service_account":                placeholder,
		"google_access_approval_organization_service_account":          placeholder,
		"google_access_approval_project_service_account":               placeholder,
		"google_access_context_manager_access_policy_iam_policy":       placeholder,
		"google_active_folder":                                         placeholder,
		"google_alloydb_cluster":                                       dataGoogleAlloydbCluster,
		"google_alloydb_instance":                                      dataGoogleAlloydbInstance,
		"google_alloydb_locations":                                     dataGoogleAlloydbLocations,
		"google_alloydb_supported_database_flags":                      dataGoogleAlloydbSupportedDatabaseFlags,
		"google_api_gateway_api_config_iam_policy":                     placeholder,
		"google_api_gateway_api_iam_policy":                            placeholder,
		"google_api_gateway_gateway_iam_policy":                        placeholder,
		"google_apigee_environment_iam_policy":                         placeholder,
		"google_app_engine_default_service_account":                    dataGoogleAppEngineDefaultServiceAccount,
		"google_apphub_application":                                    dataGoogleApphubApplication,
		"google_apphub_discovered_service":                             dataGoogleApphubDiscoveredService,
		"google_apphub_discovered_workload":                            dataGoogleApphubDiscoveredWorkload,
		"google_artifact_registry_docker_image":                        dataGoogleArtifactRegistryDockerImage,
		"google_artifact_registry_locations":                           dataGoogleArtifactRegistryLocations,
		"google_artifact_registry_repository":                          dataGoogleArtifactRegistryRepository,
		"google_artifact_registry_repository_iam_policy":               dataGoogleArtifactRegistryRepositoryIamPolicy,
		"google_backup_dr_backup":                                      dataGoogleBackupDrBackup,
		"google_backup_dr_backup_plan_association":                     dataGoogleBackupDrPlanAssociation,
		"google_backup_dr_backup_vault":                                dataGoogleBackupDrBackupVault,
		"google_backup_dr_data_source":                                 dataGoogleBackupDrDataSource,
		"google_beyondcorp_app_connection":                             dataGoogleBeyondcorpAppConnection,
		"google_beyondcorp_app_connector":                              dataGoogleBeyondcorpAppConnector,
		"google_beyondcorp_app_gateway":                                dataGoogleBeyondcorpAppGateway,
		"google_beyondcorp_application_iam_policy":                     dataGoogleBeyondcorpApplicationIamPolicy,
		"google_beyondcorp_security_gateway":                           dataGoogleBeyondcorpSecurityGateway,
		"google_beyondcorp_security_gateway_application_iam_policy":    dataGoogleBeyondcorpSecurityGatewayApplicationIamPolicy,
		"google_beyondcorp_security_gateway_iam_policy":                dataGoogleBeyondcorpSecurityGatewayIamPolicy,
		"google_bigquery_analytics_hub_data_exchange_iam_policy":       dataGoogleBigqueryHubDataExchangeIamPolicy,
		"google_bigquery_analytics_hub_listing_iam_policy":             dataGoogleBigqueryAnalyticsHubListingIamPolicy,
		"google_bigquery_connection_iam_policy":                        placeholder,
		"google_bigquery_datapolicy_data_policy_iam_policy":            dataGoogleBigqueryDatapolicyDataPolicyIamPolicy,
		"google_bigquery_dataset":                                      placeholder,
		"google_bigquery_dataset_iam_policy":                           placeholder,
		"google_bigquery_datasets":                                     placeholder,
		"google_bigquery_default_service_account":                      dataGoogleBigqueryDefaultServiceAccount,
		"google_bigquery_table":                                        placeholder,
		"google_bigquery_table_iam_policy":                             placeholder,
		"google_bigquery_tables":                                       placeholder,
		"google_bigtable_instance_iam_policy":                          dataGoogleBigtableInstanceIamPolicy,
		"google_bigtable_table_iam_policy":                             placeholder,
		"google_billing_account":                                       placeholder,
		"google_billing_account_iam_policy":                            placeholder,
		"google_binary_authorization_attestor_iam_policy":              dataGoogleBinaryAuthorizationAttestorIamPolicy,
		"google_certificate_manager_certificate_map":                   dataGoogleCertificateManagerCertificateMap,
		"google_certificate_manager_certificates":                      dataGoogleCertificateManagerCertificates,
		"google_client_config":                                         placeholder,
		"google_client_openid_userinfo":                                placeholder,
		"google_cloud_identity_group_lookup":                           placeholder,
		"google_cloud_identity_group_memberships":                      placeholder,
		"google_cloud_identity_groups":                                 placeholder,
		"google_cloud_quotas_quota_info":                               placeholder,
		"google_cloud_quotas_quota_infos":                              placeholder,
		"google_cloud_run_locations":                                   dataGoogleCloudRunLocations,
		"google_cloud_run_service":                                     dataGoogleCloudRunService,
		"google_cloud_run_service_iam_policy":                          dataGoogleCloudRunServiceIamPolicy,
		"google_cloud_run_v2_job":                                      dataGoogleCloudRunV2Job,
		"google_cloud_run_v2_job_iam_policy":                           dataGoogleCloudRunV2JobIamPolicy,
		"google_cloud_run_v2_service":                                  dataGoogleCloudRunV2Service,
		"google_cloud_run_v2_service_iam_policy":                       dataGoogleCloudRunV2ServiceIamPolicy,
		"google_cloud_run_v2_worker_pool":                              dataGoogleCloudRunV2WorkerPool,
		"google_cloud_run_v2_worker_pool_iam_policy":                   dataGoogleCloudRunWorkerPoolIamPolicy,
		"google_cloud_tasks_queue_iam_policy":                          dataGoogleCloudTasksQueueIamPolicy,
		"google_cloudbuild_trigger":                                    dataGoogleCloudbuildTrigger,
		"google_cloudbuildv2_connection_iam_policy":                    dataGoogleCloudbuildv2ConnectionIamPolicy,
		"google_clouddeploy_custom_target_type_iam_policy":             dataGoogleClouddeployCustomTargetTypeIamPolicy,
		"google_clouddeploy_delivery_pipeline_iam_policy":              dataGoogleClouddeployDeliveryPipelineIamPolicy,
		"google_clouddeploy_target_iam_policy":                         dataGoogleClouddeployTargetIamPolicy,
		"google_cloudfunctions2_function":                              dataGoogleCloudfunctionsFunction,
		"google_cloudfunctions2_function_iam_policy":                   dataGoogleCloudfunctionsFunctionIamPolicy,
		"google_cloudfunctions_function":                               dataGoogleCloudfunctionsFunction,
		"google_cloudfunctions_function_iam_policy":                    dataGoogleCloudfunctionsFunctionIamPolicy,
		"google_colab_runtime_template_iam_policy":                     dataGoogleColabRuntimeTemplateIamPolicy,
		"google_composer_environment":                                  dataGoogleComposerEnvironment,
		"google_composer_image_versions":                               dataGoogleComposerImageVersions,
		"google_composer_user_workloads_config_map":                    dataGoogleComposerUserWorkloadsConfigMap,
		"google_composer_user_workloads_secret":                        dataGoogleComposerUserWorkloadsSecret,
		"google_compute_address":                                       dataGoogleComputeAddress,
		"google_compute_addresses":                                     dataGoogleComputeAddresses,
		"google_compute_backend_bucket":                                dataGoogleComputeBackendBucket,
		"google_compute_backend_bucket_iam_policy":                     dataGoogleComputeBackendBucketIamPolicy,
		"google_compute_backend_service":                               dataGoogleComputeBackendService,
		"google_compute_backend_service_iam_policy":                    dataGoogleComputeBackendServiceIamPolicy,
		"google_compute_default_service_account":                       dataGoogleComputeDefaultServiceAccount,
		"google_compute_disk":                                          dataGoogleComputeDisk,
		"google_compute_disk_iam_policy":                               dataGoogleComputeDiskIamPolicy,
		"google_compute_forwarding_rule":                               dataGoogleComputeForwardingRule,
		"google_compute_forwarding_rules":                              dataGoogleComputeForwardingRules,
		"google_compute_global_address":                                dataGoogleComputeGlobalAddress,
		"google_compute_global_forwarding_rule":                        dataGoogleComputeGlobalForwardingRule,
		"google_compute_ha_vpn_gateway":                                dataGoogleComputeHaVpnGateway,
		"google_compute_health_check":                                  dataGoogleComputeHealthCheck,
		"google_compute_image":                                         placeholder,
		"google_compute_image_iam_policy":                              dataGoogleComputeImageIamPolicy,
		"google_compute_images":                                        dataGoogleComputeImages,
		"google_compute_instance":                                      dataGoogleComputeInstance,
		"google_compute_instance_group":                                dataGoogleComputeInstanceGroup,
		"google_compute_instance_group_manager":                        dataGoogleComputeInstanceGroupManager,
		"google_compute_instance_guest_attributes":                     dataGoogleComputeInstanceGuestAttributes,
		"google_compute_instance_iam_policy":                           placeholder,
		"google_compute_instance_serial_port":                          dataGoogleComputeInstanceSerialPort,
		"google_compute_instance_template":                             dataGoogleComputeInstanceTemplate,
		"google_compute_instance_template_iam_policy":                  dataGoogleComputeInstanceTemplateIamPolicy,
		"google_compute_instant_snapshot_iam_policy":                   dataGoogleComputeInstantSnapshotIamPolicy,
		"google_compute_lb_ip_ranges":                                  placeholder,
		"google_compute_machine_image_iam_policy":                      dataGoogleComputeMachineImageIamPolicy,
		"google_compute_machine_types":                                 dataGoogleComputeMachineTypes,
		"google_compute_network":                                       dataGoogleComputeNetwork,
		"google_compute_network_endpoint_group":                        dataGoogleComputeNetworkEndpointGroup,
		"google_compute_network_peering":                               placeholder,
		"google_compute_networks":                                      dataGoogleComputeNetworks,
		"google_compute_node_types":                                    dataGoogleComputeNodeTypes,
		"google_compute_region_backend_service":                        dataGoogleComputeRegionBackendService,
		"google_compute_region_backend_service_iam_policy":             dataGoogleComputeRegionBackendServiceIamPolicy,
		"google_compute_region_disk":                                   dataGoogleComputeRegionDisk,
		"google_compute_region_disk_iam_policy":                        dataGoogleComputeRegionDiskIamPolicy,
		"google_compute_region_instance_group":                         dataGoogleComputeRegionInstanceGroup,
		"google_compute_region_instance_group_manager":                 dataGoogleComputeRegionInstanceGroupManager,
		"google_compute_region_instance_template":                      dataGoogleComputeRegionInstanceTemplate,
		"google_compute_region_network_endpoint_group":                 dataGoogleComputeRegionNetworkEndpointGroup,
		"google_compute_region_ssl_certificate":                        placeholder,
		"google_compute_regions":                                       dataGoogleComputeRegions,
		"google_compute_reservation":                                   dataGoogleComputeReservation,
		"google_compute_resource_policy":                               dataGoogleComputeResourcePolicy,
		"google_compute_router":                                        dataGoogleComputeRouter,
		"google_compute_router_nat":                                    dataGoogleComputeRouterNat,
		"google_compute_router_status":                                 dataGoogleComputeRouterStatus,
		"google_compute_security_policy":                               dataGoogleComputeSecurityPolicy,
		"google_compute_snapshot":                                      dataGoogleComputeSnapshot,
		"google_compute_snapshot_iam_policy":                           dataGoogleComputeSnapshotIamPolicy,
		"google_compute_ssl_certificate":                               dataGoogleComputeSslCertificate,
		"google_compute_ssl_policy":                                    dataGoogleComputeSslPolicy,
		"google_compute_storage_pool_iam_policy":                       dataGoogleComputeStoragePoolIamPolicy,
		"google_compute_storage_pool_types":                            dataGoogleComputeStoragePoolTypes,
		"google_compute_subnetwork":                                    dataGoogleComputeSubnetwork,
		"google_compute_subnetwork_iam_policy":                         dataGoogleComputeSubnetworkIamPolicy,
		"google_compute_subnetworks":                                   dataGoogleComputeSubnetworks,
		"google_compute_vpn_gateway":                                   dataGoogleComputeVpnGateway,
		"google_compute_zones":                                         dataGoogleComputeZones,
		"google_container_analysis_note_iam_policy":                    dataGoogleContainerAnalysisNoteIamPolicy,
		"google_container_attached_install_manifest":                   dataGoogleContainerAttachedInstallManifest,
		"google_container_attached_versions":                           dataGoogleContainerAttachedVersions,
		"google_container_aws_versions":                                dataGoogleContainerAwsVersions,
		"google_container_azure_versions":                              dataGoogleContainerAzureVersions,
		"google_container_cluster":                                     dataGoogleContainerCluster,
		"google_container_engine_versions":                             placeholder,
		"google_container_registry_image":                              placeholder,
		"google_container_registry_repository":                         placeholder,
		"google_data_catalog_entry_group_iam_policy":                   placeholder,
		"google_data_catalog_policy_tag_iam_policy":                    placeholder,
		"google_data_catalog_tag_template_iam_policy":                  placeholder,
		"google_data_catalog_taxonomy_iam_policy":                      dataGoogleDataCatalogTaxonomyIamPolicy,
		"google_data_fusion_instance_iam_policy":                       dataGoogleDataFusionInstanceIamPolicy,
		"google_dataform_repository_iam_policy":                        dataGoogleDataformRepositoryIamPolicy,
		"google_dataplex_aspect_type_iam_policy":                       dataGoogleDataplexAspectTypeIamPolicy,
		"google_dataplex_asset_iam_policy":                             dataGoogleDataplexAssetIamPolicy,
		"google_dataplex_datascan_iam_policy":                          dataGoogleDataplexDatascanIamPolicy,
		"google_dataplex_entry_group_iam_policy":                       dataGoogleDataplexEntryGroupIamPolicy,
		"google_dataplex_entry_type_iam_policy":                        dataGoogleDataplexEntryTypeIamPolicy,
		"google_dataplex_glossary_iam_policy":                          dataGoogleDataplexGlossaryIamPolicy,
		"google_dataplex_lake_iam_policy":                              dataGoogleDataplexLakeIamPolicy,
		"google_dataplex_task_iam_policy":                              dataGoogleDataplexTaskIamPolicy,
		"google_dataplex_zone_iam_policy":                              dataGoogleDataplexZoneIamPolicy,
		"google_dataproc_autoscaling_policy_iam_policy":                placeholder,
		"google_dataproc_cluster_iam_policy":                           placeholder,
		"google_dataproc_job_iam_policy":                               placeholder,
		"google_dataproc_metastore_database_iam_policy":                dataGoogleDataprocMetastoreDatabaseIamPolicy,
		"google_dataproc_metastore_federation_iam_policy":              dataGoogleDataprocMetaStoreFederationIamPolicy,
		"google_dataproc_metastore_service":                            dataGoogleDataprocMetastoreService,
		"google_dataproc_metastore_service_iam_policy":                 dataGoogleDataprocMetastoreServiceIamPolicy,
		"google_dataproc_metastore_table_iam_policy":                   dataGoogleDataprocMetastoreTableIamPolicy,
		"google_datastream_static_ips":                                 dataGoogleDataStreamStaticIps,
		"google_dns_keys":                                              dataGoogleDNSKeys,
		"google_dns_managed_zone":                                      dataGoogleDNSManagedZone,
		"google_dns_managed_zones":                                     dataGoogleDNSManagedZones,
		"google_dns_managed_zone_iam_policy":                           dataGoogleDNSManagedZoneIamPolicy,
		"google_dns_record_set":                                        dataGoogleDNSRecordSet,
		"google_endpoints_service_consumers_iam_policy":                placeholder,
		"google_endpoints_service_iam_policy":                          dataGoogleEndpointsServiceIamPolicy,
		"google_filestore_instance":                                    dataGoogleFilestoreInstance,
		"google_folder":                                                dataGoogleFolder,
		"google_folder_iam_policy":                                     dataGoogleFolderIamPolicy,
		"google_folder_organization_policy":                            placeholder,
		"google_folders":                                               dataGoogleFolders,
		"google_gemini_repository_group_iam_policy":                    dataGoogleGeminiRepositoryGroupIamPolicy,
		"google_gke_backup_backup_plan_iam_policy":                     dataGoogleGkeBackupBackupPlanIamPolicy,
		"google_gke_backup_restore_plan_iam_policy":                    dataGoogleGkeBackupRestorePlanIamPolicy,
		"google_gke_hub_feature":                                       dataGoogleGkeHubFeature,
		"google_gke_hub_feature_iam_policy":                            dataGoogleGkeHubFeatureIamPolicy,
		"google_gke_hub_membership":                                    dataGoogleGkeHubMembership,
		"google_gke_hub_membership_iam_policy":                         dataGoogleGkeHubMembershipIamPolicy,
		"google_gke_hub_scope_iam_policy":                              dataGoogleGkeHubScopeIamPolicy,
		"google_healthcare_consent_store_iam_policy":                   dataGoogleHealthcareConsentStoreIamPolicy,
		"google_healthcare_dataset_iam_policy":                         dataGoogleHealthcareDatasetStoreIamPolicy,
		"google_healthcare_dicom_store_iam_policy":                     dataGoogleHealthcareDicomStoreIamPolicy,
		"google_healthcare_fhir_store_iam_policy":                      dataGoogleHealthcareFhirStoreIamPolicy,
		"google_healthcare_hl7_v2_store_iam_policy":                    dataGoogleHealthcareHl7V2StoreIamPolicy,
		"google_iam_policy":                                            placeholder,
		"google_iam_role":                                              placeholder,
		"google_iam_workload_identity_pool":                            dataGoogleIamWorkloadIdentityPool,
		"google_iam_workload_identity_pool_iam_policy":                 dataGoogleIamWorkloadIdentityPoolIamPolicy,
		"google_iam_workload_identity_pool_provider":                   dataGoogleIamWorkloadIdentityPoolProvider,
		"google_iap_app_engine_service_iam_policy":                     dataGoogleIapAppEngineServiceIamPolicy,
		"google_iap_app_engine_version_iam_policy":                     dataGoogleIapAppEngineVersionIamPolicy,
		"google_iap_client":                                            placeholder,
		"google_iap_tunnel_dest_group_iam_policy":                      dataGoogleIapTunnelDestGroupIamPolicy,
		"google_iap_tunnel_iam_policy":                                 dataGoogleIapTunnelIamPolicy,
		"google_iap_tunnel_instance_iam_policy":                        dataGoogleIapTunnelInstanceIamPolicy,
		"google_iap_web_backend_service_iam_policy":                    dataGoogleIapWebBackendServiceIamPolicy,
		"google_iap_web_cloud_run_service_iam_policy":                  dataGoogleIapWebCloudRunServiceIamPolicy,
		"google_iap_web_iam_policy":                                    dataGoogleIapWebIamPolicy,
		"google_iap_web_region_backend_service_iam_policy":             dataGoogleIapWebRegionBackendServiceIamPolicy,
		"google_iap_web_type_app_engine_iam_policy":                    dataGoogleIapWebTypeAppEngineIamPolicy,
		"google_iap_web_type_compute_iam_policy":                       dataGoogleIapWebTypeComputeIamPolicy,
		"google_kms_autokey_config":                                    placeholder,
		"google_kms_crypto_key":                                        dataGoogleKmsCryptoKey,
		"google_kms_crypto_key_iam_policy":                             dataGoogleKmsCryptoKeyIamPolicy,
		"google_kms_crypto_key_latest_version":                         placeholder,
		"google_kms_crypto_key_version":                                dataGoogleKmsCryptoKeyVersion,
		"google_kms_crypto_key_versions":                               placeholder,
		"google_kms_crypto_keys":                                       placeholder,
		"google_kms_ekm_connection_iam_policy":                         dataGoogleEkmConnectionIamPolicy,
		"google_kms_key_handle":                                        dataGoogleKmsKeyHandle,
		"google_kms_key_handles":                                       dataGoogleKmsKeyHandles,
		"google_kms_key_ring":                                          dataGoogleKmsKeyRing,
		"google_kms_key_ring_iam_policy":                               dataGoogleKmsKeyRingIamPolicy,
		"google_kms_key_rings":                                         dataGoogleKmsKeyRings,
		"google_kms_secret":                                            dataGoogleKmsSecret,
		"google_kms_secret_asymmetric":                                 dataGoogleKmsSecretAsymnetric,
		"google_kms_secret_ciphertext":                                 placeholder,
		"google_logging_folder_settings":                               placeholder,
		"google_logging_log_view_iam_policy":                           placeholder,
		"google_logging_organization_settings":                         placeholder,
		"google_logging_project_cmek_settings":                         dataGoogleLoggingProjectCmekSettings,
		"google_logging_project_settings":                              dataGoogleLoggingProjectSettings,
		"google_lustre_instance":                                       dataGoogleLustreInstance,
		"google_memcache_instance":                                     dataGoogleMemcacheInstance,
		"google_memorystore_instance":                                  dataGoogleMemorystoreInstance,
		"google_monitoring_app_engine_service":                         dataGoogleMonitoringAppEngineService,
		"google_monitoring_cluster_istio_service":                      dataGoogleMonitoringClusterIstioService,
		"google_monitoring_istio_canonical_service":                    dataGoogleMonitoringIstioCanonicalService,
		"google_monitoring_mesh_istio_service":                         dataGoogleMonitoringMeshIstioService,
		"google_monitoring_notification_channel":                       dataGoogleMonitoringNotificationChannel,
		"google_monitoring_uptime_check_ips":                           placeholder,
		"google_netblock_ip_ranges":                                    placeholder,
		"google_network_security_address_group_iam_policy":             dataGoogleSecurityAddressGroupIamPolicy,
		"google_notebooks_instance_iam_policy":                         dataGoogleNotebooksInstanceIamPolicy,
		"google_notebooks_runtime_iam_policy":                          dataGoogleNotebooksRuntimeIamPolicy,
		"google_oracle_database_autonomous_database":                   dataGoogleOracleDatabaseAutonomousDatabase,
		"google_oracle_database_autonomous_databases":                  dataGoogleOracleDatabaseAutonomousDatabases,
		"google_oracle_database_cloud_exadata_infrastructure":          dataGoogleOracleDatabaseCloudExadataInfrastructure,
		"google_oracle_database_cloud_exadata_infrastructures":         dataGoogleOracleDatabaseCloudExadataInfrastructures,
		"google_oracle_database_cloud_vm_cluster":                      dataGoogleOracleDatabaseCloudVMCluster,
		"google_oracle_database_cloud_vm_clusters":                     dataGoogleOracleDatabaseCloudVMClusters,
		"google_oracle_database_db_nodes":                              dataGoogleOracleDatabaseDBNodes,
		"google_oracle_database_db_servers":                            dataGoogleOracleDatabaseDBServers,
		"google_organization":                                          placeholder,
		"google_organization_iam_policy":                               dataGoogleOrganizationIamPolicy,
		"google_parameter_manager_parameter":                           dataGoogleParameterManagerParameter,
		"google_parameter_manager_parameter_version":                   dataGoogleParameterManagerParameterVersion,
		"google_parameter_manager_parameter_version_render":            dataGoogleParameterManagerParameterVersionRender,
		"google_parameter_manager_parameters":                          dataGoogleParameterManagerParameters,
		"google_parameter_manager_regional_parameter":                  dataGoogleParameterManagerRegionalParameter,
		"google_parameter_manager_regional_parameter_version":          dataGoogleParameterManagerRegionalParameterVersion,
		"google_parameter_manager_regional_parameter_version_render":   dataGoogleParameterManagerRegionalParameterVersionRender,
		"google_parameter_manager_regional_parameters":                 dataGoogleParameterManagerRegionalParameters,
		"google_privateca_ca_pool_iam_policy":                          dataGooglePrivatecaCaPoolIamPolicy,
		"google_privateca_certificate_authority":                       dataGooglePrivatecaCertificateAuthority,
		"google_privateca_certificate_template_iam_policy":             dataGooglePrivatecaCertificateTemplateIamPolicy,
		"google_privileged_access_manager_entitlement":                 dataGooglePrivilegedAccessManagerEntitlement,
		"google_project":                                               dataGoogleProject,
		"google_project_iam_custom_role":                               dataGoogleProjectIamCustomRole,
		"google_project_iam_custom_roles":                              dataGoogleProjectIamCustomRoles,
		"google_project_iam_policy":                                    dataGoogleProjectIamPolicy,
		"google_project_organization_policy":                           dataGoogleProjectOrganizationPolicy,
		"google_project_service":                                       dataGoogleProjectService,
		"google_projects":                                              placeholder,
		"google_pubsub_schema_iam_policy":                              placeholder,
		"google_pubsub_subscription":                                   dataGooglePubsubSubscription,
		"google_pubsub_subscription_iam_policy":                        dataGooglePubsubSubscriptionIamPolicy,
		"google_pubsub_topic":                                          dataGooglePubsubTopic,
		"google_pubsub_topic_iam_policy":                               dataGooglePubsubTopicIamPolicy,
		"google_redis_instance":                                        dataGoogleRedisInstance,
		"google_runtimeconfig_config_iam_policy":                       dataGoogleRuntimeconfigConfigIamPolicy,
		"google_scc_source_iam_policy":                                 dataGoogleSccSourceIamPolicy,
		"google_scc_v2_organization_source_iam_policy":                 dataGoogleSccV2OrganizationSourceIamPolicy,
		"google_secret_manager_regional_secret":                        dataGoogleSecretManagerRegionalSecret,
		"google_secret_manager_regional_secret_iam_policy":             dataGoogleSecretManagerRegionalSecretIamPolicy,
		"google_secret_manager_regional_secret_version":                dataGoogleSecretManagerRegionalSecretVersion,
		"google_secret_manager_regional_secret_version_access":         dataGoogleSecretManagerRegionalSecretVersionAccess,
		"google_secret_manager_regional_secrets":                       dataGoogleSecretManagerRegionalSecrets,
		"google_secret_manager_secret":                                 dataGoogleSecretManagerSecret,
		"google_secret_manager_secret_iam_policy":                      dataGoogleSecretManagerSecretIamPolicy,
		"google_secret_manager_secret_version":                         dataGoogleSecretManagerSecretVersion,
		"google_secret_manager_secret_version_access":                  dataGoogleManagerSecretVersionAccess,
		"google_secret_manager_secrets":                                dataGoogleSecretManagerSecrets,
		"google_secure_source_manager_instance_iam_policy":             dataGoogleSecureSourceManagerInstanceIamPolicy,
		"google_secure_source_manager_repository_iam_policy":           dataGoogleSecureSourceManagerRepositoryIamPolicy,
		"google_service_account":                                       dataGoogleServiceAccount,
		"google_service_account_access_token":                          dataGoogleServiceAccountAccessToken,
		"google_service_account_iam_policy":                            dataGoogleServiceAccountIamPolicy,
		"google_service_account_id_token":                              placeholder,
		"google_service_account_jwt":                                   dataGoogleServiceAccountJwt,
		"google_service_account_key":                                   dataGoogleServiceAccountKey,
		"google_service_accounts":                                      dataGoogleServiceAccounts,
		"google_service_directory_namespace_iam_policy":                dataGoogleServiceDirectoryNamespaceIamPolicy,
		"google_service_directory_service_iam_policy":                  dataGoogleServiceDirectoryServiceIamPolicy,
		"google_sourcerepo_repository":                                 dataGoogleSourcerepoRepository,
		"google_sourcerepo_repository_iam_policy":                      dataGoogleSourcerepoRepositoryIamPolicy,
		"google_spanner_database":                                      dataGoogleSpannerDatabase,
		"google_spanner_database_iam_policy":                           dataGoogleSpannerDatabaseIamPolicy,
		"google_spanner_instance":                                      dataGoogleSpannerInstance,
		"google_spanner_instance_iam_policy":                           dataGoogleSpannerInstanceIamPolicy,
		"google_sql_backup_run":                                        dataGoogleSQLBackupRun,
		"google_sql_ca_certs":                                          placeholder,
		"google_sql_database":                                          dataGoogleSQLDatabase,
		"google_sql_database_instance":                                 dataGoogleSQLDatabaseInstance,
		"google_sql_database_instance_latest_recovery_time":            dataGoogleSQLDatabaseInstanceLatestRecoveryTime,
		"google_sql_database_instances":                                dataGoogleSQLDatabaseInstances,
		"google_sql_databases":                                         dataGoogleSQLDatabases,
		"google_sql_tiers":                                             placeholder,
		"google_storage_bucket":                                        dataGoogleStorageBucket,
		"google_storage_bucket_iam_policy":                             dataGoogleStorageBucketIamPolicy,
		"google_storage_bucket_object":                                 dataGoogleStorageBucketObject,
		"google_storage_bucket_object_content":                         dataGoogleStorageBucketObjectContent,
		"google_storage_bucket_objects":                                dataGoogleStorageBucketObjects,
		"google_storage_buckets":                                       dataGoogleStorageBuckets,
		"google_storage_object_signed_url":                             placeholder,
		"google_storage_project_service_account":                       dataGoogleStorageProjectServiceAccount,
		"google_storage_transfer_project_service_account":              dataGoogleStorageTransferProjectServiceAccount,
		"google_tags_tag_key":                                          dataGoogleTagsTagKey,
		"google_tags_tag_key_iam_policy":                               dataGoogleTagsTagKeyIamPolicy,
		"google_tags_tag_value_iam_policy":                             dataGoogleTagsTagValueIamPolicy,
		"google_tpu_tensorflow_versions":                               dataGoogleTpuTensorflowVersions,
		"google_tpu_v2_accelerator_types":                              dataGoogleTpuV2AcceleratorTypes,
		"google_tpu_v2_runtime_versions":                               dataGoogleTpuV2RuntimeVersions,
		"google_vertex_ai_endpoint_iam_policy":                         dataGoogleVertexAiEndpointIamPolicy,
		"google_vertex_ai_feature_group_iam_policy":                    dataGoogleVertexAiFeatureGroupIamPolicy,
		"google_vertex_ai_feature_online_store_featureview_iam_policy": dataGoogleVertexAiFeatureOnlineStoreFeatureviewIamPolicy,
		"google_vertex_ai_feature_online_store_iam_policy":             dataGoogleVertexAiFeatureOnlineStoreIamPolicy,
		"google_vertex_ai_featurestore_entitytype_iam_policy":          dataGoogleVertexAiFeaturestoreEntitytypeIamPolicy,
		"google_vertex_ai_featurestore_iam_policy":                     dataGoogleVertexAiFeaturestoreIamPolicy,
		"google_vmwareengine_cluster":                                  placeholder,
		"google_vmwareengine_external_access_rule":                     placeholder,
		"google_vmwareengine_external_address":                         dataGoogleVmwareengineExternalAddress,
		"google_vmwareengine_network":                                  dataGoogleVmwareengineNetwork,
		"google_vmwareengine_network_peering":                          dataGoogleVmwareengineNetworkPeering,
		"google_vmwareengine_network_policy":                           dataGoogleVmwareengineNetworkPolicy,
		"google_vmwareengine_nsx_credentials":                          dataGoogleVmwareengineNsxCredentials,
		"google_vmwareengine_private_cloud":                            dataGoogleVmwareenginePrivateCloud,
		"google_vmwareengine_subnet":                                   dataGoogleVmwareengineSubnet,
		"google_vmwareengine_vcenter_credentials":                      dataGoogleVmwareengineVcenterCredentials,
		"google_vpc_access_connector":                                  dataGoogleVpcAccessConnector,
		"google_workbench_instance_iam_policy":                         dataGoogleWorkbenchInstanceIamPolicy,
		"google_workstations_workstation_config_iam_policy":            dataGoogleWorkstationsWorkstationConfigIamPolicy,
		"google_workstations_workstation_iam_policy":                   dataGoogleWorkstationsWorkstationIamPolicy,
		"google_tags_tag_keys":                                         dataGoogleTagsTagKeys,
		"google_tags_tag_value":                                        dataGoogleTagsTagValue,
		"google_tags_tag_values":                                       dataGoogleTagsTagValues,
		"google_storage_control_folder_intelligence_config":            dataGoogleStorageControlFolderIntelligenceConfig,
		"google_storage_control_organization_intelligence_config":      dataGoogleStorageControlOrganizationIntelligenceConfig,
		"google_storage_control_project_intelligence_config":           dataGoogleStorageControlProjectIntelligenceConfig,
		"google_site_verification_token":                               placeholder,
		"google_runtimeconfig_config":                                  dataGoogleRuntimeconfigConfig,
		"google_runtimeconfig_variable":                                dataGoogleRuntimeconfigVariable,
		"google_redis_cluster":                                         dataGoogleRedisCluster,
		"google_project_ancestry":                                      dataGoogleProjectAncestry,
		"google_access_context_manager_access_policy":                  dataGoogleAccessContextManagerAccessPolicy,
		"google_cloud_identity_group_transitive_memberships":           placeholder,
		"google_dataplex_data_quality_rules":                           dataGoogleDataplexDataQualityRules,
		"google_firebase_android_app_config":                           placeholder,
		"google_firebase_apple_app_config":                             placeholder,
		"google_firebase_web_app_config":                               placeholder,
		"google_organizations":                                         placeholder,
		"google_organization_iam_custom_role":                          dataGoogleOrganizationIamCustomRole,
		"google_organization_iam_custom_roles":                         dataGoogleOrganizationIamCustomRoles,
		"google_compute_network_attachment":                            dataGoogleComputeNetworkAttachment,
	}

	return TFLookup[result]
}
