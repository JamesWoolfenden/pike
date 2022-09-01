package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/gcp/data/iam/google_sevice_account.json
var dataGoogleServiceAccount []byte

//go:embed mapping/gcp/data/compute/google_compute_network.json
var dataGoogleComputeNetwork []byte
