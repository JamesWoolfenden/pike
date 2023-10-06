package pike

import (
	"github.com/rs/zerolog/log"
)

// GetGCPPermissions for GCP resources
func GetGCPPermissions(result ResourceV2) ([]string, error) {
	var (
		err         error
		Permissions []string
	)

	if result.TypeName == "resource" {
		Permissions, err = GetGCPResourcePermissions(result)
		if err != nil {
			return Permissions, err
		}
	} else {
		Permissions, err = GetGCPDataPermissions(result)
		if err != nil {
			return Permissions, err
		}
	}

	return Permissions, err
}

// GetGCPResourcePermissions looks up permissions required for resources
func GetGCPResourcePermissions(result ResourceV2) ([]string, error) {
	temp := GCPLookup(result.Name)

	var (
		Permissions []string
		err         error
	)

	if temp != nil {
		Permissions, err = GetPermissionMap(temp.([]byte), result.Attributes)
	} else {
		log.Printf("%s not implemented", result.Name)
	}

	return Permissions, err
}

func GCPLookup(result string) interface{} {
	TFLookup := map[string]interface{}{
		"google_artifact_registry_repository":             googleArtifactRegistryRepository,
		"google_artifact_registry_repository_iam_binding": googleArtifactRegistryRepositoryIamBinding,
		"google_artifact_registry_repository_iam_member":  googleArtifactRegistryRepositoryIamMember,
		"google_artifact_registry_repository_iam_policy":  googleArtifactRegistryRepositoryIamPolicy,
		"google_bigquery_dataset":                         googleBigqueryDataset,
		"google_bigquery_job":                             googleBigqueryJob,
		"google_bigquery_table":                           placeholder,
		"google_bigtable_instance":                        googleBigtableInstance,
		"google_cloudfunctions_function":                  googleCloudfunctionsFunction,
		"google_cloudfunctions_function_iam_member":       googleCloudfunctionsFunctionIamPolicy,
		"google_cloudfunctions_function_iam_policy":       googleCloudfunctionsFunctionIamPolicy,
		"google_compute_address":                          googleComputeAddress,
		"google_compute_firewall":                         googleComputeFirewall,
		"google_compute_global_address":                   googleComputeGlobalAddress,
		"google_compute_instance":                         googleComputeInstance,
		"google_compute_instance_template":                googleComputeInstanceTemplate,
		"google_compute_network":                          googleComputeNetwork,
		"google_compute_project_metadata_item":            googleComputeProjectMetadataItem,
		"google_compute_region_ssl_certificate":           googleComputeRegionSslCertificate,
		"google_compute_security_policy":                  googleComputeSecurityPolicy,
		"google_compute_subnetwork":                       googleComputeSubnetwork,
		"google_container_cluster":                        googleContainerCluster,
		"google_container_node_pool":                      googleContainerNodePool,
		"google_dns_managed_zone":                         googleDnsmanagedZone,
		"google_dns_policy":                               googleDnsPolicy,
		"google_dns_record_set":                           googleDnsRecordSet,
		"google_kms_crypto_key":                           googleKmsCryptoKey,
		"google_kms_crypto_key_iam_binding":               googlekmsCryptoKeyIamBinding,
		"google_kms_crypto_key_iam_member":                googlekmsCryptoKeyIamMember,
		"google_kms_crypto_key_iam_policy":                googlekmsCryptoKeyIamPolicy,
		"google_kms_key_ring":                             googleKmsKeyRing,
		"google_project_iam_binding":                      googleProjectIamBinding,
		"google_project_iam_custom_role":                  googleProjectIamCustomRole,
		"google_project_iam_member":                       googleProjectIamBinding,
		"google_pubsub_lite_reservation":                  googlePubsubLiteReservation,
		"google_pubsub_lite_subscription":                 googlePubsubLiteSubscription,
		"google_pubsub_lite_topic":                        googlePubsubLiteTopic,
		"google_pubsub_topic":                             googlePubsubTopic,
		"google_service_account":                          googleServiceAccount,
		"google_service_account_iam_binding":              googleServiceAccountIamBinding,
		"google_service_account_iam_member":               googleServiceAccountIamMember,
		"google_service_account_iam_policy":               googleServiceAccountIamPolicy,
		"google_service_account_key":                      googleServiceAccountKey,
		"google_service_networking_connection":            googleServiceNetworkingConnection,
		"google_sourcerepo_repository":                    googleSourcerepoRepository,
		"google_sql_database":                             googleSQLDatabase,
		"google_sql_database_instance":                    googleSQLDatabaseInstance,
		"google_sql_user":                                 googleSQLUser,
		"google_storage_bucket":                           googleStorageBucket,
		"google_storage_bucket_acl":                       googleStorageBucketACL,
		"google_storage_bucket_iam_binding":               googleStorageBucketIamBinding,
		"google_storage_bucket_object":                    googleStorageBucketObject,
		"google_bigtable_table":                           googleBigtableTable,
		"google_bigtable_instance_iam_policy":             googleBigTableInstanceIam,
		"google_bigtable_instance_iam_member":             googleBigTableInstanceIam,
		"google_bigtable_instance_iam_binding":            googleBigTableInstanceIam,
		"google_bigtable_table_iam_binding":               googleBigTableTableIam,
		"google_bigtable_table_iam_member":                googleBigTableTableIam,
		"google_bigtable_table_iam_policy":                googleBigTableTableIam,
		"google_pubsub_schema":                            googlePubsubSchema,
		"google_pubsub_subscription":                      googlePubsubSubscription,
		"google_project_service":                          googleProjectService,
	}

	return TFLookup[result]
}
