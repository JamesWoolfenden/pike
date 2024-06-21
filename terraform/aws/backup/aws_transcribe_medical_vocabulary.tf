resource "aws_transcribe_medical_vocabulary" "pike" {
  vocabulary_name     = "example"
  language_code       = "en-US"
  vocabulary_file_uri = "s3://${aws_s3_bucket.example.id}/${aws_s3_object.object.key}"

  tags = {
    tag1 = "value1"
    tag2 = "value3"
    pike = "permissions"
  }

  depends_on = [
    aws_s3_object.object
  ]
}
