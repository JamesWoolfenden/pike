resource "aws_iam_instance_profile" "examplea" {
  name_prefix = "imagebuilder_"
  role        = aws_iam_role.pike.name
}

resource "aws_iam_role" "pike" {
  name_prefix = "imagebuilder_"
  path        = "/"

  assume_role_policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "sts:AssumeRole",
            "Principal": {
               "Service": "ec2.amazonaws.com"
            },
            "Effect": "Allow",
            "Sid": ""
        }
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "builder" {
  role       = aws_iam_role.pike.name
  policy_arn = "arn:aws:iam::aws:policy/AWSImageBuilderFullAccess"
}


output "instance_profile" {
  value = aws_iam_instance_profile.examplea
}
