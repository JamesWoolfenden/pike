data "google_sql_tiers" "pike" {}


output "google_sql_tiers" {
  value = data.google_sql_tiers.pike
}
