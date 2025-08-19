resource "google_dialogflow_cx_generator" "pike" {
  parent        = google_dialogflow_cx_agent.pike.id
  language_code = "fr"
  display_name  = "TF Prompt generator"
  llm_model_settings {
    model       = "gemini-2.0-flash-001"
    prompt_text = "Return me some great results"
  }
  prompt_text {
    text = "Send me great results in french"
  }
  model_parameter {
    temperature = 0.55
  }
}
