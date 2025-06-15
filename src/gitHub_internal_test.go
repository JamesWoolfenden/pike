package pike

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyURL(t *testing.T) {
	type args struct {
		url string
	}
	//goland:noinspection HttpUrlsUsage
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"google", args{"www.google.com"}, true},
		{"http", args{"http://www.google.com"}, true},
		{"https", args{"https://www.google.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := verifyURL(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("verifyURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gitHubRateLimitingError_Error(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Invoke", "you are being GitHub Rate-limited"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &gitHubRateLimitingError{}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_insecureProtocolError_Error(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Invoke", "insecure protocol"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &insecureProtocolError{}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_nilResponseError_Error(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Invoke", "nil response"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &nilResponseError{}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_nonSuccessError_Error(t *testing.T) {
	type fields struct {
		response string
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Invoke", fields{"test", errors.New("test")}, "non success response test test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &nonSuccessError{
				response: tt.fields.response,
				err:      tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_verifyBranchError_Error(t *testing.T) {
	type fields struct {
		branch string
		repo   string
		owner  string
		err    error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Invoke", fields{"branchtest", "repotest", "ownertest", errors.New("test")}, "failed to verify branch branchtest repotest ownertest test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &verifyBranchError{
				branch: tt.fields.branch,
				repo:   tt.fields.repo,
				owner:  tt.fields.owner,
				err:    tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_verifyURLError_Error(t *testing.T) {
	type fields struct {
		url string
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Invoke", fields{"https://www.google.com", errors.New("test")}, "failed to verify URL https://www.google.com test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &verifyURLError{
				url: tt.fields.url,
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_workflowInvokeError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Invoke", fields{errors.New("test")}, "failed to invoke workflow test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &workflowInvokeError{
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}
