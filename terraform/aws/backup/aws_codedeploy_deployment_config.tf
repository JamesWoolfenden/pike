resource "aws_codedeploy_deployment_config" "pike" {
  deployment_config_name = "test-deployment-config"
  compute_platform       = "Server"
  minimum_healthy_hosts {
    type  = "HOST_COUNT"
    value = 2
  }
  #  traffic_routing_config {
  #    type = "TimeBasedLinear"
  #
  #    time_based_linear {
  #      interval   = 11
  #      percentage = 10
  #    }
  #  }
}
