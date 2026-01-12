resource "google_sql_source_representation_instance" "pike" {
  name             = "my-instance"
  region           = "us-central1"
  database_version = "MYSQL_8_0"
  host             = "10.20.30.40"
  port             = 3306
  username         = "some-user"
  password         = "password-for-the-user"
}
