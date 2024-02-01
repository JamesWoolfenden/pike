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

var googleApiGatewayApi []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_api_config.json
var googleApiGatewayApiConfig []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_api_config_iam.json
var googleApiGatewayApiConfigIam []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_api_iam.json
var googleApiGatewayApiIam []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_gateway.json
var googleApiGatewayGateway []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_gateway_iam.json
var googleApiGatewayGatewayIam []byte

//go:embed mapping/google/resource/spanner/google_spanner_database.json
var GoogleSpannerDatabase []byte

//go:embed mapping/google/resource/spanner/google_spanner_database_iam.json
var GoogleSpannerDatabaseIam []byte

//go:embed mapping/google/resource/spanner/google_spanner_instance.json
var GoogleSpannerInstance []byte

//go:embed mapping/google/resource/spanner/google_spanner_instance_iam.json
var GoogleSpannerInstanceIam []byte

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
