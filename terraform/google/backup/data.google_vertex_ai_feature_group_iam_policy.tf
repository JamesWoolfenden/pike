data "google_vertex_ai_feature_group_iam_policy" "pike" {
}

output "google_vertex_ai_feature_group_iam_policy" {
  value = data.google_vertex_ai_feature_group_iam_policy.pike
}
