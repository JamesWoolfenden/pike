provider "aws" {
  region  = "eu-west-2"
  profile = "basic"
}

provider "aws" {
  alias   = "central"
  region  = "us-east-1"
  profile = "basic"
}

provider "aws" {
  alias   = "ire"
  region  = "us-west-1"
  profile = "basic"
}

provider "aws" {
  alias   = "eire"
  region  = "eu-west-1"
  profile = "basic"
}
