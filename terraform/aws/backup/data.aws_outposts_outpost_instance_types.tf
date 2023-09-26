data "aws_outposts_outpost_instance_types" "pike" {
  arn = data.aws_outposts_outpost.pike.arn
}
