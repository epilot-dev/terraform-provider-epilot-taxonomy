resource "epilot-taxonomy_taxonomy_classification" "my_taxonomyclassification" {
  id = "taxonomy-slug:classification-slug"
  manifest = [
    "123e4567-e89b-12d3-a456-426614174000"
  ]
  name = "Wallbox PV"
  parents = [
    "taxonomy-slug:classification-slug"
  ]
  slug = "wallbox-pv"
}