resource "aws_controltower_landing_zone" "pike" {
  manifest_json = file("${path.module}/LandingZoneManifest.json")
  version       = "3.2"
}
