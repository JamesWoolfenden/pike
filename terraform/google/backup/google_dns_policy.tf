resource "google_dns_policy" "pike" {
  name                      = "pike"
  enable_inbound_forwarding = true

  enable_logging = true

  alternative_name_server_config {
    target_name_servers {
      ipv4_address    = "172.16.1.10"
      forwarding_path = "private"
    }
    target_name_servers {
      ipv4_address = "172.16.1.20"
    }
  }

  networks {
    network_url = google_compute_network.network-1.id
  }
  networks {
    network_url = google_compute_network.network-2.id
  }
}
