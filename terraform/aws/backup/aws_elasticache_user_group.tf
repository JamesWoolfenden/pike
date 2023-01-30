resource "aws_elasticache_user_group" "pike" {
  engine        = "REDIS"
  user_group_id = "pike"
  user_ids      = ["testuserid"]
  tags = {
    pike="permissions"
  }
}