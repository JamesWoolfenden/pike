resource "aws_location_route_calculator" "pike" {
  calculator_name = "pike"
  data_source     = "Esri"
  description     = "pikey"
  tags = {
    pike = "permissions"
    #    another="tag"
  }
}
