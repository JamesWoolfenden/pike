resource "google_biglake_table" "pike" {
  name     = "my_table"
  database = google_biglake_database.database.id
  type     = "HIVE"
  hive_options {
    table_type = "MANAGED_TABLE"
    storage_descriptor {
      location_uri  = "gs://${google_storage_bucket.bucket.name}/${google_storage_bucket_object.data_folder.name}"
      input_format  = "org.apache.hadoop.mapred.SequenceFileInputFormat"
      output_format = "org.apache.hadoop.hive.ql.io.HiveSequenceFileOutputFormat"
    }
    # Some Example Parameters.
    parameters = {
      "spark.sql.create.version"          = "3.1.3"
      "spark.sql.sources.schema.numParts" = "1"
      "transient_lastDdlTime"             = "1680894197"
      "spark.sql.partitionProvider"       = "catalog"
      "owner"                             = "John Doe"
      "spark.sql.sources.schema.part.0"   = "{\"type\":\"struct\",\"fields\":[{\"name\":\"id\",\"type\":\"integer\",\"nullable\":true,\"metadata\":{}},{\"name\":\"name\",\"type\":\"string\",\"nullable\":true,\"metadata\":{}},{\"name\":\"age\",\"type\":\"integer\",\"nullable\":true,\"metadata\":{}}]}"
      "spark.sql.sources.provider"        = "iceberg"
      "provider"                          = "iceberg"
    }
  }

}


resource "google_biglake_catalog" "catalog" {
  name     = "my_catalog"
  location = "US"
}




resource "google_storage_bucket_object" "data_folder" {
  name    = "data/"
  content = " "
  bucket  = google_storage_bucket.bucket.name
}
