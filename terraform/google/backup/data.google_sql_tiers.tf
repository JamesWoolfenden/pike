data "google_sql_tiers" "pike" {}


output "tiers" {
  value = data.google_sql_tiers.pike
}
