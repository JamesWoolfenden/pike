resource "aws_codecommit_approval_rule_template" "pike" {
  name        = "MyExampleApprovalRuleTemplate"
  description = "This is an example approval rule template"

  content = jsonencode({
    Version               = "2018-11-08"
    DestinationReferences = ["refs/heads/master"]
    Statements = [{
      Type                    = "Approvers"
      NumberOfApprovalsNeeded = 2
      ApprovalPoolMembers     = ["arn:aws:sts::680235478471:assumed-role/CodeCommitReview/*"]
    }]
  })
}
