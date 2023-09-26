package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/google/data/iam/google_service_account.json
var dataGoogleServiceAccount []byte

//go:embed mapping/google/data/compute/google_compute_network.json
var dataGoogleComputeNetwork []byte

//go:embed mapping/google/data/compute/google_compute_subnetwork.json
var dataGoogleComputeSubnetwork []byte

//go:embed mapping/google/data/compute/google_compute_zones.json
var dataGoogleComputeZones []byte

//go:embed mapping/google/data/resourcemanager/google_project.json
var dataGoogleProject []byte

//go:embed mapping/google/data/cloudkms/google_kms_key_ring.json
var dataGoogleKmsKeyRing []byte

//go:embed  mapping/google/data/cloudkms/google_kms_crypto_key.json
var dataGoogleKmsCryptoKey []byte

//go:embed  mapping/google/data/dns/google_dns_keys.json
var dataGoogleDnsKeys []byte

//go:embed  mapping/google/data/dns/google_dns_managed_zone.json
var dataGoogleDnsManagedZone []byte

//go:embed  mapping/google/data/dns/google_dns_managed_zone_iam_policy.json
var dataGoogleDnsManagedZoneIamPolicy []byte

//go:embed  mapping/google/data/dns/google_dns_record_set.json
var dataGoogleDnsRecordSet []byte

//go:embed  mapping/google/data/artifactregistry/google_artifact_registry_repository.json
var dataGoogleArtifactRegistryRepository []byte

//go:embed  mapping/google/data/artifactregistry/google_artifact_registry_repository_iam_policy.json
var dataGoogleArtifactRegistryRepositoryIamPolicy []byte
