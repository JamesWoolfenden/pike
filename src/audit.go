package pike

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// Severity ranks audit findings.
type Severity int

const (
	SevInfo Severity = iota
	SevLow
	SevMedium
	SevHigh
	SevCritical
)

var severityNames = map[Severity]string{
	SevInfo:     "INFO",
	SevLow:      "LOW",
	SevMedium:   "MEDIUM",
	SevHigh:     "HIGH",
	SevCritical: "CRITICAL",
}

func (s Severity) String() string {
	if n, ok := severityNames[s]; ok {
		return n
	}
	return "UNKNOWN"
}

// MarshalJSON renders severity as its name.
func (s Severity) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// ParseSeverity maps a name (case-insensitive) to a Severity; unknown names
// default to SevLow so a typo doesn't silently suppress everything.
func ParseSeverity(name string) Severity {
	for sev, n := range severityNames {
		if strings.EqualFold(n, name) {
			return sev
		}
	}
	return SevLow
}

// Finding is a single audit result tied to a source location.
type Finding struct {
	RuleID   string   `json:"rule_id"`
	Severity Severity `json:"severity"`
	Resource string   `json:"resource"`
	File     string   `json:"file"`
	Line     int      `json:"line"`
	Message  string   `json:"message"`
	Detail   string   `json:"detail,omitempty"`
}

// auditHandler inspects one HCL block and returns findings. baseDir is the
// directory containing the .tf file, used to resolve file() references.
type auditHandler func(block *hclsyntax.Block, baseDir string) []Finding

// auditHandlers is populated by init() in audit_{aws,gcp,azure}.go. Keys are
// "<block.Type>.<block.Labels[0]>", e.g. "resource.aws_iam_policy" or
// "data.aws_iam_policy_document".
var auditHandlers = map[string]auditHandler{}

// Audit walks dir for .tf files, dispatches IAM-relevant blocks to provider
// auditors, prints findings at or above minSeverity, and returns a non-nil
// error when any survive (so the CLI exits non-zero).
func Audit(dir string, output string, minSeverity Severity) error {
	findings, err := collectAuditFindings(dir)
	if err != nil {
		return err
	}

	filtered := findings[:0]
	for _, f := range findings {
		if f.Severity >= minSeverity {
			filtered = append(filtered, f)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		if filtered[i].File != filtered[j].File {
			return filtered[i].File < filtered[j].File
		}
		if filtered[i].Line != filtered[j].Line {
			return filtered[i].Line < filtered[j].Line
		}
		return filtered[i].RuleID < filtered[j].RuleID
	})

	renderFindings(os.Stdout, filtered, output)

	if len(filtered) > 0 {
		return &auditFindingsError{count: len(filtered)}
	}
	return nil
}

// CollectAuditFindings runs the audit rules over the Terraform in dir and
// returns every finding, unfiltered and unsorted. It is the entry point for
// callers that embed pike as a library rather than invoking the CLI.
//
// holden consumes this alongside MakePermissionBag: it needs pike's IAM
// judgements — in particular whether a Condition genuinely scopes an
// escalation-class action, which its own policy language cannot express — as
// data it can attach to its graph. Without an exported collector the only way
// to get at them is to capture Audit's stdout.
//
// Severity filtering, ordering and rendering belong to Audit, so an embedder
// can apply its own.
func CollectAuditFindings(dir string) ([]Finding, error) {
	return collectAuditFindings(dir)
}

func collectAuditFindings(dir string) ([]Finding, error) {
	files, err := GetTF(dir)
	if err != nil {
		return nil, err
	}

	var findings []Finding
	var awsAttachments []referenceRef
	var gcpPolicyDataRefs []referenceRef

	for _, file := range files {
		body, err := GetResourceBlocks(file)
		if err != nil {
			continue
		}
		baseDir := filepath.Dir(file)
		rel := file
		if r, rerr := filepath.Rel(dir, file); rerr == nil {
			rel = r
		}
		for _, block := range body.Blocks {
			if block.Type != "resource" && block.Type != "data" {
				continue
			}
			if len(block.Labels) < 2 {
				continue
			}
			key := block.Type + "." + block.Labels[0]
			h, ok := auditHandlers[key]
			if !ok {
				continue
			}
			for _, f := range h(block, baseDir) {
				f.File = rel
				findings = append(findings, f)
			}

			if block.Type != "resource" {
				continue
			}
			switch {
			case awsPolicyAttachmentTypes[block.Labels[0]]:
				awsAttachments = append(awsAttachments, referenceRef{block: block, file: rel})
			case gcpPolicyDataTypes[block.Labels[0]]:
				gcpPolicyDataRefs = append(gcpPolicyDataRefs, referenceRef{block: block, file: rel})
			}
		}
	}

	// Second pass: re-attribute findings on a resource that only carries IAM
	// content indirectly, to whatever references it in this same scan. See
	// the "reference graph" comment above.
	findings = append(findings, resolveReferenceGraph(findings, awsAttachments, "policy_arn", awsAttachmentLabels)...)
	findings = append(findings, resolveReferenceGraph(findings, gcpPolicyDataRefs, "policy_data", gcpPolicyDataLabels)...)

	return findings, nil
}

func renderFindings(w io.Writer, findings []Finding, format string) {
	if strings.EqualFold(format, "json") {
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		_ = enc.Encode(findings)
		return
	}

	if len(findings) == 0 {
		fmt.Fprintln(w, "no findings")
		return
	}

	var lastFile string
	for _, f := range findings {
		if f.File != lastFile {
			if lastFile != "" {
				fmt.Fprintln(w)
			}
			fmt.Fprintf(w, "%s\n", f.File)
			lastFile = f.File
		}
		fmt.Fprintf(w, "  %-4d %-8s %-8s %s — %s\n", f.Line, f.Severity, f.RuleID, f.Resource, f.Message)
		if f.Detail != "" {
			fmt.Fprintf(w, "       %s\n", f.Detail)
		}
	}
}

// ---- reference graph -------------------------------------------------------
//
// Some IAM content lives on one block (a customer-managed aws_iam_policy, a
// GCP data.google_iam_policy) while another block merely references it (an
// attachment's policy_arn, an authoritative policy's policy_data) — so that
// content's findings get generated in isolation from whoever actually ends
// up holding the grant. collectAuditFindings' second pass resolves those
// references back to the resource they name, and re-attributes the
// referenced resource's already-collected findings (pass one) to whatever
// references it, so e.g. a wildcard-action policy finding also shows up as
// "role X ends up with ... AWS001" on the attachment, not just on the policy
// in isolation. Resolution is scoped to resources parsed in the same
// collectAuditFindings call (one scan directory) — a policy defined in a
// module with the attachment at the root call site, or vice versa, won't be
// traced, since each is a separate GetResourceBlocks walk.

// referenceRef is a block collectAuditFindings' first pass collected
// because it might reference another resource's IAM content, held for the
// second pass's reference-graph resolution.
type referenceRef struct {
	block *hclsyntax.Block
	file  string
}

// resourceRefPattern matches the "${<traversal>}" string extractStringValue
// produces for a bare attribute reference like aws_iam_policy.wide.arn —
// the shape attrString already returns for a non-literal policy_arn/
// policy_data/role/user/group attribute.
var resourceRefPattern = regexp.MustCompile(`^\$\{([^}]+)\}$`)

// resolveResourceRef extracts the "<type>.<name>" (or "data.<type>.<name>")
// resource address a reference string names, e.g. "${aws_iam_policy.wide.arn}"
// -> "aws_iam_policy.wide". Returns ok=false for anything else (a literal
// ARN, a role name, an unresolvable expression) — this never guesses.
func resolveResourceRef(s string) (string, bool) {
	m := resourceRefPattern.FindStringSubmatch(s)
	if m == nil {
		return "", false
	}

	parts := strings.Split(m[1], ".")

	if len(parts) >= 3 && parts[0] == "data" {
		return strings.Join(parts[:3], "."), true
	}

	if len(parts) >= 2 && parts[0] != "data" {
		return strings.Join(parts[:2], "."), true
	}

	return "", false
}

// principalIdentifier returns the resource address a reference names (e.g.
// "aws_iam_role.ci" from role = aws_iam_role.ci.name), or the raw value
// otherwise — a literal name, or an unresolvable reference shown verbatim
// rather than silently dropped.
func principalIdentifier(raw string) string {
	if addr, ok := resolveResourceRef(raw); ok {
		return addr
	}
	return raw
}

// resolveReferenceGraph re-attributes findings already collected on a
// resource (indexed by its own address) to another resource that merely
// references it via refAttr — e.g. an AWS attachment's policy_arn, or a GCP
// authoritative policy's policy_data. labelsFor returns the label(s) (e.g.
// "role aws_iam_role.ci", or a fixed "authoritative policy") a derived
// finding should be attributed to for a given referencing block; an empty
// result skips that block.
func resolveReferenceGraph(findings []Finding, refs []referenceRef, refAttr string, labelsFor func(block *hclsyntax.Block) []string) []Finding {
	byAddr := map[string][]Finding{}
	for _, f := range findings {
		byAddr[f.Resource] = append(byAddr[f.Resource], f)
	}

	var derived []Finding

	for _, r := range refs {
		targetAddr, ok := resolveResourceRef(attrString(r.block, refAttr))
		if !ok {
			continue
		}

		matches := byAddr[targetAddr]
		if len(matches) == 0 {
			continue
		}

		for _, label := range labelsFor(r.block) {
			for _, f := range matches {
				derived = append(derived, Finding{
					RuleID:   f.RuleID,
					Severity: f.Severity,
					Resource: blockAddr(r.block),
					File:     r.file,
					Line:     r.block.DefRange().Start.Line,
					Message:  fmt.Sprintf("%s ends up with (via %s): %s", label, targetAddr, f.Message),
					Detail:   f.Detail,
				})
			}
		}
	}

	return derived
}

// blockAddr returns the terraform-style address of a block, e.g.
// "aws_iam_policy.app" or "data.aws_iam_policy_document.app".
func blockAddr(block *hclsyntax.Block) string {
	addr := block.Labels[0] + "." + block.Labels[1]
	if block.Type == "data" {
		return "data." + addr
	}
	return addr
}

// finding is a constructor that pulls file/line from the block's def range.
func finding(block *hclsyntax.Block, ruleID string, sev Severity, msg, detail string) Finding {
	return Finding{
		RuleID:   ruleID,
		Severity: sev,
		Resource: blockAddr(block),
		Line:     block.DefRange().Start.Line,
		Message:  msg,
		Detail:   detail,
	}
}

// attrString returns the literal string value of a named attribute on block,
// or "" if absent or not a literal.
func attrString(block *hclsyntax.Block, name string) string {
	if a, ok := block.Body.Attributes[name]; ok {
		return extractStringValue(a.Expr)
	}
	return ""
}

// attrStringList returns the literal string elements of a list-typed attribute.
func attrStringList(block *hclsyntax.Block, name string) []string {
	if a, ok := block.Body.Attributes[name]; ok {
		return extractStringList(a.Expr)
	}
	return nil
}
