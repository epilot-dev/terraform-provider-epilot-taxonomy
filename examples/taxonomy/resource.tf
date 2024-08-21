terraform {
  required_providers {
    epilot-taxonomy = {
      source  = "epilot-dev/epilot-taxonomy"
      version = "0.5.1"
    }
  }
}

variable "epilot_auth" {
  type = string
}

provider "epilot-taxonomy" {
  epilot_auth = var.epilot_auth
  server_url = "https://entity.dev.sls.epilot.io"
}


resource "epilot-taxonomy_taxonomy_classification" "my_taxo" {
  name                = "Purpose Classfication by Terraform"
  slug                = "purpose:terraform"
}
