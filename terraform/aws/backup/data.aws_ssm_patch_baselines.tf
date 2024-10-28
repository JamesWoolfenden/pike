data "aws_ssm_patch_baselines" "pike" {
}

output "aws_ssm_patch_baselines" {
  value = data.aws_ssm_patch_baselines.pike
}
