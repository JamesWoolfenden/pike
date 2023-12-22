resource "aws_connect_user_hierarchy_structure" "pike" {
  instance_id = aws_connect_instance.pike.id
  hierarchy_structure {
    level_one {
      name = "levelone"
    }
  }
}
