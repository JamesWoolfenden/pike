resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [




    //google_vertex_ai_index


    //google_vertex_ai_feature_online_store


    //google_vertex_ai_dataset


    //google_vertex_ai_endpoint


    //google_vertex_ai_endpoint_iam_policy,google_vertex_ai_endpoint_iam_member,google_vertex_ai_endpoint_iam_binding
    "aiplatform.endpoints.getIamPolicy",
    "aiplatform.endpoints.setIamPolicy",

  ]
}
