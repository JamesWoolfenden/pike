package coverage

import "testing"

func Test_coverage(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Pass"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coverage()
		})
	}
}
