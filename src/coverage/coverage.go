package coverage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	pike "github.com/jameswoolfenden/pike/src"
)

type members struct {
	DataSources []string `json:"datasources"`
	Resources   []string `json:"resources"`
}

//goland:noinspection GoUnusedFunction
func coverageAWS() error {
	data := importMembers("../parse/aws-members.json")
	missing := members{}
	target := ""

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

	Prepend := "# todo aws \n\n"

	Prepend += fmt.Sprintf("Resource percentage coverage   %3.2f \n", percent(missing.Resources, data.Resources))
	Prepend += fmt.Sprintf("Datasource percentage coverage %3.2f \n\n", percent(missing.DataSources, data.DataSources))

	target = Prepend + target
	err := os.WriteFile("aws.md", []byte(target), 0o700)
	if err != nil {
		return err
	}

	return nil
}

//goland:noinspection GoUnusedFunction
func coverageAzure() error {
	data := importMembers("../parse/azurerm-members.json")
	missing := members{}
	target := ""

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

	Prepend := "# todo azure \n\n"

	Prepend += fmt.Sprintf("Resource percentage coverage   %3.2f \n", percent(missing.Resources, data.Resources))
	Prepend += fmt.Sprintf("Datasource percentage coverage %3.2f \n\n", percent(missing.DataSources, data.DataSources))

	target = Prepend + target
	err := os.WriteFile("azure.md", []byte(target), 0o700)
	if err != nil {
		return err
	}

	return nil
}

//goland:noinspection GoUnusedFunction
func coverageGcp() error {
	data := importMembers("../parse/google-members.json")
	missing := members{}
	target := ""

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

	Prepend := "# todo google \n\n"

	Prepend += fmt.Sprintf("Resource percentage coverage   %3.2f \n", percent(missing.Resources, data.Resources))
	Prepend += fmt.Sprintf("Datasource percentage coverage %3.2f \n\n", percent(missing.DataSources, data.DataSources))

	target = Prepend + target
	err := os.WriteFile("google.md", []byte(target), 0o700)
	if err != nil {
		return err
	}

	return nil
}

func importMembers(targetMembers string) members {
	fileName, _ := filepath.Abs(targetMembers)
	file, _ := os.ReadFile(fileName)
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
