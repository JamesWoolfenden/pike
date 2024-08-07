resource "aws_bedrockagent_data_source" "pike" {
  knowledge_base_id = "EMDPPAYPZI"
  name              = "example"
  data_source_configuration {
    type = "S3"
    s3_configuration {
      bucket_arn = "arn:aws:s3:::example-bucket"
    }
  }
}
