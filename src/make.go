package pike

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/rs/zerolog/log"
)

// Make creates the required role.
func Make(directory string) (*string, error) {
	err := Scan(directory, "terraform", nil, true, true, false)
	if err != nil {
		return nil, fmt.Errorf("failed to scan directory: %w", err)
	}

	directory, err = filepath.Abs(directory)
	if err != nil {
		return nil, fmt.Errorf("failed to find path: %w", err)
	}

	policyPath, err := filepath.Abs(directory + "/.pike/")
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for policy: %w", err)
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
		log.Info().Msgf("aws role create/updated %s", arn.Value.(string))
		role := arn.Value.(string)

		return &role, nil
	}

	return nil, errors.New("no arn found in state")
}

func tfApply(policyPath string) (*tfexec.Terraform, error) {
	tfPath, err := LocateTerraform()
	if err != nil {
		return nil, err
	}

	terraform, err := tfexec.NewTerraform(policyPath, tfPath)
	if err != nil {
		return nil, err
	}

	err = terraform.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()
	err = terraform.Apply(ctx)
	if err != nil {
		return nil, fmt.Errorf("terraform apply failed %w", err)
	}

	return terraform, nil
}

// Apply executes tf using a prepared role.
func Apply(target string, region string) error {
	iamRole, err := Make(target)

	time.Sleep(5 * time.Second)

	if err != nil {
		return err
	}
	// clear any temp creds
	unSetAWSAuth()

	err = setAWSAuth(*iamRole, region)
	if err != nil {
		unSetAWSAuth()

		return err
	}

	log.Debug().Msgf("Starting terraform apply in directory: %s", target)
	defer log.Debug().Msg("Completed terraform apply")
	_, err = tfApply(target)

	if err == nil {
		log.Printf("provisioned %s", target)
	}

	unSetAWSAuth()

	return err
}

//goland:noinspection GoUnusedFunction,GoLinter
func tfPlan(policyPath string) error {
	tfPath, err := LocateTerraform()
	if err != nil {
		return err
	}

	terraform, err := tfexec.NewTerraform(policyPath, tfPath)
	if err != nil {
		return err
	}

	err = terraform.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		return err
	}

	chdir := "-chdir=" + policyPath
	cmd := exec.Command(terraform.ExecPath(), chdir, "plan", "--out", "tf.plan")

	stdout, err := cmd.Output()

	if err != nil {
		return fmt.Errorf("terraform plan generation failed: %w", err)
	}

	if len(stdout) == 0 {
		return errors.New("terraform plan output is empty")
	}

	defer os.Remove(filepath.Join(policyPath, "tf.plan"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	cmd = exec.CommandContext(ctx, terraform.ExecPath(), chdir, "show", "--json", "tf.plan")
	stdout, err = cmd.Output()

	if err != nil {
		return fmt.Errorf("terraform plan failed: %w", err)
	}

	outfile := filepath.Join(policyPath, "tf.json")
	err = os.WriteFile(outfile, stdout, 0o666)

	if err != nil {
		return fmt.Errorf("terraform show failed %w", err)
	}

	return nil
}
