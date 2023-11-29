provider "aws" {
  region  = "us-east-1"
  profile = "basic"
}

provider "aws" {
  alias   = "central"
  region  = "us-east-1"
  profile = "basic"
}
