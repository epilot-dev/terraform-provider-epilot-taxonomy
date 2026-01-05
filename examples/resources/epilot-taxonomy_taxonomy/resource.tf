resource "epilot-taxonomy_taxonomy" "my_taxonomy" {
  color   = "#FF5733"
  enabled = true
  enabled_locations = [
    {
      str                  = "...my_str..."
      taxonomy_location_id = "account"
    }
  ]
  icon      = "purpose"
  name      = "Purpose"
  order     = 10
  permanent = true
  plural    = "Purposes"
  slug      = "purpose"
  type      = "entity"
}