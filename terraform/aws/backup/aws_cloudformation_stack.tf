resource "aws_cloudformation_stack" "pike" {
  name = "pike-networking-stack"

  parameters = {
    VPCCidr = "10.0.0.0/16"
  }

  template_body = jsonencode({
    Parameters = {
      VPCCidr = {
        Type        = "String"
        Default     = "10.0.0.0/16"
        Description = "Enter the CIDR block for the VPC. Default is 10.0.0.0/16."
      }
    }

    Resources = {
      myVpc = {
        Type = "AWS::EC2::VPC"
        Properties = {
          CidrBlock = {
            "Ref" = "VPCCidr"
          }
          Tags = [
            {
              Key   = "Name"
              Value = "Primary_CF_VPC"
            }
          ]
        }
      }
    }
  })
  tags = {
    pike = "permissions"
  }
}
