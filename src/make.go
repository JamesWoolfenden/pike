package pike

import (
	"context"
	"fmt"
	"path"
	"path/filepath"
	"time"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/rs/zerolog/log"
)

// Make creates the required role.
func Make(directory string) (*string, error) {
	if directory == "" {
		return nil, &directoryNotFoundError{directory: directory}
	}

	err := Scan(directory, "terraform", nil, true, true, false, "", "", "")
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

		myValue, ok := arn.Value.(string)
		if !ok {
			return nil, &castToStringError{"arn"}
		}

		log.Info().Msgf("aws role create/updated %s", myValue)

		role, ok := arn.Value.(string)

		if !ok {
			return nil, &castToStringError{"arn"}
		}

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

	time.Sleep(5 * time.Second)

	if err != nil {
		return &makeRoleError{err}
	}
	// clear any temp credentials
	unSetAWSAuth()

	err = setAWSAuth(*iamRole, region)
	if err != nil {
		unSetAWSAuth()

		return &setAWSAuthError{err}
	}

	log.Debug().Msgf("Starting terraform apply in directory: %s", target)
	defer log.Debug().Msg("Completed terraform apply")

	_, err = tfApply(target)

	if err == nil {
		log.Info().Msgf("provisioned %s", target)
	} else {
		err = &terraformApplyError{target, err}
	}

	unSetAWSAuth()

	return err
}
