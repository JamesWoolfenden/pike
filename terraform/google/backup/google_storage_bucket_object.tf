resource "google_storage_bucket_object" "object" {
  name          = basename(var.sourcezip)
  bucket        = "pike-central"
  storage_class = "STANDARD"
  kms_key_name  = "projects/examplea/locations/us-central1/keyRings/pike-us/cryptoKeys/pike/cryptoKeyVersions/1"
  source        = var.sourcezip

}

variable "sourcezip" {
  default = "main_test.zip"
}
