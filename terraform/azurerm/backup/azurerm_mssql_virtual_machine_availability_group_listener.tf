resource "azurerm_mssql_virtual_machine_availability_group_listener" "pike_gen" {
  name                         = "listener1"
  availability_group_name      = "availabilitygroup1"
  port                         = 1433
  sql_virtual_machine_group_id = "pike"

  load_balancer_configuration {
    load_balancer_id   = "pike"
    private_ip_address = "10.0.2.11"
    probe_port         = 51572
    subnet_id          = "pike"

    sql_virtual_machine_ids = [
      azurerm_mssql_virtual_machine.example[0].id,
      azurerm_mssql_virtual_machine.example[1].id
    ]
  }

  replica {
    sql_virtual_machine_id = "pike" [0].id
    role                   = "Primary"
    commit                 = "Synchronous_Commit"
    failover               = "Automatic"
    readable_secondary     = "All"
  }

  replica {
    sql_virtual_machine_id = "pike" [1].id
    role                   = "Secondary"
    commit                 = "Asynchronous_Commit"
    failover               = "Manual"
    readable_secondary     = "No"
  }
}
