terraform {
  required_providers {
    epilot-purpose = {
      source  = "epilot-dev/epilot-purpose"
      version = "0.3.0"
    }
  }
}


variable "epilot_auth" {
  type = string
}


provider "epilot-purpose" {
  # Configuration options
  epilot_auth = var.epilot_auth
}


resource "epilot-purpose_taxonomy" "my_taxonomy" {
  create = [
    {
      name = "Nishu Created Purpose from terraform"
    },
  ]
  slug   = "purpose"
  update = []
  delete = []
}