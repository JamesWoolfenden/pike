resource "google_billing_budget" "pike" {
  billing_account = data.google_billing_account.pike.id
  display_name    = "Example Billing Budget"
  amount {
    specified_amount {
      currency_code = "USD"
      units         = "100000"
    }
  }
  threshold_rules {
    threshold_percent = 0.5
  }
}


data "google_billing_account" "pike" {
  billing_account = "013ED3-8B24F6-725D7E"
}
