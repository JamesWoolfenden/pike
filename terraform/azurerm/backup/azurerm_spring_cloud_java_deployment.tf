resource "azurerm_spring_cloud_java_deployment" "pike_gen" {
  name                = "deploy1"
  spring_cloud_app_id = "pike"
  instance_count      = 2
  jvm_options         = "-XX:+PrintGC"

  quota {
    cpu    = "2"
    memory = "4Gi"
  }

  runtime_version = "Java_11"

  environment_variables = {
    "Foo" : "Bar"
    "Env" : "Staging"
  }
}
