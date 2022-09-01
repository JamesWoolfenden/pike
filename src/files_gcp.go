package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/gcp/resource/compute/google_compute_instance.json
var googleComputeInstance []byte

//go:embed mapping/gcp/resource/storage/google_storage_bucket.json
var googleStorageBucket []byte
