package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/gcp/resource/compute/google_compute_instance.json
var googleComputeInstance []byte

//go:embed mapping/gcp/resource/storage/google_storage_bucket.json
var googleStorageBucket []byte

//go:embed mapping/gcp/resource/storage/google_storage_bucket_object.json
var googleStorageBucketObject []byte

//go:embed mapping/gcp/resource/compute/google_compute_network.json
var googleComputeNetwork []byte

//go:embed mapping/gcp/resource/compute/google_compute_subnetwork.json
var googleComputeSubnetwork []byte

//go:embed mapping/gcp/resource/compute/google_compute_address.json
var googleComputeAddress []byte

//go:embed mapping/gcp/resource/compute/google_compute_firewall.json
var googleComputeFirewall []byte

//go:embed mapping/gcp/resource/compute/google_compute_instance_template.json
var googleComputeInstanceTemplate []byte

//go:embed mapping/gcp/resource/compute/google_compute_project_metadata_item.json
var googleComputeProjectMetadataItem []byte

//go:embed mapping/gcp/resource/cloudfunctions/google_cloudfunctions_function.json
var googleCloudfunctionsFunction []byte

//go:embed mapping/gcp/resource/cloudfunctions/google_cloudfunctions_function_iam_policy.json
var googleCloudfunctionsFunctionIamPolicy []byte

//go:embed mapping/gcp/resource/iam/google_project_iam_custom_role.json
var googleProjectIamCustomRole []byte

//go:embed mapping/gcp/resource/iam/google_sevice_account.json
var googleServiceAccount []byte

//go:embed mapping/gcp/resource/cloudkms/google_kms_key_ring.json
var googleKmsKeyRing []byte

//go:embed  mapping/gcp/resource/cloudkms/google_kms_crypto_key.json
var googleKmsCryptoKey []byte

//go:embed  mapping/gcp/resource/storage/google_storage_bucket_acl.json
var googleStorageBucketACL []byte

//go:embed  mapping/gcp/resource/storage/google_storage_bucket_iam_binding.json
var googleStorageBucketIamBinding []byte

//go:embed  mapping/gcp/resource/cloudsql/google_sql_database_instance.json
var googleSQLDatabaseInstance []byte

//go:embed  mapping/gcp/resource/cloudsql/google_sql_database.json
var googleSQLDatabase []byte

//go:embed  mapping/gcp/resource/cloudsql/google_sql_user.json
var googleSQLUser []byte

//go:embed  mapping/gcp/resource/servicenetworking/google_service_networking_connection.json
var googleServiceNetworkingConnection []byte

//go:embed  mapping/gcp/resource/compute/google_compute_global_address.json
var googleComputeGlobalAddress []byte

//go:embed  mapping/gcp/resource/pubsub/google_pubsub_topic.json
var googlePubsubTopic []byte

//go:embed  mapping/gcp/resource/source/google_soucerepo_repository.json
var googleSourcerepoRepository []byte

//go:embed  mapping/gcp/resource/iam/google_service_account_iam_policy.json
var googleServiceAccountIamPolicy []byte

//go:embed  mapping/gcp/resource/iam/google_service_account_key.json
var googleServiceAccountKey []byte

//go:embed  mapping/gcp/resource/resourcemanager/google_project_iam_binding.json
var googleProjectIamBinding []byte

//go:embed  mapping/gcp/resource/container/google_container_cluster.json
var googleContainerCluster []byte

//go:embed  mapping/gcp/resource/container/google_container_node_pool.json
var googleContainerNodePool []byte

//go:embed  mapping/gcp/resource/bigquery/google_bigquery_dataset.json
var googleBigqueryDataset []byte

//go:embed  mapping/gcp/resource/bigquery/google_bigquery_job.json
var googleBigqueryJob []byte
