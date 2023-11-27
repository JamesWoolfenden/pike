resource "aws_opensearchserverless_lifecycle_policy" "pike" {
  provider = aws.central
  name     = "pike"
  type     = "retention"
  policy = jsonencode({
    "Rules" : [
      {
        "ResourceType" : "index",
        "Resource" : ["index/autoparts-inventory/*"],
        "MinIndexRetention" : "81d"
      },
      {
        "ResourceType" : "index",
        "Resource" : ["index/sales/orders*"],
        "NoMinIndexRetention" : true
      }
    ]
  })

}
