// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *TaxonomyClassificationResourceModel) ToSharedTaxonomyClassificationInput() *shared.TaxonomyClassificationInput {
	var manifest []string = []string{}
	for _, manifestItem := range r.Manifest {
		manifest = append(manifest, manifestItem.ValueString())
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	var name string
	name = r.Name.ValueString()

	var parents []string = []string{}
	for _, parentsItem := range r.Parents {
		parents = append(parents, parentsItem.ValueString())
	}
	slug := new(string)
	if !r.Slug.IsUnknown() && !r.Slug.IsNull() {
		*slug = r.Slug.ValueString()
	} else {
		slug = nil
	}
	out := shared.TaxonomyClassificationInput{
		Manifest: manifest,
		ID:       id,
		Name:     name,
		Parents:  parents,
		Slug:     slug,
	}
	return &out
}

func (r *TaxonomyClassificationResourceModel) RefreshFromSharedTaxonomyClassification(resp *shared.TaxonomyClassification) {
	if resp != nil {
		r.Manifest = []types.String{}
		for _, v := range resp.Manifest {
			r.Manifest = append(r.Manifest, types.StringValue(v))
		}
		if resp.CreatedAt != nil {
			r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		} else {
			r.CreatedAt = types.StringNull()
		}
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringValue(resp.Name)
		r.Parents = []types.String{}
		for _, v := range resp.Parents {
			r.Parents = append(r.Parents, types.StringValue(v))
		}
		r.Slug = types.StringPointerValue(resp.Slug)
		if resp.UpdatedAt != nil {
			r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		} else {
			r.UpdatedAt = types.StringNull()
		}
	}
}
