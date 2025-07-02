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

// Replace hardcoded values with constants
const (
	defaultTimeout   = 30 * time.Second
	gcpIAMAPIName    = "iam.googleapis.com"
	awsARNPrefix     = "arn:"
	gcpProjectPrefix = "projects/"
)

type invalidCloudError struct {
	arn string
}

func (e *invalidCloudError) Error() string {
	return fmt.Sprintf("Invalid Cloud: %v", e.arn)
}

// Compare IAC codebase to AWS policy.
func Compare(directory string, arn string, init bool) (bool, error) {
	var result bool

	result, err := inputValidationCompare(directory, arn)
	if err != nil {
		log.Error().Msgf("Failed to validate input %v", err)
		return false, &inputValidationError{err: err}
	}

	switch getCloudFromRole(arn) {
	case "aws":
		{
			result, err = compareAWSRole(directory, arn, init)
		}
	case "gcp":
		{
			result, err = compareGCPRole(directory, arn, init)
		}
	default:
		{
			err = &invalidCloudError{arn}
		}
	}

	return result, err
}

func getCloudFromRole(arn string) string {

	var result string

	if strings.Contains(arn, "arn:") {
		result = "aws"
	} else if strings.Contains(arn, "projects") {
		result = "gcp"
	} else {
		result = "unknown"
	}

	return result
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
	// The resource name of the role in one of the following formats:
	// `roles/{ROLE_NAME}`
	// `organizations/{ORGANIZATION_ID}/roles/{ROLE_NAME}`
	// `projects/{PROJECT_ID}/roles/{ROLE_NAME}`
	err := verifyGCPRole(arn)

	if err != nil {
		return false, &gcpRoleNotVerified{arn}
	}

	iacPolicy, err := makePermissionBag(directory, nil, init, "")
	if err != nil {
		return false, &getIAMVersionError{err}
	}

	var projectID *string
	projectID, err = GetEnv("GCP_PROJECT")

	if err != nil {
		return false, &EnvVariableNotSetError{"GCP_PROJECT"}
	}

	var API string
	API = gcpIAMAPIName

	enabled, err := isGCPAPIEnabled(*projectID, API)

	if err != nil {
		return enabled, &apiNotFoundError{API}
	}

	if !enabled {
		return enabled, &apiNotEnabledError{API}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	iamService, err := gcpiam.NewService(ctx)
	if err != nil {
		return false, &iamServiceError{err}
	}

	Roles, err := iamService.Roles.Get(arn).Context(ctx).Do()

	if Roles == nil {
		return false, &gcpRoleNotFound{arn}
	}

	if err != nil {
		return false, &gcpIAMRoleError{err}
	}

	return compareGCPPolicy(Roles, iacPolicy), nil
}

func compareGCPPolicy(Roles *gcpiam.Role, iacPolicy Sorted) bool {
	slices.Sort(Roles.IncludedPermissions)
	slices.Sort(iacPolicy.GCP)

	results := cmp.Diff(Unique(Roles.IncludedPermissions), Unique(iacPolicy.GCP))
	if results != "" {
		replacer := strings.NewReplacer("[]string{", "", "}", "")
		results = replacer.Replace(results)
		fmt.Print("Policy Comparison mismatch mismatch (-excess +needs):")
		fmt.Print(results)
		return false
	}
	return true
}

type gcpIAMRoleError struct {
	err error
}

func (m *gcpIAMRoleError) Error() string {
	return fmt.Sprintf("IAM Role Error: %v", m.err)
}

type gcpRoleNotFound struct {
	role string
}

func (e *gcpRoleNotFound) Error() string {
	return fmt.Sprintf("IAM Role Error: %v", e.role)
}

type iamServiceError struct {
	err error
}

func (m *iamServiceError) Error() string {
	return m.err.Error()
}

func isGCPAPIEnabled(projectID string, want string) (bool, error) {
	enabledAPIs, err := listEnabledAPIs(projectID)

	if err != nil || enabledAPIs == nil {
		log.Error().Msgf("Failed to list enabled APIs %v", err)
	}

	if !slices.Contains(enabledAPIs, want) {
		return false, errors.New("API not enabled")
	}

	return true, nil
}

func compareAWSRole(directory string, arn string, init bool) (bool, error) {

	// Load the Shared AWS Configuration (~/.aws/config)
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)

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

	iacPolicy, err := MakePolicy(directory, nil, init, false, "", "")
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
		log.Error().Msg("Directory cannot be empty")
		return false, &emptyDirectoryError{}
	}

	if arn == "" {
		log.Error().Msg("ARN cannot be empty")
		return false, &arnEmptyError{}
	}

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		log.Error().Msgf("Directory %s does not exist", directory)
		return false, &directoryNotFoundError{directory}
	}

	if !strings.HasPrefix(arn, awsARNPrefix) && !strings.HasPrefix(arn, gcpProjectPrefix) {
		log.Error().Msgf("Invalid ARN %s cant determine cloud in use", arn)
		return false, &invalidARNError{arn}
	}

	return true, nil
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

type arnNotVerified struct {
	arn string
}

func (e *arnNotVerified) Error() string {
	fmt.Print("ARN must be in the following format:")
	fmt.Print("arn:aws:iam::123456789012:role/role-name")
	return fmt.Sprintf("ARN %s not verified", e.arn)
}

type gcpRoleNotVerified struct {
	role string
}

func (e *gcpRoleNotVerified) Error() string {
	fmt.Print(
		`The resource name of the role in one of the following formats:
        roles/{ROLE_NAME}
        organizations/{ORGANIZATION_ID}/roles/{ROLE_NAME}
        projects/{PROJECT_ID}/roles/{ROLE_NAME}`)
	return e.role
}

func verifyGCPRole(role string) error {
	r, err := regexp.Compile("projects/(.*\\S)/roles/(.*\\S)")
	// Regex should be compiled once as package variable
	if err == nil {
		if r.MatchString(role) {
			return nil
		}
	}

	return &gcpRoleNotVerified{role}
}

func verifyAWSARN(ARN string) error {
	r, err := regexp.Compile("arn:aws:iam::(.*\\S):role/(.*\\S)")
	// Regex should be compiled once as package variable
	if err == nil {
		if r.MatchString(ARN) {
			return nil
		}
	}
	return &arnNotVerified{ARN}
}
