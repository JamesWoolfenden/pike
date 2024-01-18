resource "aws_codeguruprofiler_profiling_group" "pike" {
  name             = "example"
  compute_platform = "Default"

  agent_orchestration_config {
    profiling_enabled = true
  }
  tags = {
    pike = "permissions"
  }
}
