data "aws_ce_tags" "pike" {
  time_period {
    start = "2021-01-01"
    end   = "2022-12-01"
  }
}
