package pike

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/rs/zerolog/log"
)

// ensurePikeDir returns the absolute path of <scanPath>/.pike, creating it if
// needed. A symlinked .pike is refused: the scanned tree may be untrusted and
// the link would redirect output outside it.
func ensurePikeDir(scanPath string) (string, error) {
	newPath, err := filepath.Abs(filepath.Join(scanPath, ".pike"))
	if err != nil {
		return "", &absolutePathError{directory: scanPath, err: err}
	}

	if info, lerr := os.Lstat(newPath); lerr == nil {
		if info.Mode()&os.ModeSymlink != 0 {
			return "", &makeDirectoryError{directory: newPath, err: errSymlinkRefused}
		}
	} else if !os.IsNotExist(lerr) {
		return "", &makeDirectoryError{directory: newPath, err: lerr}
	}

	if err = os.MkdirAll(newPath, 0o750); err != nil {
		return "", &makeDirectoryError{directory: newPath, err: err}
	}

	return newPath, nil
}

// writeFileNoSymlink writes a 0600 file but refuses to write through an
// existing symlink at the target path.
func writeFileNoSymlink(name string, data []byte) error {
	if info, err := os.Lstat(name); err == nil && info.Mode()&os.ModeSymlink != 0 {
		return errSymlinkRefused
	}

	return os.WriteFile(name, data, 0o600)
}

var errSymlinkRefused = errors.New("refusing to write through a symlink")

func writeTwoRoleFile(content string, scanPath string, outFile string) error {
	if outFile == "" {
		if scanPath == "" {
			scanPath = "."
		}
		newPath, err := ensurePikeDir(scanPath)
		if err != nil {
			return err
		}
		outFile = filepath.Join(newPath, "pike.two-role.tf")
	}
	if err := writeFileNoSymlink(outFile, []byte(content)); err != nil {
		return &writeFileError{file: outFile, err: err}
	}
	log.Info().Msgf("wrote %s", outFile)
	return nil
}

func writeSplitOutput(content string, scanPath string, outFile string) error {
	if outFile == "" {
		if scanPath == "" {
			scanPath = "."
		}
		newPath, err := ensurePikeDir(scanPath)
		if err != nil {
			return err
		}
		outFile = filepath.Join(newPath, "pike.generated_policy.split.json")
	}
	if err := writeFileNoSymlink(outFile, []byte(content)); err != nil {
		return &writeFileError{file: outFile, err: err}
	}
	log.Info().Msgf("wrote %s", outFile)
	return nil
}

// WriteOutput writes out the policy as JSON or Terraform.
func WriteOutput(outPolicy OutputPolicy, outputType string, scanPath string, outFile string) error {
	d1 := []byte(outPolicy.AsString(outputType))

	if outFile == "" {
		if scanPath == "" {
			scanPath = "."
		}
		newPath, err := ensurePikeDir(scanPath)
		if err != nil {
			return err
		}

		switch strings.ToLower(outputType) {
		case terraform:
			outFile = filepath.Join(newPath, "pike.generated_policy.tf")

			if outPolicy.AWS.Terraform != "" {
				roleFile := filepath.Join(newPath, "aws_iam_role.terraform_pike.tf")
				err = writeFileNoSymlink(roleFile, roleTemplate)

				if err != nil {
					return &writeFileError{file: roleFile, err: err}
				}
			}

		case "json":
			outFile = filepath.Join(newPath, "pike.generated_policy.json")
		default:
			return &tfPolicyFormatError{}
		}
	}

	err := writeFileNoSymlink(outFile, d1)
	if err != nil {
		return &writeFileError{file: outFile, err: err}
	}

	log.Info().Msgf("wrote %s", outFile)

	return nil
}

// MakeSplitPolicy partitions a permission bag into base and escalation-class
// subsets per provider, for use with --output split.
func MakeSplitPolicy(bag Sorted) SplitPolicy {
	var sp SplitPolicy

	if len(bag.AWS) > 0 {
		set := &SplitSet{}
		for _, p := range Unique(bag.AWS) {
			if escalationAWS[p] {
				set.Escalation = append(set.Escalation, p)
			} else {
				set.Base = append(set.Base, p)
			}
		}
		sp.AWS = set
	}

	if len(bag.GCP) > 0 {
		set := &SplitSet{}
		for _, p := range Unique(bag.GCP) {
			if escalationGCP[p] {
				set.Escalation = append(set.Escalation, p)
			} else {
				set.Base = append(set.Base, p)
			}
		}
		sp.GCP = set
	}

	if len(bag.AZURE) > 0 {
		set := &SplitSet{}
		for _, p := range Unique(bag.AZURE) {
			if escalationAZURE[p] {
				set.Escalation = append(set.Escalation, p)
			} else {
				set.Base = append(set.Base, p)
			}
		}
		sp.AZURE = set
	}

	return sp
}

// GetTF return tf files in a directory.
func GetTF(dirName string) ([]string, error) {
	files, err := GetTFFiles(dirName)
	if err != nil {
		return nil, &directoryNotFoundError{dirName}
	}

	modulePath := filepath.Join(dirName, ".terraform", "modules")
	if modules, err := os.ReadDir(modulePath); err == nil {
		for _, module := range modules {
			tfFilesPath := filepath.Join(modulePath, module.Name())
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

		newFile := filepath.Join(dirName, file.Name())
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
