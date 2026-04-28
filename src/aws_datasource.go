package pike

// GetAWSDataPermissions gets permissions required for datasources.
func GetAWSDataPermissions(result ResourceV2) ([]string, error) {
	temp := AwsDataLookup(result.Name)
	if temp == nil {
		return nil, &notImplementedDatasourceError{result.Name}
	}
	return GetPermissionMap(temp, result.Attributes, result.Name)
}

func AwsDataLookup(name string) []byte {
	return awsDataMap[name]
}
