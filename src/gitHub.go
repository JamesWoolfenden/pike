package pike

import (
	"errors"
	"log"
	"strconv"

	"github.com/google/go-github/v47/github"
)

// InvokeGithubDispatchEvent uses your gitHub api key (if sufficiently enabled) to invoke a github action workflow
func InvokeGithubDispatchEvent(repository string, workflowFileName string, branch string) error {
	owner, repo, err := splitHub(repository)

	if err != nil {
		return err
	}

	ctx, client := getGithubClient()

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
		log.Print(response)
		return err
	}

	myResponse := *response.Response
	Header := myResponse.Header
	remains := Header["X-Ratelimit-Remaining"]

	if left, err := strconv.Atoi(remains[0]); err == nil {
		if left == 0 {
			return errors.New("you are being gitHub rate limited")
		}
		log.Printf("Github rate limit: %s", remains[0])
	}

	return nil
}
