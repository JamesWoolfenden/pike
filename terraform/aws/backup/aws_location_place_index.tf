resource "aws_location_place_index" "pike" {
  data_source = "Esri"
  index_name  = "pike"

  data_source_configuration {
    intended_use = "SingleUse"
  }

  description = "pikey"
  tags = {
    pike = "permissions"
    #    another="tag"
  }
}
