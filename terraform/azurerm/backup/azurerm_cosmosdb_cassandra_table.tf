resource "azurerm_cosmosdb_cassandra_table" "pike_gen" {
  name                  = "testtable"
  cassandra_keyspace_id = "pike"

  schema {
    column {
      name = "test1"
      type = "ascii"
    }

    column {
      name = "test2"
      type = "int"
    }

    partition_key {
      name = "test1"
    }
  }
}
