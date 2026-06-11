package pike

import (
	"reflect"
	"testing"
)

func Test_getPlanPermissionMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		raw        string
		attributes []string
		isReadOnly func(string) bool
		want       []string
	}{
		{
			name: "gcp attribute write filtered (issue #135)",
			raw: `[{
				"apply": [],
				"attributes": {
					"ssl_certificates": [
						"compute.regionSslCertificates.get",
						"compute.regionTargetHttpsProxies.setSslCertificates"
					],
					"tags": []
				},
				"destroy": [], "modify": [],
				"plan": ["compute.regionTargetHttpsProxies.get"]
			}]`,
			attributes: []string{"ssl_certificates"},
			isReadOnly: isGCPReadPerm,
			want: []string{
				"compute.regionSslCertificates.get",
				"compute.regionTargetHttpsProxies.get",
			},
		},
		{
			name: "aws attribute write filtered",
			raw: `[{
				"apply": [],
				"attributes": {
					"target_group_arns": [
						"autoscaling:DescribeLoadBalancerTargetGroups",
						"autoscaling:AttachLoadBalancerTargetGroups",
						"autoscaling:DetachLoadBalancerTargetGroups"
					],
					"tags": []
				},
				"destroy": [], "modify": [],
				"plan": ["autoscaling:DescribeAutoScalingGroups"]
			}]`,
			attributes: []string{"target_group_arns"},
			isReadOnly: isAWSReadPerm,
			want: []string{
				"autoscaling:DescribeLoadBalancerTargetGroups",
				"autoscaling:DescribeAutoScalingGroups",
			},
		},
		{
			name: "azure attribute write and action filtered",
			raw: `[{
				"apply": [],
				"attributes": {
					"identity": [
						"Microsoft.ManagedIdentity/userAssignedIdentities/read",
						"Microsoft.ManagedIdentity/userAssignedIdentities/assign/action",
						"Microsoft.Network/applicationGateways/write"
					],
					"tags": []
				},
				"destroy": [], "modify": [],
				"plan": ["Microsoft.Network/applicationGateways/read"]
			}]`,
			attributes: []string{"identity"},
			isReadOnly: isAzureReadPerm,
			want: []string{
				"Microsoft.ManagedIdentity/userAssignedIdentities/read",
				"Microsoft.Network/applicationGateways/read",
			},
		},
		{
			name: "plan array passes through unfiltered",
			raw: `[{
				"apply": [],
				"attributes": {"tags": []},
				"destroy": [], "modify": [],
				"plan": ["svc.res.somethingUnusual"]
			}]`,
			attributes: nil,
			isReadOnly: isGCPReadPerm,
			want:       []string{"svc.res.somethingUnusual"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := getPlanPermissionMap([]byte(tt.raw), tt.attributes, tt.name, tt.isReadOnly)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isReadPerm(t *testing.T) {
	t.Parallel()

	tests := []struct {
		fn   func(string) bool
		perm string
		want bool
	}{
		{isGCPReadPerm, "compute.instances.get", true},
		{isGCPReadPerm, "compute.instances.list", true},
		{isGCPReadPerm, "iam.serviceAccounts.getIamPolicy", true},
		{isGCPReadPerm, "compute.regionTargetHttpsProxies.setSslCertificates", false},
		{isGCPReadPerm, "iam.serviceAccounts.actAs", false},
		{isGCPReadPerm, "noverb", false},
		{isAWSReadPerm, "ec2:DescribeInstances", true},
		{isAWSReadPerm, "s3:GetObject", true},
		{isAWSReadPerm, "s3:ListBucket", true},
		{isAWSReadPerm, "acm:AddTagsToCertificate", false},
		{isAWSReadPerm, "autoscaling:AttachLoadBalancerTargetGroups", false},
		{isAzureReadPerm, "Microsoft.KeyVault/vaults/read", true},
		{isAzureReadPerm, "Microsoft.KeyVault/vaults/write", false},
		{isAzureReadPerm, "Microsoft.ManagedIdentity/userAssignedIdentities/assign/action", false},
	}

	for _, tt := range tests {
		if got := tt.fn(tt.perm); got != tt.want {
			t.Errorf("%s: got %v, want %v", tt.perm, got, tt.want)
		}
	}
}
