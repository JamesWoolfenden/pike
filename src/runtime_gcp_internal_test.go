package pike

import (
	"strings"
	"testing"
)

func Test_extractServiceAccountRef(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		perm RuntimePermission
		want string
	}{
		{
			name: "custom service account",
			perm: RuntimePermission{ResourceType: "google_cloud_run_v2_service", ResourceName: "app", ServiceAccount: "custom"},
			want: "${google_cloud_run_v2_service.app.service_account}",
		},
		{
			name: "default service account",
			perm: RuntimePermission{ResourceType: "google_cloud_run_v2_service", ResourceName: "app", ServiceAccount: "default"},
			want: "",
		},
		{
			name: "empty service account",
			perm: RuntimePermission{},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := extractServiceAccountRef(tt.perm); got != tt.want {
				t.Errorf("extractServiceAccountRef() = %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_matchesServiceAccount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		member            string
		serviceAccountRef string
		want              bool
	}{
		{
			name:              "empty ref never matches",
			member:            "serviceAccount:sa@project.iam",
			serviceAccountRef: "",
			want:              false,
		},
		{
			name:              "matching reference",
			member:            "serviceAccount:${google_cloud_run_v2_service.app.service_account}",
			serviceAccountRef: "${google_cloud_run_v2_service.app.service_account}",
			want:              true,
		},
		{
			name:              "non-matching reference",
			member:            "serviceAccount:other@project.iam",
			serviceAccountRef: "${google_cloud_run_v2_service.app.service_account}",
			want:              false,
		},
		{
			name:              "strips serviceAccount: prefix before comparing",
			member:            "serviceAccount:${google_cloud_run_v2_service.app.service_account}",
			serviceAccountRef: "${google_cloud_run_v2_service.app.service_account}",
			want:              true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := matchesServiceAccount(tt.member, tt.serviceAccountRef); got != tt.want {
				t.Errorf("matchesServiceAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateRuntimePermissions(t *testing.T) {
	t.Parallel()

	permToRole := map[string]string{
		"secretmanager.versions.access": "roles/secretmanager.secretAccessor",
		"storage.objects.get":           "roles/storage.objectViewer",
	}

	tests := []struct {
		name         string
		runtimePerms []RuntimePermission
		iamBindings  []IAMBinding
		wantLen      int
		wantStatus   string
	}{
		{
			name:         "empty permissions",
			runtimePerms: []RuntimePermission{{Permissions: []string{}}},
			wantLen:      0,
		},
		{
			name: "missing binding",
			runtimePerms: []RuntimePermission{{
				ResourceType:   "google_cloud_run_v2_service",
				ResourceName:   "app",
				ServiceAccount: "custom",
				Permissions:    []string{"secretmanager.versions.access"},
			}},
			iamBindings: []IAMBinding{},
			wantLen:     1,
			wantStatus:  "missing",
		},
		{
			name: "configured binding",
			runtimePerms: []RuntimePermission{{
				ResourceType:   "google_cloud_run_v2_service",
				ResourceName:   "app",
				ServiceAccount: "custom",
				Permissions:    []string{"secretmanager.versions.access"},
			}},
			iamBindings: []IAMBinding{{
				Role:   "roles/secretmanager.secretAccessor",
				Member: "serviceAccount:${google_cloud_run_v2_service.app.service_account}",
			}},
			wantLen:    1,
			wantStatus: "configured",
		},
		{
			name: "unknown permission not in role map",
			runtimePerms: []RuntimePermission{{
				ResourceType: "google_foo",
				ResourceName: "bar",
				Permissions:  []string{"unknown.permission.here"},
			}},
			wantLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			results := validateRuntimePermissions(tt.runtimePerms, tt.iamBindings, permToRole)
			if len(results) != tt.wantLen {
				t.Errorf("validateRuntimePermissions() len = %d, want %d", len(results), tt.wantLen)
				return
			}
			if tt.wantStatus != "" && results[0].Status != tt.wantStatus {
				t.Errorf("validateRuntimePermissions()[0].Status = %q, want %q", results[0].Status, tt.wantStatus)
			}
		})
	}
}

func Test_formatRuntimePermissions(t *testing.T) {
	t.Parallel()

	perm := RuntimePermission{
		ResourceType:   "google_cloud_run_v2_service",
		ResourceName:   "app",
		ServiceAccount: "default",
		Permissions:    []string{"secretmanager.versions.access"},
	}

	tests := []struct {
		name      string
		perms     Sorted
		prov      string
		wantText  string
		wantEmpty bool
	}{
		{
			name:      "non-GCP provider returns empty",
			perms:     Sorted{RuntimeGCP: []RuntimePermission{perm}},
			prov:      "aws",
			wantEmpty: true,
		},
		{
			name:      "no runtime perms returns empty",
			perms:     Sorted{},
			prov:      "gcp",
			wantEmpty: true,
		},
		{
			name:     "GCP provider with permissions returns output",
			perms:    Sorted{RuntimeGCP: []RuntimePermission{perm}},
			prov:     "gcp",
			wantText: "google_cloud_run_v2_service.app",
		},
		{
			name:     "google alias also works",
			perms:    Sorted{RuntimeGCP: []RuntimePermission{perm}},
			prov:     "google",
			wantText: "google_cloud_run_v2_service.app",
		},
		{
			name:     "empty provider string is GCP",
			perms:    Sorted{RuntimeGCP: []RuntimePermission{perm}},
			prov:     "",
			wantText: "google_cloud_run_v2_service.app",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := formatRuntimePermissions(tt.perms, tt.prov)
			if tt.wantEmpty && got != "" {
				t.Errorf("formatRuntimePermissions() = %q, want empty", got)
			}
			if !tt.wantEmpty && !strings.Contains(got, tt.wantText) {
				t.Errorf("formatRuntimePermissions() = %q, want substring %q", got, tt.wantText)
			}
		})
	}
}
