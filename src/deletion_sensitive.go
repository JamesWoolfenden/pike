package pike

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"strings"
	"text/template"
)

//go:embed terraform.three-role.gcp.template
var threeRoleGCPTemplate []byte

//go:embed terraform.three-role.aws.template
var threeRoleAWSTemplate []byte

//go:embed terraform.three-role.azure.template
var threeRoleAzureTemplate []byte

// DeletionRisk describes the risk of destroying a resource.
type DeletionRisk struct {
	Level      string // high, medium, low
	Resource   string // Terraform resource type
	Reason     string // Why this is deletion-sensitive
	Protection string // Suggested protection mechanism
}

// deletionSensitiveGCP maps destroy permissions to risk metadata.
var deletionSensitiveGCP = map[string]DeletionRisk{
	"container.clusters.delete": {
		Level:      "high",
		Resource:   "google_container_cluster",
		Reason:     "GKE clusters take 10+ minutes to recreate; workloads experience downtime",
		Protection: "deletion_protection = true",
	},
	"cloudsql.instances.delete": {
		Level:      "high",
		Resource:   "google_sql_database_instance",
		Reason:     "Cloud SQL instances contain stateful data; point-in-time recovery may be lost",
		Protection: "deletion_protection_enabled = true",
	},
	"storage.buckets.delete": {
		Level:      "high",
		Resource:   "google_storage_bucket",
		Reason:     "Bucket deletion destroys all objects; versioned history is lost",
		Protection: "force_destroy = false (default)",
	},
	"cloudkms.cryptoKeys.destroy": {
		Level:      "high",
		Resource:   "google_kms_crypto_key",
		Reason:     "Key destruction is irreversible after 24h; encrypted data becomes unrecoverable",
		Protection: "prevent key version destruction via IAM",
	},
	"bigtable.instances.delete": {
		Level:      "high",
		Resource:   "google_bigtable_instance",
		Reason:     "Bigtable instances contain stateful data; deletion is permanent",
		Protection: "deletion_protection = true",
	},
	"spanner.instances.delete": {
		Level:      "high",
		Resource:   "google_spanner_instance",
		Reason:     "Spanner instances contain stateful data; deletion is permanent",
		Protection: "force_destroy = false (default)",
	},
	"redis.instances.delete": {
		Level:      "medium",
		Resource:   "google_redis_instance",
		Reason:     "Redis instances may contain cached state; recreation takes 5+ minutes",
		Protection: "N/A (use lifecycle prevent_destroy)",
	},
	"compute.networks.delete": {
		Level:      "medium",
		Resource:   "google_compute_network",
		Reason:     "VPC deletion cascades to subnets, firewall rules, and peerings",
		Protection: "delete_default_routes_on_create = false, verify no dependents",
	},
	"dns.managedZones.delete": {
		Level:      "medium",
		Resource:   "google_dns_managed_zone",
		Reason:     "DNS zone deletion removes all records; propagation delay on recreation",
		Protection: "N/A (use lifecycle prevent_destroy)",
	},
}

// deletionSensitiveAWS maps destroy permissions to risk metadata.
var deletionSensitiveAWS = map[string]DeletionRisk{
	"rds:DeleteDBCluster": {
		Level:      "high",
		Resource:   "aws_rds_cluster",
		Reason:     "RDS clusters contain stateful data; recreation loses transaction history",
		Protection: "deletion_protection = true",
	},
	"rds:DeleteDBInstance": {
		Level:      "high",
		Resource:   "aws_db_instance",
		Reason:     "RDS instances contain stateful data; recreation loses transaction history",
		Protection: "deletion_protection = true",
	},
	"eks:DeleteCluster": {
		Level:      "high",
		Resource:   "aws_eks_cluster",
		Reason:     "EKS clusters take 15+ minutes to recreate; workloads experience downtime",
		Protection: "N/A (use lifecycle prevent_destroy)",
	},
	"s3:DeleteBucket": {
		Level:      "high",
		Resource:   "aws_s3_bucket",
		Reason:     "Bucket deletion destroys all objects; versioned history is lost",
		Protection: "force_destroy = false (default)",
	},
	"dynamodb:DeleteTable": {
		Level:      "high",
		Resource:   "aws_dynamodb_table",
		Reason:     "DynamoDB table deletion destroys all items; PITR backups may be lost",
		Protection: "deletion_protection_enabled = true",
	},
	"kms:ScheduleKeyDeletion": {
		Level:      "high",
		Resource:   "aws_kms_key",
		Reason:     "Key deletion is irreversible after waiting period; encrypted data becomes unrecoverable",
		Protection: "deletion_window_in_days = 30, enable_key_rotation = true",
	},
	"elasticache:DeleteCacheCluster": {
		Level:      "medium",
		Resource:   "aws_elasticache_cluster",
		Reason:     "ElastiCache deletion loses cached data; recreation takes several minutes",
		Protection: "N/A (use lifecycle prevent_destroy)",
	},
	"ec2:DeleteVpc": {
		Level:      "medium",
		Resource:   "aws_vpc",
		Reason:     "VPC deletion cascades to subnets, security groups, and route tables",
		Protection: "verify no dependents before deletion",
	},
	"route53:DeleteHostedZone": {
		Level:      "medium",
		Resource:   "aws_route53_zone",
		Reason:     "Hosted zone deletion removes all records; propagation delay on recreation",
		Protection: "N/A (use lifecycle prevent_destroy)",
	},
}

// deletionSensitiveAzure maps destroy permissions to risk metadata.
var deletionSensitiveAzure = map[string]DeletionRisk{
	"Microsoft.ContainerService/managedClusters/delete": {
		Level:      "high",
		Resource:   "azurerm_kubernetes_cluster",
		Reason:     "AKS clusters take 10+ minutes to recreate; workloads experience downtime",
		Protection: "N/A (use lifecycle prevent_destroy)",
	},
	"Microsoft.Sql/servers/delete": {
		Level:      "high",
		Resource:   "azurerm_mssql_server",
		Reason:     "SQL Server deletion removes all databases; data is permanently lost",
		Protection: "N/A (use lifecycle prevent_destroy)",
	},
	"Microsoft.Storage/storageAccounts/delete": {
		Level:      "high",
		Resource:   "azurerm_storage_account",
		Reason:     "Storage account deletion destroys all containers and blobs",
		Protection: "N/A (use lifecycle prevent_destroy, enable soft delete)",
	},
	"Microsoft.KeyVault/vaults/delete": {
		Level:      "high",
		Resource:   "azurerm_key_vault",
		Reason:     "Key Vault deletion loses secrets and keys; soft-delete recovery window is limited",
		Protection: "soft_delete_retention_days = 90, purge_protection_enabled = true",
	},
}

// deletionSensitiveSet groups deletion-sensitive permissions found per provider.
type deletionSensitiveSet struct {
	AWS   []DeletionRisk
	GCP   []DeletionRisk
	AZURE []DeletionRisk
}

func (d deletionSensitiveSet) empty() bool {
	return len(d.AWS) == 0 && len(d.GCP) == 0 && len(d.AZURE) == 0
}

// HighRiskCount returns the number of high-risk deletions.
func (d deletionSensitiveSet) HighRiskCount() int {
	count := 0
	for _, r := range d.AWS {
		if r.Level == "high" {
			count++
		}
	}
	for _, r := range d.GCP {
		if r.Level == "high" {
			count++
		}
	}
	for _, r := range d.AZURE {
		if r.Level == "high" {
			count++
		}
	}
	return count
}

// findDeletionSensitive identifies deletion-sensitive permissions in a Sorted bag.
func findDeletionSensitive(bag Sorted) deletionSensitiveSet {
	var found deletionSensitiveSet
	seen := make(map[string]bool)

	for _, p := range Unique(bag.GCP) {
		if risk, ok := deletionSensitiveGCP[p]; ok && !seen[risk.Resource] {
			found.GCP = append(found.GCP, risk)
			seen[risk.Resource] = true
		}
	}
	for _, p := range Unique(bag.AWS) {
		if risk, ok := deletionSensitiveAWS[p]; ok && !seen[risk.Resource] {
			found.AWS = append(found.AWS, risk)
			seen[risk.Resource] = true
		}
	}
	for _, p := range Unique(bag.AZURE) {
		if risk, ok := deletionSensitiveAzure[p]; ok && !seen[risk.Resource] {
			found.AZURE = append(found.AZURE, risk)
			seen[risk.Resource] = true
		}
	}
	return found
}

// WarnDeletionSensitive prints warnings for high-risk destroy permissions.
func WarnDeletionSensitive(bag Sorted) {
	found := findDeletionSensitive(bag)

	if found.empty() {
		return
	}

	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "WARNING: deletion-sensitive permissions detected — granting destroy on stateful/expensive resources.")
	fmt.Fprintln(os.Stderr, "         Use --output deletion-sensitive to generate a separate delete-only role.")

	printRisks := func(provider string, risks []DeletionRisk) {
		if len(risks) == 0 {
			return
		}
		fmt.Fprintf(os.Stderr, "\n  %s:\n", provider)
		for _, risk := range risks {
			fmt.Fprintf(os.Stderr, "    [%s] %s\n", strings.ToUpper(risk.Level), risk.Resource)
			fmt.Fprintf(os.Stderr, "           %s\n", risk.Reason)
			fmt.Fprintf(os.Stderr, "           → %s\n", risk.Protection)
		}
	}

	printRisks("GCP", found.GCP)
	printRisks("AWS", found.AWS)
	printRisks("Azure", found.AZURE)
	fmt.Fprintln(os.Stderr, "")
}

// DeletionSensitivePolicy holds permissions split into safe and deletion-sensitive.
type DeletionSensitivePolicy struct {
	// Safe contains permissions that don't involve high-risk destruction.
	Safe Sorted
	// DeletionSensitive contains only the high-risk destroy permissions.
	DeletionSensitive Sorted
	// Risks contains metadata about why each permission is sensitive.
	Risks deletionSensitiveSet
}

// ExtractDeletionSensitive splits a permission bag into safe and deletion-sensitive.
func ExtractDeletionSensitive(bag Sorted) DeletionSensitivePolicy {
	result := DeletionSensitivePolicy{
		Safe: Sorted{
			PlanAWS:      bag.PlanAWS,
			PlanGCP:      bag.PlanGCP,
			PlanAZURE:    bag.PlanAZURE,
			RuntimeGCP:   bag.RuntimeGCP,
			RuntimeAWS:   bag.RuntimeAWS,
			RuntimeAZURE: bag.RuntimeAZURE,
			IAMBindings:  bag.IAMBindings,
		},
		Risks: findDeletionSensitive(bag),
	}

	// Partition GCP permissions
	for _, p := range bag.GCP {
		if _, ok := deletionSensitiveGCP[p]; ok {
			result.DeletionSensitive.GCP = append(result.DeletionSensitive.GCP, p)
		} else {
			result.Safe.GCP = append(result.Safe.GCP, p)
		}
	}

	// Partition AWS permissions
	for _, p := range bag.AWS {
		if _, ok := deletionSensitiveAWS[p]; ok {
			result.DeletionSensitive.AWS = append(result.DeletionSensitive.AWS, p)
		} else {
			result.Safe.AWS = append(result.Safe.AWS, p)
		}
	}

	// Partition Azure permissions
	for _, p := range bag.AZURE {
		if _, ok := deletionSensitiveAzure[p]; ok {
			result.DeletionSensitive.AZURE = append(result.DeletionSensitive.AZURE, p)
		} else {
			result.Safe.AZURE = append(result.Safe.AZURE, p)
		}
	}

	return result
}

// gcpThreeRoleDetails holds template data for GCP three-role output.
type gcpThreeRoleDetails struct {
	SafeName          string
	SafeRoleID        string
	SafePermissions   string
	DeleteName        string
	DeleteRoleID      string
	DeletePermissions string
	PlanName          string
	PlanRoleID        string
	PlanPermissions   string
	Project           string
}

// awsThreeRoleDetails holds template data for AWS three-role output.
type awsThreeRoleDetails struct {
	SafeName     string
	SafePolicy   string
	DeleteName   string
	DeletePolicy string
	PlanName     string
	PlanPolicy   string
}

// azureThreeRoleDetails holds template data for Azure three-role output.
type azureThreeRoleDetails struct {
	SafeName          string
	SafePermissions   string
	DeleteName        string
	DeletePermissions string
	PlanName          string
	PlanPermissions   string
}

// MakeDeletionSensitiveOutput generates a three-role Terraform policy.
func MakeDeletionSensitiveOutput(bag Sorted, dirName string) (string, error) {
	policy := ExtractDeletionSensitive(bag)
	var s strings.Builder

	safeName, deleteName, planName := threeRoleNames("")

	// GCP
	if len(bag.GCP) > 0 {
		chunk, err := renderGCPThreeRole(policy, safeName, deleteName, planName, dirName)
		if err != nil {
			return "", err
		}
		s.WriteString(chunk)
	}

	// AWS
	if len(bag.AWS) > 0 {
		chunk, err := renderAWSThreeRole(policy, safeName, deleteName, planName)
		if err != nil {
			return "", err
		}
		s.WriteString(chunk)
	}

	// Azure
	if len(bag.AZURE) > 0 {
		chunk, err := renderAzureThreeRole(policy, safeName, deleteName, planName)
		if err != nil {
			return "", err
		}
		s.WriteString(chunk)
	}

	// Append risk details as comments
	if !policy.Risks.empty() {
		s.WriteString("\n# ============================================================\n")
		s.WriteString("# DELETION-SENSITIVE RESOURCES - Risk details and solutions\n")
		s.WriteString("# ============================================================\n")
		for _, r := range policy.Risks.GCP {
			fmt.Fprintf(&s, "# %s: %s\n", r.Resource, r.Reason)
			fmt.Fprintf(&s, "#   → Solution: %s\n", r.Protection)
		}
		for _, r := range policy.Risks.AWS {
			fmt.Fprintf(&s, "# %s: %s\n", r.Resource, r.Reason)
			fmt.Fprintf(&s, "#   → Solution: %s\n", r.Protection)
		}
		for _, r := range policy.Risks.AZURE {
			fmt.Fprintf(&s, "# %s: %s\n", r.Resource, r.Reason)
			fmt.Fprintf(&s, "#   → Solution: %s\n", r.Protection)
		}
	}

	return s.String(), nil
}

func threeRoleNames(policyName string) (safe, delete, plan string) {
	base := policyName
	if base == "" {
		base = defaultPolicyName
	}
	return base + "_safe", base + "_delete", base + "_plan"
}

func renderGCPThreeRole(policy DeletionSensitivePolicy, safeName, deleteName, planName, dirName string) (string, error) {
	project, err := getCurrentProject(dirName)
	if err != nil {
		project = defaultProject
	}

	details := gcpThreeRoleDetails{
		SafeName:          safeName,
		SafeRoleID:        safeName,
		SafePermissions:   strings.Join(Unique(policy.Safe.GCP), "\",\n    \""),
		DeleteName:        deleteName,
		DeleteRoleID:      deleteName,
		DeletePermissions: strings.Join(Unique(policy.DeletionSensitive.GCP), "\",\n    \""),
		PlanName:          planName,
		PlanRoleID:        planName,
		PlanPermissions:   strings.Join(Unique(policy.Safe.PlanGCP), "\",\n    \""),
		Project:           project,
	}

	return execThreeRoleTemplate(threeRoleGCPTemplate, details)
}

func renderAWSThreeRole(policy DeletionSensitivePolicy, safeName, deleteName, planName string) (string, error) {
	safePolicy, err := awsPolicyJSON(Unique(policy.Safe.AWS))
	if err != nil {
		return "", err
	}

	var deletePolicy string
	if len(policy.DeletionSensitive.AWS) > 0 {
		deletePolicy, err = awsPolicyJSON(Unique(policy.DeletionSensitive.AWS))
		if err != nil {
			return "", err
		}
	}

	planPolicy, err := awsPolicyJSON(Unique(policy.Safe.PlanAWS))
	if err != nil {
		return "", err
	}

	details := awsThreeRoleDetails{
		SafeName:     safeName,
		SafePolicy:   safePolicy,
		DeleteName:   deleteName,
		DeletePolicy: deletePolicy,
		PlanName:     planName,
		PlanPolicy:   planPolicy,
	}

	return execThreeRoleTemplate(threeRoleAWSTemplate, details)
}

func renderAzureThreeRole(policy DeletionSensitivePolicy, safeName, deleteName, planName string) (string, error) {
	details := azureThreeRoleDetails{
		SafeName:          safeName,
		SafePermissions:   strings.Join(Unique(policy.Safe.AZURE), "\",\n    \""),
		DeleteName:        deleteName,
		DeletePermissions: strings.Join(Unique(policy.DeletionSensitive.AZURE), "\",\n    \""),
		PlanName:          planName,
		PlanPermissions:   strings.Join(Unique(policy.Safe.PlanAZURE), "\",\n    \""),
	}

	return execThreeRoleTemplate(threeRoleAzureTemplate, details)
}

func execThreeRoleTemplate(tmplBytes []byte, data any) (string, error) {
	tmpl, err := template.New("").Parse(string(tmplBytes))
	if err != nil {
		return "", &templateParseError{err}
	}
	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, data); err != nil {
		return "", &templateExecuteError{err}
	}
	return buf.String(), nil
}
