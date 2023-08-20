locals {
  standard_tags = {
    Repo        = "my-asg-repo"
    CreatedBy   = "Terraform"
    Environment = "${var.namespace}-${var.stage}"
  }
}

#####################################################
# ASG Groups
#####################################################
resource "aws_autoscaling_group" "azbuildgroup_win" {
  count                     = var.azure_windows_asg_count
  name                      = var.azure_windows_asg_name
  min_size                  = 1
  max_size                  = var.azure_windows_asg_max_instances_day
  desired_capacity          = var.azure_windows_asg_max_instances_day
  health_check_grace_period = 30
  health_check_type         = "EC2"
  force_delete              = true
  termination_policies      = ["OldestInstance"]
  vpc_zone_identifier       = data.terraform_remote_state.vpc.outputs.private_subnets
  wait_for_capacity_timeout = "2m"
  metrics_granularity       = "1Minute"
  launch_template {
    id      = aws_launch_template.azure_win_template[0].id
    version = "$Latest"
  }
  depends_on = [
    aws_launch_template.azure_win_template[0]
  ]
  dynamic "tag" {
    for_each = local.standard_tags

    content {
      key                 = tag.key
      value               = tag.value
      propagate_at_launch = false
    }
  }
}
