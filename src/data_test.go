package pike_test

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	pike "github.com/jameswoolfenden/pike/src"
)

func TestGetResources(t *testing.T) {
	t.Parallel()

	type args struct {
		file    string
		dirName string
	}

	file, _ := filepath.Abs("../testdata/scan/examples/simple/aws_s3_bucket.pike.tf")
	moduleFile, _ := filepath.Abs("../testdata/modules/examples/local/module.local.tf")

	tests := []struct {
		name    string
		args    args
		want    []pike.ResourceV2
		wantErr bool
	}{
		{
			"empty",
			args{"", "../testdata/scan/examples/simple"},
			nil, true,
		},
		{
			"no_dir",
			args{file, ""},
			[]pike.ResourceV2{{
				"resource", "aws_s3_bucket", "pike", "aws", []string{"bucket"},
			}},
			false,
		},
		{
			"dir",
			args{file, "../testdata/scan/examples/simple"},
			[]pike.ResourceV2{{
				"resource", "aws_s3_bucket", "pike", "aws", []string{"bucket"},
			}},
			false,
		},
		{
			"module",
			args{moduleFile, "../testdata/modules/examples/local"},
			[]pike.ResourceV2{
				{
					"resource", "aws_s3_bucket", "pike", "aws", []string{"name"},
				},
				{"module", "local", "", "local", []string{"source"}},
			},
			false,
		},
		{
			"not a path",
			args{moduleFile, "../testdata/modules/examples/rubbish"},
			[]pike.ResourceV2{
				{
					"resource", "aws_s3_bucket", "pike", "aws", []string{"name"},
				},
				{"module", "local", "", "local", []string{"source"}},
			},
			false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := pike.GetResources(tt.args.file, tt.args.dirName)

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
		block   *hclsyntax.Block
		dirName string
	}

	dirName, _ := filepath.Abs("../testdata/modules/examples/local")
	block := getInitialBlock(dirName + "/module.local.tf")

	duffName, _ := filepath.Abs("../testdata/modules/examples/rubbish")
	duffBlock := getInitialBlock(duffName + "/module.local.tf")

	notLocal, _ := filepath.Abs("../testdata/modules/examples/notlocal")
	notBlock := getInitialBlock(notLocal + "/module.local.tf")

	moduleJson := make(pike.ModuleJson)
	tests := []struct {
		name    string
		args    args
		want    []pike.ResourceV2
		wantErr bool
	}{
		{
			"local",
			args{block, dirName},
			[]pike.ResourceV2{{
				"resource",
				"aws_s3_bucket",
				"pike",
				"aws",
				[]string{"name"},
			}},
			false,
		},
		{name: "rubbish", args: args{duffBlock, duffName}, wantErr: false},
		{name: "notLocal", args: args{notBlock, notLocal}, wantErr: false},
	}

	for _, tt := range tests {
		// t.Parallel()
		t.Run(tt.name, func(t *testing.T) {
			got, err := pike.GetLocalModules(tt.args.block, tt.args.dirName, moduleJson)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetResources() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLocalModules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getInitialBlock(file string) *hclsyntax.Block {
	body, _ := pike.GetResourceBlocks(file)
	if body != nil {
		blocks := body.Blocks
		block := blocks[0]

		return block
	}

	return nil
}

func TestGetModulePath(t *testing.T) {
	type args struct {
		block *hclsyntax.Block
	}

	dirName, _ := filepath.Abs("../testdata/modules/examples/local")
	block := getInitialBlock(dirName + "/module.local.tf")

	tests := []struct {
		name string
		args args
		want string
	}{
		{"basic", args{block}, "../../"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pike.GetModulePath(tt.args.block); got != tt.want {
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

	closeRange := hcl.Range{
		Filename: file, Start: hcl.Pos{Line: 12, Column: 3, Byte: 245}, End: hcl.Pos{Line: 12, Column: 4, Byte: 246},
	}
	open := hcl.Range{
		Filename: file, Start: hcl.Pos{Line: 10, Column: 13, Byte: 206}, End: hcl.Pos{Line: 10, Column: 14, Byte: 207},
	}
	typeRange := hcl.Range{
		Filename: file, Start: hcl.Pos{Line: 10, Column: 3, Byte: 196}, End: hcl.Pos{Line: 10, Column: 12, Byte: 205},
	}
	block := hclsyntax.Block{
		Type: "lifecycle", Body: &body, TypeRange: typeRange, OpenBraceRange: open, CloseBraceRange: closeRange,
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"get attributes",
			args{
				[]string{"domain_name", "validation_method", "tags", "lifecycle"},
				&block,
			},
			[]string{"domain_name", "validation_method", "tags", "lifecycle", "create_before_destroy"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pike.GetBlockAttributes(tt.args.attributes, tt.args.block); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlockAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPermission(t *testing.T) {
	t.Parallel()

	type args struct {
		result pike.ResourceV2
	}

	tests := []struct {
		name    string
		args    args
		want    pike.Sorted
		wantErr bool
	}{
		{
			name: "parse",
			args: args{
				result: pike.ResourceV2{
					TypeName:     "resource",
					Name:         "aws_acm_certificate",
					ResourceName: "pike",
					Provider:     "aws",
					Attributes: []string{
						"domain_name",
						"validation_method",
						"tags",
						"lifecycle",
						"create_before_destroy",
					},
				},
			},
			want: pike.Sorted{AWS: []string{
				"acm:AddTagsToCertificate",
				"acm:RemoveTagsFromCertificate",
				"acm:RequestCertificate",
				"acm:DescribeCertificate",
				"acm:ListTagsForCertificate",
				"acm:DeleteCertificate",
				"acm:DeleteCertificate",
			}},
		},
		{
			name: "no tags",
			args: args{
				result: pike.ResourceV2{
					TypeName:     "resource",
					Name:         "aws_acm_certificate",
					ResourceName: "pike",
					Provider:     "aws",
					Attributes: []string{
						"domain_name",
						"validation_method",
						"lifecycle",
						"create_before_destroy",
					},
				},
			},
			want: pike.Sorted{AWS: []string{
				"acm:RequestCertificate",
				"acm:DescribeCertificate",
				"acm:ListTagsForCertificate",
				"acm:DeleteCertificate",
				"acm:DeleteCertificate",
			}},
		},
		{
			name: "no-provider",
			args: args{
				result: pike.ResourceV2{
					TypeName:     "resource",
					Name:         "bogus_test",
					ResourceName: "pike",
					Provider:     "bogus",
					Attributes: []string{
						"domain_name",
						"validation_method",
						"tags",
						"lifecycle",
						"create_before_destroy",
					},
				},
			},
			want: pike.Sorted{},
		},
		{
			name: "no-iam",
			args: args{
				result: pike.ResourceV2{
					TypeName:     "resource",
					Name:         "random_string",
					ResourceName: "pike",
					Provider:     "random",
					Attributes: []string{
						"length",
						"special",
					},
				},
			},
			want: pike.Sorted{},
		},
		{name: "not-implemented", args: args{result: pike.ResourceV2{Provider: "linode"}}},
		{name: "bogus", args: args{result: pike.ResourceV2{Provider: "bogus"}}},
		{
			name: "azure",
			args: args{
				result: pike.ResourceV2{
					TypeName:     "resource",
					Name:         "azurerm_key_vault",
					ResourceName: "MyDemoApiKey",
					Provider:     "azurerm",
					Attributes:   []string{"name", "name", "resource_group"},
				},
			},
			want: pike.Sorted{
				AZURE: []string{
					"Microsoft.Resources/subscriptions/resourcegroups/read",
					"Microsoft.KeyVault/vaults/read",
					"Microsoft.KeyVault/vaults/write",
					"Microsoft.KeyVault/vaults/delete",
					"Microsoft.KeyVault/locations/deletedVaults/read",
				},
			},
		},
		{
			name: "google",
			args: args{
				result: pike.ResourceV2{
					TypeName: "resource", Name: "google_compute_instance", ResourceName: "found", Provider: "google",
					Attributes: []string{"name", "machine_type", "zone"},
				},
			},
			want: pike.Sorted{GCP: []string{
				"compute.zones.get",
				"compute.instances.create",
				"compute.instances.get",
				"compute.disks.create",
				"compute.disks.create",
				"compute.subnetworks.use",
				"compute.subnetworks.useExternalIp",
				"compute.instances.setMetadata",
				"compute.instances.delete",
				"compute.instances.get",
				"compute.instances.delete",
			}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := pike.GetPermission(tt.args.result)

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

func TestGetResourceBlocks(t *testing.T) {
	t.Parallel()

	type args struct {
		file string
	}

	empty := hcl.Range{
		Filename: "testdata/scan/examples/empty/empty.tf",
		Start:    hcl.Pos{Line: 1, Column: 1},
		End:      hcl.Pos{Line: 1, Column: 1},
	}

	random := hcl.Range{
		Filename: "testdata/scan/examples/random/random_string.pike.tf",
		Start:    hcl.Pos{Line: 1, Column: 1},
		End:      hcl.Pos{Line: 6, Column: 1, Byte: 119},
	}
	tests := []struct {
		name    string
		args    args
		want    hcl.Range
		wantErr bool
	}{
		{
			name: "pass",
			args: args{"testdata/scan/examples/random/random_string.pike.tf"},
			want: random,
		},
		{
			name: "empty",
			args: args{"testdata/scan/examples/empty/empty.tf"},
			want: empty,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := pike.GetResourceBlocks(tt.args.file)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetResourceBlocks() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got.SrcRange, tt.want) {
				t.Errorf("GetResourceBlocks() = %v, want %v", got.SrcRange, tt.want)
			}
		})
	}
}

func TestDetectBackend(t *testing.T) {
	t.Parallel()

	type args struct {
		resource  pike.ResourceV2
		block     *hclsyntax.Block
		resources []pike.ResourceV2
	}

	resource := pike.ResourceV2{
		TypeName:     "terraform",
		Name:         "",
		ResourceName: "",
		Provider:     "",
		Attributes:   nil,
	}

	convert := pike.ResourceV2{
		TypeName:     "terraform",
		Name:         "backend",
		ResourceName: "",
		Provider:     "aws",
		Attributes:   []string{"s3"},
	}

	var wanted []pike.ResourceV2

	want := pike.ResourceV2{
		TypeName:     "terraform",
		Name:         "backend",
		ResourceName: "string",
		Provider:     "aws",
		Attributes:   []string{"s3"},
	}
	wanted = append(wanted, want)
	wanted = append(wanted, convert)

	ablock := hclsyntax.Block{
		Type:   "backend",
		Labels: []string{"s3"},
		Body:   nil,
	}

	blocks := hclsyntax.Blocks{
		&ablock,
	}

	test := hclsyntax.Body{
		Attributes: nil,
		Blocks:     blocks,
		SrcRange:   hcl.Range{},
		EndRange:   hcl.Range{},
	}

	block := hclsyntax.Block{
		Type: "terraform",
		Body: &test,
	}

	emptyBlock := hclsyntax.Block{}
	item := pike.ResourceV2{
		TypeName:     "terraform",
		Name:         "backend",
		ResourceName: "string",
		Provider:     "aws",
		Attributes:   []string{"s3"},
	}

	var empty []pike.ResourceV2

	var found []pike.ResourceV2

	var nought []pike.ResourceV2

	found = append(found, item)

	tests := []struct {
		name    string
		args    args
		want    []pike.ResourceV2
		wantErr bool
	}{
		{name: "nothing", args: args{resource, &emptyBlock, empty}, want: nought, wantErr: true},
		{name: "backend", args: args{resource, &block, found}, want: wanted, wantErr: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := pike.DetectBackend(tt.args.resource, tt.args.block, tt.args.resources)

			if (err != nil) != tt.wantErr {
				t.Errorf("DetectBackend() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectBackend() got = %v, want %v", got, tt.want)
			}
		})
	}
}
