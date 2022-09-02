resource "aws_ssm_patch_baseline" "production" {
  name             = "patch-baseline"
  approved_patches = ["KB123456"]
  tags = {
    pike   = "permission"
    remove = "me"
  }
}
