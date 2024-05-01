resource "aws_bedrockagent_agent_alias" "example" {
  provider         = aws.central
  agent_alias_name = "my-agent-alias"
  agent_id         = aws_bedrockagent_agent.test.agent_id
  description      = "Test ALias"
}
