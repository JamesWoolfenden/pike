resource "random_id" "db_name_suffix" {
  byte_length = 4
}

resource "google_sql_database_instance" "main" {

  name             = "main-instance-${random_id.db_name_suffix.hex}"
  database_version = "MYSQL_5_7"
  region           = "europe-west1"
  settings {
    tier = "db-f1-micro"
    user_labels = {
      pike = "permissions"
    }
  }
  //encryption_key_name = "projects/pike-361314/locations/europe-west1/keyRings/pike/cryptoKeys/pike-key-keep"
  deletion_protection = false
}
