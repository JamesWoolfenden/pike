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
	Name string
}

func (m *repositoryFormatError) Error() string {
	return fmt.Sprintf("repository not formatted correctly %s", m.Name)
}

type gitReferenceError struct {
	Name string
}

func (m *gitReferenceError) Error() string {
	return fmt.Sprintf("git reference in module source path unsupported %s", m.Name)
}

type backendExistsError struct {
}

func (m *backendExistsError) Error() string {
	return fmt.Sprintf("no Backend found")
}
