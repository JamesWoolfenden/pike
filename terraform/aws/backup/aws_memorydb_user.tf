resource "aws_memorydb_user" "pike" {
  user_name     = "my-user"
  access_string = "on ~* &* +@all"

  authentication_mode {
    type      = "password"
    passwords = [random_password.example.result]
  }

  tags = {
    pike = "permissions"
  }
}

resource "random_password" "example" {
  length = 16
}
