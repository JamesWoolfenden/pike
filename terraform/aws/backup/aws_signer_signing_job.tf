resource "aws_signer_signing_job" "pike" {
  profile_name = aws_signer_signing_profile.pike.name

  source {
    s3 {
      bucket  = "s3-bucket-name"
      key     = "object-to-be-signed.zip"
      version = "jADjFYYYEXAMPLETszPjOmCMFDzd9dN1"
    }
  }

  destination {
    s3 {
      bucket = "s3-bucket-name"
      prefix = "signed/"
    }
  }

  ignore_signing_job_failure = true

}
