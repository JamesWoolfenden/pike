resource "aws_keyspaces_table" "pike" {
  keyspace_name = aws_keyspaces_keyspace.pike.name
  table_name    = "my_table"

  schema_definition {
    column {
      name = "Message"
      type = "ASCII"
    }

    partition_key {
      name = "Message"
    }
  }
}
