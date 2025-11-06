package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/google/data/iam/google_service_account.json
var dataGoogleServiceAccount []byte

//go:embed mapping/google/data/compute/google_compute_network.json
var dataGoogleComputeNetwork []byte

//go:embed mapping/google/data/compute/google_compute_subnetwork.json
var dataGoogleComputeSubnetwork []byte

//go:embed mapping/google/data/compute/google_compute_zones.json
var dataGoogleComputeZones []byte

//go:embed mapping/google/data/resourcemanager/google_project.json
var dataGoogleProject []byte

//go:embed mapping/google/data/cloudkms/google_kms_key_ring.json
var dataGoogleKmsKeyRing []byte

//go:embed mapping/google/data/cloudkms/google_kms_crypto_key.json
var dataGoogleKmsCryptoKey []byte

//go:embed mapping/google/data/dns/google_dns_keys.json
var dataGoogleDNSKeys []byte

//go:embed mapping/google/data/dns/google_dns_managed_zone.json
var dataGoogleDNSManagedZone []byte

//go:embed mapping/google/data/dns/google_dns_managed_zone_iam_policy.json
var dataGoogleDNSManagedZoneIamPolicy []byte

//go:embed mapping/google/data/dns/google_dns_record_set.json
var dataGoogleDNSRecordSet []byte

//go:embed mapping/google/data/artifactregistry/google_artifact_registry_repository.json
var dataGoogleArtifactRegistryRepository []byte

//go:embed mapping/google/data/artifactregistry/google_artifact_registry_repository_iam_policy.json
var dataGoogleArtifactRegistryRepositoryIamPolicy []byte

//go:embed mapping/google/data/bigquery/google_app_engine_default_service_account.json
var dataGoogleAppEngineDefaultServiceAccount []byte

//go:embed mapping/google/data/bigquery/google_bigquery_datapolicy_data_policy_iam_policy.json
var dataGoogleBigqueryDatapolicyDataPolicyIamPolicy []byte

//go:embed mapping/google/data/bigquery/google_app_engine_default_service_account.json
var dataGoogleBigqueryDefaultServiceAccount []byte

//go:embed mapping/google/data/bigtable/google_bigtable_instance_iam_policy.json
var dataGoogleBigtableInstanceIamPolicy []byte

//go:embed mapping/google/data/analyticshub/google_bigquery_analytics_hub_data_exchange_iam_policy.json
var dataGoogleBigqueryHubDataExchangeIamPolicy []byte

//go:embed mapping/google/data/analyticshub/google_bigquery_analytics_hub_listing_iam_policy.json
var dataGoogleBigqueryAnalyticsHubListingIamPolicy []byte

//go:embed mapping/google/data/cloudkms/google_kms_key_ring_iam_policy.json
var dataGoogleKmsKeyRingIamPolicy []byte

//go:embed mapping/google/data/cloudkms/google_kms_secret.json
var dataGoogleKmsSecret []byte

//go:embed mapping/google/data/cloudkms/google_kms_secret_asymmetric.json
var dataGoogleKmsSecretAsymnetric []byte

//go:embed mapping/google/data/pubsub/google_pubsub_subscription.json
var dataGooglePubsubSubscription []byte

//go:embed mapping/google/data/pubsub/google_pubsub_subscription_iam_policy.json
var dataGooglePubsubSubscriptionIamPolicy []byte

//go:embed mapping/google/data/pubsub/google_pubsub_topic.json
var dataGooglePubsubTopic []byte

//go:embed mapping/google/data/pubsub/google_pubsub_topic_iam_policy.json
var dataGooglePubsubTopicIamPolicy []byte

//go:embed mapping/google/data/spanner/google_spanner_database_iam_policy.json
var dataGoogleSpannerDatabaseIamPolicy []byte

//go:embed mapping/google/data/spanner/google_spanner_instance.json
var dataGoogleSpannerInstance []byte

//go:embed mapping/google/data/spanner/google_spanner_instance_iam_policy.json
var dataGoogleSpannerInstanceIamPolicy []byte

//go:embed mapping/google/data/storage/google_storage_bucket.json
var dataGoogleStorageBucket []byte

//go:embed mapping/google/data/storage/google_storage_bucket_iam_policy.json
var dataGoogleStorageBucketIamPolicy []byte

//go:embed mapping/google/data/storage/google_storage_bucket_object.json
var dataGoogleStorageBucketObject []byte

//go:embed mapping/google/data/storage/google_storage_bucket_object_content.json
var dataGoogleStorageBucketObjectContent []byte

//go:embed mapping/google/data/resourcemanager/google_storage_project_service_account.json
var dataGoogleStorageProjectServiceAccount []byte

//go:embed mapping/google/data/storagetransfer/google_storage_transfer_project_service_account.json
var dataGoogleStorageTransferProjectServiceAccount []byte

//go:embed mapping/google/data/aiplatform/google_vertex_ai_featurestore_entitytype_iam_policy.json
var dataGoogleVertexAiFeaturestoreEntitytypeIamPolicy []byte

//go:embed mapping/google/data/aiplatform/google_vertex_ai_featurestore_iam_policy.json
var dataGoogleVertexAiFeaturestoreIamPolicy []byte

//go:embed mapping/google/data/cloudfunctions/google_cloudfunctions_function.json
var dataGoogleCloudfunctionsFunction []byte

//go:embed mapping/google/data/cloudfunctions/google_cloudfunctions_function_iam_policy.json
var dataGoogleCloudfunctionsFunctionIamPolicy []byte

//go:embed mapping/google/data/gkebackup/google_gke_backup_backup_plan_iam_policy.json
var dataGoogleGkeBackupBackupPlanIamPolicy []byte

//go:embed mapping/google/data/gkebackup/google_gke_backup_restore_plan_iam_policy.json
var dataGoogleGkeBackupRestorePlanIamPolicy []byte

//go:embed mapping/google/data/gkehub/google_gke_hub_feature_iam_policy.json
var dataGoogleGkeHubFeatureIamPolicy []byte

//go:embed mapping/google/data/gkehub/google_gke_hub_membership_iam_policy.json
var dataGoogleGkeHubMembershipIamPolicy []byte

//go:embed mapping/google/data/gkehub/google_gke_hub_scope_iam_policy.json
var dataGoogleGkeHubScopeIamPolicy []byte

//go:embed mapping/google/data/cloudkms/google_kms_crypto_key_iam_policy.json
var dataGoogleKmsCryptoKeyIamPolicy []byte

//go:embed mapping/google/data/cloudkms/google_kms_crypto_key_version.json
var dataGoogleKmsCryptoKeyVersion []byte

//go:embed mapping/google/data/secretmanager/google_secret_manager_secret.json
var dataGoogleSecretManagerSecret []byte

//go:embed mapping/google/data/secretmanager/google_secret_manager_secret_iam_policy.json
var dataGoogleSecretManagerSecretIamPolicy []byte

//go:embed mapping/google/data/secretmanager/google_secret_manager_secret_version.json
var dataGoogleSecretManagerSecretVersion []byte

//go:embed mapping/google/data/secretmanager/google_secret_manager_secret_version_access.json
var dataGoogleManagerSecretVersionAccess []byte

//go:embed mapping/google/data/resourcemanager/google_project_service.json
var dataGoogleProjectService []byte

//go:embed mapping/google/data/iam/google_service_account_access_token.json
var dataGoogleServiceAccountAccessToken []byte

//go:embed mapping/google/data/iam/google_service_account_iam_policy.json
var dataGoogleServiceAccountIamPolicy []byte

//go:embed mapping/google/data/iam/google_service_account_jwt.json
var dataGoogleServiceAccountJwt []byte

//go:embed mapping/google/data/iam/google_service_account_key.json
var dataGoogleServiceAccountKey []byte

//go:embed mapping/google/data/redis/google_redis_instance.json
var dataGoogleRedisInstance []byte

//go:embed mapping/google/data/alloydb/google_alloydb_locations.json
var dataGoogleAlloydbLocations []byte

//go:embed mapping/google/data/alloydb/google_alloydb_supported_database_flags.json
var dataGoogleAlloydbSupportedDatabaseFlags []byte

//go:embed mapping/google/data/beyondcorp/google_beyondcorp_app_connection.json
var dataGoogleBeyondcorpAppConnection []byte

//go:embed mapping/google/data/beyondcorp/google_beyondcorp_app_connector.json
var dataGoogleBeyondcorpAppConnector []byte

//go:embed mapping/google/data/beyondcorp/google_beyondcorp_app_gateway.json
var dataGoogleBeyondcorpAppGateway []byte

//go:embed mapping/google/data/run/google_cloud_run_locations.json
var dataGoogleCloudRunLocations []byte

//go:embed mapping/google/data/run/google_cloud_run_service.json
var dataGoogleCloudRunService []byte

//go:embed mapping/google/data/run/google_cloud_run_service_iam_policy.json
var dataGoogleCloudRunServiceIamPolicy []byte

//go:embed mapping/google/data/run/google_cloud_run_v2_job.json
var dataGoogleCloudRunV2Job []byte

//go:embed mapping/google/data/run/google_cloud_run_v2_job_iam_policy.json
var dataGoogleCloudRunV2JobIamPolicy []byte

//go:embed mapping/google/data/run/google_cloud_run_v2_service.json
var dataGoogleCloudRunV2Service []byte

//go:embed mapping/google/data/run/google_cloud_run_v2_service_iam_policy.json
var dataGoogleCloudRunV2ServiceIamPolicy []byte

//go:embed mapping/google/data/compute/google_compute_address.json
var dataGoogleComputeAddress []byte

//go:embed mapping/google/data/compute/google_compute_addresses.json
var dataGoogleComputeAddresses []byte

//go:embed mapping/google/data/compute/google_compute_backend_bucket.json
var dataGoogleComputeBackendBucket []byte

//go:embed mapping/google/data/compute/google_compute_backend_bucket_iam_policy.json
var dataGoogleComputeBackendBucketIamPolicy []byte

//go:embed mapping/google/data/compute/google_compute_backend_service.json
var dataGoogleComputeBackendService []byte

//go:embed mapping/google/data/compute/google_compute_backend_service_iam_policy.json
var dataGoogleComputeBackendServiceIamPolicy []byte

//go:embed mapping/google/data/compute/google_compute_default_service_account.json
var dataGoogleComputeDefaultServiceAccount []byte

//go:embed mapping/google/data/compute/google_compute_disk.json
var dataGoogleComputeDisk []byte

//go:embed mapping/google/data/compute/google_compute_disk_iam_policy.json
var dataGoogleComputeDiskIamPolicy []byte

//go:embed mapping/google/data/compute/google_compute_forwarding_rule.json
var dataGoogleComputeForwardingRule []byte

//go:embed mapping/google/data/compute/google_compute_global_address.json
var dataGoogleComputeGlobalAddress []byte

//go:embed mapping/google/data/compute/google_compute_global_forwarding_rule.json
var dataGoogleComputeGlobalForwardingRule []byte

//go:embed mapping/google/data/compute/google_compute_ha_vpn_gateway.json
var dataGoogleComputeHaVpnGateway []byte

//go:embed mapping/google/data/compute/google_compute_health_check.json
var dataGoogleComputeHealthCheck []byte

//go:embed mapping/google/data/compute/google_compute_image_iam_policy.json
var dataGoogleComputeImageIamPolicy []byte

//go:embed mapping/google/data/compute/google_compute_instance.json
var dataGoogleComputeInstance []byte

//go:embed mapping/google/data/compute/google_compute_instance_group.json
var dataGoogleComputeInstanceGroup []byte

//go:embed mapping/google/data/compute/google_compute_instance_group_manager.json
var dataGoogleComputeInstanceGroupManager []byte

//go:embed mapping/google/data/compute/google_compute_instance_serial_port.json
var dataGoogleComputeInstanceSerialPort []byte

//go:embed mapping/google/data/compute/google_compute_instance_template.json
var dataGoogleComputeInstanceTemplate []byte

//go:embed mapping/google/data/compute/google_compute_machine_image_iam_policy.json
var dataGoogleComputeMachineImageIamPolicy []byte

//go:embed mapping/google/data/compute/google_compute_network_endpoint_group.json
var dataGoogleComputeNetworkEndpointGroup []byte

//go:embed mapping/google/data/compute/google_compute_networks.json
var dataGoogleComputeNetworks []byte

//go:embed mapping/google/data/compute/google_compute_node_types.json
var dataGoogleComputeNodeTypes []byte

//go:embed mapping/google/data/compute/google_compute_region_backend_service_iam_policy.json
var dataGoogleComputeRegionBackendServiceIamPolicy []byte

//go:embed mapping/google/data/compute/google_compute_region_disk_iam_policy.json
var dataGoogleComputeRegionDiskIamPolicy []byte

//go:embed mapping/google/data/compute/google_compute_region_instance_group.json
var dataGoogleComputeRegionInstanceGroup []byte

//go:embed mapping/google/data/compute/google_compute_region_instance_template.json
var dataGoogleComputeRegionInstanceTemplate []byte

//go:embed mapping/google/data/compute/google_compute_region_network_endpoint_group.json
var dataGoogleComputeRegionNetworkEndpointGroup []byte

//go:embed mapping/google/data/compute/google_compute_regions.json
var dataGoogleComputeRegions []byte

//go:embed mapping/google/data/compute/google_compute_resource_policy.json
var dataGoogleComputeResourcePolicy []byte

//go:embed mapping/google/data/compute/google_compute_router.json
var dataGoogleComputeRouter []byte

//go:embed mapping/google/data/compute/google_compute_router_nat.json
var dataGoogleComputeRouterNat []byte

//go:embed mapping/google/data/compute/google_compute_router_status.json
var dataGoogleComputeRouterStatus []byte

//go:embed mapping/google/data/compute/google_compute_snapshot.json
var dataGoogleComputeSnapshot []byte

//go:embed mapping/google/data/compute/google_compute_snapshot_iam_policy.json
var dataGoogleComputeSnapshotIamPolicy []byte

//go:embed mapping/google/data/compute/google_compute_ssl_certificate.json
var dataGoogleComputeSslCertificate []byte

//go:embed mapping/google/data/compute/google_compute_ssl_policy.json
var dataGoogleComputeSslPolicy []byte

//go:embed mapping/google/data/compute/google_compute_subnetwork_iam_policy.json
var dataGoogleComputeSubnetworkIamPolicy []byte

//go:embed mapping/google/data/compute/google_compute_vpn_gateway.json
var dataGoogleComputeVpnGateway []byte

//go:embed mapping/google/data/containeranalysis/google_container_analysis_note_iam_policy.json
var dataGoogleContainerAnalysisNoteIamPolicy []byte

//go:embed mapping/google/data/container/google_container_cluster.json
var dataGoogleContainerCluster []byte

//go:embed mapping/google/data/source/google_sourcerepo_repository.json
var dataGoogleSourcerepoRepository []byte

//go:embed mapping/google/data/source/google_sourcerepo_repository_iam_policy.json
var dataGoogleSourcerepoRepositoryIamPolicy []byte

//go:embed mapping/google/data/cloudsql/google_sql_database.json
var dataGoogleSQLDatabase []byte

//go:embed mapping/google/data/cloudsql/google_sql_database_instance.json
var dataGoogleSQLDatabaseInstance []byte

//go:embed mapping/google/data/cloudsql/google_sql_database_instances.json
var dataGoogleSQLDatabaseInstances []byte

//go:embed mapping/google/data/cloudsql/google_sql_databases.json
var dataGoogleSQLDatabases []byte

//go:embed mapping/google/data/cloudbuild/google_cloudbuild_trigger.json
var dataGoogleCloudbuildTrigger []byte

//go:embed mapping/google/data/cloudbuild/google_cloudbuildv2_connection_iam_policy.json
var dataGoogleCloudbuildv2ConnectionIamPolicy []byte

//go:embed mapping/google/data/monitoring/google_monitoring_istio_canonical_service.json
var dataGoogleMonitoringIstioCanonicalService []byte

//go:embed mapping/google/data/monitoring/google_monitoring_mesh_istio_service.json
var dataGoogleMonitoringMeshIstioService []byte

//go:embed mapping/google/data/monitoring/google_monitoring_notification_channel.json
var dataGoogleMonitoringNotificationChannel []byte

//go:embed mapping/google/data/tpu/google_tpu_tensorflow_versions.json
var dataGoogleTpuTensorflowVersions []byte

//go:embed mapping/google/data/tpu/google_tpu_v2_accelerator_types.json
var dataGoogleTpuV2AcceleratorTypes []byte

//go:embed mapping/google/data/tpu/google_tpu_v2_runtime_versions.json
var dataGoogleTpuV2RuntimeVersions []byte

//go:embed mapping/google/data/vmwareengine/google_vmwareengine_network.json
var dataGoogleVmwareengineNetwork []byte

//go:embed mapping/google/data/vmwareengine/google_vmwareengine_private_cloud.json
var dataGoogleVmwareenginePrivateCloud []byte

//go:embed mapping/google/data/vpcaccess/google_vpc_access_connector.json
var dataGoogleVpcAccessConnector []byte

//go:embed mapping/google/data/workstations/google_workstations_workstation_config_iam_policy.json
var dataGoogleWorkstationsWorkstationConfigIamPolicy []byte

//go:embed mapping/google/data/workstations/google_workstations_workstation_iam_policy.json
var dataGoogleWorkstationsWorkstationIamPolicy []byte

//go:embed mapping/google/data/notebooks/google_notebooks_instance_iam_policy.json
var dataGoogleNotebooksInstanceIamPolicy []byte

//go:embed mapping/google/data/notebooks/google_notebooks_runtime_iam_policy.json
var dataGoogleNotebooksRuntimeIamPolicy []byte

//go:embed mapping/google/data/secretmanager/google_secret_manager_secrets.json
var dataGoogleSecretManagerSecrets []byte

//go:embed mapping/google/data/resourcemanager/google_tags_tag_key.json
var dataGoogleTagsTagKey []byte

//go:embed mapping/google/data/monitoring/google_monitoring_cluster_istio_service.json
var dataGoogleMonitoringClusterIstioService []byte

//go:embed mapping/google/data/monitoring/google_monitoring_app_engine_service.json
var dataGoogleMonitoringAppEngineService []byte

//go:embed mapping/google/data/logging/google_logging_project_cmek_settings.json
var dataGoogleLoggingProjectCmekSettings []byte

//go:embed mapping/google/data/iap/google_iap_web_type_compute_iam_policy.json
var dataGoogleIapWebTypeComputeIamPolicy []byte

//go:embed mapping/google/data/iap/google_iap_web_type_app_engine_iam_policy.json
var dataGoogleIapWebTypeAppEngineIamPolicy []byte

//go:embed mapping/google/data/iap/google_iap_web_region_backend_service_iam_policy.json
var dataGoogleIapWebRegionBackendServiceIamPolicy []byte

//go:embed mapping/google/data/iap/google_iap_web_iam_policy.json
var dataGoogleIapWebIamPolicy []byte

//go:embed mapping/google/data/iap/google_iap_web_backend_service_iam_policy.json
var dataGoogleIapWebBackendServiceIamPolicy []byte

//go:embed mapping/google/data/iap/google_iap_tunnel_instance_iam_policy.json
var dataGoogleIapTunnelInstanceIamPolicy []byte

//go:embed mapping/google/data/iap/google_iap_tunnel_iam_policy.json
var dataGoogleIapTunnelIamPolicy []byte

//go:embed mapping/google/data/binaryauthorization/google_binary_authorization_attestor_iam_policy.json
var dataGoogleBinaryAuthorizationAttestorIamPolicy []byte

//go:embed mapping/google/data/certificatemanager/google_certificate_manager_certificate_map.json
var dataGoogleCertificateManagerCertificateMap []byte

//go:embed mapping/google/data/iap/google_iap_app_engine_version_iam_policy.json
var dataGoogleIapAppEngineVersionIamPolicy []byte

//go:embed mapping/google/data/iap/google_iap_app_engine_service_iam_policy.json
var dataGoogleIapAppEngineServiceIamPolicy []byte

//go:embed mapping/google/data/iam.googleapis.com/google_iam_workload_identity_pool_provider.json
var dataGoogleIamWorkloadIdentityPoolProvider []byte

//go:embed mapping/google/data/iam.googleapis.com/google_iam_workload_identity_pool.json
var dataGoogleIamWorkloadIdentityPool []byte

//go:embed mapping/google/data/dataplex/google_dataplex_asset_iam_policy.json
var dataGoogleDataplexAssetIamPolicy []byte

//go:embed mapping/google/data/dataplex/google_dataplex_datascan_iam_policy.json
var dataGoogleDataplexDatascanIamPolicy []byte

//go:embed mapping/google/data/datastream/google_datastream_static_ips.json
var dataGoogleDataStreamStaticIps []byte

//go:embed mapping/google/data/healthcare/google_healthcare_hl7_v2_store_iam_policy.json
var dataGoogleHealthcareHl7V2StoreIamPolicy []byte

//go:embed mapping/google/data/healthcare/google_healthcare_fhir_store_iam_policy.json
var dataGoogleHealthcareFhirStoreIamPolicy []byte

//go:embed mapping/google/data/healthcare/google_healthcare_dicom_store_iam_policy.json
var dataGoogleHealthcareDicomStoreIamPolicy []byte

//go:embed mapping/google/data/healthcare/google_healthcare_dataset_iam_policy.json
var dataGoogleHealthcareDatasetStoreIamPolicy []byte

//go:embed mapping/google/data/healthcare/google_healthcare_consent_store_iam_policy.json
var dataGoogleHealthcareConsentStoreIamPolicy []byte

//go:embed mapping/google/data/gkemulticloud/google_container_attached_versions.json
var dataGoogleContainerAttachedVersions []byte

//go:embed mapping/google/data/datafusion/google_data_fusion_instance_iam_policy.json
var dataGoogleDataFusionInstanceIamPolicy []byte

//go:embed mapping/google/data/cloudtasks/google_cloud_tasks_queue_iam_policy.json
var dataGoogleCloudTasksQueueIamPolicy []byte

//go:embed mapping/google/data/composer/google_composer_environment.json
var dataGoogleComposerEnvironment []byte

//go:embed mapping/google/data/composer/google_composer_image_versions.json
var dataGoogleComposerImageVersions []byte

//go:embed mapping/google/data/dataplex/google_dataplex_lake_iam_policy.json
var dataGoogleDataplexLakeIamPolicy []byte

//go:embed mapping/google/data/dataplex/google_dataplex_task_iam_policy.json
var dataGoogleDataplexTaskIamPolicy []byte

//go:embed mapping/google/data/dataplex/google_dataplex_zone_iam_policy.json
var dataGoogleDataplexZoneIamPolicy []byte

//go:embed mapping/google/data/metastore/google_dataproc_metastore_service_iam_policy.json
var dataGoogleDataprocMetastoreServiceIamPolicy []byte

//go:embed mapping/google/data/metastore/google_dataproc_metastore_federation_iam_policy.json
var dataGoogleDataprocMetaStoreFederationIamPolicy []byte

//go:embed mapping/google/data/metastore/google_dataproc_metastore_service.json
var dataGoogleDataprocMetastoreService []byte

//go:embed mapping/google/data/vmwareengine/google_vmwareengine_network_policy.json
var dataGoogleVmwareengineNetworkPolicy []byte

//go:embed mapping/google/data/vmwareengine/google_vmwareengine_network_policy.json
var dataGoogleVmwareengineNetworkPeering []byte

//go:embed mapping/google/data/aiplatform/google_vertex_ai_endpoint_iam_policy.json
var dataGoogleVertexAiEndpointIamPolicy []byte

//go:embed mapping/google/data/vmwareengine/google_vmwareengine_external_address.json
var dataGoogleVmwareengineExternalAddress []byte

//go:embed mapping/google/data/vmwareengine/google_vmwareengine_nsx_credentials.json
var dataGoogleVmwareengineNsxCredentials []byte

//go:embed mapping/google/data/vmwareengine/google_vmwareengine_subnet.json
var dataGoogleVmwareengineSubnet []byte

//go:embed mapping/google/data/vmwareengine/google_vmwareengine_vcenter_credentials.json
var dataGoogleVmwareengineVcenterCredentials []byte

//go:embed mapping/google/data/notebooks/google_workbench_instance_iam_policy.json
var dataGoogleWorkbenchInstanceIamPolicy []byte

//go:embed mapping/google/data/compute/google_compute_region_disk.json
var dataGoogleComputeRegionDisk []byte

//go:embed mapping/google/data/compute/google_compute_reservation.json
var dataGoogleComputeReservation []byte

//go:embed mapping/google/data/file/google_filestore_instance.json
var dataGoogleFilestoreInstance []byte

//go:embed mapping/google/data/logging/google_logging_project_settings.json
var dataGoogleLoggingProjectSettings []byte

//go:embed mapping/google/data/networksecurity/google_network_security_address_group_iam_policy.json
var dataGoogleSecurityAddressGroupIamPolicy []byte

//go:embed mapping/google/data/servicedirectory/google_service_directory_namespace_iam_policy.json
var dataGoogleServiceDirectoryNamespaceIamPolicy []byte

//go:embed mapping/google/data/servicedirectory/google_service_directory_service_iam_policy.json
var dataGoogleServiceDirectoryServiceIamPolicy []byte

//go:embed mapping/google/data/cloudsql/google_sql_backup_run.json
var dataGoogleSQLBackupRun []byte

//go:embed mapping/google/data/cloudsql/google_sql_database_instance_latest_recovery_time.json
var dataGoogleSQLDatabaseInstanceLatestRecoveryTime []byte

//go:embed mapping/google/data/datacatalog/google_data_catalog_taxonomy_iam_policy.json
var dataGoogleDataCatalogTaxonomyIamPolicy []byte

//go:embed mapping/google/data/dataform/google_dataform_repository_iam_policy.json
var dataGoogleDataformRepositoryIamPolicy []byte

//go:embed mapping/google/data/servicemanagement/google_endpoints_service_iam_policy.json
var dataGoogleEndpointsServiceIamPolicy []byte

//go:embed mapping/google/data/secretmanager/google_secret_manager_regional_secret.json
var dataGoogleSecretManagerRegionalSecret []byte

//go:embed mapping/google/data/secretmanager/google_secret_manager_regional_secret_iam_policy.json
var dataGoogleSecretManagerRegionalSecretIamPolicy []byte

//go:embed mapping/google/data/secretmanager/google_secret_manager_regional_secret_version.json
var dataGoogleSecretManagerRegionalSecretVersion []byte

//go:embed mapping/google/data/secretmanager/google_secret_manager_regional_secret_version_access.json
var dataGoogleSecretManagerRegionalSecretVersionAccess []byte

//go:embed mapping/google/data/secretmanager/google_secret_manager_secrets.json
var dataGoogleSecretManagerRegionalSecrets []byte

//go:embed mapping/google/data/iam/google_service_accounts.json
var dataGoogleServiceAccounts []byte

//go:embed mapping/google/data/spanner/google_spanner_database.json
var dataGoogleSpannerDatabase []byte

//go:embed mapping/google/data/storage/google_storage_bucket_objects.json
var dataGoogleStorageBucketObjects []byte

//go:embed mapping/google/data/storage/google_storage_buckets.json
var dataGoogleStorageBuckets []byte

//go:embed mapping/google/data/privateca/google_privateca_ca_pool_iam_policy.json
var dataGooglePrivatecaCaPoolIamPolicy []byte

//go:embed mapping/google/data/privateca/google_privateca_certificate_authority.json
var dataGooglePrivatecaCertificateAuthority []byte

//go:embed mapping/google/data/privateca/google_privateca_certificate_template_iam_policy.json
var dataGooglePrivatecaCertificateTemplateIamPolicy []byte

//go:embed mapping/google/data/iam/google_project_iam_custom_role.json
var dataGoogleProjectIamCustomRole []byte

//go:embed mapping/google/data/iam/google_project_iam_custom_roles.json
var dataGoogleProjectIamCustomRoles []byte

//go:embed mapping/google/data/privilegedaccessmanager/google_privileged_access_manager_entitlement.json
var dataGooglePrivilegedAccessManagerEntitlement []byte

//go:embed mapping/google/data/artifactregistry/google_artifact_registry_docker_image.json
var dataGoogleArtifactRegistryDockerImage []byte

//go:embed mapping/google/data/artifactregistry/google_artifact_registry_locations.json
var dataGoogleArtifactRegistryLocations []byte

//go:embed mapping/google/data/certificatemanager/google_certificate_manager_certificates.json
var dataGoogleCertificateManagerCertificates []byte

//go:embed mapping/google/data/composer/google_composer_user_workloads_config_map.json
var dataGoogleComposerUserWorkloadsConfigMap []byte

//go:embed mapping/google/data/composer/google_composer_user_workloads_secret.json
var dataGoogleComposerUserWorkloadsSecret []byte

//go:embed mapping/google/data/iam.googleapis.com/google_iam_workload_identity_pool_iam_policy.json
var dataGoogleIamWorkloadIdentityPoolIamPolicy []byte

//go:embed mapping/google/data/cloudkms/google_kms_key_handle.json
var dataGoogleKmsKeyHandle []byte

//go:embed mapping/google/data/cloudkms/google_kms_key_handles.json
var dataGoogleKmsKeyHandles []byte

//go:embed mapping/google/data/cloudkms/google_kms_key_rings.json
var dataGoogleKmsKeyRings []byte

//go:embed mapping/google/data/parametermanager/google_parameter_manager_parameter.json
var dataGoogleParameterManagerParameter []byte

//go:embed mapping/google/data/parametermanager/google_parameter_manager_parameter_version.json
var dataGoogleParameterManagerParameterVersion []byte

//go:embed mapping/google/data/parametermanager/google_parameter_manager_parameter_version_render.json
var dataGoogleParameterManagerParameterVersionRender []byte

//go:embed mapping/google/data/parametermanager/google_parameter_manager_parameters.json
var dataGoogleParameterManagerParameters []byte

//go:embed mapping/google/data/parametermanager/google_parameter_manager_regional_parameter.json
var dataGoogleParameterManagerRegionalParameter []byte

//go:embed mapping/google/data/parametermanager/google_parameter_manager_regional_parameter_version.json
var dataGoogleParameterManagerRegionalParameterVersion []byte

//go:embed mapping/google/data/parametermanager/google_parameter_manager_regional_parameter_version_render.json
var dataGoogleParameterManagerRegionalParameterVersionRender []byte

//go:embed mapping/google/data/parametermanager/google_parameter_manager_regional_parameters.json
var dataGoogleParameterManagerRegionalParameters []byte

//go:embed mapping/google/data/cloudkms/google_kms_ekm_connection_iam_policy.json
var dataGoogleEkmConnectionIamPolicy []byte

//go:embed mapping/google/data/oracledatabase/google_oracle_database_autonomous_database.json
var dataGoogleOracleDatabaseAutonomousDatabase []byte

//go:embed mapping/google/data/oracledatabase/google_oracle_database_autonomous_databases.json
var dataGoogleOracleDatabaseAutonomousDatabases []byte

//go:embed mapping/google/data/oracledatabase/google_oracle_database_cloud_exadata_infrastructure.json
var dataGoogleOracleDatabaseCloudExadataInfrastructure []byte

//go:embed mapping/google/data/oracledatabase/google_oracle_database_cloud_exadata_infrastructures.json
var dataGoogleOracleDatabaseCloudExadataInfrastructures []byte

//go:embed mapping/google/data/oracledatabase/google_oracle_database_cloud_vm_cluster.json
var dataGoogleOracleDatabaseCloudVMCluster []byte

//go:embed mapping/google/data/oracledatabase/google_oracle_database_cloud_vm_clusters.json
var dataGoogleOracleDatabaseCloudVMClusters []byte

//go:embed mapping/google/data/oracledatabase/google_oracle_database_db_nodes.json
var dataGoogleOracleDatabaseDBNodes []byte

//go:embed mapping/google/data/oracledatabase/google_oracle_database_db_servers.json
var dataGoogleOracleDatabaseDBServers []byte

//go:embed mapping/google/data/compute/google_compute_forwarding_rules.json
var dataGoogleComputeForwardingRules []byte

//go:embed mapping/google/data/compute/google_compute_images.json
var dataGoogleComputeImages []byte

//go:embed mapping/google/data/compute/google_compute_instance_guest_attributes.json
var dataGoogleComputeInstanceGuestAttributes []byte

//go:embed mapping/google/data/compute/google_compute_instance_template_iam_policy.json
var dataGoogleComputeInstanceTemplateIamPolicy []byte

//go:embed mapping/google/data/compute/google_compute_instant_snapshot_iam_policy.json
var dataGoogleComputeInstantSnapshotIamPolicy []byte

//go:embed mapping/google/data/compute/google_compute_machine_types.json
var dataGoogleComputeMachineTypes []byte

//go:embed mapping/google/data/compute/google_compute_region_backend_service.json
var dataGoogleComputeRegionBackendService []byte

//go:embed mapping/google/data/compute/google_compute_region_instance_group_manager.json
var dataGoogleComputeRegionInstanceGroupManager []byte

//go:embed mapping/google/data/compute/google_compute_security_policy.json
var dataGoogleComputeSecurityPolicy []byte

//go:embed mapping/google/data/compute/google_compute_storage_pool_iam_policy.json
var dataGoogleComputeStoragePoolIamPolicy []byte

//go:embed mapping/google/data/compute/google_compute_storage_pool_types.json
var dataGoogleComputeStoragePoolTypes []byte

//go:embed mapping/google/data/compute/google_compute_subnetworks.json
var dataGoogleComputeSubnetworks []byte

//go:embed mapping/google/data/alloydb/google_alloydb_cluster.json
var dataGoogleAlloydbCluster []byte

//go:embed mapping/google/data/alloydb/google_alloydb_instance.json
var dataGoogleAlloydbInstance []byte

//go:embed mapping/google/data/apphub/google_apphub_application.json
var dataGoogleApphubApplication []byte

//go:embed mapping/google/data/apphub/google_apphub_discovered_service.json
var dataGoogleApphubDiscoveredService []byte

//go:embed mapping/google/data/apphub/google_apphub_discovered_workload.json
var dataGoogleApphubDiscoveredWorkload []byte

//go:embed mapping/google/data/backupdr/google_backup_dr_backup.json
var dataGoogleBackupDrBackup []byte

//go:embed mapping/google/data/backupdr/google_backup_dr_backup_plan_association.json
var dataGoogleBackupDrPlanAssociation []byte

//go:embed mapping/google/data/backupdr/google_backup_dr_backup_vault.json
var dataGoogleBackupDrBackupVault []byte

//go:embed mapping/google/data/backupdr/google_backup_dr_data_source.json
var dataGoogleBackupDrDataSource []byte

//go:embed mapping/google/data/beyondcorp/google_beyondcorp_application_iam_policy.json
var dataGoogleBeyondcorpApplicationIamPolicy []byte

//go:embed mapping/google/data/beyondcorp/google_beyondcorp_security_gateway.json
var dataGoogleBeyondcorpSecurityGateway []byte

//go:embed mapping/google/data/beyondcorp/google_beyondcorp_security_gateway_application_iam_policy.json
var dataGoogleBeyondcorpSecurityGatewayApplicationIamPolicy []byte

//go:embed mapping/google/data/beyondcorp/google_beyondcorp_security_gateway_iam_policy.json
var dataGoogleBeyondcorpSecurityGatewayIamPolicy []byte

//go:embed mapping/google/data/run/google_cloud_run_v2_worker_pool.json
var dataGoogleCloudRunV2WorkerPool []byte

//go:embed mapping/google/data/run/google_cloud_run_v2_worker_pool_iam_policy.json
var dataGoogleCloudRunWorkerPoolIamPolicy []byte

//go:embed mapping/google/data/clouddeploy/google_clouddeploy_custom_target_type_iam_policy.json
var dataGoogleClouddeployCustomTargetTypeIamPolicy []byte

//go:embed mapping/google/data/clouddeploy/google_clouddeploy_delivery_pipeline_iam_policy.json
var dataGoogleClouddeployDeliveryPipelineIamPolicy []byte

//go:embed mapping/google/data/clouddeploy/google_clouddeploy_target_iam_policy.json
var dataGoogleClouddeployTargetIamPolicy []byte

//go:embed mapping/google/data/aiplatform/google_colab_runtime_template_iam_policy.json
var dataGoogleColabRuntimeTemplateIamPolicy []byte

//go:embed mapping/google/data/gkemulticloud/google_container_attached_install_manifest.json
var dataGoogleContainerAttachedInstallManifest []byte

//go:embed mapping/google/data/gkemulticloud/google_container_aws_versions.json
var dataGoogleContainerAwsVersions []byte

//go:embed mapping/google/data/gkemulticloud/google_container_azure_versions.json
var dataGoogleContainerAzureVersions []byte

//go:embed mapping/google/data/dataplex/google_dataplex_aspect_type_iam_policy.json
var dataGoogleDataplexAspectTypeIamPolicy []byte

//go:embed mapping/google/data/dataplex/google_dataplex_entry_group_iam_policy.json
var dataGoogleDataplexEntryGroupIamPolicy []byte

//go:embed mapping/google/data/dataplex/google_dataplex_entry_type_iam_policy.json
var dataGoogleDataplexEntryTypeIamPolicy []byte

//go:embed mapping/google/data/dataplex/google_dataplex_glossary_iam_policy.json
var dataGoogleDataplexGlossaryIamPolicy []byte

//go:embed mapping/google/data/metastore/google_dataproc_metastore_database_iam_policy.json
var dataGoogleDataprocMetastoreDatabaseIamPolicy []byte

//go:embed mapping/google/data/metastore/google_dataproc_metastore_table_iam_policy.json
var dataGoogleDataprocMetastoreTableIamPolicy []byte

//go:embed mapping/google/data/resourcemanager/google_folder.json
var dataGoogleFolder []byte

//go:embed mapping/google/data/resourcemanager/google_folder_iam_policy.json
var dataGoogleFolderIamPolicy []byte

//go:embed mapping/google/data/resourcemanager/google_folders.json
var dataGoogleFolders []byte

//go:embed mapping/google/data/cloudaicompanion/google_gemini_repository_group_iam_policy.json
var dataGoogleGeminiRepositoryGroupIamPolicy []byte

//go:embed mapping/google/data/gkehub/google_gke_hub_feature.json
var dataGoogleGkeHubFeature []byte

//go:embed mapping/google/data/gkehub/google_gke_hub_membership.json
var dataGoogleGkeHubMembership []byte

//go:embed mapping/google/data/iap/google_iap_tunnel_dest_group_iam_policy.json
var dataGoogleIapTunnelDestGroupIamPolicy []byte

//go:embed mapping/google/data/iap/google_iap_web_cloud_run_service_iam_policy.json
var dataGoogleIapWebCloudRunServiceIamPolicy []byte

//go:embed mapping/google/data/lustre/google_lustre_instance.json
var dataGoogleLustreInstance []byte

//go:embed mapping/google/data/memcache/google_memcache_instance.json
var dataGoogleMemcacheInstance []byte

//go:embed mapping/google/data/memorystore/google_memorystore_instance.json
var dataGoogleMemorystoreInstance []byte

//go:embed mapping/google/data/orgpolicy/google_project_organization_policy.json
var dataGoogleProjectOrganizationPolicy []byte

//go:embed mapping/google/data/resourcemanager/google_organization_iam_policy.json
var dataGoogleOrganizationIamPolicy []byte

//go:embed mapping/google/data/resourcemanager/google_project_iam_policy.json
var dataGoogleProjectIamPolicy []byte

//go:embed mapping/google/data/runtimeconfig/google_runtimeconfig_config_iam_policy.json
var dataGoogleRuntimeconfigConfigIamPolicy []byte

//go:embed mapping/google/data/securitycenter/google_scc_source_iam_policy.json
var dataGoogleSccSourceIamPolicy []byte

//go:embed mapping/google/data/securitycenter/google_scc_v2_organization_source_iam_policy.json
var dataGoogleSccV2OrganizationSourceIamPolicy []byte

//go:embed mapping/google/data/securesourcemanager/google_secure_source_manager_instance_iam_policy.json
var dataGoogleSecureSourceManagerInstanceIamPolicy []byte

//go:embed mapping/google/data/securesourcemanager/google_secure_source_manager_repository_iam_policy.json
var dataGoogleSecureSourceManagerRepositoryIamPolicy []byte

//go:embed mapping/google/data/resourcemanager/google_tags_tag_key_iam_policy.json
var dataGoogleTagsTagKeyIamPolicy []byte

//go:embed mapping/google/data/aiplatform/google_vertex_ai_feature_group_iam_policy.json
var dataGoogleVertexAiFeatureGroupIamPolicy []byte

//go:embed mapping/google/data/aiplatform/google_vertex_ai_feature_online_store_featureview_iam_policy.json
var dataGoogleVertexAiFeatureOnlineStoreFeatureviewIamPolicy []byte

//go:embed mapping/google/data/tags/google_tags_tag_value_iam_policy.json
var dataGoogleTagsTagValueIamPolicy []byte

//go:embed mapping/google/data/aiplatform/google_vertex_ai_feature_online_store_iam_policy.json
var dataGoogleVertexAiFeatureOnlineStoreIamPolicy []byte

//go:embed mapping/google/data/resourcemanager/google_tags_tag_keys.json
var dataGoogleTagsTagKeys []byte

//go:embed mapping/google/data/resourcemanager/google_tags_tag_value.json
var dataGoogleTagsTagValue []byte

//go:embed mapping/google/data/resourcemanager/google_tags_tag_values.json
var dataGoogleTagsTagValues []byte

//go:embed mapping/google/data/storage/google_storage_control_folder_intelligence_config.json
var dataGoogleStorageControlFolderIntelligenceConfig []byte

//go:embed mapping/google/data/storage/google_storage_control_organization_intelligence_config.json
var dataGoogleStorageControlOrganizationIntelligenceConfig []byte

//go:embed mapping/google/data/storage/google_storage_control_project_intelligence_config.json
var dataGoogleStorageControlProjectIntelligenceConfig []byte

//go:embed mapping/google/resource/runtimeconfig/google_runtimeconfig_config.json
var dataGoogleRuntimeconfigConfig []byte

//go:embed mapping/google/resource/runtimeconfig/google_runtimeconfig_variable.json
var dataGoogleRuntimeconfigVariable []byte

//go:embed mapping/google/data/redis/google_redis_cluster.json
var dataGoogleRedisCluster []byte

//go:embed mapping/google/data/resourcemanager/google_project_ancestry.json
var dataGoogleProjectAncestry []byte

//go:embed mapping/google/data/dns/google_dns_managed_zones.json
var dataGoogleDNSManagedZones []byte

//go:embed mapping/google/data/accesscontextmanager/google_access_context_manager_access_policy.json
var dataGoogleAccessContextManagerAccessPolicy []byte

//go:embed mapping/google/data/dataplex/google_dataplex_data_quality_rules.json
var dataGoogleDataplexDataQualityRules []byte

//go:embed mapping/google/data/iam/google_organization_iam_custom_role.json
var dataGoogleOrganizationIamCustomRole []byte

//go:embed mapping/google/data/iam/google_organization_iam_custom_roles.json
var dataGoogleOrganizationIamCustomRoles []byte

//go:embed mapping/google/data/compute/google_compute_network_attachment.json
var dataGoogleComputeNetworkAttachment []byte

//go:embed mapping/google/data/storageinsights/google_storage_insights_dataset_config.json
var dataGoogleStorageInsightsDatasetConfig []byte

//go:embed mapping/google/data/artifactregistry/google_artifact_registry_docker_images.json
var dataGoogleArtifactRegistryDockerImages []byte

//go:embed mapping/google/data/iam.googleapis.com/google_iam_workforce_pool_iam_policy.json
var dataGoogleIamWorkforcePoolIamPolicy []byte

//go:embed mapping/google/data/artifactregistry/google_artifact_registry_repositories.json
var dataGoogleArtifactRegistryRepositories []byte

//go:embed mapping/google/data/artifactregistry/google_artifact_registry_version.json
var dataGoogleArtifactRegistryVersion []byte

//go:embed mapping/google/data/artifactregistry/google_artifact_registry_versions.json
var dataGoogleArtifactRegistryVersions []byte

//go:embed mapping/google/data/artifactregistry/google_artifact_registry_npm_package.json
var dataGoogleArtifactRegistryNpmPackage []byte

//go:embed mapping/google/data/artifactregistry/google_artifact_registry_tags.json
var dataGoogleArtifactRegistryTags []byte

//go:embed mapping/google/data/iap/google_iap_web_region_forwarding_rule_service_iam_policy.json
var dataGoogleIapWebRegionForwardingRuleServiceIamPolicy []byte

//go:embed mapping/google/data/certificatemanager/google_certificate_manager_dns_authorization.json
var dataGoogleCertificateManagerDnsAuthorization []byte

//go:embed mapping/google/data/artifactregistry/google_artifact_registry_python_package.json
var dataGoogleArtifactRegistryPythonPackage []byte

//go:embed mapping/google/data/backupdr/google_backup_dr_backup_plan_associations.json
var dataGoogleBackupDrPlanAssociations []byte

//go:embed mapping/google/data/backupdr/google_backup_dr_data_source_reference.json
var dataGoogleBackupDrDataSourceReference []byte

//go:embed mapping/google/data/backupdr/google_backup_dr_data_source_references.json
var dataGoogleBackupDrDataSourceReferences []byte

//go:embed mapping/google/data/bigquery/google_bigquery_datapolicyv2_data_policy_iam_policy.json
var dataGoogleBigqueryDatapolicyv2DataPolicyIamPolicy []byte

//go:embed mapping/google/data/iap/google_iap_web_forwarding_rule_service_iam_policy.json
var dataGoogleIapForwardingRuleServiceIamPolicy []byte

//go:embed mapping/google/data/artifactregistry/google_artifact_registry_maven_artifact.json
var dataGoogleArtifactRegistryMavenArtifact []byte

//go:embed mapping/google/data/artifactregistry/google_artifact_registry_maven_artifacts.json
var dataGoogleArtifactRegistryMavenArtifacts []byte

//go:embed mapping/google/data/artifactregistry/google_artifact_registry_npm_packages.json
var dataGoogleArtifactRegistryNpmPackages []byte

//go:embed mapping/google/data/compute/google_compute_interconnect_location.json
var dataGoogleComputeInterconnectLocation []byte

//go:embed mapping/google/data/compute/google_compute_interconnect_locations.json
var dataGoogleComputeInterconnectLocations []byte

//go:embed mapping/google/data/artifactregistry/google_artifact_registry_packages.json
var dataGoogleArtifactRegistryPackages []byte
