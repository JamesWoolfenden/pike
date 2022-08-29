resource "aws_codepipeline" "example" {
  name     = "pike_is_permissions"
  role_arn = "arn:aws:iam::680235478471:role/codepipeline"
  artifact_store {
    location = "testbucketineu-west2"
    type     = "S3"
  }
  stage {
    name = "first"
    action {
      category = "Source"
      name     = "Build"
      owner    = "AWS"
      provider = "CodeCommit"
      configuration = {
        BranchName           = "master"
        PollForSourceChanges = "false"
        RepositoryName       = "cron-poll"
      }
      output_artifacts = ["SourceArtifact"]
      version          = "1"
    }
  }

  stage {
    name = "second"
    action {
      category = "Build"
      name     = "Build2"
      owner    = "AWS"
      provider = "CodeBuild"
      configuration = {
        ProjectName = "Pike"
      }
      input_artifacts  = ["SourceArtifact"]
      output_artifacts = ["BuildArtifact"]
      version          = "1"
    }
  }

  #  tags = {
  #    pike="permission"
  #  }
}
