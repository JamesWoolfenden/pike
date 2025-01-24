package pike

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v47/github"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/nacl/box"
	"golang.org/x/oauth2"
)

type awsCredentialsError struct {
	err error
}

func (e *awsCredentialsError) Error() string {
	return fmt.Sprintf("failed to get AWS credentials: %v", e.err)
}

// Remote updates a repo with AWS credentials.
func Remote(target string, repository string, region string) error {
	iamRole, err := Make(target)
	if err != nil {
		return &makeRoleError{err}
	}

	const magic = 5

	time.Sleep(magic * time.Second)

	Credentials, err := getAWSCredentials(*iamRole, region)
	if err != nil {
		return &awsCredentialsError{err}
	}

	myCredentials := Credentials.Credentials

	_, err = SetRepoSecret(repository, *myCredentials.AccessKeyId, "AWS_ACCESS_KEY_ID")
	if err != nil {
		var response *github.ErrorResponse

		errors.As(err, &response)

		log.Info().Msgf("failed to set repo secrets: %s for repository %s", response.Message, repository)

		return &setRepoSecretError{repository, err}
	}

	_, err = SetRepoSecret(repository, *myCredentials.SecretAccessKey, "AWS_SECRET_ACCESS_KEY")
	if err != nil {
		return &setRepoSecretError{repository, err}
	}

	_, err = SetRepoSecret(repository, *myCredentials.SessionToken, "AWS_SESSION_TOKEN")
	if err != nil {
		return &setRepoSecretError{repository, err}
	}

	return nil
}

// SetRepoSecret sets an encrypted GitHub action secret.
func SetRepoSecret(repository string, keyText string, keyName string) (*github.Response, error) {
	owner, repo, err := SplitHub(repository)
	if err != nil {
		return nil, &splitHubError{err: err}
	}

	keyID, publicKey, err := GetPublicKeyDetails(owner, repo)
	if err != nil {
		return nil, &getPublicKeyDetailsError{err}
	}

	encryptedBytes, err := EncryptPlaintext(keyText, publicKey)
	if err != nil {
		return nil, &encryptPlaintextError{err: err}
	}

	encryptedValue := base64.StdEncoding.EncodeToString(encryptedBytes)

	// Create an EncryptedSecret and encrypt the plaintext value into it
	eSecret := &github.EncryptedSecret{ // permit
		Name:           keyName,
		KeyID:          keyID,
		EncryptedValue: encryptedValue,
	}

	ctx, client := GetGithubClient()

	response, err := client.Actions.CreateOrUpdateRepoSecret(ctx, owner, repo, eSecret)
	if err != nil {
		return response, &updateSecretError{err: err}
	}

	return response, nil
}

// SplitHub return details from url.
func SplitHub(repository string) (string, string, error) {
	Splitter := strings.Split(repository, "/")

	var owner string

	var repo string

	switch len(Splitter) {
	case 2:
		{
			owner = Splitter[0]
			repo = Splitter[1]
		}
	case 5:
		{
			owner = Splitter[3]
			repo = Splitter[4]
		}
	default:
		return "", "", &repositoryFormatError{repository}
	}

	return owner, repo, nil
}

// GetGithubClient instantiate and return a client object for GitHub.
func GetGithubClient() (context.Context, *github.Client) {
	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		token = os.Getenv("GITHUB_API")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token}, // permit
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return ctx, client
}

// GetPublicKeyDetails obtains the public key of the owner.
func GetPublicKeyDetails(owner string, repository string) (string, string, error) {
	ctx, client := GetGithubClient()

	publicKey, _, err := client.Actions.GetRepoPublicKey(ctx, owner, repository)
	if err != nil {
		return "", "", &getPublicKeyDetailsError{err}
	}

	return publicKey.GetKeyID(), publicKey.GetKey(), nil
}

// EncryptPlaintext standard encryption.
func EncryptPlaintext(plaintext string, publicKeyB64 string) ([]byte, error) {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyB64)
	if err != nil {
		return nil, &decodeStringError{err: err}
	}

	var publicKeyBytes32 [32]byte
	copiedLen := copy(publicKeyBytes32[:], publicKeyBytes)

	if copiedLen == 0 {
		return nil, &emptyKeyError{}
	}

	plaintextBytes := []byte(plaintext)

	var encryptedBytes []byte

	cipherText, err := box.SealAnonymous(encryptedBytes, plaintextBytes, &publicKeyBytes32, nil)
	if err != nil {
		return nil, &encryptError{err: err}
	}

	return cipherText, nil
}
