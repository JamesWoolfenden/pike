package pike

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/hashicorp/terraform-exec/tfexec"
)

// Make creats the required role
func Make(directory string) error {

	err := Scan(
		directory,
		"terraform",
		"",
		true,
		true,
	)
	if err != nil {
		return err
	}

	policyPath, err := filepath.Abs(directory + "./.pike/")
	if err != nil {
		return err
	}

	tfPath, err := LocateTerraform()
	if err != nil {
		return err
	}

	tf, err := tfexec.NewTerraform(policyPath, tfPath)
	if err != nil {
		return err
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		return err
	}

	err = tf.Apply(context.Background())
	if err != nil {
		return err
	}

	state, err := tf.Show(context.Background())
	if err != nil {
		return err
	}

	if (state.Values.Outputs["arn"]) != nil {
		arn := state.Values.Outputs["arn"]
		log.Printf("aws role create/updated %s", arn.Value.(string))
		fmt.Print(arn.Value.(string))
	}

	return nil
}
