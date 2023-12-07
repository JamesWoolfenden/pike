resource "aws_load_balancer_listener_policy" "wu-tang-listener-policies-443" {
  load_balancer_name = aws_elb.pike.name
  load_balancer_port = 80

  policy_names = [
    aws_load_balancer_policy.pike.policy_name,
  ]

}
