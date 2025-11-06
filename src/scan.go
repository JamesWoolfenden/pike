package pike

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/rs/zerolog/log"
)

const tfVersion = "1.5.4"

const (
	modulesJSON  = "modules.json"
	dsStore      = ".DS_Store"
	dotTfModules = ".terraform/modules"
)

var (
	terraformMutex sync.Mutex
	initMutex      sync.Map // per-directory mutex
)

type emptyIACError struct{}

func (m *emptyIACError) Error() string {
	return "no IAC found"
}

type makePolicyError struct {
	err error
}

func (m *makePolicyError) Error() string {
	return fmt.Sprintf("failed to make policy %v", m.err)
}

type emptyScanLocationError struct{}

func (m *emptyScanLocationError) Error() string {
	return "no scan location"
}

type makeDirectoryError struct {
	directory string
	err       error
}

func (m *makeDirectoryError) Error() string {
	return fmt.Sprintf("failed to make directory %s %v", m.directory, m.err)
}

type locateTerraformError struct {
	err error
}

func (m *locateTerraformError) Error() string {
	return fmt.Sprintf("failed to find Terraform %v", m.err)
}

type terraformExecError struct {
	err error
}

func (m *terraformExecError) Error() string {
	return fmt.Sprintf("Terraform execution error %v", m.err)
}

type terraformInitError struct {
	err error
}

func (m *terraformInitError) Error() string {
	return fmt.Sprintf("Terraform init error %v", m.err)
}

type readDirectoryError struct {
	directory string
	err       error
}

func (m *readDirectoryError) Error() string {
	return fmt.Sprintf("failed to read directory %s %v", m.directory, m.err)
}

type absolutePathError struct {
	directory string
	err       error
}

func (m *absolutePathError) Error() string {
	return fmt.Sprintf("failed to get absolute path %s %v", m.directory, m.err)
}

type getTFError struct {
	directory string
	err       error
}

func (m *getTFError) Error() string {
	return fmt.Sprintf("failed to get Terraform templates %s %v", m.directory, m.err)
}

type getPolicyError struct {
	err error
}

func (m *getPolicyError) Error() string {
	return fmt.Sprintf("failed to get policy %v", m.err)
}

// Scan looks for resources in a given directory.
func Scan(dirName string, outputType string, file *string, init bool, write bool, enableResources bool, provider string, outFile string, policyName string) error {
	if dirName == "" && file == nil {
		return &emptyScanLocationError{}
	}

	OutPolicy, err := MakePolicy(dirName, file, init, enableResources, provider, policyName)
	if err != nil {
		fmt.Print(err.Error())
		return &makePolicyError{err}
	}

	if write {
		err = WriteOutput(OutPolicy, outputType, dirName, outFile)
		if err != nil {
			return &writeFileError{file: outputType, err: err}
		}
	} else {
		fmt.Print(OutPolicy.AsString(outputType)) // permit
	}

	return err
}

// WriteOutput writes out the policy as JSON or Terraform.
func WriteOutput(outPolicy OutputPolicy, outputType string, scanPath string, outFile string) error {

	var newPath string

	d1 := []byte(outPolicy.AsString(outputType))

	if outFile != "" {

	} else {
		if scanPath == "" {
			scanPath = "."
		}
		newPath, _ = filepath.Abs(path.Join(scanPath, ".pike"))

		err := os.MkdirAll(newPath, os.ModePerm)

		if err != nil {
			return &makeDirectoryError{directory: newPath, err: err}
		}

		switch strings.ToLower(outputType) {
		case terraform:
			outFile = filepath.Join(newPath, "pike.generated_policy.tf") //path.join does not work here

			if outPolicy.AWS.Terraform != "" {
				roleFile := path.Join(newPath, "aws_iam_role.terraform_pike.tf")
				err = os.WriteFile(roleFile, roleTemplate, 0o644)

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

	err := os.WriteFile(outFile, d1, 0o644)
	if err != nil {
		return &writeFileError{file: outFile, err: err}
	}

	log.Info().Msgf("wrote %s", outFile)

	return nil
}

// Init can download and install terraform if required and then terraform init your specified directory.

func Init(dirName string) (*string, []string, error) {
	// Per-directory locking
	dirMutex, _ := initMutex.LoadOrStore(dirName, &sync.Mutex{})
	mutex := dirMutex.(*sync.Mutex)
	mutex.Lock()
	defer mutex.Unlock()

	tfPath, err := LocateTerraform()
	if err != nil {
		return nil, nil, &locateTerraformError{err}
	}

	tf, err := tfexec.NewTerraform(dirName, tfPath)

	if err != nil {
		return nil, nil, &terraformExecError{err}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	err = tf.Init(ctx, tfexec.Upgrade(true))
	if err != nil {
		if errors.Is(context.DeadlineExceeded, ctx.Err()) {
			return nil, nil, fmt.Errorf("terraform init timed out after 10 minutes: %w", err)
		}
		return nil, nil, &terraformInitError{err}
	}

	log.Info().Msgf("terraform init at %s", dirName)

	modulesDir := path.Join(dirName, dotTfModules)
	modules, err := os.ReadDir(modulesDir)

	if err != nil {
		return &tfPath, nil, &readDirectoryError{directory: modulesDir, err: err}
	}

	// filter
	var found []string

	for _, module := range modules {
		if module.Name() == "modules.json" || module.Name() == ".DS_Store" {
			continue
		}

		found = append(found, module.Name())
	}

	return &tfPath, found, nil
}

// LocateTerraform finds the Terraform executable or installs it.
func LocateTerraform() (string, error) {
	terraformMutex.Lock()
	defer terraformMutex.Unlock()

	tfPath, err := exec.LookPath(terraform)

	// if you don't have tf installed, we have to install it
	if err != nil || tfPath == "" {
		log.Info().Msgf("installing Terraform %s\n", tfVersion)
		installer := &releases.ExactVersion{
			Product: product.Terraform,
			Version: version.Must(version.NewVersion(tfVersion)),
		}

		var err error

		tfPath, err = installer.Install(context.Background())
		if err != nil {
			return "", &locateTerraformError{err}
		}
	}

	return tfPath, nil
}

// MakePolicy does the guts of determining a policy from code.
func MakePolicy(dirName string, file *string, init bool, enableResources bool, provider string, policyName string) (OutputPolicy, error) {
	// Validate inputs early
	if dirName == "" && file == nil {
		return OutputPolicy{}, errors.New("either directory or file should be be set")
	}

	var output OutputPolicy

	permissionsBag, err := makePermissionBag(dirName, file, init, provider)
	if err != nil {
		return output, fmt.Errorf("failed to create permission bag: %w", err)
	}

	output, err = GetPolicy(permissionsBag, enableResources, policyName)
	if err != nil {
		return output, &getPolicyError{err: err}
	}

	return output, nil
}

// Extract common absolute path logic
func getAbsolutePath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", &absolutePathError{directory: path, err: err}
	}
	return absPath, nil
}
func makePermissionBag(dirName string, file *string, init bool, provider string) (Sorted, error) {

	var files []string

	if file == nil {
		fullPath, err := getAbsolutePath(dirName)
		if err != nil {
			return Sorted{}, err
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
			return Sorted{}, &getTFError{directory: fullPath, err: err}
		}
	} else {
		myFile, err := getAbsolutePath(*file)
		if err != nil {
			return Sorted{}, err
		}

		// is this a tfFile?
		if !(FileExists(myFile)) {
			return Sorted{}, os.ErrNotExist
		}

		files = append(files, myFile)
	}

	if len(files) == 0 {
		return Sorted{}, &emptyIACError{}
	}

	var resources []ResourceV2
	var failedFiles []string
	var criticalErrors []error

	for _, tfFile := range files {
		resource, err := GetResources(tfFile, dirName)
		if err != nil {
			failedFiles = append(failedFiles, tfFile)
			criticalErrors = append(criticalErrors, fmt.Errorf("failed to parse %s: %w", tfFile, err))
			continue
		}

		if resource != nil {
			resources = append(resources, resource...)
		}
	}

	// Fail fast if too many critical files failed
	if len(criticalErrors) > 0 {
		if len(failedFiles) > len(files)/2 { // More than 50% failed
			return Sorted{}, fmt.Errorf("critical parsing failures in %d/%d files: %v",
				len(failedFiles), len(files), criticalErrors)
		}
		log.Warn().Int("failed_files", len(failedFiles)).Msg("some terraform files failed to parse")
	}

	permissionsBag := GetPermissionBag(resources, provider)
	return permissionsBag, nil
}
func GetPermissionBag(resources []ResourceV2, provider string) Sorted {
	var permissionBag Sorted
	var newPerms Sorted

	for _, resource := range resources {
		var err error

		// implement provider filter
		if provider == "" || provider == resource.Provider {
			newPerms, err = GetPermission(resource)
		} else {
			continue
		}

		if err != nil {
			continue
		}

		switch strings.ToLower(provider) {
		case "aws":
			permissionBag.AWS = append(permissionBag.AWS, newPerms.AWS...)
		case "gcp", "google":
			permissionBag.GCP = append(permissionBag.GCP, newPerms.GCP...)
		case "azure", "azurerm":
			permissionBag.AZURE = append(permissionBag.AZURE, newPerms.AZURE...)
		case "":
			permissionBag.AWS = append(permissionBag.AWS, newPerms.AWS...)
			permissionBag.GCP = append(permissionBag.GCP, newPerms.GCP...)
			permissionBag.AZURE = append(permissionBag.AZURE, newPerms.AZURE...)
		}
	}
	return permissionBag
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
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

// GetHCLType gets the resource Name.
func GetHCLType(resourceName string) string {
	return strings.Split(resourceName, "_")[0]
}

const (
	maxFiles     = 1000
	maxFileSize  = 10 * 1024 * 1024 // 10MB
	maxResources = 50000
)

func validateLimits(files []string) error {
	if len(files) > maxFiles {
		return fmt.Errorf("too many files: %d > %d", len(files), maxFiles)
	}
	return nil
}
