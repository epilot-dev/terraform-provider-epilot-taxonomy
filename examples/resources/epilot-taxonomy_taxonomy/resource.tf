# resource "epilot-taxonomy_taxonomy" "my_taxonomy" {

# }


terraform {
  required_providers {
    epilot-taxonomy = {
      source  = "epilot-dev/epilot-taxonomy"
      version = "0.11.5"
    }
  }
}

provider "epilot-taxonomy" {
  # Configuration options
  epilot_auth = var.epilot_auth
  server_url = "https://entity.dev.sls.epilot.io"
}

variable "epilot_auth" {
  type = string
}