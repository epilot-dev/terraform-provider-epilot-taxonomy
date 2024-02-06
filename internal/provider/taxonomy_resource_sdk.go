// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/pkg/models/operations"
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *TaxonomyResourceModel) ToSharedClassificationsUpdate() *shared.ClassificationsUpdate {
	var create []shared.TaxonomyClassification = nil
	for _, createItem := range r.Create {
		createdAt := new(time.Time)
		if !createItem.CreatedAt.IsUnknown() && !createItem.CreatedAt.IsNull() {
			*createdAt, _ = time.Parse(time.RFC3339Nano, createItem.CreatedAt.ValueString())
		} else {
			createdAt = nil
		}
		id := new(string)
		if !createItem.ID.IsUnknown() && !createItem.ID.IsNull() {
			*id = createItem.ID.ValueString()
		} else {
			id = nil
		}
		name := createItem.Name.ValueString()
		var parents []string = nil
		for _, parentsItem := range createItem.Parents {
			parents = append(parents, parentsItem.ValueString())
		}
		updatedAt := new(time.Time)
		if !createItem.UpdatedAt.IsUnknown() && !createItem.UpdatedAt.IsNull() {
			*updatedAt, _ = time.Parse(time.RFC3339Nano, createItem.UpdatedAt.ValueString())
		} else {
			updatedAt = nil
		}
		create = append(create, shared.TaxonomyClassification{
			CreatedAt: createdAt,
			ID:        id,
			Name:      name,
			Parents:   parents,
			UpdatedAt: updatedAt,
		})
	}
	var delete []string = nil
	for _, deleteItem := range r.Delete {
		delete = append(delete, deleteItem.ValueString())
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	name1 := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name1 = r.Name.ValueString()
	} else {
		name1 = nil
	}
	var update []shared.TaxonomyClassification = nil
	for _, updateItem := range r.Update {
		createdAt1 := new(time.Time)
		if !updateItem.CreatedAt.IsUnknown() && !updateItem.CreatedAt.IsNull() {
			*createdAt1, _ = time.Parse(time.RFC3339Nano, updateItem.CreatedAt.ValueString())
		} else {
			createdAt1 = nil
		}
		id1 := new(string)
		if !updateItem.ID.IsUnknown() && !updateItem.ID.IsNull() {
			*id1 = updateItem.ID.ValueString()
		} else {
			id1 = nil
		}
		name2 := updateItem.Name.ValueString()
		var parents1 []string = nil
		for _, parentsItem1 := range updateItem.Parents {
			parents1 = append(parents1, parentsItem1.ValueString())
		}
		updatedAt1 := new(time.Time)
		if !updateItem.UpdatedAt.IsUnknown() && !updateItem.UpdatedAt.IsNull() {
			*updatedAt1, _ = time.Parse(time.RFC3339Nano, updateItem.UpdatedAt.ValueString())
		} else {
			updatedAt1 = nil
		}
		update = append(update, shared.TaxonomyClassification{
			CreatedAt: createdAt1,
			ID:        id1,
			Name:      name2,
			Parents:   parents1,
			UpdatedAt: updatedAt1,
		})
	}
	out := shared.ClassificationsUpdate{
		Create:  create,
		Delete:  delete,
		Enabled: enabled,
		Name:    name1,
		Update:  update,
	}
	return &out
}

func (r *TaxonomyResourceModel) RefreshFromOperationsUpdateClassificationsForTaxonomyResponseBody(resp *operations.UpdateClassificationsForTaxonomyResponseBody) {
	if len(r.Created) > len(resp.Created) {
		r.Created = r.Created[:len(resp.Created)]
	}
	for createdCount, createdItem := range resp.Created {
		var created1 TaxonomyClassification
		if createdItem.CreatedAt != nil {
			created1.CreatedAt = types.StringValue(createdItem.CreatedAt.Format(time.RFC3339Nano))
		} else {
			created1.CreatedAt = types.StringNull()
		}
		created1.ID = types.StringPointerValue(createdItem.ID)
		created1.Name = types.StringValue(createdItem.Name)
		created1.Parents = nil
		for _, v := range createdItem.Parents {
			created1.Parents = append(created1.Parents, types.StringValue(v))
		}
		if createdItem.UpdatedAt != nil {
			created1.UpdatedAt = types.StringValue(createdItem.UpdatedAt.Format(time.RFC3339Nano))
		} else {
			created1.UpdatedAt = types.StringNull()
		}
		if createdCount+1 > len(r.Created) {
			r.Created = append(r.Created, created1)
		} else {
			r.Created[createdCount].CreatedAt = created1.CreatedAt
			r.Created[createdCount].ID = created1.ID
			r.Created[createdCount].Name = created1.Name
			r.Created[createdCount].Parents = created1.Parents
			r.Created[createdCount].UpdatedAt = created1.UpdatedAt
		}
	}
	r.Deleted = nil
	for _, deletedItem := range resp.Deleted {
		var deleted1 types.String
		deleted1Result, _ := json.Marshal(deletedItem)
		deleted1 = types.StringValue(string(deleted1Result))
		r.Deleted = append(r.Deleted, deleted1)
	}
	if len(r.Updated) > len(resp.Updated) {
		r.Updated = r.Updated[:len(resp.Updated)]
	}
	for updatedCount, updatedItem := range resp.Updated {
		var updated1 TaxonomyClassification
		if updatedItem.CreatedAt != nil {
			updated1.CreatedAt = types.StringValue(updatedItem.CreatedAt.Format(time.RFC3339Nano))
		} else {
			updated1.CreatedAt = types.StringNull()
		}
		updated1.ID = types.StringPointerValue(updatedItem.ID)
		updated1.Name = types.StringValue(updatedItem.Name)
		updated1.Parents = nil
		for _, v := range updatedItem.Parents {
			updated1.Parents = append(updated1.Parents, types.StringValue(v))
		}
		if updatedItem.UpdatedAt != nil {
			updated1.UpdatedAt = types.StringValue(updatedItem.UpdatedAt.Format(time.RFC3339Nano))
		} else {
			updated1.UpdatedAt = types.StringNull()
		}
		if updatedCount+1 > len(r.Updated) {
			r.Updated = append(r.Updated, updated1)
		} else {
			r.Updated[updatedCount].CreatedAt = updated1.CreatedAt
			r.Updated[updatedCount].ID = updated1.ID
			r.Updated[updatedCount].Name = updated1.Name
			r.Updated[updatedCount].Parents = updated1.Parents
			r.Updated[updatedCount].UpdatedAt = updated1.UpdatedAt
		}
	}
}
