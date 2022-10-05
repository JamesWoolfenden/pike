package pike

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt" //nolint:goimports
	"log"
	"os" //nolint:goimports
	"strings"
	"time"

	"github.com/google/go-github/v47/github"
	"golang.org/x/crypto/nacl/box"
	"golang.org/x/oauth2"
)

// Remote updates a repo with AWS credentials
func Remote(target string, repository string, region string) error {
	iamRole, err := Make(target)

	time.Sleep(5 * time.Second)

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
		return err
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

// SetRepoSecret does what it is named
func SetRepoSecret(repository string, keyText string, keyName string) (*github.Response, error) {

	Splitter := strings.Split(repository, "/")

	if len(Splitter) != 2 {
		errString := fmt.Sprintf("repository not formatted correctly %s", repository)
		return nil, errors.New(errString)
	}

	owner := Splitter[0]
	repo := Splitter[1]

	keyID, publicKey, err := getPublicKeyDetails(owner, repo)

	if err != nil {
		return nil, err
	}

	encryptedBytes, err := encryptPlaintext(keyText, publicKey)

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

	ctx, client := getGithubClient()

	response, err := client.Actions.CreateOrUpdateRepoSecret(ctx, owner, repo, eSecret)

	if err != nil {
		return response, err
	}
	return response, nil
}

func getGithubClient() (context.Context, *github.Client) {
	token := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return ctx, client
}

func getPublicKeyDetails(owner, repository string) (keyID, pkValue string, err error) {
	ctx, client := getGithubClient()

	publicKey, response, err := client.Actions.GetRepoPublicKey(ctx, owner, repository)
	if err != nil {
		log.Print(&response)
		return keyID, pkValue, err
	}

	return publicKey.GetKeyID(), publicKey.GetKey(), err
}

func encryptPlaintext(plaintext, publicKeyB64 string) ([]byte, error) {
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
