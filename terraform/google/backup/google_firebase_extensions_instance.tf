#resource "google_firebase_extensions_instance" "pike" {
#  provider = google-beta
#  project = google_firebase_project.pike.project
#  instance_id = "storage-resize-images"
#  config {
#    extension_ref = "firebase/storage-resize-images"
#    extension_version = "0.1.37"
#
#    # The following params apply to the firebase/storage-resize-images extension.
#    # Different extensions may have different params
#    params = {
#      DELETE_ORIGINAL_FILE = false
#      MAKE_PUBLIC          = false
#      IMAGE_TYPE           = false
#      IS_ANIMATED          = true
#      FUNCTION_MEMORY      = 1024
#      DO_BACKFILL          = false
#      IMG_SIZES            = "200x200"
#      IMG_BUCKET           = google_storage_bucket.images.name
#      LOCATION             = ""
#    }
#
#    system_params = {
#      "firebaseextensions.v1beta.function/maxInstances"               = 3000
#      "firebaseextensions.v1beta.function/memory"                     = 256
#      "firebaseextensions.v1beta.function/minInstances"               = 0
#      "firebaseextensions.v1beta.function/vpcConnectorEgressSettings" = "VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED"
#    }
#
#    allowed_event_types = [
#      "firebase.extensions.storage-resize-images.v1.complete"
#    ]
#
#    eventarc_channel = "projects/pike-gcp/locations/europe-west1/channels/firebase"
#  }
#  depends_on = [
#    google_firebase_project.pike
#  ]
#}
#
#
#resource "google_storage_bucket" "images" {
#  provider                    = google-beta
#  project                     = "pike-gcp"
#  name                        = "bucket-id-forfirebasejgw"
#  location                    = "EU"
#  uniform_bucket_level_access = true
#
#  # Delete all objects when the bucket is deleted
#  force_destroy = true
#}
