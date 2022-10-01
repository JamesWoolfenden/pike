resource "aws_amplify_app" "pike" {
  name       = "pike"
  repository = "https://github.com/hortonworks/simple-yarn-app"

  auto_branch_creation_patterns = []
  # The default build_spec added by the Amplify Console for React.
  build_spec = <<-EOT
    version: 0.1
    frontend:
      phases:
        preBuild:
          commands:
            - yarn install
        build:
          commands:
            - yarn run build
      artifacts:
        baseDirectory: build
        files:
          - '**/*'
      cache:
        paths:
          - node_modules/**/*
  EOT

  enable_auto_branch_creation = false
  enable_basic_auth           = false
  enable_branch_auto_build    = false
  enable_branch_auto_deletion = false
  iam_service_role_arn        = "arn:aws:iam::680235478471:role/amplifyconsole-backend-role"

  access_token = "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"
  # The default rewrites and redirects added by the Amplify Console.
  custom_rule {
    source = "/<*>"
    status = "404"
    target = "/index.html"
  }

  environment_variables = {
    ENV = "test"
  }

  tags = {
    pike = "permissions"
  }
}

variable "GITHUB_TOKEN" {}
