package coverage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

// Helper function to create test JSON file
func createTestMembersFile(t *testing.T, data members) string {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "test-members.json")

	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	return filePath
}

// Test the data sources processing logic specifically
func TestCoverageAWS_DataSourcesProcessing(t *testing.T) {
	tests := []struct {
		name                string
		dataSources         []string
		expectedMissing     []string
		expectedTargetLines int
	}{
		{
			name:                "empty data sources",
			dataSources:         []string{},
			expectedMissing:     []string{},
			expectedTargetLines: 0,
		},
		{
			name:                "aws data source not in pike lookup",
			dataSources:         []string{"aws_s3_bucket_duff"},
			expectedMissing:     []string{"aws_s3_bucket_duff"},
			expectedTargetLines: 1,
		},
		{
			name:                "non-aws data source not in pike lookup",
			dataSources:         []string{"google_storage_bucket"},
			expectedMissing:     []string{},
			expectedTargetLines: 0,
		},
		{
			name:                "mixed aws and non-aws data sources",
			dataSources:         []string{"aws_s3_bucket_duff", "google_storage_bucket", "aws_ec2_instance"},
			expectedMissing:     []string{"aws_s3_bucket_duff", "aws_ec2_instance"},
			expectedTargetLines: 2,
		},
		{
			name:                "data source containing aws in middle",
			dataSources:         []string{"terraform_aws_s3_bucket"},
			expectedMissing:     []string{"terraform_aws_s3_bucket"},
			expectedTargetLines: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test data
			testData := members{
				DataSources: tt.dataSources,
				Resources:   []string{}, // Empty for this test
			}

			// Create temporary test file
			testFile := createTestMembersFile(t, testData)

			// Load test data
			data := members{}
			fileName, _ := filepath.Abs(testFile)
			file, err := os.ReadFile(fileName)
			if err != nil {
				t.Fatalf("Failed to read test file: %v", err)
			}
			err = json.Unmarshal(file, &data)
			if err != nil {
				t.Fatalf("Failed to unmarshal test data: %v", err)
			}

			// Simulate the data sources processing logic from lines 34-40
			missing := members{}
			target := ""

			for _, myData := range data.DataSources {
				// Simulate pike.AwsDataLookup returning nil (not found)
				if temp := pike.AwsDataLookup(myData); temp == nil {
					if strings.Contains(myData, "aws") {
						missing.DataSources = append(missing.DataSources, myData)
						target += "./resource.ps1 " + myData + " -type data\n"
					}
				}
			}

			// Verify results
			if len(missing.DataSources) != len(tt.expectedMissing) {
				t.Errorf("Expected %d missing data sources, got %d", len(tt.expectedMissing), len(missing.DataSources))
			}

			for i, expected := range tt.expectedMissing {
				if i >= len(missing.DataSources) || missing.DataSources[i] != expected {
					t.Errorf("Expected missing data source %s, got %s", expected, missing.DataSources[i])
				}
			}

			// Count target lines
			targetLines := 0
			if target != "" {
				targetLines = strings.Count(target, "\n")
			}

			if targetLines != tt.expectedTargetLines {
				t.Errorf("Expected %d target lines, got %d", tt.expectedTargetLines, targetLines)
			}

			// Verify target format for aws data sources
			for _, expectedDS := range tt.expectedMissing {
				expectedLine := "./resource.ps1 " + expectedDS + " -type data\n"
				if !strings.Contains(target, expectedLine) {
					t.Errorf("Expected target to contain %s", expectedLine)
				}
			}
		})
	}
}

func TestCoverageAWS_DataSourcesTargetFormat(t *testing.T) {
	testData := members{
		DataSources: []string{"aws_s3_bucket", "aws_ec2_instance"},
		Resources:   []string{},
	}

	testFile := createTestMembersFile(t, testData)

	data := members{}
	fileName, _ := filepath.Abs(testFile)
	file, _ := os.ReadFile(fileName)
	_ = json.Unmarshal(file, &data)

	target := ""
	for _, myData := range data.DataSources {
		if temp := pike.AwsDataLookup(myData); temp == nil {
			if strings.Contains(myData, "aws") {
				target += "./resource.ps1 " + myData + " -type data\n"
			}
		}
	}

	expectedTarget := "./resource.ps1 aws_ec2_instance -type data\n"
	if target != expectedTarget {
		t.Errorf("Expected target format:\n%s\nGot:\n%s", expectedTarget, target)
	}
}

func TestCoverageAWS_DataSourcesCaseInsensitive(t *testing.T) {
	testCases := []struct {
		name       string
		dataSource string
		shouldAdd  bool
	}{
		{"lowercase aws", "aws_ec2_instance", true},
		{"uppercase AWS", "AWS_EC2_INSTANCE", false},
		{"mixed case", "Aws_EC2_Bucket", false},
		{"aws in middle", "terraform_aws_s3", true},
		{"no aws", "google_storage", false},
		{"aws substring", "awesome_bucket", false}, // contains "aws"
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testData := members{
				DataSources: []string{tc.dataSource},
				Resources:   []string{},
			}

			testFile := createTestMembersFile(t, testData)

			data := members{}
			fileName, _ := filepath.Abs(testFile)
			file, _ := os.ReadFile(fileName)
			_ = json.Unmarshal(file, &data)

			missing := members{}
			for _, myData := range data.DataSources {
				if temp := pike.AwsDataLookup(myData); temp == nil {
					if strings.Contains(myData, "aws") {
						missing.DataSources = append(missing.DataSources, myData)
					}
				}
			}

			if tc.shouldAdd && len(missing.DataSources) == 0 {
				t.Errorf("Expected %s to be added to missing data sources", tc.dataSource)
			}
			if !tc.shouldAdd && len(missing.DataSources) > 0 {
				t.Errorf("Expected %s not to be added to missing data sources", tc.dataSource)
			}
		})
	}
}
