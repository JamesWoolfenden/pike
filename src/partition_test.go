package pike

import (
	"testing"
)

func TestExtractPartitionFromARN(t *testing.T) {
	tests := []struct {
		name      string
		arn       string
		want      string
		wantError bool
	}{
		{
			name:      "standard commercial partition",
			arn:       "arn:aws:iam::123456789012:role/MyRole",
			want:      "aws",
			wantError: false,
		},
		{
			name:      "china partition",
			arn:       "arn:aws-cn:iam::123456789012:role/MyRole",
			want:      "aws-cn",
			wantError: false,
		},
		{
			name:      "govcloud partition",
			arn:       "arn:aws-us-gov:iam::123456789012:role/MyRole",
			want:      "aws-us-gov",
			wantError: false,
		},
		{
			name:      "user identity in china",
			arn:       "arn:aws-cn:iam::123456789012:user/testuser",
			want:      "aws-cn",
			wantError: false,
		},
		{
			name:      "group identity in govcloud",
			arn:       "arn:aws-us-gov:iam::123456789012:group/testgroup",
			want:      "aws-us-gov",
			wantError: false,
		},
		{
			name:      "assumed role in china partition",
			arn:       "arn:aws-cn:sts::123456789012:assumed-role/MyRole/session-name",
			want:      "aws-cn",
			wantError: false,
		},
		{
			name:      "empty arn",
			arn:       "",
			want:      "",
			wantError: true,
		},
		{
			name:      "invalid arn format - no colons",
			arn:       "invalid",
			want:      "",
			wantError: true,
		},
		{
			name:      "invalid arn - missing prefix",
			arn:       "notarn:aws:iam::123456789012:role/MyRole",
			want:      "",
			wantError: true,
		},
		{
			name:      "arn with trailing slashes",
			arn:       "arn:aws:iam::123456789012:role/MyRole/",
			want:      "aws",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractPartitionFromARN(tt.arn)
			if (err != nil) != tt.wantError {
				t.Errorf("ExtractPartitionFromARN() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractPartitionFromARN() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNonCommercialPartition(t *testing.T) {
	tests := []struct {
		name      string
		partition string
		want      bool
	}{
		{"standard aws", "aws", false},
		{"china partition", "aws-cn", true},
		{"govcloud partition", "aws-us-gov", true},
		{"empty string", "", false},
		{"custom partition", "aws-custom", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsNonCommercialPartition(tt.partition)
			if got != tt.want {
				t.Errorf("IsNonCommercialPartition() got = %v, want %v", got, tt.want)
			}
		})
	}
}
