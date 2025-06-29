package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/google/resource/compute/google_compute_instance.json
var googleComputeInstance []byte

//go:embed mapping/google/resource/storage/google_storage_bucket.json
var googleStorageBucket []byte

//go:embed mapping/google/resource/storage/google_storage_bucket_object.json
var googleStorageBucketObject []byte

//go:embed mapping/google/resource/compute/google_compute_network.json
var googleComputeNetwork []byte

//go:embed mapping/google/resource/compute/google_compute_subnetwork.json
var googleComputeSubnetwork []byte

//go:embed mapping/google/resource/compute/google_compute_address.json
var googleComputeAddress []byte

//go:embed mapping/google/resource/compute/google_compute_firewall.json
var googleComputeFirewall []byte

//go:embed mapping/google/resource/compute/google_compute_instance_template.json
var googleComputeInstanceTemplate []byte

//go:embed mapping/google/resource/compute/google_compute_project_metadata_item.json
var googleComputeProjectMetadataItem []byte

//go:embed mapping/google/resource/cloudfunctions/google_cloudfunctions_function.json
var googleCloudfunctionsFunction []byte

//go:embed mapping/google/resource/cloudfunctions/google_cloudfunctions_function_iam_policy.json
var googleCloudfunctionsFunctionIamPolicy []byte

//go:embed mapping/google/resource/iam/google_project_iam_custom_role.json
var googleProjectIamCustomRole []byte

//go:embed mapping/google/resource/iam/google_service_account.json
var googleServiceAccount []byte

//go:embed mapping/google/resource/cloudkms/google_kms_key_ring.json
var googleKmsKeyRing []byte

//go:embed mapping/google/resource/cloudkms/google_kms_crypto_key.json
var googleKmsCryptoKey []byte

//go:embed mapping/google/resource/storage/google_storage_bucket_acl.json
var googleStorageBucketACL []byte

//go:embed mapping/google/resource/storage/google_storage_bucket_iam_binding.json
var googleStorageBucketIamBinding []byte

//go:embed mapping/google/resource/cloudsql/google_sql_database_instance.json
var googleSQLDatabaseInstance []byte

//go:embed mapping/google/resource/cloudsql/google_sql_database.json
var googleSQLDatabase []byte

//go:embed mapping/google/resource/cloudsql/google_sql_user.json
var googleSQLUser []byte

//go:embed mapping/google/resource/servicenetworking/google_service_networking_connection.json
var googleServiceNetworkingConnection []byte

//go:embed mapping/google/resource/compute/google_compute_global_address.json
var googleComputeGlobalAddress []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_topic.json
var googlePubsubTopic []byte

//go:embed mapping/google/resource/source/google_soucerepo_repository.json
var googleSourcerepoRepository []byte

//go:embed mapping/google/resource/iam/google_service_account_iam_policy.json
var googleServiceAccountIamPolicy []byte

//go:embed mapping/google/resource/iam/google_service_account_key.json
var googleServiceAccountKey []byte

//go:embed mapping/google/resource/resourcemanager/google_project_iam_binding.json
var googleProjectIamBinding []byte

//go:embed mapping/google/resource/container/google_container_cluster.json
var googleContainerCluster []byte

//go:embed mapping/google/resource/container/google_container_node_pool.json
var googleContainerNodePool []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_dataset.json
var googleBigqueryDataset []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_job.json
var googleBigqueryJob []byte

//go:embed mapping/google/resource/artifactregistry/google_artifact_registry_repository.json
var googleArtifactRegistryRepository []byte

//go:embed mapping/google/resource/artifactregistry/google_artifact_registry_repository_iam_binding.json
var googleArtifactRegistryRepositoryIamBinding []byte

//go:embed mapping/google/resource/artifactregistry/google_artifact_registry_repository_iam_member.json
var googleArtifactRegistryRepositoryIamMember []byte

//go:embed mapping/google/resource/artifactregistry/google_artifact_registry_repository_iam_policy.json
var googleArtifactRegistryRepositoryIamPolicy []byte

//go:embed mapping/google/resource/pubsublite/google_pubsub_lite_reservation.json
var googlePubsubLiteReservation []byte

//go:embed mapping/google/resource/pubsublite/google_pubsub_lite_subscription.json
var googlePubsubLiteSubscription []byte

//go:embed mapping/google/resource/pubsublite/google_pubsub_lite_topic.json
var googlePubsubLiteTopic []byte

//go:embed mapping/google/resource/compute/google_compute_security_policy.json
var googleComputeSecurityPolicy []byte

//go:embed mapping/google/resource/cloudkms/google_kms_crypto_key_iam_binding.json
var googlekmsCryptoKeyIamBinding []byte

//go:embed mapping/google/resource/cloudkms/google_kms_crypto_key_iam_member.json
var googlekmsCryptoKeyIamMember []byte

//go:embed mapping/google/resource/cloudkms/google_kms_crypto_key_iam_policy.json
var googlekmsCryptoKeyIamPolicy []byte

//go:embed mapping/google/resource/dns/google_dns_managed_zone.json
var googleDnsmanagedZone []byte

//go:embed mapping/google/resource/dns/google_dns_policy.json
var googleDNSPolicy []byte

//go:embed mapping/google/resource/dns/google_dns_record_set.json
var googleDNSRecordSet []byte

//go:embed mapping/google/resource/iam/google_service_account_iam_binding.json
var googleServiceAccountIamBinding []byte

//go:embed mapping/google/resource/iam/google_service_account_iam_member.json
var googleServiceAccountIamMember []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_instance.json
var googleBigtableInstance []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_instance.json
var googleComputeRegionSslCertificate []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_table.json
var googleBigtableTable []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_instance_iam.json
var googleBigTableInstanceIam []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_table_iam.json
var googleBigTableTableIam []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_schema.json
var googlePubsubSchema []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_subscription.json
var googlePubsubSubscription []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_topic_iam_binding.json
var googlePubsubTopicIam []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_secret.json
var googleSecretManagerSecret []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_secret_iam_binding.json
var googleSecretManagerSecretIam []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_secret_version.json
var googleSecretManagerSecretVersion []byte

//go:embed mapping/google/resource/redis/google_redis_instance.json
var googleRedisInstance []byte

//go:embed mapping/google/resource/resourcemanager/google_project_service.json
var googleProjectService []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_job.json
var googleCloudRunV2Job []byte

//go:embed mapping/google/resource/cloudscheduler/google_cloud_scheduler_job.json
var googleCloudSchedulerJob []byte

//go:embed mapping/google/resource/storage/google_storage_bucket_access_control.json
var googleStorageBucketAccessControl []byte

//go:embed mapping/google/resource/storage/google_storage_bucket_iam_member.json
var googleStorageBucketIamMember []byte

//go:embed mapping/google/resource/storage/google_storage_bucket_iam_policy.json
var googleStorageBucketIamPolicy []byte

//go:embed mapping/google/resource/storage/google_storage_default_object_access_control.json
var googleStorageDefaultObjectAccessControl []byte

//go:embed mapping/google/resource/storage/google_storage_default_object_acl.json
var googleStorageDefaultObjectACL []byte

//go:embed mapping/google/resource/storage/google_storage_hmac_key.json
var googleStorageHmacKey []byte

//go:embed mapping/google/resource/storage/google_storage_insights_report_config.json
var googleStorageInsightsReportConfig []byte

//go:embed mapping/google/resource/storage/google_storage_object_access_control.json
var googleStorageObjectAccessControl []byte

//go:embed mapping/google/resource/cloudbuild/google_cloudbuild_trigger.json
var googleCloudbuildTrigger []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_endpoint.json
var googleServiceDirectoryEndpoint []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_namespace.json
var googleServiceDirectoryNamespace []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_namespace_iam_binding.json
var googleServiceDirectoryNamespaceIamBinding []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_namespace_iam_member.json
var googleServiceDirectoryNamespaceIamMember []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_namespace_iam_policy.json
var googleServiceDirectoryNamespaceIamPolicy []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_service.json
var googleServiceDirectoryService []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_service_iam_binding.json
var googleServiceDirectoryServiceIamBinding []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_service_iam_member.json
var googleServiceDirectoryServiceIamMember []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_service_iam_policy.json
var googleServiceDirectoryServiceIamPolicy []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_access_level.json
var googleAccessContextManagerAccessLevel []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_access_levels.json
var googleAccessContextManagerAccessLevels []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_access_policy.json
var googleAccessContextManagerAccessPolicy []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_access_policy_iam.json
var googleAccessContextManagerAccessPolicyIam []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_authorized_orgs_desc.json
var googleAccessContextManagerAuthorizedOrgsDesc []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_gcp_user_access_binding.json
var googleAccessContextManagerGcpUserAccessBinding []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_service_perimeter.json
var googleAccessContextManagerServicePerimeter []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_service_perimeters.json
var googleAccessContextManagerServicePerimeters []byte

//go:embed mapping/google/resource/alloydb/google_alloydb_backup.json
var googleAlloydbBackup []byte

//go:embed mapping/google/resource/alloydb/google_alloydb_cluster.json
var googleAlloydbCluster []byte

//go:embed mapping/google/resource/alloydb/google_alloydb_instance.json
var googleAlloydbInstance []byte

//go:embed mapping/google/resource/alloydb/google_alloydb_user.json
var googleAlloydbUser []byte

//go:embed mapping/google/resource/firebase/google_firebase_android_app.json
var googleFirebaseAndroidApp []byte

//go:embed mapping/google/resource/firebase/google_firebase_apple_app.json
var googleFirebaseAppleApp []byte

//go:embed mapping/google/resource/firebasedatabase/google_firebase_database_instance.json
var googleFirebaseDatabaseInstance []byte

//go:embed mapping/google/resource/firebasehosting/google_firebase_hosting_site.json
var googleFirebaseHostingSite []byte

//go:embed mapping/google/resource/firebase/google_firebase_project.json
var googleFirebaseProject []byte

//go:embed mapping/google/resource/firebasestorage/google_firebase_storage_bucket.json
var googleFirebaseStorageBucket []byte

//go:embed mapping/google/resource/firebase/google_firebase_android_app.json
var googleFirebaseWebApp []byte

//go:embed mapping/google/resource/firebaserules/google_firebaserules_release.json
var googleFirebaserulesRelease []byte

//go:embed mapping/google/resource/firebaserules/google_firebaserules_ruleset.json
var googleFirebaserulesRuleset []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_app_profile.json
var googleBigtableAppProfile []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_gc_policy.json
var googleBigtableGcPolicy []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_api.json
var googleAPIGatewayAPI []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_api_config.json
var googleAPIGatewayAPIConfig []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_api_config_iam.json
var googleAPIGatewayAPIConfigIam []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_api_iam.json
var googleAPIGatewayAPIIam []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_gateway.json
var googleAPIGatewayGateway []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_gateway_iam.json
var googleAPIGatewayGatewayIam []byte

//go:embed mapping/google/resource/spanner/google_spanner_database.json
var googleSpannerDatabase []byte

//go:embed mapping/google/resource/spanner/google_spanner_database_iam.json
var googleSpannerDatabaseIam []byte

//go:embed mapping/google/resource/spanner/google_spanner_instance.json
var googleSpannerInstance []byte

//go:embed mapping/google/resource/spanner/google_spanner_instance_iam.json
var googleSpannerInstanceIam []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_service.json
var googleCloudRunV2Service []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_job_iam.json
var googleCloudRunV2JobIam []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_service_iam.json
var googleCloudRunV2ServiceIam []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_dataset.json
var googleVertexAiDataset []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_endpoint.json
var googleVertexAiEndpoint []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_endpoint_iam.json
var googleVertexAiEndpointIam []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_group.json
var googleVertexAiFeatureGroup []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_group_feature.json
var googleVertexAiFeatureGroupFeature []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store.json
var googleVertexAiFeatureOnlineStore []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store_featureview.json
var googleVertexAiFeatureOnlineStoreFeatureview []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_featurestore.json
var googleVertexAiFeaturestore []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_featurestore_entitytype.json
var googleVertexAiFeaturestoreEntitytype []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_featurestore_entitytype_feature.json
var googleVertexAiFeaturestoreEntitytypeFeature []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_featurestore_entitytype_iam.json
var googleVertexAiFeaturestoreEntitytypeIam []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_featurestore_iam.json
var googleVertexAiFeaturestoreIam []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_index.json
var googleVertexAiIndex []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_index_endpoint.json
var googleVertexAiIndexEndpoint []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_metadata_store.json
var googleVertexAiMetadataStore []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_tensorboard.json
var googleVertexAiTensorboard []byte

//go:embed mapping/google/resource/analyticshub/google_bigquery_analytics_hub_data_exchange.json
var googleBigqueryAnalyticsHubDataExchange []byte

//go:embed mapping/google/resource/analyticshub/google_bigquery_analytics_hub_data_exchange_iam.json
var googleBigqueryAnalyticsHubDataExchangeIam []byte

//go:embed mapping/google/resource/analyticshub/google_bigquery_analytics_hub_listing.json
var googleBigqueryAnalyticsHubListing []byte

//go:embed mapping/google/resource/analyticshub/google_bigquery_analytics_hub_listing_iam.json
var googleBigqueryAnalyticsHubListingIam []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_bi_reservation.json
var googleBigqueryBiReservation []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_capacity_commitment.json
var googleBigqueryCapacityCommitment []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_connection.json
var googleBigqueryConnection []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_connection_iam.json
var googleBigqueryConnectionIam []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_data_transfer_config.json
var googleBigqueryDataTransferConfig []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_dataset_iam.json
var googleBigqueryDatasetIam []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_reservation.json
var googleBigqueryReservation []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_reservation_assignment.json
var googleBigqueryReservationAssignment []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_routine.json
var googleBigqueryRoutine []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_table_iam.json
var googleBigqueryTableIam []byte

//go:embed mapping/google/resource/composer/google_composer_environment.json
var googleComposerEnvironment []byte

//go:embed mapping/google/resource/iam/google_iam_workload_identity_pool.json
var googleIamWorkloadIdentityPool []byte

//go:embed mapping/google/resource/iam/google_iam_workload_identity_pool_provider.json
var googleIamWorkloadIdentityPoolProvider []byte

//go:embed mapping/google/resource/iam/google_project_iam_audit_config.json
var googleProjectIamAuditConfig []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_table.json
var googleBigQueryTable []byte

//go:embed mapping/google/resource/backend/gcs.json
var gcsBackend []byte

//go:embed mapping/google/resource/privateca/google_privateca_ca_pool.json
var googlePrivateCaPool []byte

//go:embed mapping/google/resource/privateca/google_privateca_ca_pool_iam_binding.json
var googlePrivateCaPoolIamBinding []byte

//go:embed mapping/google/resource/privateca/google_privateca_ca_pool_iam_member.json
var googlePrivateCaPoolIamMember []byte

//go:embed mapping/google/resource/privateca/google_privateca_ca_pool_iam_policy.json
var googlePrivateCaPoolIamPolicy []byte

//go:embed mapping/google/resource/privateca/google_privateca_certificate_template.json
var googlePrivatecaCertificateTemplate []byte

//go:embed mapping/google/resource/privateca/google_privateca_certificate_template_iam_binding.json
var googlePrivatecaCertificateTemplateIamBinding []byte

//go:embed mapping/google/resource/privateca/google_privateca_certificate_template_iam_member.json
var googlePrivatecaCertificateTemplateIamMember []byte

//go:embed mapping/google/resource/privateca/google_privateca_certificate_template_iam_policy.json
var googlePrivatecaCertificateTemplateIamPolicy []byte

//go:embed mapping/google/resource/privilegedaccessmanager/google_privileged_access_manager_entitlement.json
var googlePrivilegedAccessManagerEntitlement []byte

//go:embed mapping/google/resource/run/google_cloud_run_domain_mapping.json
var googleCloudRunDomainMapping []byte

//go:embed mapping/google/resource/run/google_cloud_run_service.json
var googleCloudRunService []byte

//go:embed mapping/google/resource/run/google_cloud_run_service_iam_binding.json
var googleCloudRunServiceIamBinding []byte

//go:embed mapping/google/resource/run/google_cloud_run_service_iam_member.json
var googleCloudRunServiceIamMember []byte

//go:embed mapping/google/resource/run/google_cloud_run_service_iam_policy.json
var googleCloudRunServiceIamPolicy []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_worker_pool.json
var googleCloudRunV2WorkerPool []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_worker_pool_iam_binding.json
var googleCloudRunV2WorkerPoolIamBinding []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_worker_pool_iam_member.json
var googleCloudRunV2WorkerPoolIamMember []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_worker_pool_iam_policy.json
var googleCloudRunV2WorkerPoolIamPolicy []byte

//go:embed mapping/google/resource/datacatalog/google_bigquery_datapolicy_data_policy.json
var googleBigqueryDatapolicyDataPolicy []byte

//go:embed mapping/google/resource/datacatalog/google_bigquery_datapolicy_data_policy_iam_binding.json
var googleBigqueryDatapolicyDataPolicyIamBinding []byte

//go:embed mapping/google/resource/datacatalog/google_bigquery_datapolicy_data_policy_iam_member.json
var googleBigqueryDatapolicyDataPolicyIamMember []byte

//go:embed mapping/google/resource/datacatalog/google_bigquery_datapolicy_data_policy_iam_policy.json
var googleBigqueryDatapolicyDataPolicyIamPolicy []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_entry.json
var googleDataCatalogEntry []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_entry_group.json
var googleDataCatalogEntryGroup []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_entry_group_iam_member.json
var googleDataCatalogEntryGroupIamMember []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_entry_group_iam_policy.json
var googleDataCatalogEntryGroupIamPolicy []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_policy_tag.json
var googleDataCatalogPolicyTag []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_tag_template_iam_member.json
var googleDataCatalogPolicyTagIamMember []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_tag_template_iam_policy.json
var googleDataCatalogPolicyTagIamPolicy []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_tag.json
var googleDataCatalogTag []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_tag_template.json
var googleDataCatalogTagTemplate []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_tag_template_iam_policy.json
var googleDataCatalogTagTemplateIamPolicy []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_taxonomy.json
var googleDataCatalogTaxonomy []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_policy_tag_iam_policy.json
var googleDataCatalogTaxonomyIamBinding []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_taxonomy_iam_member.json
var googleDataCatalogTaxonomyIamMember []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_taxonomy_iam_policy.json
var googleDataCatalogTaxonomyIamPolicy []byte

//go:embed mapping/google/resource/cloudfunctions/google_cloudfunctions2_function.json
var googleCloudfunction2Function []byte

//go:embed mapping/google/resource/cloudfunctions/google_cloudfunctions2_function_iam_binding.json
var googleCloudfunction2FunctionIamBinding []byte

//go:embed mapping/google/resource/cloudfunctions/google_cloudfunctions2_function_iam_member.json
var googleCloudfunction2FunctionIamMember []byte

//go:embed mapping/google/resource/cloudfunctions/google_cloudfunctions2_function_iam_policy.json
var googleCloudfunction2FunctionIamPolicy []byte

//go:embed mapping/google/resource/cloudfunctions/google_cloudfunctions_function_iam_binding.json
var googleCloudfunctionsFunctionIamBinding []byte

//go:embed mapping/google/resource/cloudkms/google_kms_crypto_key_version.json
var googleKmsCryptoKeyVersions []byte

//go:embed mapping/google/resource/cloudkms/google_kms_key_handle.json
var googleKmsKeyHandle []byte

//go:embed mapping/google/resource/cloudkms/google_kms_key_ring_iam_binding.json
var googleKmsKeyRingIamBinding []byte

//go:embed mapping/google/resource/cloudkms/google_kms_key_ring_iam_member.json
var googleKmsKeyRingIamMember []byte

//go:embed mapping/google/resource/cloudkms/google_kms_key_ring_iam_policy.json
var googleKmsKeyRingIamPolicy []byte

//go:embed mapping/google/resource/cloudkms/google_kms_key_ring_import_job.json
var googleKmsKeyRingImportJob []byte

//go:embed mapping/google/resource/cloudkms/google_kms_secret_ciphertext.json
var googleKmsSecretCiphertext []byte

//go:embed mapping/google/resource/compute/google_project_usage_export_bucket.json
var googleProjectUsageExportBucket []byte

//go:embed mapping/google/resource/iam/google_default_service_accounts.json
var googleDefaultServiceAccounts []byte

//go:embed mapping/google/resource/iam/google_project_default_service_accounts.json
var googleProjectDefaultServiceAccounts []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_regional_secret.json
var googleSecretManagerRegionalSecret []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_regional_secret_iam_binding.json
var googleSecretManagerRegionalSecretsIamBinding []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_regional_secret_iam_member.json
var googleSecretManagerRegionalSecretIamMember []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_regional_secret_iam_policy.json
var googleSecretManagerRegionalSecretIamPolicy []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_regional_secret_version.json
var googleSecretManagerRegionalSecretVersion []byte

//go:embed mapping/google/resource/spanner/google_spanner_backup_schedule.json
var googleSpannerBackupSchedule []byte

//go:embed mapping/google/resource/spanner/google_spanner_instance_config.json
var googleSpannerInstanceConfig []byte

//go:embed mapping/google/resource/spanner/google_spanner_instance_partition.json
var googleSpannerInstancePartition []byte

//go:embed mapping/google/resource/biglake/google_biglake_catalog.json
var googleBiglakeCatalog []byte

//go:embed mapping/google/resource/biglake/google_biglake_database.json
var googleBiglakeDatabase []byte

//go:embed mapping/google/resource/biglake/google_biglake_table.json
var googleBiglakeTable []byte

//go:embed mapping/google/resource/analyticshub/google_bigquery_analytics_hub_listing_subscription.json
var googleBigqueryAnalyticsHubListingSubscription []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_dataset_access.json
var googleBigqueryDatasetAccess []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_row_access_policy.json
var googleBigqueryRowAccessPolicy []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_authorized_view.json
var googleBigtableAuthorizedView []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_logical_view.json
var googleBigtableLogicalView []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_materialized_view.json
var googleBigtableMaterializedView []byte

//go:embed mapping/google/resource/cloudbuild/google_cloudbuildv2_connection.json
var googleCloudbuildv2Connection []byte

//go:embed mapping/google/resource/cloudbuild/google_cloudbuildv2_connection_iam_binding.json
var googleCloudbuildv2ConnectionIamBinding []byte

//go:embed mapping/google/resource/cloudbuild/google_cloudbuildv2_connection_iam_member.json
var googleCloudbuildv2ConnectionIamMember []byte

//go:embed mapping/google/resource/cloudbuild/google_cloudbuildv2_connection_iam_policy.json
var googleCloudbuildv2ConnectionIamPolicy []byte

//go:embed mapping/google/resource/cloudbuild/google_cloudbuildv2_repository.json
var googleCloudbuildv2Repository []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_asset.json
var googleDataplexAsset []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_asset_iam_binding.json
var googleDataplexAssetIamBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_asset_iam_member.json
var googleDataplexAssetIamMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_asset_iam_policy.json
var googleDataplexAssetIamPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_datascan.json
var googleDataplexDatascan []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_datascan_iam_binding.json
var googleDataplexDatascanIamBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_datascan_iam_member.json
var googleDataplexDatascanIamMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_datascan_iam_policy.json
var googleDataplexDatascanIamPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_glossary.json
var googleDataplexGlossary []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_glossary_category.json
var googleDataplexGlossaryCategory []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_glossary_term.json
var googleDataplexGlossaryTerm []byte

//go:embed mapping/google/resource/apigee/google_apigee_environment_iam_binding.json
var googleApigeeEnvironmentIamBinding []byte

//go:embed mapping/google/resource/apigee/google_apigee_environment_iam_member.json
var googleApigeeEnvironmentIamMember []byte

//go:embed mapping/google/resource/apigee/google_apigee_environment_iam_policy.json
var googleApigeeEnvironmentIamPolicy []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_application_iam_binding.json
var googleBeyondcorpApplicationIamBinding []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_application_iam_member.json
var googleBeyondcorpApplicationIamMember []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_application_iam_policy.json
var googleBeyondcorpApplicationIamPolicy []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_security_gateway_application_iam_binding.json
var googleBeyondcorpSecurityGatewayApplicationIamBinding []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_security_gateway_application_iam_member.json
var googleBeyondcorpSecurityGatewayApplicationIamMember []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_security_gateway_application_iam_policy.json
var googleBeyondcorpSecurityGatewayApplicationIamPolicy []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_security_gateway_iam_binding.json
var googleBeyondcorpSecurityGatewayIamBinding []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_security_gateway_iam_member.json
var googleBeyondcorpSecurityGatewayIamMember []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_security_gateway_iam_policy.json
var googleBeyondcorpSecurityGatewayIamPolicy []byte

//go:embed mapping/google/resource/billing/google_billing_account_iam_binding.json
var googleBillingAccountIamBinding []byte

//go:embed mapping/google/resource/billing/google_billing_account_iam_member.json
var googleBillingAccountIamMember []byte

//go:embed mapping/google/resource/billing/google_billing_account_iam_policy.json
var googleBillingAccountIamPolicy []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_schema_iam_binding.json
var googlePubsubSchemaIamBinding []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_schema_iam_member.json
var googlePubsubSchemaIamMember []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_schema_iam_policy.json
var googlePubsubSchemaIamPolicy []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_subscription_iam_binding.json
var googlePubsubSubscriptionIamBinding []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_subscription_iam_member.json
var googlePubsubSubscriptionIamMember []byte

//go:embed mapping/google/resource/securesourcemanager/google_secure_source_manager_instance_iam_binding.json
var googleSecureSourceManagerInstanceIamBinding []byte

//go:embed mapping/google/resource/securesourcemanager/google_secure_source_manager_instance_iam_member.json
var googleSecureSourceManagerInstanceIamMember []byte

//go:embed mapping/google/resource/securesourcemanager/google_secure_source_manager_instance_iam_policy.json
var googleSecureSourceManagerInstanceIamPolicy []byte

//go:embed mapping/google/resource/securesourcemanager/google_secure_source_manager_repository_iam_binding.json
var googleSecureSourceManagerRepositoryIamBinding []byte

//go:embed mapping/google/resource/securesourcemanager/google_secure_source_manager_repository_iam_member.json
var googleSecureSourceManagerRepositoryIamMember []byte

//go:embed mapping/google/resource/securesourcemanager/google_secure_source_manager_repository_iam_policy.json
var googleSecureSourceManagerRepositoryIamPolicy []byte

//go:embed mapping/google/resource/source/google_sourcerepo_repository_iam_binding.json
var googleSourcerepoRepositoryIamBinding []byte

//go:embed mapping/google/resource/source/google_sourcerepo_repository_iam_member.json
var googleSourcerepoRepositoryIamMember []byte

//go:embed mapping/google/resource/source/google_sourcerepo_repository_iam_policy.json
var googleSourcerepoRepositoryIamPolicy []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_key_iam_binding.json
var googleTagsTagKeyIamBinding []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_key_iam_member.json
var googleTagsTagKeyIamMember []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_key_iam_policy.json
var googleTagsTagKeyIamPolicy []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_value_iam_binding.json
var googleTagsTagValueIamBinding []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_value_iam_member.json
var googleTagsTagValueIamMember []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_value_iam_policy.json
var googleTagsTagValueIamPolicy []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_group_iam_binding.json
var googleVertexAIFeatureGroupIamBinding []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_group_iam_member.json
var googleVertexAIFeatureGroupIamMember []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_group_iam_policy.json
var googleVertexAIFeatureGroupIamPolicy []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store_featureview_iam_binding.json
var googleVertexAIFeatureOnlineStoreFeatureviewIamBinding []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store_featureview_iam_member.json
var googleVertexAIFeatureOnlineStoreFeatureviewIamMember []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store_featureview_iam_policy.json
var googleVertexAIFeatureOnlineStoreFeatureviewIamPolicy []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store_iam_binding.json
var googleVertexAIFeatureOnlineStoreIamBinding []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store_iam_member.json
var googleVertexAIFeatureOnlineStoreIamMember []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store_iam_policy.json
var googleVertexAIFeatureOnlineStoreIamPolicy []byte

//go:embed mapping/google/resource/notebooks/google_workbench_instance_iam_binding.json
var googleWorkbenchInstanceIamBinding []byte

//go:embed mapping/google/resource/notebooks/google_workbench_instance_iam_member.json
var googleWorkbenchInstanceIamMember []byte

//go:embed mapping/google/resource/notebooks/google_workbench_instance_iam_policy.json
var googleWorkbenchInstanceIamPolicy []byte

//go:embed mapping/google/resource/workstations/google_workstations_workstation_config_iam_binding.json
var googleWorkstationsWorkstationConfigIamBinding []byte

//go:embed mapping/google/resource/workstations/google_workstations_workstation_config_iam_member.json
var googleWorkstationsWorkstationConfigIamMember []byte

//go:embed mapping/google/resource/workstations/google_workstations_workstation_config_iam_policy.json
var googleWorkstationsWorkstationConfigIamPolicy []byte

//go:embed mapping/google/resource/workstations/google_workstations_workstation_iam_binding.json
var googleWorkstationsWorkstationIamBinding []byte

//go:embed mapping/google/resource/workstations/google_workstations_workstation_iam_member.json
var googleWorkstationsWorkstationIamMember []byte

//go:embed mapping/google/resource/compute/google_compute_instance_template_iam_binding.json
var googleComputeInstanceTemplateIamBinding []byte

//go:embed mapping/google/resource/compute/google_compute_instance_template_iam_member.json
var googleComputeInstanceTemplateIamMember []byte

//go:embed mapping/google/resource/compute/google_compute_instance_template_iam_policy.json
var googleComputeInstanceTemplateIamPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_instant_snapshot_iam_binding.json
var googleComputeInstanceSnapshotIamBinding []byte

//go:embed mapping/google/resource/compute/google_compute_instant_snapshot_iam_member.json
var googleComputeInstanceSnapshotIamMember []byte

//go:embed mapping/google/resource/compute/google_compute_instant_snapshot_iam_policy.json
var googleComputeInstanceSnapshotIamPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_machine_image_iam_binding.json
var googleComputeMachineImageIamBinding []byte

//go:embed mapping/google/resource/compute/google_compute_machine_image_iam_member.json
var googleComputeMachineImageIamMember []byte

//go:embed mapping/google/resource/compute/google_compute_machine_image_iam_policy.json
var googleComputeMachineImageIamPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_region_backend_service_iam_binding.json
var googleComputeRegionBackendServiceIamBinding []byte

//go:embed mapping/google/resource/compute/google_compute_region_backend_service_iam_member.json
var googleComputeRegionBackendServiceIamMember []byte

//go:embed mapping/google/resource/compute/google_compute_region_backend_service_iam_policy.json
var googleComputeRegionBackendServiceIamPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_region_disk_iam_binding.json
var googleComputeRegionDiskIamBinding []byte

//go:embed mapping/google/resource/compute/google_compute_region_disk_iam_member.json
var googleComputeRegionDiskIamMember []byte

//go:embed mapping/google/resource/compute/google_compute_region_disk_iam_policy.json
var googleComputeRegionDiskIamPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_snapshot_iam_binding.json
var googleComputeSnapshotIamBinding []byte

//go:embed mapping/google/resource/compute/google_compute_snapshot_iam_member.json
var googleComputeSnapshotIamMember []byte

//go:embed mapping/google/resource/compute/google_compute_snapshot_iam_policy.json
var googleComputeSnapshotIamPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_storage_pool_iam_binding.json
var googleComputeStoragePoolIamBinding []byte

//go:embed mapping/google/resource/compute/google_compute_storage_pool_iam_member.json
var googleComputeStoragePoolIamMember []byte

//go:embed mapping/google/resource/compute/google_compute_storage_pool_iam_policy.json
var googleComputeStoragePoolIamPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_subnetwork_iam_binding.json
var googleComputeSubnetworkIamBinding []byte

//go:embed mapping/google/resource/compute/google_compute_subnetwork_iam_member.json
var googleComputeSubnetworkIamMember []byte

//go:embed mapping/google/resource/compute/google_compute_subnetwork_iam_policy.json
var googleComputeSubnetworkIamPolicy []byte

//go:embed mapping/google/resource/containeranalysis/google_container_analysis_note_iam_binding.json
var googleContainerAnalysisNoteIamBinding []byte

//go:embed mapping/google/resource/containeranalysis/google_container_analysis_note_iam_member.json
var googleContainerAnalysisNoteIamMember []byte

//go:embed mapping/google/resource/containeranalysis/google_container_analysis_note_iam_policy.json
var googleContainerAnalysisNoteIamPolicy []byte

//go:embed mapping/google/resource/dataform/google_dataform_repository_iam_binding.json
var googleDataformRepositoryIamBinding []byte

//go:embed mapping/google/resource/dataform/google_dataform_repository_iam_member.json
var googleDataformRepositoryIamMember []byte

//go:embed mapping/google/resource/dataform/google_dataform_repository_iam_policy.json
var googleDataformRepositoryIamPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_aspect_type_iam_binding.json
var googleDataplexAspectTypeIamBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_aspect_type_iam_member.json
var googleDataplexAspectTypeIamMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_aspect_type_iam_policy.json
var googleDataplexAspectTypeIamPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry_group_iam_binding.json
var googleDataplexEntryGroupIamBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry_group_iam_member.json
var googleDataplexEntryGroupIamMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry_group_iam_policy.json
var googleDataplexEntryGroupIamPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry_group_iam_binding.json
var googleDataplexEntryTypeIamBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry_group_iam_member.json
var googleDataplexEntryTypeIamMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry_group_iam_policy.json
var googleDataplexEntryTypeIamPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_glossary_iam_binding.json
var googleDataplexGlossaryIamBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_glossary_iam_member.json
var googleDataplexGlossaryIamMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_glossary_iam_policy.json
var googleDataplexGlossaryIamPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_lake_iam_binding.json
var googleDataplexLakeIamBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_lake_iam_member.json
var googleDataplexLakeIamMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_lake_iam_policy.json
var googleDataplexLakeIamPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_task_iam_member.json
var googleDataplexTaskIamMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_task_iam_policy.json
var googleDataplexTaskIamPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_zone_iam_binding.json
var googleDataplexZoneIamBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_zone_iam_member.json
var googleDataplexZoneIamMember []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_autoscaling_policy_iam_binding.json
var googleDataprocAutoscalingPolicyIamBinding []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_autoscaling_policy_iam_member.json
var googleDataprocAutoscalingPolicyIamMember []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_autoscaling_policy_iam_binding.json
var googleDataprocAutoscalingPolicyIamPolicy []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_cluster_iam_binding.json
var googleDataprocClusterIamBinding []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_cluster_iam_member.json
var googleDataprocClusterIamMember []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_cluster_iam_policy.json
var googleDataprocClusterIamPolicy []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_job_iam_binding.json
var googleDataprocJobIamBinding []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_job_iam_member.json
var googleDataprocJobIamMember []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_job_iam_policy.json
var googleDataprocJobIamPolicy []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_database_iam_binding.json
var googleDataprocMetastoreDatabaseIamBinding []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_database_iam_member.json
var googleDataprocMetastoreDatabaseIamMember []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_database_iam_policy.json
var googleDataprocMetastoreDatabaseIamPolicy []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_federation_iam_binding.json
var googleDataprocMetastoreFederationIamBinding []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_federation_iam_member.json
var googleDataprocMetastoreFederationIamMember []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_federation_iam_policy.json
var googleDataprocMetastoreFederationIamPolicy []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_federation_iam_binding.json
var googleDataprocMetastoreServiceIamBinding []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_service_iam_member.json
var googleDataprocMetastoreServiceIamMember []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_federation_iam_policy.json
var googleDataprocMetastoreServiceIamPolicy []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_federation_iam_binding.json
var googleDataprocMetastoreTableIamBinding []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_table_iam_member.json
var googleDataprocMetastoreTableIamMember []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_table_iam_policy.json
var googleDataprocMetastoreTableIamPolicy []byte

//go:embed mapping/google/resource/binaryauthorization/google_binary_authorization_attestor_iam_binding.json
var googleBinaryAuthorizationAttestorIamBinding []byte

//go:embed mapping/google/resource/binaryauthorization/google_binary_authorization_attestor_iam_member.json
var googleBinaryAuthorizationAttestorIamMember []byte

//go:embed mapping/google/resource/binaryauthorization/google_binary_authorization_attestor_iam_policy.json
var googleBinaryAuthorizationAttestorIamPolicy []byte

//go:embed mapping/google/resource/cloudtasks/google_cloud_tasks_queue_iam_binding.json
var googleCloudTasksQueueIamBinding []byte

//go:embed mapping/google/resource/cloudtasks/google_cloud_tasks_queue_iam_member.json
var googleCloudTasksQueueIamMember []byte

//go:embed mapping/google/resource/cloudtasks/google_cloud_tasks_queue_iam_policy.json
var googleCloudTasksQueueIamPolicy []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_custom_target_type_iam_binding.json
var googleClouddeployCustomTargetTypeIamBinding []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_custom_target_type_iam_member.json
var googleClouddeployCustomTargetTypeIamMember []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_custom_target_type_iam_policy.json
var googleClouddeployCustomTargetTypeIamPolicy []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_delivery_pipeline_iam_binding.json
var googleClouddeployDeliveryPipelineIamBinding []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_delivery_pipeline_iam_member.json
var googleClouddeployDeliveryPipelineIamMember []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_delivery_pipeline_iam_policy.json
var googleClouddeployDeliveryPipelineIamPolicy []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_target_iam_binding.json
var googleClouddeployTargetIamBinding []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_target_iam_member.json
var googleClouddeployTargetIamMember []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_target_iam_policy.json
var googleClouddeployTargetIamPolicy []byte

//go:embed mapping/google/resource/aiplatform/google_colab_runtime_template_iam_member.json
var googleColabRuntimeTemplateIamBinding []byte

//go:embed mapping/google/resource/aiplatform/google_colab_runtime_template_iam_member.json
var googleColabRuntimeTemplateIamMember []byte

//go:embed mapping/google/resource/aiplatform/google_colab_runtime_template_iam_policy.json
var googleColabRuntimeTemplateIamPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_backend_bucket_iam_binding.json
var googleComputeBackendBucketIamBinding []byte

//go:embed mapping/google/resource/compute/google_compute_backend_bucket_iam_member.json
var googleComputeBackendBucketIamMember []byte

//go:embed mapping/google/resource/compute/google_compute_backend_bucket_iam_policy.json
var googleComputeBackendBucketIamPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_backend_service_iam_binding.json
var googleComputeBackendServiceIamBinding []byte

//go:embed mapping/google/resource/compute/google_compute_backend_service_iam_member.json
var googleComputeBackendServiceIamMember []byte

//go:embed mapping/google/resource/compute/google_compute_backend_service_iam_policy.json
var googleComputeBackendServiceIamPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_disk_iam_binding.json
var googleComputeDiskIamBinding []byte

//go:embed mapping/google/resource/compute/google_compute_disk_iam_member.json
var googleComputeDiskIamMember []byte

//go:embed mapping/google/resource/compute/google_compute_image_iam_binding.json
var googleComputeImageIamBinding []byte

//go:embed mapping/google/resource/compute/google_compute_image_iam_member.json
var googleComputeImageIamMember []byte

//go:embed mapping/google/resource/compute/google_compute_image_iam_policy.json
var googleComputeImageIamPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_instance_iam_binding.json
var googleComputeInstanceIamBinding []byte

//go:embed mapping/google/resource/compute/google_compute_instance_iam_member.json
var googleComputeInstanceIamMember []byte

//go:embed mapping/google/resource/compute/google_compute_instance_iam_policy.json
var googleComputeInstanceIamPolicy []byte

//go:embed mapping/google/resource/datafusion/google_data_fusion_instance_iam_binding.json
var googleDataFusionInstanceIamBinding []byte

//go:embed mapping/google/resource/datafusion/google_data_fusion_instance_iam_member.json
var googleDataFusionInstanceIamMember []byte

//go:embed mapping/google/resource/datafusion/google_data_fusion_instance_iam_policy.json
var googleDataFusionInstanceIamPolicy []byte

//go:embed mapping/google/resource/dns/google_dns_managed_zone_iam_binding.json
var googleDNSManagedZoneIamBinding []byte

//go:embed mapping/google/resource/dns/google_dns_managed_zone_iam_member.json
var googleDNSManagedZoneIamMember []byte

//go:embed mapping/google/resource/dns/google_dns_managed_zone_iam_policy.json
var googleDNSManagedZoneIamPolicy []byte

//go:embed mapping/google/resource/servicemanagement/google_endpoints_service_consumers_iam_binding.json
var googleEndpointsServiceConsumersIamBinding []byte

//go:embed mapping/google/resource/servicemanagement/google_endpoints_service_consumers_iam_member.json
var googleEndpointsServiceConsumersIamMember []byte

//go:embed mapping/google/resource/servicemanagement/google_endpoints_service_consumers_iam_policy.json
var googleEndpointsServiceConsumersIamPolicy []byte

//go:embed mapping/google/resource/servicemanagement/google_endpoints_service_iam_binding.json
var googleEndpointsServiceIamBinding []byte

//go:embed mapping/google/resource/servicemanagement/google_endpoints_service_iam_member.json
var googleEndpointsServiceIamMember []byte

//go:embed mapping/google/resource/servicemanagement/google_endpoints_service_iam_policy.json
var googleEndpointsServiceIamPolicy []byte

//go:embed mapping/google/resource/resourcemanager/google_folder_iam_binding.json
var googleFolderIamBinding []byte

//go:embed mapping/google/resource/resourcemanager/google_folder_iam_member.json
var googleFolderIamMember []byte

//go:embed mapping/google/resource/resourcemanager/google_folder_iam_policy.json
var googleFolderIamPolicy []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_repository_group_iam_binding.json
var googleGeminiRepositoryGroupIamBinding []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_repository_group_iam_member.json
var googleGeminiRepositoryGroupIamMember []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_repository_group_iam_policy.json
var googleGeminiRepositoryGroupIamPolicy []byte

//go:embed mapping/google/resource/gkebackup/google_gke_backup_backup_plan_iam_binding.json
var googleGkeBackupBackupPlanIamBinding []byte

//go:embed mapping/google/resource/gkebackup/google_gke_backup_backup_plan_iam_member.json
var googleGkeBackupBackupPlanIamMember []byte

//go:embed mapping/google/resource/gkebackup/google_gke_backup_backup_plan_iam_policy.json
var googleGkeBackupBackupPlanIamPolicy []byte

//go:embed mapping/google/resource/gkebackup/google_gke_backup_restore_plan_iam_binding.json
var googleGkeBackupRestorePlanIamBinding []byte

//go:embed mapping/google/resource/gkebackup/google_gke_backup_restore_plan_iam_member.json
var googleGkeBackupRestorePlanIamMember []byte

//go:embed mapping/google/resource/gkebackup/google_gke_backup_restore_plan_iam_policy.json
var googleGkeBackupRestorePlanIamPolicy []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_feature_iam_binding.json
var googleGkeHubFeatureIamBinding []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_feature_iam_member.json
var googleGkeHubFeatureIamMember []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_feature_iam_policy.json
var googleGkeHubFeatureIamPolicy []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_membership_iam_binding.json
var googleGkeHubMembershipIamBinding []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_membership_iam_member.json
var googleGkeHubMembershipIamMember []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_membership_iam_policy.json
var googleGkeHubMembershipIamPolicy []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_scope_iam_binding.json
var googleGkeHubScopeIamBinding []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_scope_iam_member.json
var googleGkeHubScopeIamMember []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_scope_iam_policy.json
var googleGkeHubScopeIamPolicy []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_consent_store_iam_binding.json
var googleHealthcareConsentStoreIamBinding []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_consent_store_iam_member.json
var googleHealthcareConsentStoreIamMember []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_consent_store_iam_policy.json
var googleHealthcareConsentStoreIamPolicy []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_dataset_iam_binding.json
var googleHealthcareDatasetIamBinding []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_dataset_iam_member.json
var googleHealthcareDatasetIamMember []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_dataset_iam_policy.json
var googleHealthcareDatasetIamPolicy []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_dicom_store_iam_binding.json
var googleHealthcareDicomStoreIamBinding []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_dicom_store_iam_member.json
var googleHealthcareDicomStoreIamMember []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_dicom_store_iam_policy.json
var googleHealthcareDicomStoreIamPolicy []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_fhir_store_iam_binding.json
var googleHealthcareFhirStoreIamBinding []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_fhir_store_iam_member.json
var googleHealthcareFhirStoreIamMember []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_fhir_store_iam_policy.json
var googleHealthcareFhirStoreIamPolicy []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_hl7_v2_store_iam_binding.json
var googleHealthcareHl7V2StoreIamBinding []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_hl7_v2_store_iam_member.json
var googleHealthcareHl7V2StoreIamMember []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_hl7_v2_store_iam_policy.json
var googleHealthcareHl7V2StoreIamPolicy []byte

//go:embed mapping/google/resource/iam.googleapis.com/google_iam_workload_identity_pool_iam_binding.json
var googleIamWorkloadIdentityPoolIamBinding []byte

//go:embed mapping/google/resource/iam.googleapis.com/google_iam_workload_identity_pool_iam_member.json
var googleIamWorkloadIdentityPoolIamMember []byte

//go:embed mapping/google/resource/iam.googleapis.com/google_iam_workload_identity_pool_iam_policy.json
var googleIamWorkloadIdentityPoolIamPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_app_engine_service_iam_binding.json
var googleIapAppEngineServiceIamBinding []byte

//go:embed mapping/google/resource/iap/google_iap_app_engine_service_iam_member.json
var googleIapAppEngineServiceIamMember []byte

//go:embed mapping/google/resource/iap/google_iap_app_engine_service_iam_policy.json
var googleIapAppEngineServiceIamPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_app_engine_version_iam_binding.json
var googleIapAppEngineVersionIamBinding []byte

//go:embed mapping/google/resource/iap/google_iap_app_engine_version_iam_member.json
var googleIapAppEngineVersionIamMember []byte

//go:embed mapping/google/resource/iap/google_iap_app_engine_version_iam_policy.json
var googleIapAppEngineVersionIamPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_dest_group_iam_binding.json
var googleIapTunnelDestGroupIamBinding []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_dest_group_iam_member.json
var googleIapTunnelDestGroupIamMember []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_dest_group_iam_policy.json
var googleIapTunnelDestGroupIamPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_iam_binding.json
var googleIapTunnelIamBinding []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_iam_member.json
var googleIapTunnelIamMember []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_iam_policy.json
var googleIapTunnelIamPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_instance_iam_binding.json
var googleIapTunnelInstanceIamBinding []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_instance_iam_member.json
var googleIapTunnelInstanceIamMember []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_instance_iam_policy.json
var googleIapTunnelInstanceIamPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_web_backend_service_iam_binding.json
var googleIapWebBackendServiceIamBinding []byte

//go:embed mapping/google/resource/iap/google_iap_web_backend_service_iam_member.json
var googleIapWebBackendServiceIamMember []byte

//go:embed mapping/google/resource/iap/google_iap_web_backend_service_iam_policy.json
var googleIapWebBackendServiceIamPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_web_cloud_run_service_iam_binding.json
var googleIapWebCloudRunServiceIamBinding []byte

//go:embed mapping/google/resource/iap/google_iap_web_cloud_run_service_iam_member.json
var googleIapWebCloudRunServiceIamMember []byte

//go:embed mapping/google/resource/iap/google_iap_web_cloud_run_service_iam_policy.json
var googleIapWebCloudRunServiceIamPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_web_iam_binding.json
var googleIapWebIamBinding []byte

//go:embed mapping/google/resource/iap/google_iap_web_iam_member.json
var googleIapWebIamMember []byte

//go:embed mapping/google/resource/iap/google_iap_web_iam_policy.json
var googleIapWebIamPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_web_region_backend_service_iam_binding.json
var googleIapWebRegionBackendServiceIamBinding []byte

//go:embed mapping/google/resource/iap/google_iap_web_region_backend_service_iam_member.json
var googleIapWebRegionBackendServiceIamMember []byte

//go:embed mapping/google/resource/iap/google_iap_web_region_backend_service_iam_policy.json
var googleIapWebRegionBackendServiceIamPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_web_type_app_engine_iam_binding.json
var googleIapWebTypeAppEngineIamBinding []byte

//go:embed mapping/google/resource/iap/google_iap_web_type_app_engine_iam_member.json
var googleIapWebTypeAppEngineIamMember []byte

//go:embed mapping/google/resource/iap/google_iap_web_type_app_engine_iam_policy.json
var googleIapWebTypeAppEngineIamPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_web_type_compute_iam_binding.json
var googleIapWebTypeComputeIamBinding []byte

//go:embed mapping/google/resource/iap/google_iap_web_type_compute_iam_member.json
var googleIapWebTypeComputeIamMember []byte

//go:embed mapping/google/resource/iap/google_iap_web_type_compute_iam_policy.json
var googleIapWebTypeComputeIamPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_disk_iam_policy.json
var googleComputeDiskIamPolicy []byte

//go:embed mapping/google/resource/kms/google_kms_ekm_connection_iam_binding.json
var googleKmsEkmConnectionIamBinding []byte

//go:embed mapping/google/resource/kms/google_kms_ekm_connection_iam_member.json
var googleKmsEkmConnectionIamMember []byte

//go:embed mapping/google/resource/kms/google_kms_ekm_connection_iam_policy.json
var googleKmsEkmConnectionIamPolicy []byte

//go:embed mapping/google/resource/logging/google_logging_log_view_iam_binding.json
var googleLoggingLogViewIamBinding []byte

//go:embed mapping/google/resource/logging/google_logging_log_view_iam_member.json
var googleLoggingLogViewIamMember []byte

//go:embed mapping/google/resource/networksecurity/google_network_security_address_group_iam_binding.json
var googleNetworkSecurityAddressGroupIamBinding []byte

//go:embed mapping/google/resource/networksecurity/google_network_security_address_group_iam_member.json
var googleNetworkSecurityAddressGroupIamMember []byte

//go:embed mapping/google/resource/networksecurity/google_network_security_address_group_iam_policy.json
var googleNetworkSecurityAddressGroupIamPolicy []byte

//go:embed mapping/google/resource/notebooks/google_notebooks_instance_iam_binding.json
var googleNotebooksInstanceIamBinding []byte

//go:embed mapping/google/resource/notebooks/google_notebooks_instance_iam_member.json
var googleNotebooksInstanceIamMember []byte

//go:embed mapping/google/resource/notebooks/google_notebooks_instance_iam_policy.json
var googleNotebooksInstanceIamPolicy []byte

//go:embed mapping/google/resource/notebooks/google_notebooks_runtime_iam_binding.json
var googleNotebooksRuntimeIamBinding []byte

//go:embed mapping/google/resource/notebooks/google_notebooks_runtime_iam_member.json
var googleNotebooksRuntimeIamMember []byte

//go:embed mapping/google/resource/logging/google_logging_log_view_iam_policy.json
var googleLoggingLogViewIamPolicy []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_subscription_iam_policy.json
var googlePubsubSubscriptionIamPolicy []byte

//go:embed mapping/google/resource/workstations/google_workstations_workstation_iam_policy.json
var googleWorkstationsWorkstationIamPolicy []byte

//go:embed mapping/google/resource/resourcemanager/google_organization_iam_binding.json
var googleOrganizationIamBinding []byte

//go:embed mapping/google/resource/resourcemanager/google_organization_iam_member.json
var googleOrganizationIamMember []byte

//go:embed mapping/google/resource/resourcemanager/google_organization_iam_policy.json
var googleOrganizationIamPolicy []byte

//go:embed mapping/google/resource/resourcemanager/google_project_iam_policy.json
var googleProjectIamPolicy []byte

//go:embed mapping/google/resource/notebooks/google_notebooks_runtime_iam_policy.json
var googleNotebooksRuntimeIamPolicy []byte

//go:embed mapping/google/resource/runtimeconfig/google_runtimeconfig_config_iam_binding.json
var googleRuntimeconfigConfigIamBinding []byte

//go:embed mapping/google/resource/runtimeconfig/google_runtimeconfig_config_iam_member.json
var googleRuntimeconfigConfigIamMember []byte

//go:embed mapping/google/resource/runtimeconfig/google_runtimeconfig_config_iam_policy.json
var googleRuntimeconfigConfigIamPolicy []byte

//go:embed mapping/google/resource/securitycenter/google_scc_source_iam_binding.json
var googleSccSourceIamBinding []byte

//go:embed mapping/google/resource/securitycenter/google_scc_source_iam_member.json
var googleSccSourceIamMember []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_organization_source_iam_binding.json
var googleSccV2OrganizationSourceIamBinding []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_organization_source_iam_member.json
var googleSccV2OrganizationSourceIamMember []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_organization_source_iam_policy.json
var googleSccV2OrganizationSourceIamPolicy []byte

//go:embed mapping/google/resource/storage/google_storage_managed_folder_iam_binding.json
var googleStorageManagedFolderIamBinding []byte

//go:embed mapping/google/resource/storage/google_storage_managed_folder_iam_member.json
var googleStorageManagedFolderIamMember []byte

//go:embed mapping/google/resource/storage/google_storage_managed_folder_iam_policy.json
var googleStorageManagedFolderIamPolicy []byte

//go:embed mapping/google/resource/securitycenter/google_scc_source_iam_policy.json
var googleSccSourceIamPolicy []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_location_tag_binding.json
var googleTagsLocationTagBinding []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_binding.json
var googleTagsTagBinding []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_key.json
var googleTagsTagKey []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_value.json
var googleTagsTagValue []byte
