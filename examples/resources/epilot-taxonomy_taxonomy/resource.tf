resource "epilot-taxonomy_taxonomy" "my_taxonomy" {
  create = [
    {
      created_at = "2022-11-01T07:59:50.446Z"
      id         = "693ca07e-87a4-435b-b803-31eeb618b4f0"
      name       = "Wallbox PV"
      parents = [
        "375edde2-a483-43e1-b5cc-bd4e9e6dcc1a",
      ]
      updated_at = "2022-06-30T12:03:25.687Z"
    },
  ]
  delete = [
    "82cde8ed-628d-4119-b60a-1c728a071b1b",
  ]
  enabled = false
  name    = "Kent Schaden"
  slug    = "...my_slug..."
  update = [
    {
      created_at = "2022-03-09T09:06:58.643Z"
      id         = "8ad44015-2b92-484c-939c-60b7b233d8f8"
      name       = "Wallbox PV"
      parents = [
        "5546a98b-9dea-423a-a1ab-e6c7452af4ee",
      ]
      updated_at = "2022-12-02T16:32:24.955Z"
    },
  ]
}