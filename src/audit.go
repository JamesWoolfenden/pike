package pike

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

func collectAuditFindings(dir string) ([]Finding, error) {
	files, err := GetTF(dir)
	if err != nil {
		return nil, err
	}

	var findings []Finding
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
		}
	}
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
