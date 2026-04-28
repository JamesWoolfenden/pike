resource "azurerm_monitor_data_collection_rule" "pike_gen" {
  name                        = "example-rule"
  resource_group_name         = "pike"
  location                    = "pike"
  data_collection_endpoint_id = "pike"

  destinations {
    log_analytics {
      workspace_resource_id = "pike"
      name                  = "example-destination-log"
    }

    event_hub {
      event_hub_id = "pike"
      name         = "example-destination-eventhub"
    }

    storage_blob {
      storage_account_id = "pike"
      container_name     = "pike"
      name               = "example-destination-storage"
    }

    azure_monitor_metrics {
      name = "example-destination-metrics"
    }
  }

  data_flow {
    streams      = ["Microsoft-InsightsMetrics"]
    destinations = ["example-destination-metrics"]
  }

  data_flow {
    streams      = ["Microsoft-InsightsMetrics", "Microsoft-Syslog", "Microsoft-Perf"]
    destinations = ["example-destination-log"]
  }

  data_flow {
    streams       = ["Custom-MyTableRawData"]
    destinations  = ["example-destination-log"]
    output_stream = "Microsoft-Syslog"
    transform_kql = "source | project TimeGenerated = Time, Computer, Message = AdditionalContext"
  }

  data_sources {
    syslog {
      facility_names = ["*"]
      log_levels     = ["*"]
      name           = "example-datasource-syslog"
      streams        = ["Microsoft-Syslog"]
    }

    iis_log {
      streams         = ["Microsoft-W3CIISLog"]
      name            = "example-datasource-iis"
      log_directories = ["C:\\Logs\\W3SVC1"]
    }

    log_file {
      name          = "example-datasource-logfile"
      format        = "text"
      streams       = ["Custom-MyTableRawData"]
      file_patterns = ["C:\\JavaLogs\\*.log"]
      settings {
        text {
          record_start_timestamp_format = "ISO 8601"
        }
      }
    }

    performance_counter {
      streams                       = ["Microsoft-Perf", "Microsoft-InsightsMetrics"]
      sampling_frequency_in_seconds = 60
      counter_specifiers            = ["Processor(*)\\% Processor Time"]
      name                          = "example-datasource-perfcounter"
    }

    windows_event_log {
      streams        = ["Microsoft-WindowsEvent"]
      x_path_queries = ["*![System/Level=1]"]
      name           = "example-datasource-wineventlog"
    }

    extension {
      streams            = ["Microsoft-WindowsEvent"]
      input_data_sources = ["example-datasource-wineventlog"]
      extension_name     = "example-extension-name"
      extension_json = jsonencode({
        a = 1
        b = "hello"
      })
      name = "example-datasource-extension"
    }
  }

  stream_declaration {
    stream_name = "Custom-MyTableRawData"
    column {
      name = "Time"
      type = "datetime"
    }
    column {
      name = "Computer"
      type = "string"
    }
    column {
      name = "AdditionalContext"
      type = "string"
    }
  }

  identity {
    type         = "UserAssigned"
    identity_ids = ["pike"]
  }

  description = "data collection rule example"
  tags = {
    foo = "bar"
  }
  depends_on = [
    azurerm_log_analytics_solution.example
  ]
}
