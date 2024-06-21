package coverage

import (
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func Test_coverageAws(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "pass", wantErr: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := coverageAWS(); (err != nil) != tt.wantErr {
				t.Errorf("coverageAWS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_percent(t *testing.T) {
	t.Parallel()

	type args struct {
		missing []string
		data    []string
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Pass", args{[]string{"a", "b"}, []string{"a", "b", "c"}}, 33.33333333333334},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := percent(tt.args.missing, tt.args.data)

			if !pike.AlmostEqual(got, tt.want) {
				t.Errorf("percent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coverageAzure(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "pass", wantErr: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if err := coverageAzure(); (err != nil) != tt.wantErr {
				t.Errorf("coverageAzure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_coverageGcp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "pass", wantErr: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if err := coverageGcp(); (err != nil) != tt.wantErr {
				t.Errorf("coverageGcp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
