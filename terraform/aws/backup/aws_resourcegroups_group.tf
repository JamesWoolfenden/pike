resource "aws_resourcegroups_group" "pike" {
  name = "pike"

  resource_query {
    query = <<JSON
    {
        "ResourceTypeFilters": [
          "AWS::EC2::Instance"
        ],
        "TagFilters": [
          {
            "Key": "Stage",
            "Values": [
              "Test"
            ]
          }
        ]
      }
JSON
  }
  tags = {
    pike = "permissions"
  }
}
