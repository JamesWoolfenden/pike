package pike

// GetAZUREDataPermissions gets permissions required for datasources.
func GetAZUREDataPermissions(result ResourceV2) ([]string, error) {
	temp := AzureDataLookup(result.Name)
	if temp == nil {
		return nil, &notImplementedDatasourceError{Name: result.Name}
	}
	return GetPermissionMap(temp, result.Attributes, result.Name)
}

func AzureDataLookup(name string) []byte {
	return azureDataMap[name]
}
