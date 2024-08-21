---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "epilot-taxonomy_taxonomy_classification Data Source - terraform-provider-epilot-taxonomy"
subcategory: ""
description: |-
  TaxonomyClassification DataSource
---

# epilot-taxonomy_taxonomy_classification (Data Source)

TaxonomyClassification DataSource

## Example Usage

```terraform
data "epilot-taxonomy_taxonomy_classification" "my_taxonomyclassification" {
  classification_slug = "purpose:<name>"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `created_at` (String)
- `id` (String) The ID of this resource.
- `name` (String)
- `parents` (List of String)
- `slug` (String) URL-friendly identifier for the classification
- `updated_at` (String)

