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
			false,
		},
		{
			"pass not empty",
			args{identity, morePolicy},
			PolicyDiff{nil, []string{"s3:*", "s3-object-lambda:*"}},
			false,
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
			PolicyDiff{[]string{"s3:*", "s3-object-lambda:*"}, nil},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompareAllow(tt.args.identity, tt.args.policy)
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

func Test_contains(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"found", args{[]string{"foo", "bar"}, "foo"}, true},
		{"not found", args{[]string{"foo", "bar"}, "bart"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
