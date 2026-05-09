package pike

import (
	"context"
	"fmt"
	"path"
	"path/filepath"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/rs/zerolog/log"
)

const (
	iamRetryAttempts  = 12
	iamRetryBaseSleep = 5 * time.Second
	iamRetryMaxSleep  = 30 * time.Second
)

var iamDenied = regexp.MustCompile(`(?i)AccessDenied|AccessDeniedException|UnauthorizedOperation|is not authorized to perform`)

// Make creates the required role.
func Make(directory string) (*string, error) {
	if directory == "" {
		return nil, &directoryNotFoundError{directory: directory}
	}

	err := Scan(directory, "terraform", nil, true, true, false, "", "", "", false)
	if err != nil {
		return nil, fmt.Errorf("failed to scan directory: %w", err)
	}

	directory, err = filepath.Abs(directory)
	if err != nil {
		return nil, &absolutePathError{directory, err}
	}

	policyPath, err := filepath.Abs(path.Join(directory, ".pike"))
	if err != nil {
		return nil, &absolutePathError{directory, err}
	}

	tf, err := tfApply(policyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to apply terraform: %w", err)
	}

	state, err := tf.Show(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to show terraform state: %w", err)
	}

	if (state.Values.Outputs["arn"]) != nil {
		arn := state.Values.Outputs["arn"]

		role, ok := arn.Value.(string)
		if !ok {
			return nil, &castToStringError{"arn"}
		}

		log.Info().Msgf("aws role create/updated %s", role)

		return &role, nil
	}

	return nil, &arnNotFoundInStateError{}
}

type castToStringError struct {
	value string
}

func (e *castToStringError) Error() string {
	return fmt.Sprint("cannot convert ", e.value, " to a string")
}

type arnNotFoundInStateError struct{}

func (e *arnNotFoundInStateError) Error() string {
	return "no arn found in state"
}

type roleNotFoundInStateError struct {
	provider string
}

func (e *roleNotFoundInStateError) Error() string {
	return fmt.Sprintf("no role found in state for provider: %s", e.provider)
}

// MakeGCP creates the required GCP IAM custom role.
func MakeGCP(directory string) (*string, error) {
	if directory == "" {
		return nil, &directoryNotFoundError{directory: directory}
	}

	err := Scan(directory, "terraform", nil, true, true, false, "google", "", "", false)
	if err != nil {
		return nil, fmt.Errorf("failed to scan directory: %w", err)
	}

	directory, err = filepath.Abs(directory)
	if err != nil {
		return nil, &absolutePathError{directory, err}
	}

	policyPath, err := filepath.Abs(path.Join(directory, ".pike"))
	if err != nil {
		return nil, &absolutePathError{directory, err}
	}

	tf, err := tfApply(policyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to apply terraform: %w", err)
	}

	state, err := tf.Show(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to show terraform state: %w", err)
	}

	for _, resource := range state.Values.RootModule.Resources {
		if resource.Type == "google_project_iam_custom_role" {
			if id, ok := resource.AttributeValues["id"].(string); ok {
				log.Info().Msgf("GCP role created/updated: %s", id)
				return &id, nil
			}
		}
	}

	return nil, &roleNotFoundInStateError{provider: "gcp"}
}

// MakeAzure creates the required Azure role definition.
func MakeAzure(directory string) (*string, error) {
	if directory == "" {
		return nil, &directoryNotFoundError{directory: directory}
	}

	err := Scan(directory, "terraform", nil, true, true, false, "azurerm", "", "", false)
	if err != nil {
		return nil, fmt.Errorf("failed to scan directory: %w", err)
	}

	directory, err = filepath.Abs(directory)
	if err != nil {
		return nil, &absolutePathError{directory, err}
	}

	policyPath, err := filepath.Abs(path.Join(directory, ".pike"))
	if err != nil {
		return nil, &absolutePathError{directory, err}
	}

	tf, err := tfApply(policyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to apply terraform: %w", err)
	}

	state, err := tf.Show(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to show terraform state: %w", err)
	}

	for _, resource := range state.Values.RootModule.Resources {
		if resource.Type == "azurerm_role_definition" {
			if id, ok := resource.AttributeValues["role_definition_resource_id"].(string); ok {
				log.Info().Msgf("Azure role definition created/updated: %s", id)
				return &id, nil
			}
		}
	}

	return nil, &roleNotFoundInStateError{provider: "azure"}
}

func tfApply(policyPath string) (*tfexec.Terraform, error) {
	tfPath, err := LocateTerraform()
	if err != nil {
		return nil, &locateTerraformError{err}
	}

	terraform, err := tfexec.NewTerraform(policyPath, tfPath)
	if err != nil {
		return nil, &terraformNewError{err: err}
	}

	err = terraform.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		return nil, &terraformInitError{err}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	err = terraform.Apply(ctx)
	if err != nil {
		return nil, &terraformApplyError{err: err, target: policyPath}
	}

	return terraform, nil
}

// Apply executes tf using a prepared role.
func Apply(target string, region string) error {
	iamRole, err := Make(target)
	if err != nil {
		return &makeRoleError{err}
	}

	// clear any temp credentials (best-effort: unsetenv failures on supported
	// OSes are pathological, so we log and continue rather than abort)
	if unsetErr := unSetAWSAuth(); unsetErr != nil {
		log.Warn().Err(unsetErr).Msg("failed to clear AWS env vars before auth")
	}

	err = setAWSAuth(*iamRole, region)
	if err != nil {
		if unsetErr := unSetAWSAuth(); unsetErr != nil {
			log.Warn().Err(unsetErr).Msg("failed to clear AWS env vars after setAWSAuth error")
		}

		return &setAWSAuthError{err}
	}

	log.Debug().Msgf("Starting terraform apply in directory: %s", target)
	defer log.Debug().Msg("Completed terraform apply")

	err = tfApplyWithIAMRetry(target)

	if err == nil {
		log.Info().Msgf("provisioned %s", target)
	} else {
		err = &terraformApplyError{target, err}
	}

	if unsetErr := unSetAWSAuth(); unsetErr != nil {
		log.Warn().Err(unsetErr).Msg("failed to clear AWS env vars after apply")
	}

	return err
}

// tfApplyWithIAMRetry runs tfApply, retrying with linear backoff while the
// failure looks like IAM eventual consistency on a just-created role. Any
// other error returns immediately.
func tfApplyWithIAMRetry(target string) error {
	var err error
	for i := 1; i <= iamRetryAttempts; i++ {
		_, err = tfApply(target)
		if err == nil {
			return nil
		}
		if !iamDenied.MatchString(err.Error()) {
			return err
		}
		if i == iamRetryAttempts {
			break
		}
		sleep := min(iamRetryBaseSleep*time.Duration(i), iamRetryMaxSleep)
		log.Info().Msgf("IAM not yet consistent (attempt %d/%d), retrying in %v", i, iamRetryAttempts, sleep)
		time.Sleep(sleep)
	}
	return fmt.Errorf("still IAM-denied after %d attempts: %w", iamRetryAttempts, err)
}
