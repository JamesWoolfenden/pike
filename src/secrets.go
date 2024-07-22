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

// Remote updates a repo with AWS credentials.
func Remote(target string, repository string, region string) error {
	iamRole, err := Make(target)

	const magic = 5

	time.Sleep(magic * time.Second)

	if err != nil {
		return err
	}

	Creds, err2 := getAWSCredentials(*iamRole, region)

	if err2 != nil {
		return err2
	}

	myCredentials := Creds.Credentials

	_, err = SetRepoSecret(repository, *myCredentials.AccessKeyId, "AWS_ACCESS_KEY_ID")

	if err != nil {
		var response *github.ErrorResponse
		errors.As(err, &response)
		log.Printf("failed to set repo secrets: %s for repository %s", response.Message, repository)

		return fmt.Errorf("failed to set repo secrets: %s for repository %s", response.Message, repository)
	}

	_, err = SetRepoSecret(repository, *myCredentials.SecretAccessKey, "AWS_SECRET_ACCESS_KEY")
	if err != nil {
		return err
	}

	_, err = SetRepoSecret(repository, *myCredentials.SessionToken, "AWS_SESSION_TOKEN")

	if err != nil {
		return err
	}

	return nil
}

// SetRepoSecret sets an encrypted gitHub action secret.
func SetRepoSecret(repository string, keyText string, keyName string) (*github.Response, error) {
	owner, repo, err2 := SplitHub(repository)
	if err2 != nil {
		return nil, err2
	}

	keyID, publicKey, err := GetPublicKeyDetails(owner, repo)
	if err != nil {
		return nil, err
	}

	encryptedBytes, err := EncryptPlaintext(keyText, publicKey)
	if err != nil {
		return nil, err
	}

	encryptedValue := base64.StdEncoding.EncodeToString(encryptedBytes)

	// Create an EncryptedSecret and encrypt the plaintext value into it
	eSecret := &github.EncryptedSecret{
		Name:           keyName,
		KeyID:          keyID,
		EncryptedValue: encryptedValue,
	}

	ctx, client := GetGithubClient()

	response, err := client.Actions.CreateOrUpdateRepoSecret(ctx, owner, repo, eSecret)
	if err != nil {
		return response, err
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
		errString := fmt.Sprintf("repository not formatted correctly %s", repository)
		return "", "", errors.New(errString)
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
		&oauth2.Token{AccessToken: token},
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
		return "", "", err
	}

	return publicKey.GetKeyID(), publicKey.GetKey(), nil
}

// EncryptPlaintext standard encryption.
func EncryptPlaintext(plaintext string, publicKeyB64 string) ([]byte, error) {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyB64)
	if err != nil {
		return nil, err
	}

	var publicKeyBytes32 [32]byte
	copiedLen := copy(publicKeyBytes32[:], publicKeyBytes)

	if copiedLen == 0 {
		return nil, fmt.Errorf("could not convert publicKey to bytes")
	}

	plaintextBytes := []byte(plaintext)

	var encryptedBytes []byte

	cipherText, err := box.SealAnonymous(encryptedBytes, plaintextBytes, &publicKeyBytes32, nil)
	if err != nil {
		return nil, err
	}

	return cipherText, nil
}
