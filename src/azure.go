package pike

// GetAZUREPermissions for AZURE resources.
func GetAZUREPermissions(result ResourceV2) ([]string, error) {
	if result.TypeName == resource {
		return GetAZUREResourcePermissions(result)
	}
	return GetAZUREDataPermissions(result)
}

// GetAZUREResourcePermissions looks up permissions required for resources.
func GetAZUREResourcePermissions(result ResourceV2) ([]string, error) {
	temp := AzureLookup(result.Name)
	if temp == nil {
		return nil, &notImplementedResourceError{result.Name}
	}
	return GetPermissionMap(temp, result.Attributes, result.Name)
}

func AzureLookup(name string) []byte {
	return azureResourceMap[name]
}
