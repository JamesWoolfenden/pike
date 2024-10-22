data "aws_prometheus_default_scraper_configuration" "pike" {
}

output "aws_prometheus_default_scraper_configuration" {
  value = data.aws_prometheus_default_scraper_configuration.pike
}
