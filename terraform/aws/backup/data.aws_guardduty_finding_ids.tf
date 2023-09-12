data "aws_guardduty_finding_ids" "pike" {
  detector_id = data.aws_guardduty_detector.pike.id
}
