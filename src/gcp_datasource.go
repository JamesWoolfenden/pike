package pike

// GetGCPDataPermissions gets permissions required for datasources.
func GetGCPDataPermissions(result ResourceV2) ([]string, error) {
	temp := GCPDataLookup(result.Name)
	if temp == nil {
		return nil, &notImplementedDatasourceError{result.Name}
	}
	return GetPermissionMap(temp, result.Attributes, result.Name)
}

func GCPDataLookup(name string) []byte {
	return gcpDataMap[name]
}
