package pike

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// ParsedDir is a set of Terraform files that have already been read and
// parsed. Embedders that hold the file contents (or want to run both the
// permission scan and the audit over one directory) build one of these once
// and reuse it, so each file is read and parsed a single time.
type ParsedDir struct {
	// Dir is the directory the files belong to. It is used to find
	// .terraform/modules, resolve local module sources and make audit paths
	// relative.
	Dir string
	// Files lists the file paths in a stable order.
	Files []string
	// Bodies holds the parsed body of every file that parsed successfully.
	Bodies map[string]*hclsyntax.Body
	// Errors holds the read or parse error of every file that failed.
	Errors map[string]error
}

// parseSource parses HCL source that has already been read.
func parseSource(src []byte, name string) (*hclsyntax.Body, error) {
	parsedFile, diags := hclparse.NewParser().ParseHCL(src, name)
	if diags != nil {
		return nil, diags
	}

	return parsedFile.Body.(*hclsyntax.Body), nil
}

// ParseSources parses file contents the caller has already read. sources maps
// a file path to its bytes; nothing is read from disk. dir is the directory
// the paths are relative to.
func ParseSources(dir string, sources map[string][]byte) *ParsedDir {
	parsed := &ParsedDir{
		Dir:    dir,
		Files:  make([]string, 0, len(sources)),
		Bodies: make(map[string]*hclsyntax.Body, len(sources)),
		Errors: map[string]error{},
	}

	for name := range sources {
		parsed.Files = append(parsed.Files, name)
	}

	sort.Strings(parsed.Files)

	for _, name := range parsed.Files {
		body, err := parseSource(sources[name], name)
		if err != nil {
			parsed.Errors[name] = err
			continue
		}

		parsed.Bodies[name] = body
	}

	return parsed
}

// ParseDir finds the .tf files in dir (including downloaded modules) and
// reads and parses each of them once.
func ParseDir(dir string) (*ParsedDir, error) {
	files, err := GetTF(dir)
	if err != nil {
		return nil, err
	}

	return parseFiles(dir, files), nil
}

func parseFiles(dir string, files []string) *ParsedDir {
	parsed := &ParsedDir{
		Dir:    dir,
		Files:  files,
		Bodies: make(map[string]*hclsyntax.Body, len(files)),
		Errors: map[string]error{},
	}

	for _, file := range files {
		body, err := GetResourceBlocks(file)
		if err != nil {
			parsed.Errors[file] = err
			continue
		}

		parsed.Bodies[file] = body
	}

	return parsed
}

// PermissionBag returns the IAM permissions the parsed Terraform needs. It is
// the equivalent of MakePermissionBag without any file I/O for the parsed
// files. provider may be empty for all providers.
func (p *ParsedDir) PermissionBag(provider string, suppressDeprecated bool) (Sorted, error) {
	if len(p.Files) == 0 {
		return Sorted{}, &emptyIACError{}
	}

	moduleJSON, moduleErr := GetModuleJson(filepath.Join(p.Dir, ".terraform", "modules"))

	var resources []ResourceV2
	var failedFiles []string
	var criticalErrors []error

	var allIAMBindings []IAMBinding
	awsDefaultTags := false

	for _, tfFile := range p.Files {
		if err, failed := p.Errors[tfFile]; failed {
			failedFiles = append(failedFiles, tfFile)
			criticalErrors = append(criticalErrors, fmt.Errorf("failed to parse %s: %w", tfFile, err))

			continue
		}

		body := p.Bodies[tfFile]

		resource, err := extractResources(body, p.Dir, moduleJSON, moduleErr)
		if err != nil {
			failedFiles = append(failedFiles, tfFile)
			criticalErrors = append(criticalErrors, fmt.Errorf("failed to parse %s: %w", tfFile, err))

			continue
		}

		resources = append(resources, resource...)

		if body != nil {
			allIAMBindings = append(allIAMBindings, ExtractIAMBindings(body)...)
			awsDefaultTags = awsDefaultTags || HasAWSDefaultTags(body)
		}
	}

	return finishPermissionBag(p.Files, resources, allIAMBindings, awsDefaultTags, failedFiles, criticalErrors, provider, suppressDeprecated)
}

// AuditFindings runs the audit rules over the parsed Terraform. It is the
// equivalent of CollectAuditFindings without re-reading the files. Files that
// failed to parse are skipped.
func (p *ParsedDir) AuditFindings() ([]Finding, error) {
	return auditParsed(p)
}

// Analyse runs the permission scan and the audit over one parse of the
// directory, so callers that need both do not walk and parse it twice.
func (p *ParsedDir) Analyse(provider string, suppressDeprecated bool) (Sorted, []Finding, error) {
	bag, err := p.PermissionBag(provider, suppressDeprecated)
	if err != nil {
		return Sorted{}, nil, err
	}

	findings, err := p.AuditFindings()
	if err != nil {
		return Sorted{}, nil, err
	}

	return bag, findings, nil
}
