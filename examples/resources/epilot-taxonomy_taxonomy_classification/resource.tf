# resource "epilot-taxonomy_taxonomy_classification" "my_taxonomyclassification" {
#   id = "taxonomy-slug:classification-slug"
#   manifest = [
#     "123e4567-e89b-12d3-a456-426614174000"
#   ]
#   name = "Wallbox PV"
#   parents = [
#     "taxonomy-slug:classification-slug"
#   ]
#   slug = "wallbox-pv"
# }

terraform {
  required_providers {
    epilot-taxonomy = {
      source  = "epilot-dev/epilot-taxonomy"
      version = "0.8.1"
    }
  }
}

provider "epilot-taxonomy" {
  epilot_auth = "eyJraWQiOiJHUkc3WTl0SHRkdVZ4UmtSMGdDbThtME9Fdmh4YWt4U2dCUHBCQ2hHbURnPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiIzYmEwZmUxMi04YzU3LTRiYjktOWNmZi1mNDliY2NkMmVhYTgiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLmV1LWNlbnRyYWwtMS5hbWF6b25hd3MuY29tXC9ldS1jZW50cmFsLTFfV1JLVWFBWnhHIiwiY3VzdG9tOml2eV9vcmdfaWQiOiI5MDY2NjEiLCJjb2duaXRvOnVzZXJuYW1lIjoibi5nb2VsQGVwaWxvdC5jbG91ZCIsImN1c3RvbTppdnlfdXNlcl9pZCI6IjExMDAwMDM3Iiwib3JpZ2luX2p0aSI6IjgyOTJhNDlmLWRlODUtNDBmZS04ZDMyLTY3MjJhZDg1YWQwYiIsImF1ZCI6IjdwcWg1Zm41amtrNWxyZmlqcG52ZGlvcGQ0IiwiZXZlbnRfaWQiOiI4NzMyNGYzNy1kOTI1LTQ5YTItYjRjNi00MWRhMTU4MzU4ZjYiLCJ0b2tlbl91c2UiOiJpZCIsImF1dGhfdGltZSI6MTc1ODEyNjA1OSwiZXhwIjoxNzU4MTI5NjU5LCJpYXQiOjE3NTgxMjYwNTksImp0aSI6ImMzNmUyNTg5LTFkYTItNGYxOC1iNTM5LWQzYTc2ZTIzY2FkYSIsImVtYWlsIjoibi5nb2VsQGVwaWxvdC5jbG91ZCJ9.u5IHmhe6eqCrvsNFjZR80UJfIFmgbHVASnM1nhOkrksXF9OcHTO3SBuS73OO7XfAoHaeyN3crItGUbcTmYZakYvyA3ogxsq2M3k2gbfPM30Z1slETC6R9mrtpn9gBSZD9yzOhgZrlbP1ltXbrV46-tiptoioa44GTEk-DuEkKplXUEyxqse2svLrU_fzhE1cOHXF9Gzb4JrlxA-RbAa9R-0B0ORJ3gXGO-X8772kkDVGvyh9YlJfqmCvdvdlzUiu2l59FkbLW-cQNxZ_NCov56if_mk2wGVJWzoiHx4VsxyF-Km2NbHVz0hROGswt7pfDwiS4fOjUuVdCJx_eFBieg"
  server_url = "https://entity.dev.sls.epilot.io"
}


# import {
#   to = epilot-taxonomy_taxonomy_classification.my_label
#   id = "joe-test-family:nishu-blueprint-label"
# }