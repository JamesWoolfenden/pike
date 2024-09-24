data "aws_bedrockagent_agent_versions" "pike" {
  agent_id = "example-id"
}

output "aws_bedrockagent_agent_versions" {
  value = data.aws_bedrockagent_agent_versions.pike
}
