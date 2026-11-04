# Deletion-Sensitive Permissions Warning

**Status: Implemented**

## Overview

Some Terraform resources are expensive, time-consuming, or impossible to recreate once destroyed. Pike now warns when `destroy` permissions for these resources are present in a scan, helping teams understand the blast radius of IAM permissions they're granting.

## Usage

```bash
# See warnings with solutions (automatic on every scan)
pike scan -d ./terraform

# Generate three-role split (safe + delete + plan)
pike scan -d ./terraform --sensitive
```

## Output

The `--sensitive` flag generates three Terraform roles:

- `terraform_pike_safe` — day-to-day operations (create, read, update) without high-risk deletes
- `terraform_pike_delete` — high-risk destruction permissions only (break-glass)
- `terraform_pike_plan` — read-only permissions for `terraform plan`

## Future: Schema Change

Add optional `deletion_sensitive` field to resource JSON mappings:

```json
[
  {
    "apply": ["container.clusters.create", ...],
    "destroy": ["container.clusters.delete"],
    "deletion_sensitive": {
      "risk": "high",
      "reason": "GKE clusters take 10+ minutes to recreate; workloads experience downtime",
      "protection": "deletion_protection = true"
    }
  }
]
```

### Risk Levels

| Level | Criteria | Examples |
|-------|----------|----------|
| `high` | Data loss, 10+ min recreation, production impact | GKE/EKS clusters, Cloud SQL, RDS, storage buckets with data |
| `medium` | Expensive recreation, networking disruption | VPCs with many subnets, load balancers, DNS zones |
| `low` | Moderate inconvenience | Service accounts with many bindings, custom roles |

## Implementation

### New File: `src/deletion_sensitive.go`

```go
package pike

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
}

var deletionSensitiveAWS = map[string]DeletionRisk{
    "rds:DeleteDBCluster": {
        Level:      "high",
        Resource:   "aws_rds_cluster",
        Reason:     "RDS clusters contain stateful data; recreation loses transaction history",
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
}

type DeletionRisk struct {
    Level      string // high, medium, low
    Resource   string // Terraform resource type
    Reason     string // Why this is deletion-sensitive
    Protection string // Suggested protection mechanism
}

// WarnDeletionSensitive prints warnings for high-risk destroy permissions.
func WarnDeletionSensitive(bag Sorted) {
    var found []DeletionRisk

    for _, p := range Unique(bag.GCP) {
        if risk, ok := deletionSensitiveGCP[p]; ok {
            found = append(found, risk)
        }
    }
    for _, p := range Unique(bag.AWS) {
        if risk, ok := deletionSensitiveAWS[p]; ok {
            found = append(found, risk)
        }
    }

    if len(found) == 0 {
        return
    }

    fmt.Fprintln(os.Stderr, "WARNING: deletion-sensitive permissions detected — granting destroy on stateful/expensive resources.")
    fmt.Fprintln(os.Stderr, "         Consider enabling deletion protection or using lifecycle { prevent_destroy = true }.")
    fmt.Fprintln(os.Stderr, "")

    for _, risk := range found {
        fmt.Fprintf(os.Stderr, "  [%s] %s\n", strings.ToUpper(risk.Level), risk.Resource)
        fmt.Fprintf(os.Stderr, "         %s\n", risk.Reason)
        fmt.Fprintf(os.Stderr, "         Protection: %s\n", risk.Protection)
    }
}
```

### Integration in `scan.go`

```go
func Scan(...) error {
    // ... existing code ...

    if !noLint {
        WarnEscalation(bag)
        WarnDeletionSensitive(bag)  // <-- Add here

        if strict && (!findEscalation(bag).empty() || len(findEscalationCombos(bag)) > 0) {
            return &escalationClassError{findEscalation(bag)}
        }
    }
    // ...
}
```

## Example Output

```
$ pike scan -d ./terraform

WARNING: deletion-sensitive permissions detected — granting destroy on stateful/expensive resources.
         Consider enabling deletion protection or using lifecycle { prevent_destroy = true }.

  [HIGH] google_container_cluster
         GKE clusters take 10+ minutes to recreate; workloads experience downtime
         Protection: deletion_protection = true

  [HIGH] google_sql_database_instance
         Cloud SQL instances contain stateful data; point-in-time recovery may be lost
         Protection: deletion_protection_enabled = true
```

## Future Enhancements

1. **`--strict-delete` flag**: Exit non-zero when high-risk delete permissions are present
2. **JSON mapping integration**: Move risk metadata into JSON files for easier maintenance
3. **Risk tiers in policy output**: Group permissions by risk level in the generated policy
4. **Audit mode**: List all deletion-sensitive resources in a codebase, regardless of permissions

## Resources to Classify

### High Risk (stateful, expensive recreation)
- GKE/EKS/AKS clusters
- Cloud SQL/RDS/Azure SQL instances
- Storage buckets (GCS/S3/Azure Blob)
- KMS keys
- Pub/Sub topics with subscriptions
- BigQuery datasets with tables
- Firestore databases
- Spanner instances

### Medium Risk (networking disruption, moderate recreation time)
- VPCs with many subnets/peerings
- Load balancers with backends
- DNS managed zones
- VPN gateways
- Cloud NAT
- Service accounts with many bindings

### Low Risk (inconvenient but recoverable)
- Custom IAM roles
- Firewall rules
- Cloud Scheduler jobs
- Monitoring alert policies
