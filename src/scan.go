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

// Runtime detects runtime IAM permissions needed by service accounts.
func Runtime(dirName string, outputType string, file *string, init bool, provider string) error {
	if dirName == "" && file == nil {
		return &emptyScanLocationError{}
	}

	permissionsBag, err := makePermissionBag(dirName, file, init, provider)
	if err != nil {
		return fmt.Errorf("failed to create permission bag: %w", err)
	}

	// Output runtime permissions
	output := formatRuntimePermissions(permissionsBag, provider)
	if output == "" {
		fmt.Println("No runtime permissions detected.")
		return nil
	}

	fmt.Print(output)
	return nil
}

// formatRuntimePermissions formats runtime permissions for output.
func formatRuntimePermissions(perms Sorted, provider string) string {
	var output string

	switch strings.ToLower(provider) {
	case "gcp", "google":
		if len(perms.RuntimeGCP) > 0 {
			output += formatGCPRuntimeWithValidation(perms.RuntimeGCP, perms.IAMBindings)
		}
	case "aws":
		// TODO: Implement AWS runtime permission formatting
		if len(perms.RuntimeAWS) > 0 {
			output += "=== AWS Runtime Permissions ===\n"
			output += "(AWS runtime permission detection not yet implemented)\n\n"
		}
	case "azure", "azurerm":
		// TODO: Implement Azure runtime permission formatting
		if len(perms.RuntimeAZURE) > 0 {
			output += "=== Azure Runtime Permissions ===\n"
			output += "(Azure runtime permission detection not yet implemented)\n\n"
		}
	case "":
		// Show all
		if len(perms.RuntimeGCP) > 0 {
			output += formatGCPRuntimeWithValidation(perms.RuntimeGCP, perms.IAMBindings)
		}
		if len(perms.RuntimeAWS) > 0 {
			output += "=== AWS Runtime Permissions ===\n"
			output += "(AWS runtime permission detection not yet implemented)\n\n"
		}
		if len(perms.RuntimeAZURE) > 0 {
			output += "=== Azure Runtime Permissions ===\n"
			output += "(Azure runtime permission detection not yet implemented)\n\n"
		}
	}

	return output
}

// ValidationResult tracks the status of an IAM binding requirement.
type ValidationResult struct {
	ResourceType   string
	ResourceName   string
	ServiceAccount string
	Role           string
	Permissions    []string
	Status         string // "configured", "missing", "partial"
	ExistingMember string // If configured, what member string is used
}

// validateRuntimePermissions checks if required runtime permissions have IAM bindings.
func validateRuntimePermissions(runtimePerms []RuntimePermission, iamBindings []IAMBinding, permissionToRole map[string]string) []ValidationResult {
	var results []ValidationResult

	// Build a map of role -> IAM bindings for quick lookup
	roleBindings := make(map[string][]IAMBinding)
	for _, binding := range iamBindings {
		roleBindings[binding.Role] = append(roleBindings[binding.Role], binding)
	}

	// Check each runtime permission requirement
	for _, runtimePerm := range runtimePerms {
		if len(runtimePerm.Permissions) == 0 {
			continue
		}

		// Group permissions by role
		rolesNeeded := make(map[string][]string)
		for _, perm := range runtimePerm.Permissions {
			if role, ok := permissionToRole[perm]; ok {
				rolesNeeded[role] = append(rolesNeeded[role], perm)
			}
		}

		// Check each role
		for role, perms := range rolesNeeded {
			result := ValidationResult{
				ResourceType:   runtimePerm.ResourceType,
				ResourceName:   runtimePerm.ResourceName,
				ServiceAccount: runtimePerm.ServiceAccount,
				Role:           role,
				Permissions:    perms,
				Status:         "missing",
			}

			// Check if this role has an IAM binding
			if bindings, exists := roleBindings[role]; exists {
				// Check if any binding references this service account
				serviceAccountRef := extractServiceAccountRef(runtimePerm)
				for _, binding := range bindings {
					if matchesServiceAccount(binding.Member, serviceAccountRef, runtimePerm) {
						result.Status = "configured"
						result.ExistingMember = binding.Member
						break
					}
				}
			}

			results = append(results, result)
		}
	}

	return results
}

// extractServiceAccountRef creates a service account reference from runtime permission.
func extractServiceAccountRef(runtimePerm RuntimePermission) string {
	if runtimePerm.ServiceAccount == "custom" {
		return fmt.Sprintf("${%s.%s.service_account}", runtimePerm.ResourceType, runtimePerm.ResourceName)
	}
	return "" // default service account - can't match
}

// matchesServiceAccount checks if an IAM binding member matches a service account reference.
func matchesServiceAccount(member string, serviceAccountRef string, runtimePerm RuntimePermission) bool {
	if serviceAccountRef == "" {
		// Default service account - we can't reliably match
		return false
	}

	// Check for exact match (for Terraform references)
	memberRef := strings.TrimPrefix(member, "serviceAccount:")
	return strings.Contains(memberRef, serviceAccountRef)
}

// formatGCPRuntimeWithValidation generates output with validation results.
func formatGCPRuntimeWithValidation(runtimePerms []RuntimePermission, iamBindings []IAMBinding) string {
	if len(runtimePerms) == 0 {
		return ""
	}

	// Permission to role mapping
	permissionToRole := buildPermissionToRoleMap()

	// Validate runtime permissions against IAM bindings
	validationResults := validateRuntimePermissions(runtimePerms, iamBindings, permissionToRole)

	var output strings.Builder

	// Group results by resource
	resourceResults := make(map[string][]ValidationResult)
	for _, result := range validationResults {
		key := fmt.Sprintf("%s.%s", result.ResourceType, result.ResourceName)
		resourceResults[key] = append(resourceResults[key], result)
	}

	// Process each resource
	for _, runtimePerm := range runtimePerms {
		if len(runtimePerm.Permissions) == 0 {
			continue
		}

		key := fmt.Sprintf("%s.%s", runtimePerm.ResourceType, runtimePerm.ResourceName)
		results := resourceResults[key]

		// Header for this resource
		output.WriteString(fmt.Sprintf("# Resource: %s.%s\n", runtimePerm.ResourceType, runtimePerm.ResourceName))

		// Warning about default service account
		if runtimePerm.ServiceAccount == "default" {
			output.WriteString("# ⚠️  WARNING: Using default service account (broad permissions).\n")
			output.WriteString("#     Best practice: Define a dedicated service account with least-privilege access.\n")
		} else if runtimePerm.ServiceAccount == "custom" {
			output.WriteString(fmt.Sprintf("# ℹ️  Service account: Defined in resource (reference as ${%s.%s.service_account})\n",
				runtimePerm.ResourceType, runtimePerm.ResourceName))
		}
		output.WriteString("#\n")

		// Validation status
		configured := 0
		missing := 0
		for _, result := range results {
			if result.Status == "configured" {
				configured++
			} else {
				missing++
			}
		}

		if configured > 0 && missing == 0 {
			output.WriteString("# ✅ IAM Status: All required permissions are configured\n")
		} else if configured > 0 {
			output.WriteString(fmt.Sprintf("# ⚠️  IAM Status: %d configured, %d missing\n", configured, missing))
		} else {
			output.WriteString(fmt.Sprintf("# ❌ IAM Status: No IAM bindings found (%d missing)\n", missing))
		}
		output.WriteString("#\n")

		// Show validation details
		if len(results) > 0 {
			output.WriteString("# IAM Binding Requirements:\n")
			for _, result := range results {
				if result.Status == "configured" {
					output.WriteString(fmt.Sprintf("#   ✅ %s - CONFIGURED\n", result.Role))
					output.WriteString(fmt.Sprintf("#      Member: %s\n", result.ExistingMember))
				} else {
					output.WriteString(fmt.Sprintf("#   ❌ %s - MISSING\n", result.Role))
					output.WriteString("#      Permissions: ")
					output.WriteString(strings.Join(result.Permissions, ", "))
					output.WriteString("\n")
				}
			}
			output.WriteString("#\n")
		}

		// Generate Terraform code for missing bindings
		missingResults := []ValidationResult{}
		for _, result := range results {
			if result.Status == "missing" {
				missingResults = append(missingResults, result)
			}
		}

		if len(missingResults) > 0 {
			output.WriteString("# Suggested IAM bindings to add:\n#\n")

			for _, result := range missingResults {
				roleID := strings.ReplaceAll(strings.Split(result.Role, "/")[1], ".", "_")
				resourceID := strings.ReplaceAll(runtimePerm.ResourceName, "-", "_")
				bindingName := fmt.Sprintf("runtime_%s_%s", resourceID, roleID)

				output.WriteString(fmt.Sprintf("resource \"google_project_iam_member\" \"%s\" {\n", bindingName))
				output.WriteString("  project = var.project_id\n")
				output.WriteString(fmt.Sprintf("  role    = \"%s\"\n", result.Role))

				if runtimePerm.ServiceAccount == "custom" {
					output.WriteString(fmt.Sprintf("  member  = \"serviceAccount:${%s.%s.service_account}\"\n",
						runtimePerm.ResourceType, runtimePerm.ResourceName))
				} else {
					output.WriteString("  member  = \"serviceAccount:YOUR_SERVICE_ACCOUNT_EMAIL\"  # TODO: Replace with actual service account\n")
				}

				output.WriteString("}\n\n")
			}
		}

		// List all detected permissions
		output.WriteString("# All permissions detected:\n")
		uniquePerms := Unique(runtimePerm.Permissions)
		for _, perm := range uniquePerms {
			output.WriteString(fmt.Sprintf("#   - %s\n", perm))
		}
		output.WriteString("\n")
	}

	return output.String()
}

// buildPermissionToRoleMap creates the GCP permission to role mapping.
func buildPermissionToRoleMap() map[string]string {
	return map[string]string{
		// Secret Manager
		"secretmanager.versions.access": "roles/secretmanager.secretAccessor",

		// CloudSQL
		"cloudsql.instances.connect": "roles/cloudsql.client",
		"cloudsql.instances.get":     "roles/cloudsql.client",

		// Cloud Storage
		"storage.objects.get":    "roles/storage.objectViewer",
		"storage.objects.create": "roles/storage.objectCreator",
		"storage.objects.list":   "roles/storage.objectViewer",
		"storage.objects.delete": "roles/storage.objectAdmin",
		"storage.buckets.get":    "roles/storage.objectViewer",

		// Pub/Sub
		"pubsub.topics.publish":        "roles/pubsub.publisher",
		"pubsub.subscriptions.consume": "roles/pubsub.subscriber",

		// BigQuery
		"bigquery.datasets.get":   "roles/bigquery.dataViewer",
		"bigquery.tables.get":     "roles/bigquery.dataViewer",
		"bigquery.tables.getData": "roles/bigquery.dataViewer",
		"bigquery.jobs.create":    "roles/bigquery.jobUser",

		// Artifact Registry
		"artifactregistry.repositories.downloadArtifacts": "roles/artifactregistry.reader",

		// Logging & Monitoring
		"logging.logEntries.create":    "roles/logging.logWriter",
		"monitoring.timeSeries.create": "roles/monitoring.metricWriter",

		// IAM
		"iam.serviceAccounts.getAccessToken": "roles/iam.workloadIdentityUser",
		"iam.serviceAccounts.actAs":          "roles/iam.serviceAccountUser",

		// KMS
		"cloudkms.cryptoKeyVersions.useToDecrypt": "roles/cloudkms.cryptoKeyDecrypter",

		// Compute
		"compute.addresses.use": "roles/compute.networkUser",

		// Cloud Functions
		"cloudfunctions.functions.invoke": "roles/cloudfunctions.invoker",

		// Cloud Run
		"run.routes.invoke": "roles/run.invoker",

		// Composer
		"composer.environments.get": "roles/composer.worker",

		// Cloud Tasks
		"cloudtasks.tasks.create": "roles/cloudtasks.enqueuer",

		// Firestore
		"firestore.documents.get":    "roles/datastore.user",
		"firestore.documents.create": "roles/datastore.user",

		// Datastore
		"datastore.entities.get":    "roles/datastore.user",
		"datastore.entities.create": "roles/datastore.user",

		// Spanner
		"spanner.databases.read":  "roles/spanner.databaseReader",
		"spanner.sessions.create": "roles/spanner.databaseReader",
		"spanner.databases.write": "roles/spanner.databaseUser",

		// App Engine
		"appengine.applications.get": "roles/appengine.appViewer",

		// Artifact Registry (upload)
		"artifactregistry.repositories.uploadArtifacts": "roles/artifactregistry.writer",

		// Bigtable
		"bigtable.tables.readRows":   "roles/bigtable.reader",
		"bigtable.tables.mutateRows": "roles/bigtable.user",

		// Workflows
		"workflows.executions.create": "roles/workflows.invoker",

		// Eventarc
		"eventarc.events.receiveEvent": "roles/eventarc.eventReceiver",

		// Metastore
		"metastore.tables.get":  "roles/metastore.metadataViewer",
		"metastore.tables.list": "roles/metastore.metadataViewer",

		// Cloud Run (services management for Cloud Build)
		"run.services.get":    "roles/run.developer",
		"run.services.update": "roles/run.developer",

		// Cloud Functions (management for Cloud Build)
		"cloudfunctions.functions.get":    "roles/cloudfunctions.developer",
		"cloudfunctions.functions.update": "roles/cloudfunctions.developer",

		// BigQuery (additional permissions)
		"bigquery.tables.create": "roles/bigquery.dataEditor",
		"bigquery.tables.update": "roles/bigquery.dataEditor",

		// Dataflow
		"dataflow.jobs.get": "roles/dataflow.worker",
	}
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

		err := os.MkdirAll(newPath, 0o750)

		if err != nil {
			return &makeDirectoryError{directory: newPath, err: err}
		}

		switch strings.ToLower(outputType) {
		case terraform:
			outFile = filepath.Join(newPath, "pike.generated_policy.tf") //path.join does not work here

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

	permissionsBag := GetPermissionBag(resources, provider)
	permissionsBag.IAMBindings = allIAMBindings
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
			permissionBag.RuntimeAWS = append(permissionBag.RuntimeAWS, newPerms.RuntimeAWS...)
		case "gcp", "google":
			permissionBag.GCP = append(permissionBag.GCP, newPerms.GCP...)
			permissionBag.RuntimeGCP = append(permissionBag.RuntimeGCP, newPerms.RuntimeGCP...)
		case "azure", "azurerm":
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
