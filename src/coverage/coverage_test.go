package coverage

import "testing"

func Test_coverageAws(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "pass", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := coverageAWS(); (err != nil) != tt.wantErr {
				t.Errorf("coverageAWS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_percent(t *testing.T) {
	type args struct {
		missing []string
		data    []string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := percent(tt.args.missing, tt.args.data); got != tt.want {
				t.Errorf("percent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coverageAzure(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "pass", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := coverageAzure(); (err != nil) != tt.wantErr {
				t.Errorf("coverageAzure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
