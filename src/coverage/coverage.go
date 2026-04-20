package coverage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	pike "github.com/jameswoolfenden/pike/src"
)

// members mirrors the shape written by parse/parse.go's provider struct.
// The JSON tag for DataSources MUST stay camelCase "dataSources" — that's
// what Parse serialises. An earlier version of this struct had the tag as
// lowercase "datasources", which silently unmarshalled to an empty slice
// and made every coverage run under-report datasource coverage by 100%.
//
// DeprecatedResources and DeprecatedData were added in April 2026 when
// pike parse started capturing Block.Deprecated off the provider schema.
// Older members JSON files won't have these keys — json.Unmarshal leaves
// them nil and the deprecated section of the report is omitted cleanly.
type members struct {
	DataSources         []string          `json:"dataSources"`
	Resources           []string          `json:"resources"`
	DeprecatedResources map[string]string `json:"deprecatedResources,omitempty"`
	DeprecatedData      map[string]string `json:"deprecatedData,omitempty"`
}

//goland:noinspection GoUnusedFunction
func coverageAWS() error {
	absolute, err := filepath.Abs("../../parse/aws-members.json")

	if err != nil {
		return &fileWriteError{err}
	}

	data := importMembers(absolute)

	missing := members{}
	target := "```shell\n"

	for _, myData := range data.Resources {
		if temp := pike.AwsLookup(myData); temp == nil {
			if strings.Contains(myData, "aws") {
				missing.Resources = append(missing.Resources, myData)
				target += "./resource.ps1 " + myData + "\n"
			}
		}
	}

	for _, myData := range data.DataSources {
		if temp := pike.AwsDataLookup(myData); temp == nil {
			if strings.Contains(myData, "aws") {
				missing.DataSources = append(missing.DataSources, myData)
				target += "./resource.ps1 " + myData + " -type data\n"
			}
		}
	}
	target += "```\n"

	Prepend := resourceTable(missing, data, "AWS") + deprecatedSection(data)

	target = Prepend + target
	err = os.WriteFile("aws.md", []byte(target), 0o600)
	if err != nil {
		return &fileWriteError{err}
	}

	return nil
}

type fileWriteError struct {
	err error
}

func (e *fileWriteError) Error() string {
	return e.err.Error()
}

//goland:noinspection GoUnusedFunction
func coverageAzure() error {

	absolute, err := filepath.Abs("../../parse/azurerm-members.json")

	if err != nil {
		return &fileWriteError{err}
	}

	data := importMembers(absolute)

	missing := members{}
	target := "```shell\n"

	for _, myData := range data.Resources {
		if temp := pike.AzureLookup(myData); temp == nil {
			missing.Resources = append(missing.Resources, myData)
			target += "./resource.ps1 " + myData + "\n"
		}
	}

	for _, myData := range data.DataSources {
		if temp := pike.AzureDataLookup(myData); temp == nil {
			missing.DataSources = append(missing.DataSources, myData)
			target += "./resource.ps1 " + myData + " -type data\n"
		}
	}
	target += "```\n"

	Prepend := resourceTable(missing, data, "Azure") + deprecatedSection(data)
	target = Prepend + target
	err = os.WriteFile("azure.md", []byte(target), 0o600)

	if err != nil {
		return &fileWriteError{err}
	}

	return nil
}

//goland:noinspection GoUnusedFunction
func coverageGcp() error {
	absolute, err := filepath.Abs("../../parse/google-members.json")

	if err != nil {
		return &fileWriteError{err}
	}

	data := importMembers(absolute)
	missing := members{}
	target := "```shell\n"

	for _, myData := range data.Resources {
		if temp := pike.GCPLookup(myData); temp == nil {
			missing.Resources = append(missing.Resources, myData)
			target += "./resource.ps1 " + myData + "\n"
		}
	}

	for _, myData := range data.DataSources {
		if temp := pike.GCPDataLookup(myData); temp == nil {
			missing.DataSources = append(missing.DataSources, myData)
			target += "./resource.ps1 " + myData + " -type data\n"
		}
	}
	target += "```\n"

	Prepend := resourceTable(missing, data, "Google") + deprecatedSection(data)

	target = Prepend + target
	err = os.WriteFile("google.md", []byte(target), 0o600)

	if err != nil {

		return &fileWriteError{err}
	}

	return nil
}

func resourceTable(missing members, data members, cloud string) string {
	Prepend := fmt.Sprintf("# %s Resource Status\n\n", cloud)
	Prepend += fmt.Sprintf("| Terraform  | Coverage %% | Resources | Total Resources |\n")
	Prepend += "|------------|------------|-----------|-----------------|\n"
	Prepend += fmt.Sprintf("| Resources  | %3.2f      | %5d       | %5d            |\n",
		percent(missing.Resources, data.Resources),
		len(data.Resources)-len(missing.Resources), len(data.Resources))
	Prepend += fmt.Sprintf("| Datasource | %3.2f      | %5d       | %5d             |\n\n",
		percent(missing.DataSources, data.DataSources),
		len(data.DataSources)-len(missing.DataSources), len(data.DataSources))
	return Prepend
}

// deprecatedSection renders a markdown block listing the resources and
// datasources that the provider's latest schema marks as deprecated. It
// returns the empty string when nothing is deprecated so older members
// files (pre-April-2026 format, no deprecated fields) produce the same
// output they always did.
//
// Output shape:
//
//	## Deprecated
//
//	{N resources, M datasources} flagged as deprecated in the latest ...
//
//	### Deprecated Resources
//	| Resource | Note |
//	...
//
//	### Deprecated Data Sources
//	| Data Source | Note |
//	...
//
// Tables are emitted only when non-empty so single-sided deprecations
// (e.g. google resource removed but data form still live) don't leave
// an empty header.
func deprecatedSection(data members) string {
	if len(data.DeprecatedResources) == 0 && len(data.DeprecatedData) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("## Deprecated\n\n")
	b.WriteString(fmt.Sprintf(
		"%d resources and %d datasources are flagged as deprecated in the latest provider schema. "+
			"Users pinned to an older provider major may already be affected when they upgrade.\n\n",
		len(data.DeprecatedResources), len(data.DeprecatedData),
	))

	if len(data.DeprecatedResources) > 0 {
		b.WriteString("### Deprecated Resources\n\n")
		b.WriteString(deprecationTable("Resource", data.DeprecatedResources))
	}
	if len(data.DeprecatedData) > 0 {
		b.WriteString("### Deprecated Data Sources\n\n")
		b.WriteString(deprecationTable("Data Source", data.DeprecatedData))
	}

	return b.String()
}

// deprecationTable renders a two-column markdown table with stable
// (sorted) row order. An empty Note column just reads "—" so the table
// doesn't collapse visually when a provider set Deprecated=true without
// supplying a message.
func deprecationTable(header string, entries map[string]string) string {
	names := make([]string, 0, len(entries))
	for n := range entries {
		names = append(names, n)
	}
	sort.Strings(names)

	var b strings.Builder
	fmt.Fprintf(&b, "| %s | Note |\n", header)
	b.WriteString("|---|---|\n")
	for _, n := range names {
		note := entries[n]
		if note == "" {
			note = "—"
		}
		// Escape pipes in the note so they don't break the table cell.
		note = strings.ReplaceAll(note, "|", "\\|")
		// Collapse newlines — some provider descriptions are multi-line.
		note = strings.ReplaceAll(note, "\n", " ")
		fmt.Fprintf(&b, "| %s | %s |\n", n, note)
	}
	b.WriteString("\n")
	return b.String()
}

func importMembers(targetMembers string) members {
	fileName, _ := filepath.Abs(targetMembers)
	file, _ := os.ReadFile(fileName) // #nosec G304 -- Test helper reading data files
	data := members{}

	_ = json.Unmarshal(file, &data)

	return data
}

func percent(missing []string, data []string) float64 {
	var source float64

	var target float64

	source = float64(len(missing))

	target = float64(len(data))

	return 100 - (source/target)*100
}
