package pike

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v47/github"
	"github.com/rs/zerolog/log"
)

const lastOK = 299

type verifyURLError struct {
	url string
	err error
}

func (m *verifyURLError) Error() string {
	return fmt.Sprintf("failed to verify URL %s %v", m.url, m.err)
}

type verifyBranchError struct {
	branch string
	repo   string
	owner  string
	err    error
}

func (m *verifyBranchError) Error() string {
	return fmt.Sprintf("failed to verify branch %s %s %s %v", m.branch, m.repo, m.owner, m.err)
}

type nilResponseError struct{}

func (m *nilResponseError) Error() string {
	return "nil response"
}

type nonSuccessError struct {
	response string
	err      error
}

func (m *nonSuccessError) Error() string {
	return fmt.Sprintf("non success response %s %v", m.response, m.err)
}

type workflowInvokeError struct {
	err error
}

func (m *workflowInvokeError) Error() string {
	return fmt.Sprintf("failed to invoke workflow %v", m.err)
}

type gitHubRateLimitingError struct{}

func (m *gitHubRateLimitingError) Error() string {
	return "you are being GitHub Rate-limited"
}

type insecureProtocolError struct{}

func (m *insecureProtocolError) Error() string {
	return "insecure protocol"
}

// InvokeGithubDispatchEvent uses your GitHub api key (if sufficiently enabled) to invoke a GitHub action workflow.
func InvokeGithubDispatchEvent(repository string, workflowFileName string, branch string) error {
	owner, repo, err := SplitHub(repository)
	if err != nil {
		log.Print(err)

		return &splitHubError{err}
	}

	url := "https://api.github.com/repos/" + owner + "/" + repo + "/actions/workflows/" + workflowFileName

	err = verifyURL(url)
	if err != nil {
		log.Error().Err(err)

		return &verifyURLError{url, err}
	}

	ctx, client := GetGithubClient()

	err = verifyBranch(client, owner, repo, branch)
	if err != nil {
		log.Error().Err(err)

		return &verifyBranchError{branch, repo, owner, err}
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

	if response == nil {
		return &nilResponseError{}
	}

	if response.StatusCode > lastOK {
		return &nonSuccessError{response.Status, err}
	}

	if err != nil {
		log.Info().Msgf("invoke failed %s", response.Response.Status)

		return &workflowInvokeError{err: err}
	}

	myResponse := *response.Response
	Header := myResponse.Header
	remains := Header["X-Ratelimit-Remaining"]

	if len(remains) != 0 {
		if left, err := strconv.Atoi(remains[0]); err == nil {
			if left == 0 {
				return &gitHubRateLimitingError{}
			}

			log.Info().Msgf("Invoked: Github rate limit remaining: %s", remains[0])
		}
	}

	return nil
}

type listBranchesError struct {
	err error
}

func (m *listBranchesError) Error() string {
	return fmt.Sprintf("failed to list branches %v", m.err)
}

// verifyBranch checks that a branch exists in a repo.
func verifyBranch(client *github.Client, owner string, repo string, branch string) error {
	ctx := context.Background()
	branches, _, err := client.Repositories.ListBranches(ctx, owner, repo, nil)

	if err != nil {
		return &listBranchesError{err}
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

	return &branchNotFoundError{branch, repo}
}

type branchNotFoundError struct {
	branch string
	repo   string
}

func (m *branchNotFoundError) Error() string {
	return fmt.Sprintf("branch %s not found for repo %s", m.branch, m.repo)
}

// verifyURL tests a url.
func verifyURL(url string) error {
	if //goland:noinspection HttpUrlsUsage
	strings.Contains(strings.ToLower(url), "http://") {
		return &insecureProtocolError{}
	}

	resp, err := http.Get(url) //nolint:gosec

	if resp == nil {
		return &nilResponseError{}
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode > lastOK {
		return &nonSuccessError{response: strconv.Itoa(resp.StatusCode), err: err}
	}

	if err != nil {
		log.Info().Msgf("failed to reach %s for %s", url, resp.Status)

		return &verifyURLError{url, err}
	}

	return nil
}
