data "google_dataplex_asset_iam_policy" "pike" {
  asset         = "pike"
  dataplex_zone = "pike"
  lake          = "pike"
}
