resource "azurerm_container_registry_task" "pike_gen" {
  name                  = "example-task"
  container_registry_id = "pike"
  platform {
    os = "Linux"
  }
  docker_step {
    dockerfile_path = "Dockerfile"
    context_path    = "https://github.com/<username>/<repository>#<branch>:<folder>"
    image_names     = ["helloworld:{{.Run.ID}}"]
  }
}
