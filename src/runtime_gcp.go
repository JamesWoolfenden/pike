package pike

import (
	"fmt"
	"strings"
)

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

func formatRuntimePermissions(perms Sorted, prov string) string {
	p := strings.ToLower(prov)
	if p != "gcp" && p != "google" && p != "" {
		return ""
	}

	if len(perms.RuntimeGCP) > 0 {
		return formatGCPRuntimeWithValidation(perms.RuntimeGCP, perms.IAMBindings)
	}

	return ""
}

// validateRuntimePermissions checks if required runtime permissions have IAM bindings.
func validateRuntimePermissions(runtimePerms []RuntimePermission, iamBindings []IAMBinding, permissionToRole map[string]string) []ValidationResult {
	var results []ValidationResult

	roleBindings := make(map[string][]IAMBinding)
	for _, binding := range iamBindings {
		roleBindings[binding.Role] = append(roleBindings[binding.Role], binding)
	}

	for _, runtimePerm := range runtimePerms {
		if len(runtimePerm.Permissions) == 0 {
			continue
		}

		rolesNeeded := make(map[string][]string)
		for _, perm := range runtimePerm.Permissions {
			if role, ok := permissionToRole[perm]; ok {
				rolesNeeded[role] = append(rolesNeeded[role], perm)
			}
		}

		for role, perms := range rolesNeeded {
			result := ValidationResult{
				ResourceType:   runtimePerm.ResourceType,
				ResourceName:   runtimePerm.ResourceName,
				ServiceAccount: runtimePerm.ServiceAccount,
				Role:           role,
				Permissions:    perms,
				Status:         "missing",
			}

			if bindings, exists := roleBindings[role]; exists {
				serviceAccountRef := extractServiceAccountRef(runtimePerm)
				for _, binding := range bindings {
					if matchesServiceAccount(binding.Member, serviceAccountRef) {
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

func extractServiceAccountRef(runtimePerm RuntimePermission) string {
	if runtimePerm.ServiceAccount == "custom" {
		return fmt.Sprintf("${%s.%s.service_account}", runtimePerm.ResourceType, runtimePerm.ResourceName)
	}
	return ""
}

func matchesServiceAccount(member string, serviceAccountRef string) bool {
	if serviceAccountRef == "" {
		return false
	}
	memberRef := strings.TrimPrefix(member, "serviceAccount:")
	return strings.Contains(memberRef, serviceAccountRef)
}

// formatGCPRuntimeWithValidation generates output with validation results.
func formatGCPRuntimeWithValidation(runtimePerms []RuntimePermission, iamBindings []IAMBinding) string {
	if len(runtimePerms) == 0 {
		return ""
	}

	permissionToRole := buildPermissionToRoleMap()
	validationResults := validateRuntimePermissions(runtimePerms, iamBindings, permissionToRole)

	var output strings.Builder

	resourceResults := make(map[string][]ValidationResult)
	for _, result := range validationResults {
		key := fmt.Sprintf("%s.%s", result.ResourceType, result.ResourceName)
		resourceResults[key] = append(resourceResults[key], result)
	}

	for _, runtimePerm := range runtimePerms {
		if len(runtimePerm.Permissions) == 0 {
			continue
		}

		key := fmt.Sprintf("%s.%s", runtimePerm.ResourceType, runtimePerm.ResourceName)
		results := resourceResults[key]

		fmt.Fprintf(&output, "# Resource: %s.%s\n", runtimePerm.ResourceType, runtimePerm.ResourceName)

		switch runtimePerm.ServiceAccount {
		case "default":
			output.WriteString("# ⚠️  WARNING: Using default service account (broad permissions).\n")
			output.WriteString("#     Best practice: Define a dedicated service account with least-privilege access.\n")
		case "custom":
			fmt.Fprintf(&output, "# ℹ️  Service account: Defined in resource (reference as ${%s.%s.service_account})\n",
				runtimePerm.ResourceType, runtimePerm.ResourceName)
		}
		output.WriteString("#\n")

		configured, missing := 0, 0
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
			fmt.Fprintf(&output, "# ⚠️  IAM Status: %d configured, %d missing\n", configured, missing)
		} else {
			fmt.Fprintf(&output, "# ❌ IAM Status: No IAM bindings found (%d missing)\n", missing)
		}
		output.WriteString("#\n")

		if len(results) > 0 {
			output.WriteString("# IAM Binding Requirements:\n")
			for _, result := range results {
				if result.Status == "configured" {
					fmt.Fprintf(&output, "#   ✅ %s - CONFIGURED\n", result.Role)
					fmt.Fprintf(&output, "#      Member: %s\n", result.ExistingMember)
				} else {
					fmt.Fprintf(&output, "#   ❌ %s - MISSING\n", result.Role)
					output.WriteString("#      Permissions: ")
					output.WriteString(strings.Join(result.Permissions, ", "))
					output.WriteString("\n")
				}
			}
			output.WriteString("#\n")
		}

		var missingResults []ValidationResult
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

				fmt.Fprintf(&output, "resource \"google_project_iam_member\" \"%s\" {\n", bindingName)
				output.WriteString("  project = var.project_id\n")
				fmt.Fprintf(&output, "  role    = \"%s\"\n", result.Role)
				if runtimePerm.ServiceAccount == "custom" {
					fmt.Fprintf(&output, "  member  = \"serviceAccount:${%s.%s.service_account}\"\n",
						runtimePerm.ResourceType, runtimePerm.ResourceName)
				} else {
					output.WriteString("  member  = \"serviceAccount:YOUR_SERVICE_ACCOUNT_EMAIL\"  # TODO: Replace with actual service account\n")
				}
				output.WriteString("}\n\n")
			}
		}

		output.WriteString("# All permissions detected:\n")
		for _, perm := range Unique(runtimePerm.Permissions) {
			fmt.Fprintf(&output, "#   - %s\n", perm)
		}
		output.WriteString("\n")
	}

	return output.String()
}

// buildPermissionToRoleMap maps GCP permissions to the predefined roles that grant them.
func buildPermissionToRoleMap() map[string]string {
	return map[string]string{
		"secretmanager.versions.access":                   "roles/secretmanager.secretAccessor",
		"cloudsql.instances.connect":                      "roles/cloudsql.client",
		"cloudsql.instances.get":                          "roles/cloudsql.client",
		"storage.objects.get":                             "roles/storage.objectViewer",
		"storage.objects.create":                          "roles/storage.objectCreator",
		"storage.objects.list":                            "roles/storage.objectViewer",
		"storage.objects.delete":                          "roles/storage.objectAdmin",
		"storage.buckets.get":                             "roles/storage.objectViewer",
		"pubsub.topics.publish":                           "roles/pubsub.publisher",
		"pubsub.subscriptions.consume":                    "roles/pubsub.subscriber",
		"bigquery.datasets.get":                           "roles/bigquery.dataViewer",
		"bigquery.tables.get":                             "roles/bigquery.dataViewer",
		"bigquery.tables.getData":                         "roles/bigquery.dataViewer",
		"bigquery.jobs.create":                            "roles/bigquery.jobUser",
		"bigquery.tables.create":                          "roles/bigquery.dataEditor",
		"bigquery.tables.update":                          "roles/bigquery.dataEditor",
		"artifactregistry.repositories.downloadArtifacts": "roles/artifactregistry.reader",
		"artifactregistry.repositories.uploadArtifacts":   "roles/artifactregistry.writer",
		"logging.logEntries.create":                       "roles/logging.logWriter",
		"monitoring.timeSeries.create":                    "roles/monitoring.metricWriter",
		"iam.serviceAccounts.getAccessToken":              "roles/iam.workloadIdentityUser",
		"iam.serviceAccounts.actAs":                       "roles/iam.serviceAccountUser",
		"cloudkms.cryptoKeyVersions.useToDecrypt":         "roles/cloudkms.cryptoKeyDecrypter",
		"compute.addresses.use":                           "roles/compute.networkUser",
		"cloudfunctions.functions.invoke":                 "roles/cloudfunctions.invoker",
		"cloudfunctions.functions.get":                    "roles/cloudfunctions.developer",
		"cloudfunctions.functions.update":                 "roles/cloudfunctions.developer",
		"run.routes.invoke":                               "roles/run.invoker",
		"run.services.get":                                "roles/run.developer",
		"run.services.update":                             "roles/run.developer",
		"composer.environments.get":                       "roles/composer.worker",
		"cloudtasks.tasks.create":                         "roles/cloudtasks.enqueuer",
		"firestore.documents.get":                         "roles/datastore.user",
		"firestore.documents.create":                      "roles/datastore.user",
		"datastore.entities.get":                          "roles/datastore.user",
		"datastore.entities.create":                       "roles/datastore.user",
		"spanner.databases.read":                          "roles/spanner.databaseReader",
		"spanner.sessions.create":                         "roles/spanner.databaseReader",
		"spanner.databases.write":                         "roles/spanner.databaseUser",
		"appengine.applications.get":                      "roles/appengine.appViewer",
		"bigtable.tables.readRows":                        "roles/bigtable.reader",
		"bigtable.tables.mutateRows":                      "roles/bigtable.user",
		"workflows.executions.create":                     "roles/workflows.invoker",
		"eventarc.events.receiveEvent":                    "roles/eventarc.eventReceiver",
		"metastore.tables.get":                            "roles/metastore.metadataViewer",
		"metastore.tables.list":                           "roles/metastore.metadataViewer",
		"dataflow.jobs.get":                               "roles/dataflow.worker",
	}
}
