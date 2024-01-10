resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_servicequotas_template_association
          "servicequotas:AssociateServiceQuotaTemplate",
          "servicequotas:DisassociateServiceQuotaTemplate",
          "organizations:EnableAWSServiceAccess",

          //aws_servicequotas_template
          "servicequotas:PutServiceQuotaIncreaseRequestIntoTemplate",
          "servicequotas:DeleteServiceQuotaIncreaseRequestFromTemplate",
          "servicequotas:GetServiceQuotaIncreaseRequestFromTemplate",

          //aws_servicecatalog_tag_option_resource_association
          "servicecatalog:AssociateTagOptionWithResource",
          "servicecatalog:DisassociateTagOptionFromResource",

          //aws_servicecatalog_tag_option
          "servicecatalog:CreateTagOption",
          "servicecatalog:DeleteTagOption",
          "servicecatalog:UpdateTagOption",
          "servicecatalog:DescribeTagOption",

          //aws_servicecatalog_service_action
          "servicecatalog:CreateServiceAction",
          "servicecatalog:DescribeServiceAction",
          "servicecatalog:DeleteServiceAction",
          "servicecatalog:UpdateServiceAction",
          "ssm:DescribeDocument",

          //aws_servicecatalog_provisioned_product
          "servicecatalog:ProvisionProduct",
          "servicecatalog:TagResource",
          "servicecatalog:UntagResource",

          //aws_servicecatalog_product_portfolio_association
          "servicecatalog:AssociateProductWithPortfolio",
          "servicecatalog:DisassociateProductFromPortfolio",
          "iam:GetUser",

          //aws_servicecatalog_product
          "servicecatalog:CreateProduct",
          "servicecatalog:DescribeProduct",
          "servicecatalog:DeleteProduct",
          "servicecatalog:UpdateProduct",

          //aws_servicecatalog_principal_portfolio_association
          "servicecatalog:AssociatePrincipalWithPortfolio",
          "servicecatalog:DisassociatePrincipalFromPortfolio",

          //aws_servicecatalog_portfolio
          "servicecatalog:CreatePortfolio",
          "servicecatalog:DescribePortfolio",
          "servicecatalog:DeletePortfolio",
          "servicecatalog:UpdatePortfolio",

          //aws_servicecatalog_organizations_access
          "servicecatalog:EnableAWSOrganizationsAccess",
          "servicecatalog:DisableAWSOrganizationsAccess",
          "servicecatalog:GetAWSOrganizationsAccessStatus",
          "iam:CreateServiceLinkedRole",

          //aws_servicecatalog_budget_resource_association
          "servicecatalog:AssociateBudgetWithResource",
          "servicecatalog:DisassociateBudgetFromResource",

          //aws_servicecatalog_portfolio_share
          "servicecatalog:CreatePortfolioShare",
          "servicecatalog:UpdatePortfolioShare",
          "servicecatalog:DeletePortfolioShare",

        ],
        "Resource" : "*",
      }
    ]
  })
  tags = {
    pike      = "permissions"
    createdby = "JamesWoolfenden"
  }
}

resource "aws_iam_role_policy_attachment" "basic" {
  role       = aws_iam_role.basic.name
  policy_arn = aws_iam_policy.basic.arn
}

resource "aws_iam_user_policy_attachment" "basic" {
  # checkov:skip=CKV_AWS_40: By design
  user       = "basic"
  policy_arn = aws_iam_policy.basic.arn
}
