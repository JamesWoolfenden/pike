resource "aws_elasticache_user" "pike" {
  user_id       = "basic"
  user_name     = "default"
  access_string = "on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember"
  engine        = "REDIS"
  passwords     = ["password123456789"]
  tags = {
    pike = "permissions"
    #    another="tag"
  }
}
