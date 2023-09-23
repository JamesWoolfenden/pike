data "aws_outposts_outposts" "pike" {
  site_id = data.aws_outposts_site.pike.id
}
