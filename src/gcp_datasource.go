package pike

// GetGCPDataPermissions gets permissions required for datasources.
func GetGCPDataPermissions(result ResourceV2) ([]string, error) {
	var (
		Permissions []string
		err         error
	)

	if temp := GCPDataLookup(result.Name); temp != nil {
		Permissions, err = GetPermissionMap(temp, result.Attributes, result.Name)
	} else {
		return nil, &notImplementedDatasourceError{result.Name}
	}

	return Permissions, err
}

func GCPDataLookup(name string) []byte {
	return gcpDataMap[name]
}
