resource "aws_dms_certificate" "pike" {
  certificate_id  = "test-dms-certificate-tf"
  certificate_pem = file("cert.pem")
  tags = {
    pike = "permission"
  }
}
