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

//go:embed mapping/gcp/resource/compute/google_compute_address.json
var googleComputeAddress []byte

//go:embed mapping/gcp/resource/compute/google_compute_firewall.json
var googleComputeFirewall []byte

//go:embed mapping/gcp/resource/iam/google_project_iam_custom_role.json
var googleProjectIamCustomRole []byte

//go:embed mapping/gcp/resource/iam/google_sevice_account.json
var googleServiceAccount []byte

//go:embed mapping/gcp/resource/cloudkms/google_kms_key_ring.json
var googleKmsKeyRing []byte

//go:embed  mapping/gcp/resource/cloudkms/google_kms_crypto_key.json
var googleKmsCryptoKey []byte

//go:embed  mapping/gcp/resource/storage/google_storage_bucket_acl.json
var googleStorageBucketAcl []byte

//go:embed  mapping/gcp/resource/storage/google_storage_bucket_iam_binding.json
var googleStorageBucketIamBinding []byte
