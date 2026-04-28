resource "azurerm_network_connection_monitor" "pike_gen" {
  name               = "example-Monitor"
  network_watcher_id = "pike"
  location           = "pike"

  endpoint {
    name               = "source"
    target_resource_id = "pike"

    filter {
      item {
        address = "pike"
        type    = "AgentAddress"
      }

      type = "Include"
    }
  }

  endpoint {
    name    = "destination"
    address = "terraform.io"
  }

  test_configuration {
    name                      = "tcpName"
    protocol                  = "Tcp"
    test_frequency_in_seconds = 60

    tcp_configuration {
      port = 80
    }
  }

  test_group {
    name                     = "exampletg"
    destination_endpoints    = ["destination"]
    source_endpoints         = ["source"]
    test_configuration_names = ["tcpName"]
  }

  notes = "examplenote"

  output_workspace_resource_ids = ["pike"]

  depends_on = ["pike"]
}
