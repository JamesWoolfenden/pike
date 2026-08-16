package pike

// GetAZUREPermissions for AZURE resources.
func GetAZUREPermissions(result ResourceV2) ([]string, error) {
	if result.TypeName == resource || result.TypeName == terraform {
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

func getAZUREPlanPermissions(result ResourceV2) ([]string, error) {
	switch result.TypeName {
	case resource, terraform:
		temp := AzureLookup(result.Name)
		if temp == nil {
			return nil, &notImplementedResourceError{result.Name}
		}
		return getPlanPermissionMap(temp, result.Attributes, result.Name, isAzureReadPerm)
	case data:
		temp := AzureDataLookup(result.Name)
		if temp == nil {
			return nil, &notImplementedDatasourceError{Name: result.Name}
		}
		return getPlanPermissionMap(temp, result.Attributes, result.Name, isAzureReadPerm)
	default:
		return nil, nil
	}
}

func AzureLookup(name string) []byte {
	return azureResourceMap[name]
}
