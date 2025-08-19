resource "google_dialogflow_cx_security_settings" "pike" {
  display_name          = "dialogflowcx-security-settings"
  location              = "global"
  purge_data_types      = []
  retention_window_days = 7
}
