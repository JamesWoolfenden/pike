resource "aws_sagemaker_workteam" "pike" {
  description    = "my description"
  workforce_name = aws_sagemaker_workforce.pike.id
  workteam_name  = "pike"

  member_definition {
    oidc_member_definition {
      groups = ["example"]
    }
  }

  tags = {
    pike = "permissions"
  }
}
