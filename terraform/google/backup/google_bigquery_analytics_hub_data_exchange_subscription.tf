resource "google_bigquery_analytics_hub_data_exchange_subscription" "pike" {
  provider = google-beta
  project  = google_bigquery_dataset.subscription.project #Subscriber's project
  location = "us"

  data_exchange_project  = google_bigquery_analytics_hub_data_exchange.subscription.project
  data_exchange_location = google_bigquery_analytics_hub_data_exchange.subscription.location
  data_exchange_id       = google_bigquery_analytics_hub_data_exchange.subscription.data_exchange_id

  subscription_id    = "my_subscription_id"
  subscriber_contact = "testuser@example.com"

  destination_dataset {
    location = "us"

    dataset_reference {
      project_id = google_bigquery_dataset.subscription.project #Subscriber's project
      dataset_id = "subscribed_dest_dataset"
    }
    friendly_name = "Subscribed Destination Dataset"
    description   = "Destination dataset for subscription"
    labels = {
      environment = "development"
      owner       = "team-a"
    }
  }

  refresh_policy = "ON_READ"
}
