resource "aws_dlm_lifecycle_policy" "pike" {
  description        = "pike"
  execution_role_arn = "arn:aws:iam::680235478471:role/service-role/AWSDataLifecycleManagerDefaultRole"
  policy_details {
    resource_types = ["VOLUME"]

    schedule {
      name = "2 weeks of daily snapshots"

      create_rule {
        interval      = 24
        interval_unit = "HOURS"
        times         = ["23:45"]
      }

      retain_rule {
        count = 14
      }

      tags_to_add = {
        SnapshotCreator = "DLM"
      }

      copy_tags = false
    }

    target_tags = {
      Snapshot = "true"
    }
  }
  state = "ENABLED"

  tags = {
    pike   = "permissions"
    delete = "me"
  }
}
