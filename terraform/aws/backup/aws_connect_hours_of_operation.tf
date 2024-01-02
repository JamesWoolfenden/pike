resource "aws_connect_hours_of_operation" "pike" {
  name        = "Office Hours"
  description = "Monday office hours"
  time_zone   = "EST"
  instance_id = aws_connect_instance.pike.id
  config {
    day = "MONDAY"

    end_time {
      hours   = 23
      minutes = 8
    }

    start_time {
      hours   = 8
      minutes = 0
    }
  }

  config {
    day = "TUESDAY"

    end_time {
      hours   = 21
      minutes = 0
    }

    start_time {
      hours   = 9
      minutes = 0
    }
  }

  tags = {
    pike    = "permissions"
    another = "tag"
  }
}
