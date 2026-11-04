package pike

import (
	"strings"
	"testing"
)

func TestFindDeletionSensitive(t *testing.T) {
	tests := []struct {
		name     string
		bag      Sorted
		wantGCP  int
		wantAWS  int
		wantHigh int
	}{
		{
			name: "GKE cluster deletion",
			bag: Sorted{
				GCP: []string{
					"container.clusters.create",
					"container.clusters.delete",
					"container.clusters.get",
				},
			},
			wantGCP:  1,
			wantHigh: 1,
		},
		{
			name: "Multiple GCP resources",
			bag: Sorted{
				GCP: []string{
					"container.clusters.delete",
					"cloudsql.instances.delete",
					"storage.buckets.delete",
				},
			},
			wantGCP:  3,
			wantHigh: 3,
		},
		{
			name: "AWS RDS cluster",
			bag: Sorted{
				AWS: []string{
					"rds:CreateDBCluster",
					"rds:DeleteDBCluster",
					"rds:DescribeDBClusters",
				},
			},
			wantAWS:  1,
			wantHigh: 1,
		},
		{
			name: "No deletion-sensitive permissions",
			bag: Sorted{
				GCP: []string{
					"container.clusters.create",
					"container.clusters.get",
					"compute.instances.get",
				},
			},
			wantGCP:  0,
			wantHigh: 0,
		},
		{
			name: "Mixed providers",
			bag: Sorted{
				GCP: []string{"container.clusters.delete"},
				AWS: []string{"rds:DeleteDBCluster"},
			},
			wantGCP:  1,
			wantAWS:  1,
			wantHigh: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := findDeletionSensitive(tt.bag)

			if len(found.GCP) != tt.wantGCP {
				t.Errorf("GCP count = %d, want %d", len(found.GCP), tt.wantGCP)
			}
			if len(found.AWS) != tt.wantAWS {
				t.Errorf("AWS count = %d, want %d", len(found.AWS), tt.wantAWS)
			}
			if found.HighRiskCount() != tt.wantHigh {
				t.Errorf("HighRiskCount = %d, want %d", found.HighRiskCount(), tt.wantHigh)
			}
		})
	}
}

func TestExtractDeletionSensitive(t *testing.T) {
	bag := Sorted{
		GCP: []string{
			"container.clusters.create",
			"container.clusters.delete",
			"container.clusters.get",
			"cloudsql.instances.delete",
			"compute.instances.get",
		},
		PlanGCP: []string{
			"container.clusters.get",
		},
	}

	policy := ExtractDeletionSensitive(bag)

	// Check safe permissions don't include deletion-sensitive ones
	for _, p := range policy.Safe.GCP {
		if p == "container.clusters.delete" || p == "cloudsql.instances.delete" {
			t.Errorf("Safe.GCP should not contain deletion-sensitive permission: %s", p)
		}
	}

	// Check deletion-sensitive permissions are extracted
	if len(policy.DeletionSensitive.GCP) != 2 {
		t.Errorf("DeletionSensitive.GCP = %d, want 2", len(policy.DeletionSensitive.GCP))
	}

	// Plan permissions should be preserved in safe
	if len(policy.Safe.PlanGCP) != 1 {
		t.Errorf("Safe.PlanGCP = %d, want 1", len(policy.Safe.PlanGCP))
	}
}

func TestMakeDeletionSensitiveOutput(t *testing.T) {
	bag := Sorted{
		GCP: []string{
			"container.clusters.create",
			"container.clusters.delete",
			"container.clusters.get",
		},
		PlanGCP: []string{
			"container.clusters.get",
		},
	}

	output, err := MakeDeletionSensitiveOutput(bag, "")
	if err != nil {
		t.Fatalf("MakeDeletionSensitiveOutput failed: %v", err)
	}

	// Check output contains actual Terraform resources
	if !strings.Contains(output, "google_project_iam_custom_role") {
		t.Error("Output should contain google_project_iam_custom_role resource")
	}
	if !strings.Contains(output, "terraform_pike_safe") {
		t.Error("Output should contain terraform_pike_safe role")
	}
	if !strings.Contains(output, "terraform_pike_delete") {
		t.Error("Output should contain terraform_pike_delete role")
	}
	if !strings.Contains(output, "container.clusters.delete") {
		t.Error("Output should contain container.clusters.delete permission")
	}
	if !strings.Contains(output, "deletion_protection = true") {
		t.Error("Output should contain protection suggestion in comments")
	}
}

func TestDeletionSensitiveSetEmpty(t *testing.T) {
	empty := deletionSensitiveSet{}
	if !empty.empty() {
		t.Error("Empty set should return true for empty()")
	}

	notEmpty := deletionSensitiveSet{
		GCP: []DeletionRisk{{Level: "high", Resource: "test"}},
	}
	if notEmpty.empty() {
		t.Error("Non-empty set should return false for empty()")
	}
}
