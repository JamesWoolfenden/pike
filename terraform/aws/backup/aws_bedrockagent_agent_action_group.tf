resource "aws_bedrockagent_agent_action_group" "pike" {
  provider          = aws.central
  action_group_name = "pike"
  agent_id          = aws_bedrockagent_agent.test.agent_id
  agent_version     = "DRAFT"
}
