resource "aws_internetmonitor_monitor" "pike" {
  monitor_name                  = "pike"
  traffic_percentage_to_monitor = 1
  tags = {
    pike    = "permissions"
    another = "tag"
  }
}
