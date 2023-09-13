package pike

import (
	"fmt"
)

// GetGCPDataPermissions gets permissions required for datasources
func GetGCPDataPermissions(result ResourceV2) ([]string, error) {
	temp := GCPDataLookup(result.Name)

	var (
		Permissions []string
		err         error
	)

	if temp != nil {
		Permissions, err = GetPermissionMap(temp.([]byte), result.Attributes)
	} else {
		return nil, fmt.Errorf("data.%s not implemented", result.Name)
	}

	return Permissions, err
}

func GCPDataLookup(result string) interface{} {
	TFLookup := map[string]interface{}{
		"google_service_account":    dataGoogleServiceAccount,
		"google_compute_image":      placeholder,
		"google_compute_network":    dataGoogleComputeNetwork,
		"google_compute_subnetwork": dataGoogleComputeSubnetwork,
		"google_compute_zones":      dataGoogleComputeZones,
		"google_project":            dataGoogleProject,
		"google_iam_policy":         placeholder,
		"google_iam_role":           placeholder,
		"google_kms_crypto_key":     dataGoogleKmsCryptoKey,
		"google_kms_key_ring":       dataGoogleKmsKeyRing,
	}
	return TFLookup[result]
}
