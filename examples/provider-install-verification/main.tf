terraform {
  required_providers {
    terraformtest = {
      source = "hashicorp.com/edu/terraformtest"
    }
  }
}

provider "terraformtest" {}

data "terraformtest_coffees" "example" {}
