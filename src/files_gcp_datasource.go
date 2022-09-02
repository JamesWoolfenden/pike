package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/gcp/data/iam/google_sevice_account.json
var dataGoogleServiceAccount []byte

//go:embed mapping/gcp/data/compute/google_compute_network.json
var dataGoogleComputeNetwork []byte

//go:embed mapping/gcp/data/compute/google_compute_subnetwork.json
var dataGoogleComputeSubnetwork []byte

//go:embed mapping/gcp/data/compute/google_compute_zones.json
var dataGoogleComputeZones []byte

//go:embed mapping/gcp/data/resourcemanager/google_project.json
var dataGoogleProject []byte
