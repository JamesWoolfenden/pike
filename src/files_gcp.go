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

//go:embed  mapping/google/resource/cloudkms/google_kms_crypto_key.json
var googleKmsCryptoKey []byte

//go:embed  mapping/google/resource/storage/google_storage_bucket_acl.json
var googleStorageBucketACL []byte

//go:embed  mapping/google/resource/storage/google_storage_bucket_iam_binding.json
var googleStorageBucketIamBinding []byte

//go:embed  mapping/google/resource/cloudsql/google_sql_database_instance.json
var googleSQLDatabaseInstance []byte

//go:embed  mapping/google/resource/cloudsql/google_sql_database.json
var googleSQLDatabase []byte

//go:embed  mapping/google/resource/cloudsql/google_sql_user.json
var googleSQLUser []byte

//go:embed  mapping/google/resource/servicenetworking/google_service_networking_connection.json
var googleServiceNetworkingConnection []byte

//go:embed  mapping/google/resource/compute/google_compute_global_address.json
var googleComputeGlobalAddress []byte

//go:embed  mapping/google/resource/pubsub/google_pubsub_topic.json
var googlePubsubTopic []byte

//go:embed  mapping/google/resource/source/google_soucerepo_repository.json
var googleSourcerepoRepository []byte

//go:embed  mapping/google/resource/iam/google_service_account_iam_policy.json
var googleServiceAccountIamPolicy []byte

//go:embed  mapping/google/resource/iam/google_service_account_key.json
var googleServiceAccountKey []byte

//go:embed  mapping/google/resource/resourcemanager/google_project_iam_binding.json
var googleProjectIamBinding []byte

//go:embed  mapping/google/resource/container/google_container_cluster.json
var googleContainerCluster []byte

//go:embed  mapping/google/resource/container/google_container_node_pool.json
var googleContainerNodePool []byte

//go:embed  mapping/google/resource/bigquery/google_bigquery_dataset.json
var googleBigqueryDataset []byte

//go:embed  mapping/google/resource/bigquery/google_bigquery_job.json
var googleBigqueryJob []byte

//go:embed  mapping/google/resource/artifactregistry/google_artifact_registry_repository.json
var googleArtifactRegistryRepository []byte

//go:embed  mapping/google/resource/artifactregistry/google_artifact_registry_repository_iam_binding.json
var googleArtifactRegistryRepositoryIamBinding []byte

//go:embed  mapping/google/resource/artifactregistry/google_artifact_registry_repository_iam_member.json
var googleArtifactRegistryRepositoryIamMember []byte

//go:embed  mapping/google/resource/artifactregistry/google_artifact_registry_repository_iam_policy.json
var googleArtifactRegistryRepositoryIamPolicy []byte

//go:embed  mapping/google/resource/pubsublite/google_pubsub_lite_reservation.json
var googlePubsubLiteReservation []byte

//go:embed  mapping/google/resource/pubsublite/google_pubsub_lite_subscription.json
var googlePubsubLiteSubscription []byte

//go:embed  mapping/google/resource/pubsublite/google_pubsub_lite_topic.json
var googlePubsubLiteTopic []byte

//go:embed  mapping/google/resource/compute/google_compute_security_policy.json
var googleComputeSecurityPolicy []byte

//go:embed  mapping/google/resource/cloudkms/google_kms_crypto_key_iam_binding.json
var googlekmsCryptoKeyIamBinding []byte

//go:embed  mapping/google/resource/cloudkms/google_kms_crypto_key_iam_member.json
var googlekmsCryptoKeyIamMember []byte

//go:embed  mapping/google/resource/cloudkms/google_kms_crypto_key_iam_policy.json
var googlekmsCryptoKeyIamPolicy []byte

//go:embed  mapping/google/resource/dns/google_dns_managed_zone.json
var googleDnsmanagedZone []byte

//go:embed  mapping/google/resource/dns/google_dns_policy.json
var googleDnsPolicy []byte

//go:embed  mapping/google/resource/dns/google_dns_record_set.json
var googleDnsRecordSet []byte

//go:embed  mapping/google/resource/iam/google_service_account_iam_binding.json
var googleServiceAccountIamBinding []byte

//go:embed  mapping/google/resource/iam/google_service_account_iam_member.json
var googleServiceAccountIamMember []byte

//go:embed  mapping/google/resource/bigtable/google_bigtable_instance.json
var googleBigtableInstance []byte

//go:embed  mapping/google/resource/bigtable/google_bigtable_instance.json
var googleComputeRegionSslCertificate []byte

//go:embed  mapping/google/resource/bigtable/google_bigtable_table.json
var googleBigtableTable []byte

//go:embed  mapping/google/resource/bigtable/google_bigtable_instance_iam.json
var googleBigTableInstanceIam []byte

//go:embed  mapping/google/resource/bigtable/google_bigtable_table_iam.json
var googleBigTableTableIam []byte

//go:embed  mapping/google/resource/pubsub/google_pubsub_schema.json
var googlePubsubSchema []byte

//go:embed  mapping/google/resource/pubsub/google_pubsub_subscription.json
var googlePubsubSubscription []byte
