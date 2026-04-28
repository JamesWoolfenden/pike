resource "azurerm_api_management_email_template" "pike_gen" {
  template_name       = "ConfirmSignUpIdentityDefault"
  resource_group_name = "pike"
  api_management_name = "pike"
  subject             = "Customized confirmation email for your new $OrganizationName API account"
  body                = <<EOF
<!DOCTYPE html >
<html>
<head>
  <meta charset="UTF-8" />
  <title>Customized Letter Title</title>
</head>
<body>
  <p style="font-size:12pt;font-family:'Segoe UI'">Dear $DevFirstName $DevLastName,</p>
</body>
</html>
EOF
}
