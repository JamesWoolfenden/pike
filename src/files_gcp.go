package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/gcp/resource/google_compute_instance.json
var googleComputeInstance []byte
