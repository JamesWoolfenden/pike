resource "aws_iam_role" "test" {
  assume_role_policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Principal" : { "AWS" : "arn:aws:iam::680235478471:root" },
          "Action" : "sts:AssumeRole",
        }
      ]
    }
  )
  //  tags = {name="policytest"}
  #  inline_policy {
  #    name = "my_inline_policy"
  #
  #    policy = jsonencode({
  #      Version = "2012-10-17"
  #      Statement = [
  #        {
  #          Action   = ["ec2:Describe*"]
  #          Effect   = "Allow"
  #          Resource = "*"
  #        },
  #      ]
  #    })
  #  }
  description = "a policy bigger"
  name        = "test2"
  //permissions_boundary = "arn:aws:iam::aws:policy/AmazonEC2FullAccess"
  //path="/service-role/"
  //max_session_duration = 3600
  //managed_policy_arns=["arn:aws:iam::aws:policy/AmazonGlacierReadOnlyAccess"]

  //arn:aws:iam::aws:policy/AmazonGlacierReadOnlyAccess
  //force_detach_policies = true
}
