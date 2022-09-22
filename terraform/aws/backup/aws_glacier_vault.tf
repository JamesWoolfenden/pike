resource "aws_glacier_vault" "pike" {
  name = "pike"

  notification {
    sns_topic = "arn:aws:sns:eu-west-2:680235478471:pike"
    events    = ["ArchiveRetrievalCompleted", "InventoryRetrievalCompleted"]
  }

  access_policy = <<EOF
{
    "Version":"2012-10-17",
    "Statement":[
       {
          "Sid": "add-read-only-perm",
          "Principal": "*",
          "Effect": "Allow",
          "Action": [
             "glacier:InitiateJob",
             "glacier:GetJobOutput"
          ],
          "Resource": "arn:aws:glacier:eu-west-2:680235478471:vaults/pike"
       }
    ]
}
EOF

  tags = {
    pike = "permissions"
    // delete="me"
  }
}
