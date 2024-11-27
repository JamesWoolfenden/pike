package pike

import "fmt"

type notImplementedResourceError struct {
	Name string
}

func (m *notImplementedResourceError) Error() string {
	return fmt.Sprintf("not implemented %s", m.Name)
}

type notImplementedDatasourceError struct {
	Name string
}

func (m *notImplementedDatasourceError) Error() string {
	return fmt.Sprintf("data.%s not implemented", m.Name)
}

type unknownPermissionError struct {
	Name string
}

func (m *unknownPermissionError) Error() string {
	return fmt.Sprintf("unknown permission resource type %s", m.Name)
}

type repositoryFormatError struct {
	name string
}

func (m *repositoryFormatError) Error() string {
	return fmt.Sprintf("repository not formatted correctly %s", m.name)
}

type gitReferenceError struct {
	name string
}

func (m *gitReferenceError) Error() string {
	return fmt.Sprintf("git reference in module source path unsupported %s", m.name)
}

type backendExistsError struct {
}

func (m *backendExistsError) Error() string {
	return "backend already exists"
}

type mappingsEmpty struct{}

func (m *mappingsEmpty) Error() string {
	return "mappings are empty"
}

type invalidJsonError struct{}

func (m *invalidJsonError) Error() string {
	return "invalid json, was empty or corrupt"
}

type zeroLengthAttributesError struct {
	resource string
}

func (m *zeroLengthAttributesError) Error() string {
	return fmt.Sprintf("no attributes provided for resource %s", m.resource)
}

type emptyTypeNameError struct{}

func (m *emptyTypeNameError) Error() string {
	return "TypeName cannot be empty"
}

type emptyNameError struct{}

func (m *emptyNameError) Error() string {
	return "Name cannot be empty"
}

type assertionFailedError struct {
	message string
	err     error
}

func (m *assertionFailedError) Error() string {
	return fmt.Sprintf("assertion failed: %s %v", m.message, m.err)
}

type getAWSResourcePermissionsError struct {
	err error
}

func (m *getAWSResourcePermissionsError) Error() string {
	return fmt.Sprintf("failed to get AWS resource permissions %v", m.err)
}

type getAWSDataPermissionsError struct {
	err error
}

func (m *getAWSDataPermissionsError) Error() string {
	return fmt.Sprintf("failed to get AWS data permissions %v", m.err)
}

type unmarshallJsonError struct {
	err      error
	resource string
}

func (m *unmarshallJsonError) Error() string {
	return fmt.Sprintf("failed to unmarshal json %v for %s", m.err, m.resource)
}

type attributesFieldMissingError struct {
}

func (m *attributesFieldMissingError) Error() string {
	return "attributes field missing"
}

type assertionError struct {
	message string
}

func (m *assertionError) Error() string {
	return fmt.Sprintf("assertion failed for: %s", m.message)
}

type templateParseError struct {
	err error
}

func (m *templateParseError) Error() string {
	return fmt.Sprintf("failed to parse template %v", m.err)
}

type templateExecuteError struct {
	err error
}

func (m *templateExecuteError) Error() string {
	return fmt.Sprintf("failed to execute template %v", m.err)
}
