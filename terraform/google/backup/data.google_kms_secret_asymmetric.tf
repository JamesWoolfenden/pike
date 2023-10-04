data "google_kms_secret_asymmetric" "pike" {
  provider           = google-beta
  crypto_key_version = "projects/pike-gcp/locations/europe-west2/keyRings/pike/cryptoKeys/pike/cryptoKeyVersions/1"
  crc32              = "12c59e54"
  ciphertext         = <<EOT
    M7nUoba9EGVTu2LjNjBKGdGVBYjyS/i/AY+4yQMQF0Qf/RfUfX31Jw6+VO9OuThq
    ylu/7ihX9XD4bM7yYdXnMv9p1OHQUlorSBSbb/J6n1W9UJhcp6um8Tw8/Isx4f75
    4PskYS6f8Y2ItliGt1/A9iR5BTgGtJBwOxMlgoX2Ggq+Nh4E5SbdoaE5o6CO1nBx
    eIPsPEebQ6qC4JehQM3IGuV/lrm58+hZhaXAqNzX1cEYyAt5GYqJIVCiI585SUYs
    wRToGyTgaN+zthF0HP9IWlR4Am4LmJ/1OcePTnYw11CkU8wNRbDzVAzogwNH+rXr
    LTmf7hxVjBm6bBSVSNFcBKAXFlllubSfIeZ5hgzGqn54OmSf6odO12L5JxllddHc
    yAd54vWKs2kJtnsKV2V4ZdkI0w6y1TeI67baFZDNGo6qsCpFMPnvv7d46Pg2VOp1
    J6Ivner0NnNHE4MzNmpZRk8WXMwqq4P/gTiT7F/aCX6oFCUQ4AWPQhJYh2dkcOmL
    IP+47Veb10aFn61F1CJwpmOOiGNXKdDT1vK8CMnnwhm825K0q/q9Zqpzc1+1ae1z
    mSqol1zCoa88CuSN6nTLQlVnN/dzfrGbc0boJPaM0iGhHtSzHk4SWg84LhiJB1q9
    A9XFJmOVdkvRY9nnz/iVLAdd0Q3vFtLqCdUYsNN2yh4=
  EOT
}
