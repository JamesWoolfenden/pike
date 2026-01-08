data "google_project" "project" {
}

resource "google_essential_contacts_contact" "contact" {
  parent                              = data.google_project.project.id
  email                               = "foo@bar.com"
  language_tag                        = "en-GB"
  notification_category_subscriptions = ["ALL"]
}
