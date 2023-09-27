resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "resourcemanager.projects.get",

    //schema
    "pubsub.schemas.attach",
    //"pubsub.schemas.commit",
    "pubsub.schemas.create",
    "pubsub.schemas.delete",
    "pubsub.schemas.get",
    "pubsub.schemas.list",
    "pubsub.schemas.listRevisions",
    //"pubsub.schemas.rollback",
    //"pubsub.schemas.validate",

    //topics
    "pubsub.topics.create",
    "pubsub.topics.delete",
    "pubsub.topics.get",
    "pubsub.topics.list",

    //subscription
    //"pubsub.subscriptions.consume",
    "pubsub.subscriptions.create",
    "pubsub.subscriptions.delete",
    "pubsub.subscriptions.get",
    "pubsub.subscriptions.list",
    "pubsub.subscriptions.update",
    "pubsub.topics.attachSubscription",
    "pubsub.topics.detachSubscription",
  ]
}
