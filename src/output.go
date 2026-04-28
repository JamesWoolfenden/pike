package pike

import (
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"

	"github.com/rs/zerolog/log"
)

// WriteOutput writes out the policy as JSON or Terraform.
func WriteOutput(outPolicy OutputPolicy, outputType string, scanPath string, outFile string) error {
	d1 := []byte(outPolicy.AsString(outputType))

	if outFile == "" {
		if scanPath == "" {
			scanPath = "."
		}
		newPath, err := filepath.Abs(path.Join(scanPath, ".pike"))
		if err != nil {
			return &absolutePathError{directory: scanPath, err: err}
		}

		err = os.MkdirAll(newPath, 0o750)

		if err != nil {
			return &makeDirectoryError{directory: newPath, err: err}
		}

		switch strings.ToLower(outputType) {
		case terraform:
			outFile = filepath.Join(newPath, "pike.generated_policy.tf") // path.join does not work here

			if outPolicy.AWS.Terraform != "" {
				roleFile := path.Join(newPath, "aws_iam_role.terraform_pike.tf")
				err = os.WriteFile(roleFile, roleTemplate, 0o600)

				if err != nil {
					return &writeFileError{file: roleFile, err: err}
				}
			}

		case "json":
			outFile = path.Join(newPath, "pike.generated_policy.json")
		default:
			return &tfPolicyFormatError{}
		}
	}

	err := os.WriteFile(outFile, d1, 0o600)
	if err != nil {
		return &writeFileError{file: outFile, err: err}
	}

	log.Info().Msgf("wrote %s", outFile)

	return nil
}

// GetTF return tf files in a directory.
func GetTF(dirName string) ([]string, error) {
	files, err := GetTFFiles(dirName)
	if err != nil {
		return nil, &directoryNotFoundError{dirName}
	}

	modulePath := path.Join(dirName, dotTfModules)
	if modules, err := os.ReadDir(modulePath); err == nil {
		for _, module := range modules {
			tfFilesPath := path.Join(modulePath, module.Name())
			moreFiles, _ := GetTFFiles(tfFilesPath)
			files = append(files, moreFiles...)
		}
	}

	return files, nil
}

// GetTFFiles get tf files in directory.
func GetTFFiles(dirName string) ([]string, error) {
	rawFiles, err := os.ReadDir(dirName)
	if err != nil {
		return nil, &readDirectoryError{dirName, err}
	}

	var files []string

	for _, file := range rawFiles {
		fileExtension := filepath.Ext(file.Name())

		if fileExtension != ".tf" {
			continue
		}

		newFile := path.Join(dirName, file.Name())
		files = append(files, newFile)
	}

	return files, nil
}

// StringInSlice looks for item in slice.
func StringInSlice(a string, list []string) bool {
	return slices.Contains(list, a)
}

// GetHCLType gets the resource Name.
func GetHCLType(resourceName string) string {
	return strings.Split(resourceName, "_")[0]
}
