package pike_test

import (
	"os"
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func FuzzGetPermissionMap(f *testing.F) {
	f.Add([]byte(`[{"attributes":{"create":["s3:PutObject"],"read":["s3:GetObject"]}}]`), "create", "aws_s3_bucket")
	f.Add([]byte(`[{"attributes":{"delete":["ec2:TerminateInstances"]}}]`), "delete", "aws_instance")
	f.Add([]byte(`[]`), "create", "aws_s3_bucket")
	f.Add([]byte(`{}`), "create", "aws_s3_bucket")
	f.Add([]byte(``), "create", "aws_s3_bucket")

	f.Fuzz(func(t *testing.T, data []byte, attribute string, resource string) {
		//nolint:errcheck
		pike.GetPermissionMap(data, []string{attribute}, resource)
	})
}

func FuzzGetResourceBlocks(f *testing.F) {
	f.Add([]byte(`resource "aws_s3_bucket" "example" { bucket = "my-bucket" }`))
	f.Add([]byte(`resource "aws_instance" "example" { ami = "ami-123" instance_type = "t2.micro" }`))
	f.Add([]byte(``))

	f.Fuzz(func(t *testing.T, data []byte) {
		tmp, err := os.CreateTemp("", "fuzz-*.tf")
		if err != nil {
			t.Skip()
		}
		defer os.Remove(tmp.Name())
		if _, err := tmp.Write(data); err != nil {
			tmp.Close()
			t.Skip()
		}
		tmp.Close()

		//nolint:errcheck
		pike.GetResourceBlocks(tmp.Name())
	})
}
