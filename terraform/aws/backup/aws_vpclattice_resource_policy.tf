resource "aws_vpclattice_resource_policy" "pike" {
  resource_arn = aws_vpclattice_service_network.pike.arn

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Sid    = "test-pol-principals-6"
      Effect = "Allow"
      Principal = {
        "AWS" = "arn:aws:iam::680235478471:root"
      }
      Action = [
        "vpc-lattice:CreateServiceNetworkVpcAssociation",
        "vpc-lattice:CreateServiceNetworkServiceAssociation",
        "vpc-lattice:GetServiceNetwork"
      ]
      Resource = aws_vpclattice_service_network.pike.arn
    }]
  })
}
