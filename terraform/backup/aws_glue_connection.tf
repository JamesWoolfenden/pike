resource "aws_glue_connection" "example" {
  connection_properties = {
    JDBC_CONNECTION_URL = "jdbc:mysql://example.com/exampledatabase"
    PASSWORD            = "examplepassword"
    USERNAME            = "exampleusername"
  }

  name = "example"
  tags = {
    pike = "permissions"
  }
}
