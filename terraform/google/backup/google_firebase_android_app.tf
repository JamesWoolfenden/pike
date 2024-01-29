resource "google_firebase_android_app" "pike" {
  provider      = google-beta
  display_name  = "Display Name Basic"
  package_name  = "android.package.app"
  sha1_hashes   = ["2145bdf698b8715039bd0e83f2069bed435ac21c"]
  sha256_hashes = ["2145bdf698b8715039bd0e83f2069bed435ac21ca1b2c3d4e5f6123456789abc"]

}
