resource "aws_ssm_document" "foo" {
  name          = "test_document"
  document_type = "Command"

  content = <<DOC
  {
    "schemaVersion": "1.2",
    "description": "Check IP configuration of a Linux instance.",
    "parameters": {

    },
    "runtimeConfig": {
      "aws:runShellScript": {
        "properties": [
          {
            "id": "0.aws:runShellScript",
            "runCommand": ["ifconfig"]
          }
        ]
      }
    }
  }
DOC
  #  tags = {
  #    pike="permissions"
  #  }
}
