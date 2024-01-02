resource "aws_connect_contact_flow" "pike" {
  instance_id = aws_connect_instance.pike.id
  name        = "Test"
  description = "Test Contact Flow Description"
  type        = "CONTACT_FLOW"
  content = jsonencode({
    Version     = "2019-10-30"
    StartAction = "12345678-1234-1234-1234-123456789012"
    Actions = [
      {
        Identifier = "12345678-1234-1234-1234-123456789012"
        Type       = "MessageParticipant"

        Transitions = {
          NextAction = "abcdef-abcd-abcd-abcd-abcdefghijkl"
          Errors     = []
          Conditions = []
        }

        Parameters = {
          Text = "Thanks for calling the sample flow!"
        }
      },
      {
        Identifier  = "abcdef-abcd-abcd-abcd-abcdefghijkl"
        Type        = "DisconnectParticipant"
        Transitions = {}
        Parameters  = {}
      }
    ]
  })
  tags = {
    pike    = "permissions"
    another = "tag"
  }
}
