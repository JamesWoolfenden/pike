package pike

// GetAZUREPermissions for AZURE resources.
func GetAZUREPermissions(result ResourceV2) ([]string, error) {
	var (
		err         error
		Permissions []string
	)

	if result.TypeName == resource {
		Permissions, err = GetAZUREResourcePermissions(result)
		if err != nil {
			return Permissions, err
		}
	} else {
		Permissions, err = GetAZUREDataPermissions(result)
		if err != nil {
			return Permissions, err
		}
	}

	return Permissions, err
}

// GetAZUREResourcePermissions looks up permissions required for resources.
func GetAZUREResourcePermissions(result ResourceV2) ([]string, error) {
	var (
		Permissions []string
		err         error
	)

	if temp := AzureLookup(result.Name); temp != nil {
		Permissions, err = GetPermissionMap(temp, result.Attributes, result.Name)
	} else {
		return nil, &notImplementedResourceError{result.Name}
	}

	return Permissions, err
}

func AzureLookup(name string) []byte {
	return azureResourceMap[name]
}
