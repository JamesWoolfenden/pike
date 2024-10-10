resource "aws_quicksight_template_alias" "pike" {
  alias_name              = "example-alias"
  template_id             = aws_quicksight_template.pike.template_id
  template_version_number = aws_quicksight_template.pike.version_number
}
