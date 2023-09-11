package pike

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v47/github"
	"github.com/rs/zerolog/log"
)

// InvokeGithubDispatchEvent uses your gitHub api key (if sufficiently enabled) to invoke a gitHub action workflow
func InvokeGithubDispatchEvent(repository string, workflowFileName string, branch string) error {
	owner, repo, err := SplitHub(repository)
	if err != nil {
		log.Print(err)

		return fmt.Errorf("failed to SplitHub %w", err)
	}

	url := "https://api.github.com/repos/" + owner + "/" + repo + "/actions/workflows/" + workflowFileName

	err2 := VerifyURL(url)
	if err2 != nil {
		log.Print(err2)

		return err2
	}

	ctx, client := GetGithubClient()

	err3 := VerifyBranch(client, owner, repo, branch)
	if err3 != nil {
		log.Print(err3)

		return err3
	}

	event := github.CreateWorkflowDispatchEventRequest{
		Ref: branch,
	}

	response, err := client.Actions.CreateWorkflowDispatchEventByFileName(
		ctx,
		owner,
		repo,
		workflowFileName,
		event)
	if err != nil {
		log.Printf("invoke failed %s", response.Response.Status)

		return err
	}

	myResponse := *response.Response
	Header := myResponse.Header
	remains := Header["X-Ratelimit-Remaining"]

	if left, err := strconv.Atoi(remains[0]); err == nil {
		if left == 0 {
			return errors.New("you are being gitHub rate limited")
		}
		log.Printf("Invoked: Github rate limit remaining: %s", remains[0])
	}

	return nil
}

// VerifyBranch checks that a branch exists in a repo
func VerifyBranch(client *github.Client, owner string, repo string, branch string) error {
	ctx := context.Background()
	branches, _, err := client.Repositories.ListBranches(ctx, owner, repo, nil)

	if err != nil {
		return err
	}

	found := false

	for _, stem := range branches {
		temp := *stem
		if *temp.Name == branch {
			found = true
		}
	}
	if found {
		return nil
	}
	return errors.New("branch " + branch + " not found for " + repo)
}

// VerifyURL tests a url
func VerifyURL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("failed to reach %s for %s", url, resp.Status)

		return err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("non ok status code %s for %s", resp.Status, url)
		return errors.New(resp.Status)
	}

	return nil
}
