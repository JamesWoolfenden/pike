resource "aws_elb" "pike" {
  name               = "wu-tang"
  availability_zones = ["eu-west-2a"]

  listener {
    instance_port     = 80
    instance_protocol = "http"
    lb_port           = 80
    lb_protocol       = "http"
    #    ssl_certificate_id = "arn:aws:acm:eu-west-2:680235478471:certificate/ac27a014-f902-4fba-a2ce-37320747646c"
  }

  tags = {
    Name = "wu-tang"
  }
}

resource "aws_load_balancer_policy" "pike" {
  load_balancer_name = aws_elb.pike.name
  policy_name        = "pike-ca-pubkey-policy"
  policy_type_name   = "PublicKeyPolicyType"

  # The public key of a CA certificate file can be extracted with:
  # $ cat wu-tang-ca.pem | openssl x509 -pubkey -noout | grep -v '\-\-\-\-' | tr -d '\n' > wu-tang-pubkey
  policy_attribute {
    name  = "PublicKey"
    value = file("example.pub")
  }
}
