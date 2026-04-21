package pike

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/jameswoolfenden/pike/internal/provider"
	"github.com/jameswoolfenden/pike/internal/tfinstall"
	"github.com/rs/zerolog/log"
)

const (
	modulesJSON  = "modules.json"
	dsStore      = ".DS_Store"
	dotTfModules = ".terraform/modules"
)

var initMutex sync.Map // per-directory mutex

// Scan looks for resources in a given directory.
func Scan(dirName string, outputType string, file *string, init bool, write bool, enableResources bool, provider string, outFile string, policyName string, suppressDeprecated bool) error {
	if dirName == "" && file == nil {
		return &emptyScanLocationError{}
	}

	OutPolicy, err := MakePolicy(dirName, file, init, enableResources, provider, policyName, suppressDeprecated)
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

// Runtime detects runtime IAM permissions needed by service accounts.
//
// Only GCP is supported today. AWS/Azure providers are rejected with
// unsupportedRuntimeProviderError.
func Runtime(dirName string, outputType string, file *string, init bool, prov string) error {
	if dirName == "" && file == nil {
		return &emptyScanLocationError{}
	}

	switch provider.Normalise(prov) {
	case provider.Google, "":
		// supported - fall through to the main path below
	default:
		return &unsupportedRuntimeProviderError{provider: prov}
	}

	permissionsBag, err := makePermissionBag(dirName, file, init, prov, false)
	if err != nil {
		return fmt.Errorf("failed to create permission bag: %w", err)
	}

	// Output runtime permissions
	output := formatRuntimePermissions(permissionsBag, prov)
	if output == "" {
		fmt.Println("No runtime permissions detected.")
		return nil
	}

	fmt.Print(output)
	return nil
}

// Init can download and install terraform if required and then terraform init your specified directory.

func Init(dirName string) (*string, []string, error) {
	// Per-directory locking. initMutex should only ever hold *sync.Mutex,
	// but use a comma-ok assertion so a future refactor that stores a
	// different value type fails loudly with a real error instead of a
	// panic deep in the stack.
	dirMutex, _ := initMutex.LoadOrStore(dirName, &sync.Mutex{})
	mutex, ok := dirMutex.(*sync.Mutex)
	if !ok {
		return nil, nil, fmt.Errorf("initMutex corruption: expected *sync.Mutex for %q, got %T", dirName, dirMutex)
	}
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
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
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
// LocateTerraform finds the Terraform executable or installs it.
// The search and install logic lives in internal/tfinstall; this wrapper
// preserves the locateTerraformError type for callers that inspect it.
func LocateTerraform() (string, error) {
	p, err := tfinstall.LocateTerraform()
	if err != nil {
		return "", &locateTerraformError{err}
	}

	return p, nil
}

// MakePolicy does the guts of determining a policy from code.
func MakePolicy(dirName string, file *string, init bool, enableResources bool, provider string, policyName string, suppressDeprecated bool) (OutputPolicy, error) {
	// Validate inputs early
	if dirName == "" && file == nil {
		return OutputPolicy{}, errors.New("either directory or file should be be set")
	}

	var output OutputPolicy

	permissionsBag, err := makePermissionBag(dirName, file, init, provider, suppressDeprecated)
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

func makePermissionBag(dirName string, file *string, init bool, provider string, suppressDeprecated bool) (Sorted, error) {

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

	var allIAMBindings []IAMBinding

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

		// Also extract IAM bindings from this file
		body, err := GetResourceBlocks(tfFile)
		if err == nil && body != nil {
			bindings := ExtractIAMBindings(body)
			allIAMBindings = append(allIAMBindings, bindings...)
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

	permissionsBag := GetPermissionBag(resources, provider, suppressDeprecated)
	permissionsBag.IAMBindings = allIAMBindings
	return permissionsBag, nil
}
func GetPermissionBag(resources []ResourceV2, prov string, suppressDeprecated bool) Sorted {
	var permissionBag Sorted
	var newPerms Sorted

	// Track which (provider, resource-type, kind) tuples we've already
	// warned about so scanning the same deprecated resource declared five
	// times in the user's terraform only logs once. Keyed on a composite
	// string because that's cheaper than a nested map and readability here
	// is not the priority.
	warned := map[string]struct{}{}

	for _, resource := range resources {
		var err error

		// implement provider filter
		if prov == "" || prov == resource.Provider {
			newPerms, err = GetPermission(resource)
		} else {
			continue
		}

		if !suppressDeprecated {
			warnIfDeprecated(resource, warned)
		}

		if err != nil {
			continue
		}

		switch provider.Normalise(prov) {
		case provider.AWS:
			permissionBag.AWS = append(permissionBag.AWS, newPerms.AWS...)
			permissionBag.RuntimeAWS = append(permissionBag.RuntimeAWS, newPerms.RuntimeAWS...)
		case provider.Google:
			permissionBag.GCP = append(permissionBag.GCP, newPerms.GCP...)
			permissionBag.RuntimeGCP = append(permissionBag.RuntimeGCP, newPerms.RuntimeGCP...)
		case provider.Azure:
			permissionBag.AZURE = append(permissionBag.AZURE, newPerms.AZURE...)
			permissionBag.RuntimeAZURE = append(permissionBag.RuntimeAZURE, newPerms.RuntimeAZURE...)
		case "":
			permissionBag.AWS = append(permissionBag.AWS, newPerms.AWS...)
			permissionBag.GCP = append(permissionBag.GCP, newPerms.GCP...)
			permissionBag.AZURE = append(permissionBag.AZURE, newPerms.AZURE...)
			permissionBag.RuntimeAWS = append(permissionBag.RuntimeAWS, newPerms.RuntimeAWS...)
			permissionBag.RuntimeGCP = append(permissionBag.RuntimeGCP, newPerms.RuntimeGCP...)
			permissionBag.RuntimeAZURE = append(permissionBag.RuntimeAZURE, newPerms.RuntimeAZURE...)
		}
	}
	return permissionBag
}
