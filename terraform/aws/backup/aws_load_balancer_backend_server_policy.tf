resource "aws_load_balancer_backend_server_policy" "pike" {

  load_balancer_name = aws_elb.pike.name
  instance_port      = 80

  policy_names = [
    aws_load_balancer_policy.pike.policy_name,
  ]

}
