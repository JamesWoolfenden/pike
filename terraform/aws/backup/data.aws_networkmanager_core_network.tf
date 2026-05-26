data "aws_networkmanager_core_network" "pike" {
  core_network_id = "core-network-01234567890abcdef"
}

output "aws_networkmanager_core_network" {
  value = data.aws_networkmanager_core_network.pike
}
