resource "google_network_security_address_group" "pike" {
  name        = "pike-address-group"
  parent      = "projects/pike-477416"
  description = "Pike test address group - updated"
  location    = "europe-west2"
  items       = ["10.0.0.0/24"]
  type        = "IPV4"
  capacity    = 100
}
