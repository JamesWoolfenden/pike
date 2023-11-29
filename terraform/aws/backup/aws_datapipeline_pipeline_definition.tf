resource "aws_datapipeline_pipeline_definition" "pike" {
  pipeline_id = aws_datapipeline_pipeline.pike.id
  pipeline_object {
    id   = "Default"
    name = "Default"
    field {
      key          = "workerGroup"
      string_value = "workerGroup"
    }
  }
  pipeline_object {
    id   = "Schedule"
    name = "Schedule"
    field {
      key          = "startDateTime"
      string_value = "2012-12-12T00:00:00"
    }
    field {
      key          = "type"
      string_value = "Schedule"
    }
    field {
      key          = "period"
      string_value = "1 hour"
    }
    field {
      key          = "endDateTime"
      string_value = "2012-12-21T18:00:00"
    }
  }
  pipeline_object {
    id   = "SayHello"
    name = "SayHello"
    field {
      key          = "type"
      string_value = "ShellCommandActivity"
    }
    field {
      key          = "command"
      string_value = "echo hello"
    }
    field {
      key          = "parent"
      string_value = "Default"
    }
    field {
      key          = "schedule"
      string_value = "Schedule"
    }
  }
}
