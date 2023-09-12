package coverage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	pike "github.com/jameswoolfenden/pike/src"
)

func coverage() error {
	type members struct {
		DataSources []string `json:"datasources"`
		Resources   []string `json:"resources"`
	}

	fileName, _ := filepath.Abs("../parse/aws-members.json")
	file, _ := os.ReadFile(fileName)
	data := members{}
	missing := members{}

	target := ""

	_ = json.Unmarshal(file, &data)

	for _, myData := range data.Resources {
		if temp := pike.AwsLookup(myData); temp == nil {
			missing.Resources = append(missing.Resources, myData)
			target += "./resource.ps1 " + myData + "\n"
		}
	}

	for _, myData := range data.DataSources {
		if temp := pike.AwsDataLoookup(myData); temp == nil {
			missing.DataSources = append(missing.DataSources, myData)
			target += "./resource.ps1 " + myData + " -type data\n"
		}
	}

	Prepend := "# todo aws \n\n"

	Prepend += fmt.Sprintf("Resource percentage coverage   %3.2f \n", percent(missing.Resources, data.Resources))
	Prepend += fmt.Sprintf("Datasource percentage coverage %3.2f \n\n", percent(missing.DataSources, data.DataSources))

	target = Prepend + target
	err := os.WriteFile("aws.md", []byte(target), 0700)
	if err != nil {
		return err
	}

	return nil
}

func percent(missing []string, data []string) float64 {
	var source float64
	var target float64

	source = float64(len(missing))

	target = float64(len(data))

	return 100 - (source/target)*100
}
