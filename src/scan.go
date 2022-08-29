package pike

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Scan looks for resources in a given directory
func Scan(dirname string, output string, file string) error {
	Policy, err := MakePolicy(dirname, output, file)
	if err != nil {
		return err
	}

	fmt.Print(Policy)
	return err
}

// MakePolicy does the guts of determining a policy from code
func MakePolicy(dirname string, output string, file string) (string, error) {
	var files []string

	if file == "" {
		fullPath, err := filepath.Abs(dirname)

		if err != nil {
			return "", err
		}

		files, err = GetTF(fullPath)
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

		resource, err := GetResources(file)

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
		PermissionBag.GCP = append(PermissionBag.GCP, newPerms.AWS...)
	}

	Policy, err2 := GetPolicy(PermissionBag, output)
	if err2 != nil {
		return "", err2
	}
	return Policy, nil
}

// GetTF return tf files in a directory
func GetTF(dirName string) ([]string, error) {
	rawFiles, err := os.ReadDir(dirName)
	var files []string
	for _, file := range rawFiles {
		if file.IsDir() {

			if file.Name() == ".terraform" || file.Name() == ".git" {
				continue
			}
			newDirName := dirName + "/" + file.Name()
			moreFiles, err := GetTF(newDirName)
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
