package pike

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/google/go-cmp/cmp"
	"github.com/rs/zerolog/log"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	gcpiam "google.golang.org/api/iam/v1"
	"google.golang.org/api/serviceusage/v1"
)

// Compare IAC codebase to AWS policy.
func Compare(directory string, arn string, init bool) (bool, error) {
	var result bool

	valid, err := inputValidationCompare(directory, arn)
	if err != nil {
		return valid, &inputValidationError{err}
	}

	switch *getCloudFromRole(arn) {
	case "arn":
		{
			result, err = compareAWSRole(directory, arn, init)
		}
	case "gcp":
		{
			result, err = compareGCPRole(directory, arn, init)
		}
	}

	return result, err
}

func getCloudFromRole(arn string) *string {
	var result string

	if strings.Contains(arn, "arn:") {
		result = "aws"
		return &result
	}

	if strings.Contains(arn, "roles") {
		result = "gcp"
		return &result
	}

	return nil
}

type apiNotFoundError struct {
	API string
}

func (m *apiNotFoundError) Error() string {
	return fmt.Sprintf("API %s not found", m.API)
}

type apiNotEnabledError struct {
	API string
}

func (m *apiNotEnabledError) Error() string {
	return fmt.Sprintf("API %s not enabled", m.API)
}

func compareGCPRole(directory string, arn string, init bool) (bool, error) {
	iacPolicy, err := MakePermissionBag(directory, nil, init, "")
	if err != nil {
		return false, &getIAMVersionError{err}
	}

	ctx := context.Background()

	var projectID string
	temp, err := GetEnv("GCP_PROJECT")

	projectID = *temp

	if err != nil {
		return false, &EnvVariableNotSetError{"GCP_PROJECT"}
	}

	var API string
	API = "iam.googleapis.com"

	enabled, err := isGCPAPIEnabled(projectID, API)

	if err != nil {
		return enabled, &apiNotFoundError{API}
	}

	if !enabled {
		return enabled, &apiNotEnabledError{API}
	}

	iamService, err := gcpiam.NewService(ctx)
	if err != nil {
		log.Error().Msgf("Failed to create IAM Service %v", err)
	}

	// The resource name of the role in one of the following formats:
	// `roles/{ROLE_NAME}`
	// `organizations/{ORGANIZATION_ID}/roles/{ROLE_NAME}`
	// `projects/{PROJECT_ID}/roles/{ROLE_NAME}`
	err = VerifyGCPRole(arn)

	if err != nil {
		return false, &GcpRoleNotVerified{arn}
	}

	Roles, err := iamService.Roles.Get(arn).Context(ctx).Do()

	if Roles == nil || err != nil {
		log.Error().Msgf("Failed to get role %v", err)
	}

	//reflect.DeepEqual(iacPolicy.GCP, Roles.IncludedPermissions)
	results := cmp.Diff(Roles.IncludedPermissions, iacPolicy.GCP)
	if results != "" {
		results = strings.Replace(results, "[]string{", "", -1)
		results = strings.Replace(results, "}", "", -1)
		fmt.Print("Policy Comparison mismatch mismatch (-excess +needs):")
		fmt.Print(results)
		return true, nil
	}
	return false, nil
}

func isGCPAPIEnabled(projectID string, want string) (bool, error) {
	enabledAPIs, err := listEnabledAPIs(projectID)

	if err != nil {
		log.Error().Msgf("Failed to list enabled APIs %v", err)
	}

	if !slices.Contains(enabledAPIs, want) {
		return false, errors.New("API not enabled")
	}

	return true, nil
}

func compareAWSRole(directory string, arn string, init bool) (bool, error) {
	// Load the Shared AWS Configuration (~/.aws/config)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return false, &awsConfigError{err}
	}

	client := iam.NewFromConfig(cfg)

	version, err := getVersion(client, arn)
	if err != nil {
		return false, &getVersionError{err}
	}

	policy, err := getPolicyVersion(client, arn, *version)
	if err != nil {
		return false, &getPolicyVersionError{err}
	}

	iacPolicy, err := MakePolicy(directory, nil, init, false, "")
	if err != nil {
		return false, &getIAMVersionError{err}
	}

	sorted, err := sortActions(iacPolicy.AWS.JSONOut)
	if err != nil {
		return false, &sortActionsError{iacPolicy.AWS.JSONOut}
	}

	// iam versus iac
	fmt.Printf("IAM Policy %s versus Local %s \n", arn, directory)

	return compareIAMPolicy(*policy, *sorted)
}

func inputValidationCompare(directory string, arn string) (bool, error) {
	if directory == "" {
		return false, &emptyDirectoryError{}
	}

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return false, &directoryNotFoundError{directory}
	}

	if arn == "" {
		return false, &arnEmptyError{}
	}

	if !strings.HasPrefix(arn, "arn:") && !strings.HasPrefix(arn, "roles/") {
		return false, &invalidARNError{arn}
	}

	return false, nil
}

type compareDifferenceError struct {
	err error
}

func (m *compareDifferenceError) Error() string {
	return fmt.Sprintf("compare difference failed: %v", m.err)
}

// compareIAMPolicy takes two IAM policies and compares.
func compareIAMPolicy(policy string, oldPolicy string) (bool, error) {
	differ := diff.New()
	compare, err := differ.Compare([]byte(policy), []byte(oldPolicy))

	if err != nil {
		return false, &compareDifferenceError{err}
	}

	if compare.Modified() {
		return showDifferences(policy, compare)
	}

	return true, nil
}

type formatterError struct {
	err error
}

func (m *formatterError) Error() string {
	return fmt.Sprintf("formatter failed: %v", m.err)
}

func showDifferences(policy string, compare diff.Diff) (bool, error) {
	var aJSON map[string]interface{}
	err := json.Unmarshal([]byte(policy), &aJSON)

	if err != nil {
		return false, &marshallPolicyError{err}
	}

	myConfig := formatter.AsciiFormatterConfig{
		ShowArrayIndex: true,
		Coloring:       true,
	}

	myFormatter := formatter.NewAsciiFormatter(aJSON, myConfig)
	diffString, err := myFormatter.Format(compare)

	if err != nil {
		return false, &formatterError{err}
	}

	fmt.Print(diffString)

	return false, nil
}

func listEnabledAPIs(projectID string) ([]string, error) {
	ctx := context.Background()
	serviceUsageService, err := serviceusage.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create Service Usage client: %v", err)
	}

	parent := fmt.Sprintf("projects/%s", projectID)
	req := serviceUsageService.Services.List(parent).Filter("state:ENABLED")

	var services []string
	if err := req.Pages(ctx, func(page *serviceusage.ListServicesResponse) error {
		for _, service := range page.Services {
			services = append(services, service.Config.Name)
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to list services: %v", err)
	}

	return services, nil
}

type GcpRoleNotVerified struct {
	role string
}

func (e *GcpRoleNotVerified) Error() string {
	log.Info().Msg(
		`The resource name of the role in one of the following formats:")
	roles/{ROLE_NAME}
	organizations/{ORGANIZATION_ID}/roles/{ROLE_NAME}
	projects/{PROJECT_ID}/roles/{ROLE_NAME}`)
	return e.role
}

func VerifyGCPRole(role string) error {
	r, err := regexp.Compile("projects/(.*\\S)/roles/(.*\\S)")
	if err == nil {
		if r.MatchString(role) {
			return nil
		}
	}
	return &GcpRoleNotVerified{role}
}
