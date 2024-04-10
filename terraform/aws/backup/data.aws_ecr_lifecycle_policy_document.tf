data "aws_ecr_lifecycle_policy_document" "pike" {
  rule {
    priority    = 1
    description = "This is a test."

    selection {
      tag_status      = "tagged"
      tag_prefix_list = ["prod"]
      count_type      = "imageCountMoreThan"
      count_number    = 100
    }
  }
}

output "aws_ecr_lifecycle_policy_document" {
  value = data.aws_ecr_lifecycle_policy_document.pike
}
