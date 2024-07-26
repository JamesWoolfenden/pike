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
		return nil, err
	}

	directory, err = filepath.Abs(directory)
	if err != nil {
		return nil, fmt.Errorf("failed to find path %w", err)
	}

	policyPath, err := filepath.Abs(directory + "/.pike/")
	if err != nil {
		return nil, err
	}

	tf, err2 := tfApply(policyPath)
	if err2 != nil {
		return nil, err2
	}

	state, err := tf.Show(context.Background())
	if err != nil {
		return nil, err
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

	err = terraform.Apply(context.Background())
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

	cmd = exec.Command(terraform.ExecPath(), chdir, "show", "--json", "tf.plan")
	stdout, err = cmd.Output()

	if err != nil {
		fmt.Println(err.Error())

		return err
	}

	outfile := filepath.Join(policyPath, "tf.json")
	err = os.WriteFile(outfile, stdout, 0o666)

	if err != nil {
		return fmt.Errorf("terraform show failed %w", err)
	}

	return nil
}
