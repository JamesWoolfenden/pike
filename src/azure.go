package pike

// GetAZUREPermissions for AZURE resources.
func GetAZUREPermissions(result ResourceV2) ([]string, error) {
	if result.TypeName == resource || result.TypeName == terraform {
		return GetAZUREResourcePermissions(result)
	}
	return GetAZUREDataPermissions(result)
}

// privateEndpointApprovalPermissions maps azurerm_private_endpoint's
// subresource_names values to the resource provider's private-endpoint
// connection approval action. Only entries verified against the target
// resource provider's actual operation list belong here (e.g. via
// `az provider operation show --namespace <provider>`) — the action's shape
// is provider-specific, not a guessable pattern, so an unverified guess here
// is worse than the "unknown, add manually" fallback.
var privateEndpointApprovalPermissions = map[string]string{
	"redisCache": "Microsoft.Cache/redis/PrivateEndpointConnectionsApproval/action",
}

// azurePrivateEndpointPermissions adds the target-specific private-endpoint
// connection approval permission based on subresource_names, instead of the
// single hardcoded (and mostly wrong) permission the base mapping used to
// apply to every azurerm_private_endpoint regardless of target. Unrecognised
// or absent subresource_names add no permission here, matching the
// "unknown, add manually" fallback rather than a guessed one.
func azurePrivateEndpointPermissions(result ResourceV2, found []string) []string {
	for _, name := range result.AttributeValues["subresource_names"] {
		if permission, ok := privateEndpointApprovalPermissions[name]; ok && !Contains(found, permission) {
			found = append(found, permission)
		}
	}

	return found
}

// GetAZUREResourcePermissions looks up permissions required for resources.
func GetAZUREResourcePermissions(result ResourceV2) ([]string, error) {
	temp := AzureLookup(result.Name)
	if temp == nil {
		return nil, &notImplementedResourceError{result.Name}
	}

	found, err := GetPermissionMap(temp, result.Attributes, result.Name)
	if err != nil {
		return found, err
	}

	if result.Name == "azurerm_private_endpoint" {
		found = azurePrivateEndpointPermissions(result, found)
	}

	return found, nil
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
