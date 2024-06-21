resource "aws_transcribe_vocabulary_filter" "pike" {
  vocabulary_filter_name = "example"
  language_code          = "en-GB"
  words                  = ["cars", "bucket", "pike"]

  tags = {
    tag1 = "value1"
    tag2 = "value3"
    pike = "permissions"
  }
}
