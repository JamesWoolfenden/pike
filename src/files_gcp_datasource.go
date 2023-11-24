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
