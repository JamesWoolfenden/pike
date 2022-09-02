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
