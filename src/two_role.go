package pike

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"strings"
	"text/template"
)

//go:embed terraform.two-role.aws.template
var twoRoleAWSTemplate []byte

//go:embed terraform.two-role.gcp.template
var twoRoleGCPTemplate []byte

//go:embed terraform.two-role.azure.template
var twoRoleAzureTemplate []byte

type awsTwoRoleDetails struct {
	ApplyName   string
	ApplyPolicy string
	PlanName    string
	PlanPolicy  string
}

type gcpTwoRoleDetails struct {
	ApplyName        string
	ApplyRoleID      string
	ApplyPermissions string
	PlanName         string
	PlanRoleID       string
	PlanPermissions  string
	Project          string
}

type azureTwoRoleDetails struct {
	ApplyName        string
	ApplyPermissions string
	PlanName         string
	PlanPermissions  string
}

// WriteTwoRoleOutput renders the full two-role Terraform setup (roles +
// policies + attachments) from the raw permission bag for write mode.
func WriteTwoRoleOutput(bag Sorted, policyName string, dirName string) (string, error) {
	var s strings.Builder

	applyName, planName := twoRoleNames(policyName)

	if len(bag.AWS) > 0 {
		chunk, err := renderAWSTwoRole(bag, applyName, planName)
		if err != nil {
			return "", err
		}
		s.WriteString(chunk)
	}

	if len(bag.GCP) > 0 {
		chunk, err := renderGCPTwoRole(bag, applyName, planName, dirName)
		if err != nil {
			return "", err
		}
		s.WriteString(chunk)
	}

	if len(bag.AZURE) > 0 {
		chunk, err := renderAzureTwoRole(bag, applyName, planName)
		if err != nil {
			return "", err
		}
		s.WriteString(chunk)
	}

	return s.String(), nil
}

func twoRoleNames(policyName string) (apply, plan string) {
	base := policyName
	if base == "" {
		base = defaultPolicyName
	}
	return base + "_apply", base + "_plan"
}

func renderAWSTwoRole(bag Sorted, applyName, planName string) (string, error) {
	applyPolicy, err := awsPolicyJSON(Unique(bag.AWS))
	if err != nil {
		return "", err
	}
	planPolicy, err := awsPolicyJSON(Unique(bag.PlanAWS))
	if err != nil {
		return "", err
	}
	details := awsTwoRoleDetails{
		ApplyName:   applyName,
		ApplyPolicy: applyPolicy,
		PlanName:    planName,
		PlanPolicy:  planPolicy,
	}
	return execTwoRoleTemplate(twoRoleAWSTemplate, details)
}

func renderGCPTwoRole(bag Sorted, applyName, planName, dirName string) (string, error) {
	project, err := getCurrentProject(dirName)
	if err != nil {
		project = defaultProject
	}
	details := gcpTwoRoleDetails{
		ApplyName:        applyName,
		ApplyRoleID:      applyName,
		ApplyPermissions: strings.Join(Unique(bag.GCP), "\",\n    \""),
		PlanName:         planName,
		PlanRoleID:       planName,
		PlanPermissions:  strings.Join(Unique(bag.PlanGCP), "\",\n    \""),
		Project:          project,
	}
	return execTwoRoleTemplate(twoRoleGCPTemplate, details)
}

func renderAzureTwoRole(bag Sorted, applyName, planName string) (string, error) {
	details := azureTwoRoleDetails{
		ApplyName:        applyName,
		ApplyPermissions: strings.Join(Unique(bag.AZURE), "\",\n    \""),
		PlanName:         planName,
		PlanPermissions:  strings.Join(Unique(bag.PlanAZURE), "\",\n    \""),
	}
	return execTwoRoleTemplate(twoRoleAzureTemplate, details)
}

func awsPolicyJSON(permissions []string) (string, error) {
	p, err := NewAWSPolicy(permissions, false)
	if err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func execTwoRoleTemplate(tmplBytes []byte, data any) (string, error) {
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
