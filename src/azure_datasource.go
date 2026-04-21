package pike

// GetAZUREDataPermissions gets permissions required for datasources.
func GetAZUREDataPermissions(result ResourceV2) ([]string, error) {
	var (
		Permissions []string
		err         error
	)

	if temp := AzureDataLookup(result.Name); temp != nil {
		Permissions, err = GetPermissionMap(temp, result.Attributes, result.Name)
	} else {
		return nil, &notImplementedDatasourceError{Name: result.Name}
	}

	return Permissions, err
}

func AzureDataLookup(name string) []byte {
	return azureDataMap[name]
}
