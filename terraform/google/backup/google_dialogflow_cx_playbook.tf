resource "google_dialogflow_cx_playbook" "pike" {
  parent        = google_dialogflow_cx_agent.pike.id
  display_name  = "Example Display Name"
  goal          = "Example Goal"
  playbook_type = "ROUTINE"
  instruction {
    steps {
      text = "step 1"
      steps = jsonencode([
        {
          "text" : "step 1 1"
        },
        {
          "text" : "step 1 2",
          "steps" : [
            {
              "text" : "step 1 2 1"
            },
            {
              "text" : "step 1 2 2"
            }
          ]
        },
        {
          "text" : "step 1 3"
        }
      ])
    }
    steps {
      text = "step 2"
    }
    steps {
      text = "step 3"
    }
  }
}
