package pike

import (
	"reflect"
	"sort"
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
		"AWS001 aws_iam_policy.wildcard_jsonencode",
		"AWS001 data.aws_iam_policy_document.doc",
		"AWS002 aws_iam_policy.from_file",
		"AWS002 aws_iam_role_policy.service_wildcard",
		"AWS003 aws_iam_role.with_inline",
		"AWS003 aws_iam_user_policy.escalation_heredoc",
		"AWS004 aws_iam_user_policy.escalation_heredoc",
		"AWS005 aws_iam_group_policy.not_action",
		"AWS006 aws_iam_role_policy_attachment.admin",
		"AWS006 aws_iam_user_policy_attachment.s3full",
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
