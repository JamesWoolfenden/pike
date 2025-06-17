package pike

import (
	"errors"
	"reflect"
	"testing"

	Identity "github.com/jameswoolfenden/identity/src"
	"github.com/stretchr/testify/assert"
)

func Test_compareAllowError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"invoke", fields{errors.New("basic")}, "compare allow error basic",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &compareAllowError{
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_getIAMError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"invoke", fields{errors.New("test")}, "get IAM error test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &getIAMError{
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_identityParseError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"invoke", fields{errors.New("test")}, "Identity parsing error test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &identityParseError{
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

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

func TestPolicyDiff_Empty(t *testing.T) {
	tests := []struct {
		name string
		pd   PolicyDiff
		want bool
	}{
		{
			name: "both nil",
			pd:   PolicyDiff{nil, nil},
			want: true,
		},
		{
			name: "empty slices",
			pd:   PolicyDiff{[]string{}, []string{}},
			want: true,
		},
		{
			name: "over nil under empty",
			pd:   PolicyDiff{nil, []string{}},
			want: true,
		},
		{
			name: "over empty under nil",
			pd:   PolicyDiff{[]string{}, nil},
			want: true,
		},
		{
			name: "over with content",
			pd:   PolicyDiff{[]string{"s3:GetObject"}, nil},
			want: false,
		},
		{
			name: "under with content",
			pd:   PolicyDiff{nil, []string{"s3:PutObject"}},
			want: false,
		},
		{
			name: "both with content",
			pd:   PolicyDiff{[]string{"s3:GetObject"}, []string{"s3:PutObject"}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := len(tt.pd.Over) == 0 && len(tt.pd.Under) == 0
			if got != tt.want {
				t.Errorf("PolicyDiff empty check = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicyDiff_Equal(t *testing.T) {
	tests := []struct {
		name     string
		first    PolicyDiff
		second   PolicyDiff
		wantSame bool
	}{
		{
			name:     "identical empty",
			first:    PolicyDiff{},
			second:   PolicyDiff{},
			wantSame: true,
		},
		{
			name:     "identical with content",
			first:    PolicyDiff{[]string{"s3:GetObject"}, []string{"s3:PutObject"}},
			second:   PolicyDiff{[]string{"s3:GetObject"}, []string{"s3:PutObject"}},
			wantSame: true,
		},
		{
			name:     "different over",
			first:    PolicyDiff{[]string{"s3:GetObject"}, []string{"s3:PutObject"}},
			second:   PolicyDiff{[]string{"s3:ListBucket"}, []string{"s3:PutObject"}},
			wantSame: false,
		},
		{
			name:     "different under",
			first:    PolicyDiff{[]string{"s3:GetObject"}, []string{"s3:PutObject"}},
			second:   PolicyDiff{[]string{"s3:GetObject"}, []string{"s3:DeleteObject"}},
			wantSame: false,
		},
		{
			name:     "different lengths",
			first:    PolicyDiff{[]string{"s3:GetObject"}, []string{"s3:PutObject"}},
			second:   PolicyDiff{[]string{"s3:GetObject"}, []string{"s3:PutObject", "s3:DeleteObject"}},
			wantSame: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reflect.DeepEqual(tt.first, tt.second)
			if got != tt.wantSame {
				t.Errorf("PolicyDiff equality check = %v, want %v", got, tt.wantSame)
			}
		})
	}
}
