resource "aws_location_map" "pike" {
  description = "pikey"
  configuration {
    style = "VectorHereBerlin"
  }

  map_name = "pike"
  tags = {
    pike = "permissions"
    #    another="tag"
  }
}
