package pike

import "log"

// GetGCPDataPermissions gets permissions required for datasources
func GetGCPDataPermissions(result ResourceV2) []string {

	TFLookup := map[string]interface{}{}

	var Permissions []string

	temp := TFLookup[result.Name]
	if temp != nil {
		Permissions = GetPermissionMap(TFLookup[result.Name].([]byte), result.Attributes)
	} else {
		log.Printf("data.%s not implemented", result.Name)
	}

	return Permissions
}
