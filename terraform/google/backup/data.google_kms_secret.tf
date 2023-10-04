data "google_kms_secret" "pike" {
  crypto_key = "projects/pike-gcp/locations/europe-west2/keyRings/pike/cryptoKeys/pike"
  ciphertext = "CiQAqD+xX4SXOSziF4a8JYvq4spfAuWhhYSNul33H85HnVtNQW4SOgDu2UZ46dQCRFl5MF6ekabviN8xq+F+2035ZJ85B+xTYXqNf4mZs0RJitnWWuXlYQh6axnnJYu3kDU="
}
