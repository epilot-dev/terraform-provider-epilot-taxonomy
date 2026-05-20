resource "epilot-taxonomy_taxonomy" "my_taxonomy" {
  color   = "#FF5733"
  enabled = true
  enabled_locations = [
    "..."
  ]
  icon      = "purpose"
  name      = "Purpose"
  order     = 10
  permanent = false
  plural    = "Purposes"
  slug      = "purpose"
  type      = "entity"
}