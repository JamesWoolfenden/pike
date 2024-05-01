data "aws_bedrock_foundation_model" "example" {
  provider = aws.central
  model_id = "amazon.titan-text-express-v1"
}

resource "aws_bedrock_custom_model" "pike" {
  provider              = aws.central
  custom_model_name     = "example-model"
  job_name              = "example-job-1"
  base_model_identifier = data.aws_bedrock_foundation_model.example.model_arn
  role_arn              = aws_iam_role.bedrock.arn

  hyperparameters = {
    "epochCount"              = "1"
    "batchSize"               = "1"
    "learningRate"            = "0.005"
    "learningRateWarmupSteps" = "0"
  }

  output_data_config {
    s3_uri = "s3://${aws_s3_bucket.output.id}/data/"
  }

  training_data_config {
    s3_uri = "s3://${aws_s3_bucket.training.id}/data/train.jsonl"
  }
}

resource "aws_s3_bucket" "training" {
  bucket = "piketrainingjgw"
}

resource "aws_s3_bucket" "output" {
  bucket = "pikeoutputjgw"
}


resource "aws_iam_role" "bedrock" {
  assume_role_policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Principal" : { "AWS" : "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root" },
          "Action" : "sts:AssumeRole",
        }
      ]
    }
  )
}


data "aws_caller_identity" "current" {}
