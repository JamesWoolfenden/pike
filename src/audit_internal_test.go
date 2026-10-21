package pike

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func collectRuleHits(t *testing.T, dir string) []string {
	t.Helper()
	findings, err := collectAuditFindings(dir)
	if err != nil {
		t.Fatalf("collectAuditFindings(%q): %v", dir, err)
	}
	var hits []string
	for _, f := range findings {
		hits = append(hits, f.RuleID+" "+f.Resource)
	}
	sort.Strings(hits)
	return hits
}

func TestAudit_GCP(t *testing.T) {
	t.Parallel()
	got := collectRuleHits(t, "testdata/audit/gcp")
	want := []string{
		"GCP001 google_project_iam_member.owner",
		"GCP001 data.google_iam_policy.x",
		"GCP001 google_project_iam_policy.authoritative",
		"GCP002 google_project_iam_binding.editor_public",
		"GCP003 google_project_iam_policy.authoritative",
		"GCP004 google_folder_iam_member.impersonate",
		"GCP005 google_project_iam_custom_role.esc",
		"GCP006 google_project_iam_binding.editor_public",
	}
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got  %v\nwant %v", got, want)
	}
}

func TestAudit_Azure(t *testing.T) {
	t.Parallel()
	got := collectRuleHits(t, "testdata/audit/azure")
	want := []string{
		"AZURE001 azurerm_role_assignment.owner",
		"AZURE002 azurerm_role_assignment.contributor",
		"AZURE002 azurerm_role_assignment.contributor_lowercase",
		"AZURE003 azurerm_role_definition.wide",
		"AZURE004 azurerm_role_definition.wide",
	}
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got  %v\nwant %v", got, want)
	}
}

func TestAudit_AWS(t *testing.T) {
	t.Parallel()
	got := collectRuleHits(t, "testdata/audit/aws")
	want := []string{
		"AWS000 aws_iam_policy.interpolated",
		"AWS001 aws_iam_policy.wide",
		"AWS001 aws_iam_role_policy_attachment.wide_to_ci",
		"AWS001 aws_iam_policy.wildcard_jsonencode",
		"AWS001 data.aws_iam_policy_document.doc",
		"AWS002 aws_iam_policy.from_file",
		"AWS002 aws_iam_role_policy.iam_wildcard",
		"AWS002 aws_iam_role_policy.service_wildcard",
		"AWS003 aws_iam_role.with_inline",
		"AWS003 aws_iam_role_policy.region_condition_escalation",
		"AWS003 aws_iam_user_policy.attach_user_policy",
		"AWS003 aws_iam_user_policy.escalation_heredoc",
		"AWS004 aws_iam_role_policy.iam_wildcard",
		"AWS004 aws_iam_role_policy.region_condition_escalation",
		"AWS004 aws_iam_user_policy.attach_user_policy",
		"AWS004 aws_iam_user_policy.escalation_heredoc",
		"AWS005 aws_iam_group_policy.not_action",
		"AWS006 aws_iam_role_policy_attachment.admin",
		"AWS006 aws_iam_user_policy_attachment.s3full",
		"AWS009 aws_iam_role.with_inline",
		"AWS007 aws_iam_role.wildcard_trust",
		"AWS008 aws_iam_role.root_wildcard_trust",
		"AWS010 aws_iam_role.oidc_trust_no_sub",
	}
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got  %v\nwant %v", got, want)
	}
}

func TestAudit_Clean(t *testing.T) {
	t.Parallel()
	got := collectRuleHits(t, "testdata/audit/clean")
	if len(got) != 0 {
		t.Errorf("expected no findings, got %v", got)
	}
}

func TestAudit_AWSConditionSuppressesEscalation(t *testing.T) {
	t.Parallel()
	// data.aws_iam_policy_document.doc statement[1] has iam:PutRolePolicy with
	// a condition block; AWS004 must not fire for it.
	findings, err := collectAuditFindings("testdata/audit/aws")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range findings {
		if f.RuleID == "AWS004" && f.Resource == "data.aws_iam_policy_document.doc" {
			t.Errorf("AWS004 fired on a statement with a condition: %+v", f)
		}
	}
}

func TestAudit_AWSNonScopingConditionDoesNotSuppressEscalation(t *testing.T) {
	t.Parallel()
	// aws_iam_role_policy.region_condition_escalation grants iam:PassRole
	// guarded only by a Condition on aws:RequestedRegion, which has no
	// bearing on escalation risk. AWS004 must still fire.
	findings, err := collectAuditFindings("testdata/audit/aws")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range findings {
		if f.RuleID == "AWS004" && f.Resource == "aws_iam_role_policy.region_condition_escalation" {
			return
		}
	}
	t.Errorf("AWS004 did not fire for iam:PassRole guarded only by a non-scoping condition (aws:RequestedRegion)")
}

func TestAudit_AWSTrustPolicyConditionsSuppressFindings(t *testing.T) {
	t.Parallel()
	// wildcard_trust_scoped (Principal AWS:"*" + aws:PrincipalOrgID) and
	// oidc_trust_scoped (Federated + token.actions.githubusercontent.com:sub)
	// must not trigger AWS007/AWS008/AWS010 — the whole point of the
	// condition exceptions is that these specific, deliberate patterns don't
	// false-positive.
	findings, err := collectAuditFindings("testdata/audit/aws")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range findings {
		if f.Resource == "aws_iam_role.wildcard_trust_scoped" || f.Resource == "aws_iam_role.oidc_trust_scoped" {
			t.Errorf("expected no findings for %s, got %s", f.Resource, f.RuleID)
		}
	}
}

func TestAudit_AWSAttachmentGraph(t *testing.T) {
	t.Parallel()
	// aws_iam_policy.wide (Action "*") is attached to aws_iam_role.ci via
	// aws_iam_role_policy_attachment.wide_to_ci's policy_arn reference.
	// AWS001 must fire on the attachment too, attributed to the role, not
	// just on the policy in isolation.
	findings, err := collectAuditFindings("testdata/audit/aws")
	if err != nil {
		t.Fatal(err)
	}

	var derived *Finding
	for i := range findings {
		if findings[i].RuleID == "AWS001" && findings[i].Resource == "aws_iam_role_policy_attachment.wide_to_ci" {
			derived = &findings[i]
			break
		}
	}

	if derived == nil {
		t.Fatal("expected an AWS001 finding attributed to aws_iam_role_policy_attachment.wide_to_ci, found none")
	}

	if !strings.Contains(derived.Message, "role aws_iam_role.ci") {
		t.Errorf("expected message to attribute the finding to role aws_iam_role.ci, got %q", derived.Message)
	}

	if !strings.Contains(derived.Message, "aws_iam_policy.wide") {
		t.Errorf("expected message to name the source policy aws_iam_policy.wide, got %q", derived.Message)
	}
}

func TestAudit_GCPPolicyDataGraph(t *testing.T) {
	t.Parallel()
	// data.google_iam_policy.x (roles/owner) is applied by
	// google_project_iam_policy.authoritative via its policy_data
	// reference. GCP001 must fire on the resource too, not just on the
	// data source in isolation.
	findings, err := collectAuditFindings("testdata/audit/gcp")
	if err != nil {
		t.Fatal(err)
	}

	var derived *Finding
	for i := range findings {
		if findings[i].RuleID == "GCP001" && findings[i].Resource == "google_project_iam_policy.authoritative" {
			derived = &findings[i]
			break
		}
	}

	if derived == nil {
		t.Fatal("expected a GCP001 finding attributed to google_project_iam_policy.authoritative, found none")
	}

	if !strings.Contains(derived.Message, "data.google_iam_policy.x") {
		t.Errorf("expected message to name the source data source data.google_iam_policy.x, got %q", derived.Message)
	}
}

func TestResolveResourceRef(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		in       string
		wantAddr string
		wantOK   bool
	}{
		{"resource attribute reference", "${aws_iam_policy.wide.arn}", "aws_iam_policy.wide", true},
		{"data source attribute reference", "${data.aws_iam_policy_document.doc.json}", "data.aws_iam_policy_document.doc", true},
		{"gcp policy_data reference", "${data.google_iam_policy.x.policy_data}", "data.google_iam_policy.x", true},
		{"literal managed policy ARN", "arn:aws:iam::aws:policy/AdministratorAccess", "", false},
		{"literal customer ARN", "arn:aws:iam::123456789012:policy/custom", "", false},
		{"empty string", "", "", false},
		{"bare literal role name", "ci", "", false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			addr, ok := resolveResourceRef(tt.in)
			if ok != tt.wantOK || addr != tt.wantAddr {
				t.Errorf("resolveResourceRef(%q) = (%q, %v), want (%q, %v)", tt.in, addr, ok, tt.wantAddr, tt.wantOK)
			}
		})
	}
}

func TestAudit_AWSServiceWildcardCoveringEscalationAction(t *testing.T) {
	t.Parallel()
	// aws_iam_role_policy.iam_wildcard grants "iam:*", which trivially
	// includes every escalation-class iam: action; AWS004 must fire
	// alongside AWS002, not just AWS002.
	findings, err := collectAuditFindings("testdata/audit/aws")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range findings {
		if f.RuleID == "AWS004" && f.Resource == "aws_iam_role_policy.iam_wildcard" {
			return
		}
	}
	t.Errorf("AWS004 did not fire for iam:* (a service-wide wildcard covering escalation-class actions)")
}

func TestParseSeverity(t *testing.T) {
	t.Parallel()
	tests := map[string]Severity{
		"info":     SevInfo,
		"LOW":      SevLow,
		"Medium":   SevMedium,
		"high":     SevHigh,
		"CRITICAL": SevCritical,
		"bogus":    SevLow,
	}
	for in, want := range tests {
		if got := ParseSeverity(in); got != want {
			t.Errorf("ParseSeverity(%q) = %v, want %v", in, got, want)
		}
	}
}
