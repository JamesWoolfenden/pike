package pike

//goland:noinspection GoLinter

// GetAWSDataPermissions gets permissions required for datasource's.
//
//goland:noinspection GoLinter
func GetAWSDataPermissions(result ResourceV2) ([]string, error) {
	var (
		Permissions []string
		err         error
	)

	if temp := AwsDataLookup(result.Name); temp != nil {
		Permissions, err = GetPermissionMap(temp, result.Attributes, result.Name)
	} else {
		return nil, &notImplementedDatasourceError{result.Name}
	}

	return Permissions, err
}

func AwsDataLookup(name string) []byte {
	return awsDataMap[name]
}
