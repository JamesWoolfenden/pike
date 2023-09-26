data "aws_outposts_outpost_instance_type" "pike" {
  arn                      = data.aws_outposts_outpost.pike.arn
  preferred_instance_types = ["m5.large", "m5.4xlarge"]
}

data "aws_outposts_outpost" "pike" {

}
