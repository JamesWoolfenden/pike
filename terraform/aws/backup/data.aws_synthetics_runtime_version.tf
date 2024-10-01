data "aws_synthetics_runtime_version" "pike" {
  prefix = "syn-nodejs-puppeteer"
  latest = true
}

output "aws_synthetics_runtime_version" {
  value = data.aws_synthetics_runtime_version.pike
}
