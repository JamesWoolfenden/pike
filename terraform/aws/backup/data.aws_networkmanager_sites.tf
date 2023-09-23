data "aws_networkmanager_sites" "pike" {
  global_network_id = "s-324234234"

  tags = {
    Env = "test"
  }
}
