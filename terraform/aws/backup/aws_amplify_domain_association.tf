resource "aws_amplify_domain_association" "pike" {
  app_id      = aws_amplify_app.pike.id
  domain_name = "example.com"

  sub_domain {
    branch_name = aws_amplify_branch.pike.branch_name
    prefix      = ""
  }


  sub_domain {
    branch_name = aws_amplify_branch.pike.branch_name
    prefix      = "www"
  }
}
