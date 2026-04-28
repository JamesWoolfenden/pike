resource "azurerm_stream_analytics_function_javascript_uda" "pike_gen" {
  name                    = "example-javascript-function"
  stream_analytics_job_id = "pike"

  script = <<SCRIPT
function main() {
    this.init = function () {
        this.state = 0;
    }

    this.accumulate = function (value, timestamp) {
        this.state += value;
    }

    this.computeResult = function () {
        return this.state;
    }
}
SCRIPT

  input {
    type = "bigint"
  }

  output {
    type = "bigint"
  }
}
