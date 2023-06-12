package pike_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/go-github/v47/github"
	pike "github.com/jameswoolfenden/pike/src"
)

func TestSetRepoSecret(t *testing.T) {
	t.Parallel()

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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := pike.SetRepoSecret(tt.args.repository, tt.args.keyText, tt.args.keyName)
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
	t.Parallel()

	tests := []struct {
		name  string
		want  context.Context
		want1 *github.Client
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, got1 := pike.GetGithubClient()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGithubClient() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetGithubClient() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getPublicKeyDetails(t *testing.T) {
	t.Parallel()

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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotKeyID, gotPkValue, err := pike.GetPublicKeyDetails(tt.args.owner, tt.args.repository)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPublicKeyDetails() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if gotKeyID != tt.wantKeyID {
				t.Errorf("GetPublicKeyDetails() gotKeyID = %v, want %v", gotKeyID, tt.wantKeyID)
			}
			if gotPkValue != tt.wantPkValue {
				t.Errorf("GetPublicKeyDetails() gotPkValue = %v, want %v", gotPkValue, tt.wantPkValue)
			}
		})
	}
}

func Test_encryptPlaintext(t *testing.T) {
	t.Parallel()

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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := pike.EncryptPlaintext(tt.args.plaintext, tt.args.publicKeyB64)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptPlaintext() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncryptPlaintext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemote(t *testing.T) {
	t.Parallel()

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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := pike.Remote(tt.args.target, tt.args.repository, tt.args.region); (err != nil) != tt.wantErr {
				t.Errorf("Remote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_splitHub(t *testing.T) {
	t.Parallel()

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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, got1, err := pike.SplitHub(tt.args.repository)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitHub() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("SplitHub() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SplitHub() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
