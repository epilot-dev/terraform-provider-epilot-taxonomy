terraform {
  required_providers {
    epilot-taxonomy = {
      source  = "epilot-dev/epilot-taxonomy"
      version = "0.3.0"
    }
  }
}


variable "epilot_auth" {
  type = string
}


provider "epilot-taxonomy" {
  # Configuration options
  epilot_auth = var.epilot_auth
}


resource "epilot-taxonomy_taxonomy" "my_taxonomy" {
  create = [
    {
      name = "Nishu Created Purpose from terraform"
    },
  ]
  slug   = "purpose"
  update = []
  delete = []
}