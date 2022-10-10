package pike

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/go-github/v47/github"
)

func TestSetRepoSecret(t *testing.T) {
	type args struct {
		repository string
		keyText    string
		keyName    string
	}
	tests := []struct {
		name    string
		args    args
		want    *github.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetRepoSecret(tt.args.repository, tt.args.keyText, tt.args.keyName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetRepoSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRepoSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGithubClient(t *testing.T) {
	tests := []struct {
		name  string
		want  context.Context
		want1 *github.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getGithubClient()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGithubClient() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getGithubClient() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getPublicKeyDetails(t *testing.T) {
	type args struct {
		owner      string
		repository string
	}
	tests := []struct {
		name        string
		args        args
		wantKeyID   string
		wantPkValue string
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKeyID, gotPkValue, err := getPublicKeyDetails(tt.args.owner, tt.args.repository)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPublicKeyDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKeyID != tt.wantKeyID {
				t.Errorf("getPublicKeyDetails() gotKeyID = %v, want %v", gotKeyID, tt.wantKeyID)
			}
			if gotPkValue != tt.wantPkValue {
				t.Errorf("getPublicKeyDetails() gotPkValue = %v, want %v", gotPkValue, tt.wantPkValue)
			}
		})
	}
}

func Test_encryptPlaintext(t *testing.T) {
	type args struct {
		plaintext    string
		publicKeyB64 string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encryptPlaintext(tt.args.plaintext, tt.args.publicKeyB64)
			if (err != nil) != tt.wantErr {
				t.Errorf("encryptPlaintext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encryptPlaintext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemote(t *testing.T) {
	type args struct {
		target     string
		repository string
		region     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Remote(tt.args.target, tt.args.repository, tt.args.region); (err != nil) != tt.wantErr {
				t.Errorf("Remote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_splitHub(t *testing.T) {
	type args struct {
		repository string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := splitHub(tt.args.repository)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitHub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("splitHub() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("splitHub() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
