resource "aws_db_proxy_default_target_group" "pike" {
  db_proxy_name = aws_db_proxy.pike.name

  connection_pool_config {
    connection_borrow_timeout    = 120
    init_query                   = "SET x=1, y=2"
    max_connections_percent      = 100
    max_idle_connections_percent = 50
    session_pinning_filters      = ["EXCLUDE_VARIABLE_SETS"]
  }

}
