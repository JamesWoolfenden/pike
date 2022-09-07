package pike

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

const tfVersion = "1.2.3"

// Scan looks for resources in a given directory
func Scan(dirName string, output string, file string, init bool, excludes *cli.StringSlice) error {

	Policy, err := MakePolicy(dirName, output, file, init, excludes)
	if err != nil {
		return err
	}

	fmt.Print(Policy)
	return err
}

// Init can download and install terraform if required and then terraform init your specified directory
func Init(dirName string) (string, error) {

	tfPath, _ := exec.LookPath("terraform")

	//if you don't have tf installed we have to install it
	if tfPath == "" {
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

	tf, err := tfexec.NewTerraform(dirName, tfPath)
	if err != nil {
		return "", err
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		return "", err
	}

	log.Printf("terraform init at %s\n", dirName)
	return tfPath, err
}

// MakePolicy does the guts of determining a policy from code
func MakePolicy(dirName string, output string, file string, init bool, excludes *cli.StringSlice) (string, error) {
	var files []string

	if file == "" {
		fullPath, err := filepath.Abs(dirName)

		if err != nil {
			return "", err
		}
		if init {
			_, err := Init(dirName)
			if err != nil {
				return "", err
			}
		}

		files, err = GetTF(fullPath, true, excludes)
		if err != nil {
			return "", err
		}
	} else {
		myFile, err := filepath.Abs(file)

		if err != nil {
			return "", err
		}

		// is this a file
		if !(fileExists(myFile)) {
			return "", os.ErrNotExist
		}

		files = append(files, myFile)
	}

	var resources []ResourceV2
	for _, file := range files {

		resource, err := GetResources(file, dirName)

		if err != nil {
			// parse the other files
			log.Print(err)
		}
		if resource != nil {
			resources = append(resources, resource...)
		}
	}
	var PermissionBag Sorted

	for _, resource := range resources {
		newPerms, err := GetPermission(resource)

		if err != nil {
			return "", err
		}

		PermissionBag.AWS = append(PermissionBag.AWS, newPerms.AWS...)
		PermissionBag.GCP = append(PermissionBag.GCP, newPerms.GCP...)
	}

	Policy, err2 := GetPolicy(PermissionBag, output)
	if err2 != nil {
		return "", err2
	}
	return Policy, nil
}

// GetTF return tf files in a directory
func GetTF(dirName string, recurse bool, excludes *cli.StringSlice) ([]string, error) {
	rawFiles, err := os.ReadDir(dirName)
	var files []string
	for _, file := range rawFiles {
		if file.IsDir() && recurse {
			var excludeDir []string
			if excludes != nil {
				excludeDir = excludes.Value()
			}

			excludeDir = append(excludeDir, ".git", ".external_modules", ".pike")
			if stringInSlice(file.Name(), excludeDir) {
				continue
			}

			newDirName := dirName + "/" + file.Name()
			moreFiles, err := GetTF(newDirName, false, nil)
			if err == nil {
				if moreFiles != nil {
					files = append(files, moreFiles...)
				}
			}
		}

		fileExtension := filepath.Ext(file.Name())

		if fileExtension != ".tf" {
			continue
		}
		files = append(files, dirName+"/"+file.Name())
	}

	if err != nil {
		return nil, err
	}
	return files, nil
}

func stringInSlice(a string, list []string) bool {
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
