package pike

import (
	"fmt"
)

// GetGCPDataPermissions gets permissions required for datasources
func GetGCPDataPermissions(result ResourceV2) ([]string, error) {
	temp := GCPDataLookup(result.Name)

	var (
		Permissions []string
		err         error
	)

	if temp != nil {
		Permissions, err = GetPermissionMap(temp.([]byte), result.Attributes)
	} else {
		return nil, fmt.Errorf("data.%s not implemented", result.Name)
	}

	return Permissions, err
}

func GCPDataLookup(result string) interface{} {
	TFLookup := map[string]interface{}{
		"google_api_gateway_api_config_iam_policy":               placeholder,
		"google_api_gateway_api_iam_policy":                      placeholder,
		"google_api_gateway_gateway_iam_policy":                  placeholder,
		"google_app_engine_default_service_account":              dataGoogleAppEngineDefaultServiceAccount,
		"google_artifact_registry_repository":                    dataGoogleArtifactRegistryRepository,
		"google_artifact_registry_repository_iam_policy":         dataGoogleArtifactRegistryRepositoryIamPolicy,
		"google_bigquery_analytics_hub_data_exchange_iam_policy": dataGoogleBigqueryHubDataExchangeIamPolicy,
		"google_bigquery_analytics_hub_listing_iam_policy":       dataGoogleBigqueryAnalyticsHubListingIamPolicy,
		"google_bigquery_connection_iam_policy":                  placeholder,
		"google_bigquery_datapolicy_data_policy_iam_policy":      dataGoogleBigqueryDatapolicyDataPolicyIamPolicy,
		"google_bigquery_dataset_iam_policy":                     placeholder,
		"google_bigquery_default_service_account":                dataGoogleBigqueryDefaultServiceAccount,
		"google_bigquery_table_iam_policy":                       placeholder,
		"google_bigtable_instance_iam_policy":                    dataGoogleBigtableInstanceIamPolicy,
		"google_bigtable_table_iam_policy":                       placeholder,
		"google_compute_image":                                   placeholder,
		"google_compute_network":                                 dataGoogleComputeNetwork,
		"google_compute_subnetwork":                              dataGoogleComputeSubnetwork,
		"google_compute_zones":                                   dataGoogleComputeZones,
		"google_dns_keys":                                        dataGoogleDnsKeys,
		"google_dns_managed_zone":                                dataGoogleDnsManagedZone,
		"google_dns_managed_zone_iam_policy":                     dataGoogleDnsManagedZoneIamPolicy,
		"google_dns_record_set":                                  dataGoogleDnsRecordSet,
		"google_iam_policy":                                      placeholder,
		"google_iam_role":                                        placeholder,
		"google_kms_crypto_key":                                  dataGoogleKmsCryptoKey,
		"google_kms_key_ring":                                    dataGoogleKmsKeyRing,
		"google_kms_key_ring_iam_policy":                         dataGoogleKmsKeyRingIamPolicy,
		"google_kms_secret":                                      dataGoogleKmsSecret,
		"google_kms_secret_asymmetric":                           dataGoogleKmsSecretAsymnetric,
		"google_kms_secret_ciphertext":                           placeholder,
		"google_project":                                         dataGoogleProject,
		"google_projects":                                        placeholder,
		"google_pubsub_subscription":                             dataGooglePubsubSubscription,
		"google_pubsub_subscription_iam_policy":                  dataGooglePubsubSubscriptionIamPolicy,
		"google_pubsub_topic":                                    dataGooglePubsubTopic,
		"google_pubsub_topic_iam_policy":                         dataGooglePubsubTopicIamPolicy,
		"google_service_account":                                 dataGoogleServiceAccount,
		"google_spanner_database_iam_policy":                     dataGoogleSpannerDatabaseIamPolicy,
		"google_spanner_instance":                                dataGoogleSpannerInstance,
		"google_spanner_instance_iam_policy":                     dataGoogleSpannerInstanceIamPolicy,
		"google_storage_bucket":                                  dataGoogleStorageBucket,
		"google_storage_bucket_iam_policy":                       dataGoogleStorageBucketIamPolicy,
		"google_storage_bucket_object":                           dataGoogleStorageBucketObject,
		"google_storage_bucket_object_content":                   dataGoogleStorageBucketObjectContent,
		"google_storage_object_signed_url":                       placeholder,
		"google_storage_project_service_account":                 dataGoogleStorageProjectServiceAccount,
		"google_storage_transfer_project_service_account":        dataGoogleStorageTransferProjectServiceAccount,
		"google_vertex_ai_featurestore_entitytype_iam_policy":    dataGoogleVertexAiFeaturestoreEntitytypeIamPolicy,
		"google_vertex_ai_featurestore_iam_policy":               dataGoogleVertexAiFeaturestoreIamPolicy,
		"google_cloudfunctions2_function":                        dataGoogleCloudfunctionsFunction,
		"google_cloudfunctions2_function_iam_policy":             dataGoogleCloudfunctionsFunctionIamPolicy,
		"google_cloudfunctions_function":                         dataGoogleCloudfunctionsFunction,
		"google_cloudfunctions_function_iam_policy":              dataGoogleCloudfunctionsFunctionIamPolicy,
		"google_gke_backup_backup_plan_iam_policy":               dataGoogleGkeBackupBackupPlanIamPolicy,
		"google_gke_backup_restore_plan_iam_policy":              dataGoogleGkeBackupRestorePlanIamPolicy,
		"google_gke_hub_feature_iam_policy":                      dataGoogleGkeHubFeatureIamPolicy,
		"google_gke_hub_membership_iam_policy":                   dataGoogleGkeHubMembershipIamPolicy,
		"google_gke_hub_scope_iam_policy":                        dataGoogleGkeHubScopeIamPolicy,
		"google_kms_crypto_key_iam_policy":                       dataGoogleKmsCryptoKeyIamPolicy,
		"google_kms_crypto_key_version":                          dataGoogleKmsCryptoKeyVersion,
		"google_secret_manager_secret":                           dataGoogleSecretManagerSecret,
		"google_secret_manager_secret_iam_policy":                dataGoogleSecretManagerSecretIamPolicy,
		"google_secret_manager_secret_version":                   dataGoogleSecretManagerSecretVersion,
		"google_secret_manager_secret_version_access":            dataGoogleManagerSecretVersionAccess,
		"google_client_config":                                   placeholder,
		"google_client_openid_userinfo":                          placeholder,
		"google_project_service":                                 dataGoogleProjectService,
		"google_service_account_access_token":                    dataGoogleServiceAccountAccessToken,
		"google_service_account_iam_policy":                      dataGoogleServiceAccountIamPolicy,
		"google_service_account_id_token":                        placeholder,
		"google_service_account_jwt":                             dataGoogleServiceAccountJwt,
		"google_service_account_key":                             dataGoogleServiceAccountKey,
		"google_redis_instance":                                  dataGoogleRedisInstance,
		"google_alloydb_locations":                               dataGoogleAlloydbLocations,
		"google_alloydb_supported_database_flags":                dataGoogleAlloydbSupportedDatabaseFlags,
		"google_apigee_environment_iam_policy":                   placeholder,
		"google_beyondcorp_app_connection":                       dataGoogleBeyondcorpAppConnection,
		"google_beyondcorp_app_connector":                        dataGoogleBeyondcorpAppConnector,
		"google_beyondcorp_app_gateway":                          dataGoogleBeyondcorpAppGateway,
		"google_cloud_run_locations":                             dataGoogleCloudRunLocations,
		"google_cloud_run_service":                               dataGoogleCloudRunService,
		"google_cloud_run_service_iam_policy":                    dataGoogleCloudRunServiceIamPolicy,
		"google_cloud_run_v2_job":                                dataGoogleCloudRunV2Job,
		"google_cloud_run_v2_job_iam_policy":                     dataGoogleCloudRunV2JobIamPolicy,
		"google_cloud_run_v2_service":                            dataGoogleCloudRunV2Service,
		"google_cloud_run_v2_service_iam_policy":                 dataGoogleCloudRunV2ServiceIamPolicy,
	}
	return TFLookup[result]
}
