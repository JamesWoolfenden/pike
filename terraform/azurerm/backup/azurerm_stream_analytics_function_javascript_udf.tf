resource "azurerm_stream_analytics_function_javascript_udf" "pike_gen" {
  name                      = "example-javascript-function"
  stream_analytics_job_name = "pike"
  resource_group_name       = "pike"

  script = <<SCRIPT
function getRandomNumber(in) {
  return in;
}
SCRIPT


  input {
    type = "bigint"
  }

  output {
    type = "bigint"
  }
}
