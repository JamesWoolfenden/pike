//go:build auth
// +build auth

package pike

import (
	"reflect"
	"testing"

	Identity "github.com/jameswoolfenden/identity/src"
)

func TestCompareAllow(t *testing.T) {
	type args struct {
		identity Identity.IAM
		policy   Identity.Policy
	}

	var identity Identity.IAM
	var moreIdentity Identity.IAM
	moreIdentity.Policies = make([]Identity.Policy, 1)
	moreIdentity.Policies[0].Statements = make([]Identity.Statement, 1)
	moreIdentity.Policies[0].Statements[0] = Identity.Statement{
		Sid:      "",
		Effect:   "Allow",
		Action:   []string{"s3:*", "s3-object-lambda:*"},
		Resource: []string{"*"},
	}
	var policy Identity.Policy

	statements := make([]Identity.Statement, 1)
	statements[0] = Identity.Statement{
		Sid:      "",
		Effect:   "Allow",
		Action:   []string{"s3:*", "s3-object-lambda:*"},
		Resource: []string{"*"},
	}

	var morePolicy Identity.Policy
	morePolicy.Statements = statements

	tests := []struct {
		name    string
		args    args
		want    PolicyDiff
		wantErr bool
	}{
		{
			"pass empty",
			args{identity, policy},
			PolicyDiff{},
			true,
		},
		{
			"pass not empty",
			args{identity, morePolicy},
			PolicyDiff{nil, nil},
			true,
		},
		{
			"pass",
			args{moreIdentity, morePolicy},
			PolicyDiff{},
			false,
		},
		{
			"different",
			args{moreIdentity, policy},
			PolicyDiff{nil, nil},
			true,
		},
		//todo more testcases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := compareAllow(tt.args.identity, tt.args.policy)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompareAllow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompareAllow() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInspect(t *testing.T) {
	type args struct {
		directory string
		init      bool
	}

	tests := []struct {
		name    string
		args    args
		want    PolicyDiff
		wantErr bool
	}{
		//	{"Pass", args{"../terraform/aws/backup", false}, []string{"foo", "bar"}, false},
		{"no dir", args{"../terraform/aws/nodir", false}, PolicyDiff{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Inspect(tt.args.directory, tt.args.init)
			if (err != nil) != tt.wantErr {
				t.Errorf("Inspect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inspect() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInspectExtended(t *testing.T) {
	type args struct {
		directory string
		init      bool
	}

	myDiff := PolicyDiff{
		Over: []string{"ssm:DescribePatchBaselines"},
		Under: []string{
			"dynamodb:DeleteItem", "dynamodb:DescribeTable", "dynamodb:GetItem", "dynamodb:PutItem",
			"s3:DeleteObject", "s3:GetObject", "s3:ListBucket", "s3:PutObject",
		},
	}

	tests := []struct {
		name    string
		args    args
		want    PolicyDiff
		wantErr bool
	}{
		//{
		//	name: "empty directory",
		//	args: args{
		//		directory: "",
		//		init:      false,
		//	},
		//	want:    PolicyDiff{},
		//	wantErr: true,
		//},
		{
			// its comparing
			name: "init true",
			args: args{
				directory: "../terraform/aws",
				init:      true,
			},
			want:    myDiff,
			wantErr: false,
		},
		{
			name: "directory with spaces",
			args: args{
				directory: "../terraform/aws/test dir",
				init:      false,
			},
			want:    PolicyDiff{},
			wantErr: true,
		},
		{
			name: "relative path",
			args: args{
				directory: "./test",
				init:      false,
			},
			want:    PolicyDiff{},
			wantErr: true,
		},
		{
			name: "absolute path",
			args: args{
				directory: "/absolute/path/test",
				init:      false,
			},
			want:    PolicyDiff{},
			wantErr: true,
		},
		{
			name: "directory with special chars",
			args: args{
				directory: "../terraform/aws/test@#$",
				init:      false,
			},
			want:    PolicyDiff{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Inspect(tt.args.directory, tt.args.init)
			if (err != nil) != tt.wantErr {
				t.Errorf("Inspect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inspect() got = %v, want %v", got, tt.want)
			}
		})
	}
}
