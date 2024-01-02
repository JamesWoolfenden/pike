#resource "aws_connect_vocabulary" "pike" {
#  instance_id   = aws_connect_instance.pike.arn
#  name          = "example"
#  content       = "Phrase\tIPA\tSoundsLike\tDisplayAs\nLos-Angeles\t\t\tLos Angeles\nF.B.I.\tɛ f b i aɪ\t\tFBI\nEtienne\t\teh-tee-en\t"
#  language_code = "en-US"
#  tags = {
#    pike   = "permissions"
#    "Key1" = "Value1"
#  }
#}
