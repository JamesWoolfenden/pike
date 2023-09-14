resource "google_storage_bucket" "static-site" {
  name                        = "pike-image-store-com"
  location                    = "EU"
  force_destroy               = true
  storage_class               = "STANDARD"
  uniform_bucket_level_access = true
  labels = {
    pike = "permissions"
  }

  #  website {
  #    main_page_suffix = "index.html"
  #    not_found_page   = "404.html"
  #  }
  #  cors {
  #    origin          = ["http://image-store.com"]
  #    method          = ["GET", "HEAD", "PUT", "POST", "DELETE"]
  #    response_header = ["*"]
  #    max_age_seconds = 3600
  #  }
}
