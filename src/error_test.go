package pike

import "testing"

func Test_notImplementedResourceError_Error(t *testing.T) {
	type fields struct {
		Name string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Pass", fields{Name: "test"}, "not implemented test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &notImplementedResourceError{
				Name: tt.fields.Name,
			}
			if got := m.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_notImplementedDatasourceError_Error(t *testing.T) {
	type fields struct {
		Name string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"pass", fields{Name: "test"}, "data.test not implemented"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &notImplementedDatasourceError{
				Name: tt.fields.Name,
			}
			if got := m.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unknownPermissionError_Error(t *testing.T) {
	type fields struct {
		Name string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"pass", fields{Name: "test"}, "unknown permission resource type test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &unknownPermissionError{
				Name: tt.fields.Name,
			}
			if got := m.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
