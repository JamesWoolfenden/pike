resource "aws_m2_application" "pike" {
  name        = "Example"
  engine_type = "bluage"
  definition {
    content = <<EOF
    {
  "definition": {
    "listeners": [
      {
        "port": 8196,
        "type": "http"
      }
    ],
    "ba-application": {
      "app-location": "PlanetsDemo-v1.zip"
    }
  },
  "source-locations": [
    {
      "source-id": "s3-source",
      "source-type": "s3",
      "properties": {
        "s3-bucket": "example-bucket",
"s3-key-prefix": "v1"
      }
    }
  ],
  "template-version": "2.0"
}

EOF
  }
}
