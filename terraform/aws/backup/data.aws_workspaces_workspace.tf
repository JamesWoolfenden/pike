data "aws_workspaces_workspace" "pike" {
  workspace_id = "ws-cj5xcxsz5"
}


output "aws_workspaces_workspace" {
  value = data.aws_workspaces_workspace.pike
}
