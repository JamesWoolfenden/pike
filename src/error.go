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
