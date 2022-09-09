package pike

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/hashicorp/hcl/v2"

	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func TestGetResources(t *testing.T) {
	type args struct {
		file    string
		dirName string
	}
	tests := []struct {
		name    string
		args    args
		want    []ResourceV2
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetResources(tt.args.file, tt.args.dirName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetResources() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetResources() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLocalModules(t *testing.T) {
	type args struct {
		block     *hclsyntax.Block
		dirName   string
		Resources []ResourceV2
	}

	file, _ := filepath.Abs("E:\\Code\\pike\\testdata\\modules\\examples\\local\\module.local.tf")
	dirName := "e:/code/pike/testdata/modules/examples/local"
	block := getInitialBlock(file)

	duffName := "e:/code/pike/testdata/modules/examples/rubbish"
	duffBlock := getInitialBlock("e:/code/pike/testdata/modules/examples/rubbish/module.local.tf")

	notLocal := "e:/code/pike/testdata/modules/examples/notlocal"
	notBlock := getInitialBlock("e:/code/pike/testdata/modules/examples/notlocal/module.local.tf")

	tests := []struct {
		name    string
		args    args
		want    []ResourceV2
		wantErr bool
	}{
		{"local",
			args{block, dirName, nil},
			[]ResourceV2{{
				"resource",
				"aws_s3_bucket",
				"pike",
				"aws",
				[]string{"name"}}},
			false},
		{"rubbish", args{duffBlock, duffName, nil}, nil, true},
		{"notlocal", args{notBlock, notLocal, nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getLocalModules(tt.args.block, tt.args.dirName, tt.args.Resources)
			if (err != nil) != tt.wantErr {
				t.Errorf("getLocalModules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLocalModules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getInitialBlock(file string) *hclsyntax.Block {
	body, _ := GetResourceBlocks(file)
	blocks := body.Blocks
	block := blocks[0]
	return block
}

func TestGetModulePath(t *testing.T) {
	type args struct {
		block *hclsyntax.Block
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{ },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetModulePath(tt.args.block); got != tt.want {
				t.Errorf("GetModulePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBlockAttributes(t *testing.T) {
	type args struct {
		attributes []string
		block      *hclsyntax.Block
	}
	file, _ := filepath.Abs("terraform\\aws\\backup/aws_acm_certificate.tf")

	attribute := hclsyntax.Attribute{Name: "create_before_destroy"}
	attributes := hclsyntax.Attributes{
		"create_before_destroy": &attribute,
	}
	body := hclsyntax.Body{Attributes: attributes}

	closeRange := hcl.Range{Filename: file, Start: hcl.Pos{Line: 12, Column: 3, Byte: 245}, End: hcl.Pos{Line: 12, Column: 4, Byte: 246}}
	open := hcl.Range{Filename: file, Start: hcl.Pos{Line: 10, Column: 13, Byte: 206}, End: hcl.Pos{Line: 10, Column: 14, Byte: 207}}
	typeRange := hcl.Range{Filename: file, Start: hcl.Pos{Line: 10, Column: 3, Byte: 196}, End: hcl.Pos{Line: 10, Column: 12, Byte: 205}}
	block := hclsyntax.Block{Type: "lifecycle", Body: &body, TypeRange: typeRange, OpenBraceRange: open, CloseBraceRange: closeRange}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{"get attributes", args{
			[]string{"domain_name", "validation_method", "tags", "lifecycle"},
			&block},
			[]string{"domain_name", "validation_method", "tags", "lifecycle", "create_before_destroy"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBlockAttributes(tt.args.attributes, tt.args.block); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlockAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPermission(t *testing.T) {
	type args struct {
		result ResourceV2
	}
	tests := []struct {
		name    string
		args    args
		want    Sorted
		wantErr bool
	}{
		{"parse",
			args{
				ResourceV2{
					"resource",
					"aws_acm_certificate",
					"pike",
					"aws",
					[]string{
						"domain_name",
						"validation_method",
						"tags",
						"lifecycle",
						"create_before_destroy",
					}}},
			Sorted{[]string{
				"acm:AddTagsToCertificate",
				"acm:RemoveTagsFromCertificate",
				"acm:RequestCertificate",
				"acm:DescribeCertificate",
				"acm:ListTagsForCertificate",
				"acm:DeleteCertificate",
				"acm:DeleteCertificate"}, nil, nil},
			false},
		{"no tags",
			args{
				ResourceV2{
					"resource",
					"aws_acm_certificate",
					"pike",
					"aws",
					[]string{
						"domain_name",
						"validation_method",
						"lifecycle",
						"create_before_destroy",
					}}},
			Sorted{[]string{
				"acm:RequestCertificate",
				"acm:DescribeCertificate",
				"acm:ListTagsForCertificate",
				"acm:DeleteCertificate",
				"acm:DeleteCertificate"}, nil, nil},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPermission(tt.args.result)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPermission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}
