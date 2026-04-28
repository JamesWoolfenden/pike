resource "azurerm_stream_analytics_output_function" "pike_gen" {
  name                      = "exampleoutput"
  resource_group_name       = "pike"
  stream_analytics_job_name = "pike"
  function_app              = "pike"
  function_name             = "examplefunctionname"
}
