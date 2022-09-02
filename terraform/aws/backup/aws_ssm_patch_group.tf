resource "aws_ssm_patch_group" "patchgroup" {
  baseline_id = aws_ssm_patch_baseline.production.id
  patch_group = "patch-group-name"
}
