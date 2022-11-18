resource "aws_backup_framework" "pike" {
  name        = "pike"
  description = "pike"

  control {
    name = "BACKUP_PLAN_MIN_FREQUENCY_AND_MIN_RETENTION_CHECK"

    input_parameter {
      name  = "requiredFrequencyUnit"
      value = "days"
    }
    input_parameter {
      name  = "requiredFrequencyValue"
      value = "1"
    }
    input_parameter {
      name  = "requiredRetentionDays"
      value = "35"
    }
  }
  control {
    name = "BACKUP_RECOVERY_POINT_MINIMUM_RETENTION_CHECK"

    input_parameter {
      name  = "requiredRetentionDays"
      value = "35"
    }
  }
  control {
    name = "BACKUP_RECOVERY_POINT_ENCRYPTED"
  }
  control {
    name = "BACKUP_RECOVERY_POINT_MANUAL_DELETION_DISABLED"
  }
  control {
    name = "BACKUP_RESOURCES_PROTECTED_BY_BACKUP_PLAN"

    scope {
      compliance_resource_types = [
        "Aurora",
        "DynamoDB",
        "EBS",
        "EC2",
        "EFS",
        "FSx",
        "RDS",
        "S3",
        "Storage Gateway",
        "VirtualMachine",
      ]
    }
  }
  control {
    name = "BACKUP_RESOURCES_PROTECTED_BY_CROSS_ACCOUNT"

    scope {
      compliance_resource_types = [
        "Aurora",
        "DynamoDB",
        "EBS",
        "EC2",
        "EFS",
        "FSx",
        "RDS",
        "S3",
        "Storage Gateway",
        "VirtualMachine",
      ]
    }
  }
  control {
    name = "BACKUP_RESOURCES_PROTECTED_BY_CROSS_REGION"

    scope {
      compliance_resource_types = [
        "Aurora",
        "DynamoDB",
        "EBS",
        "EC2",
        "EFS",
        "FSx",
        "RDS",
        "S3",
        "Storage Gateway",
        "VirtualMachine",
      ]
    }
  }
  tags = {
    "pike" = "Permissions"
  }
}
