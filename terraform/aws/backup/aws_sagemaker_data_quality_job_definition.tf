resource "aws_sagemaker_data_quality_job_definition" "pike" {
  name = "my-data-quality-job-definition"

  data_quality_app_specification {
    image_uri = data.aws_sagemaker_prebuilt_ecr_image.pike.registry_path
  }
  data_quality_job_input {
    endpoint_input {
      endpoint_name = aws_sagemaker_endpoint.pike.name
    }
  }
  data_quality_job_output_config {
    monitoring_outputs {
      s3_output {
        s3_uri = "https://${aws_s3_bucket.pike.bucket_regional_domain_name}/output"
      }
    }
  }
  job_resources {
    cluster_config {
      instance_count    = 1
      instance_type     = "ml.t3.medium"
      volume_size_in_gb = 20
    }
  }
  role_arn = aws_iam_role.example.arn
}


data "aws_sagemaker_prebuilt_ecr_image" "pike" {
  repository_name = "blazingtext"
}
