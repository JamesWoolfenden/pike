data "aws_servicequotas_templates" "pike" {
  provider = aws.usa
  region   = "us-east-1"
}

provider "aws" {
  alias  = "usa"
  region = "us-east-1"
}

output "templates" {
  value = data.aws_servicequotas_templates.pike
}
