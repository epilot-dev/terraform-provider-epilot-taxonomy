terraform {
  required_providers {
    epilot-taxonomy = {
      source  = "epilot-dev/epilot-taxonomy"
      version = "0.12.0"
    }
  }
}

provider "epilot-taxonomy" {
  server_url = "..." # Optional
}