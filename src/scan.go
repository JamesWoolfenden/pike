package pike

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-version"

	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/rs/zerolog/log"
)

const tfVersion = "1.5.4"

type emptyIACError struct{}

func (m *emptyIACError) Error() string {
	return "no IAC found"
}

// Scan looks for resources in a given directory.
func Scan(dirName string, output string, file *string, init bool, write bool, enableResources bool) error {
	OutPolicy, err := MakePolicy(dirName, file, init, enableResources)
	if err != nil {
		return err
	}

	if write {
		err2 := WriteOutput(OutPolicy, output, dirName)
		if err2 != nil {
			return err2
		}
	} else {
		fmt.Print(OutPolicy.AsString(output))
	}

	return err
}

// WriteOutput writes out the policy as json or terraform.
func WriteOutput(outPolicy OutputPolicy, output, location string) error {
	newPath, _ := filepath.Abs(location + "/.pike")
	err := os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		return err
	}

	var outFile string

	d1 := []byte(outPolicy.AsString(output))

	switch strings.ToLower(output) {
	case terraform:
		outFile = newPath + "/pike.generated_policy.tf"

		if outPolicy.AWS.Terraform != "" {
			err = os.WriteFile(newPath+"/aws_iam_role.terraform_pike.tf", roleTemplate, 0o644)
		}

		if err != nil {
			return err
		}
	case "json":
		outFile = newPath + "/pike.generated_policy.json"
	default:
		return errors.New("output format supports only json and terraform")
	}

	err = os.WriteFile(outFile, d1, 0o644)

	if err != nil {
		return err
	}

	return nil
}

// Init can download and install terraform if required and then terraform init your specified directory.
func Init(dirName string) (*string, []string, error) {
	tfPath, err := LocateTerraform()
	if err != nil {
		return nil, nil, err
	}

	tf, err := tfexec.NewTerraform(dirName, tfPath)
	if err != nil {
		return nil, nil, err
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		return nil, nil, fmt.Errorf("init failed %w", err)
	}

	log.Printf("terraform init at %s", dirName)

	modules, err := os.ReadDir(dirName + "/" + ".terraform/modules")

	// filter
	var found []string

	for _, module := range modules {
		if module.Name() == "modules.json" || module.Name() == ".DS_Store" {
			continue
		}

		found = append(found, module.Name())
	}

	if err != nil {
		return &tfPath, nil, err
	}

	return &tfPath, found, err
}

// LocateTerraform finds the Terraform executable or installs it
func LocateTerraform() (string, error) {
	tfPath, err := exec.LookPath(terraform)

	// if you don't have tf installed, we have to install it
	if err != nil || tfPath == "" {
		log.Printf("installing Terraform %s\n", tfVersion)
		installer := &releases.ExactVersion{
			Product: product.Terraform,
			Version: version.Must(version.NewVersion(tfVersion)),
		}

		var err error

		tfPath, err = installer.Install(context.Background())
		if err != nil {
			return "", err
		}
	}

	return tfPath, nil
}

// MakePolicy does the guts of determining a policy from code.
func MakePolicy(dirName string, file *string, init bool, EnableResources bool) (OutputPolicy, error) {
	var (
		files  []string
		Output OutputPolicy
	)

	if file == nil {
		fullPath, err := filepath.Abs(dirName)
		if err != nil {
			return Output, err
		}

		if init {
			_, modules, err := Init(fullPath)
			if err != nil {
				log.Printf("modules not found at %s", dirName)
			}

			for _, module := range modules {
				log.Printf("downloaded %s", module)
			}
		}

		files, err = GetTF(fullPath)

		if err != nil {
			return Output, err
		}
	} else {
		myFile, err := filepath.Abs(*file)
		if err != nil {
			return Output, err
		}

		// is this a tfFile?
		if !(FileExists(myFile)) {
			return Output, os.ErrNotExist
		}

		files = append(files, myFile)
	}

	if len(files) == 0 {
		return Output, &emptyIACError{}
	}

	var resources []ResourceV2

	for _, tfFile := range files {
		resource, err := GetResources(tfFile, dirName)
		if err != nil {
			// parse the other files
			log.Print(err)
		}

		if resource != nil {
			resources = append(resources, resource...)
		}
	}

	var PermissionBag Sorted

	var newPerms Sorted

	for _, resource := range resources {
		var err error
		newPerms, err = GetPermission(resource)

		if err != nil {
			continue
		}

		PermissionBag.AWS = append(PermissionBag.AWS, newPerms.AWS...)
		PermissionBag.GCP = append(PermissionBag.GCP, newPerms.GCP...)
		PermissionBag.AZURE = append(PermissionBag.AZURE, newPerms.AZURE...)
	}

	Output, err2 := GetPolicy(PermissionBag, EnableResources)
	if err2 != nil {
		return Output, err2
	}

	return Output, nil
}

// GetTF return tf files in a directory
func GetTF(dirName string) ([]string, error) {
	files, err := GetTFFiles(dirName)
	if err != nil {
		return nil, fmt.Errorf("folder %s can't be found, may not be local path", dirName)
	}

	modulePath := dirName + "/.terraform/modules"
	if modules, err := os.ReadDir(modulePath); err == nil {
		for _, module := range modules {
			moreFiles, _ := GetTFFiles(modulePath + "/" + module.Name())
			files = append(files, moreFiles...)
		}
	}

	return files, nil
}

// GetTFFiles get tf files in directory
func GetTFFiles(dirName string) ([]string, error) {
	rawFiles, err := os.ReadDir(dirName)

	var files []string

	for _, file := range rawFiles {
		fileExtension := filepath.Ext(file.Name())

		if fileExtension != ".tf" {
			continue
		}

		files = append(files, dirName+"/"+file.Name())
	}

	return files, err
}

// StringInSlice looks for item in slice
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

// GetHCLType gets the resource Name
func GetHCLType(resourceName string) string {
	return strings.Split(resourceName, "_")[0]
}
