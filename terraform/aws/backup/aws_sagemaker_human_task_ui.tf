resource "aws_sagemaker_human_task_ui" "pike" {
  human_task_ui_name = "example"

  ui_template {
    content = file("sagemaker-human-task-ui-template.html")
  }

  tags = {
    pike = "permission"
  }
}
