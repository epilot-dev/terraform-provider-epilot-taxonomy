// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/internal/utils"
	"time"
)

type TaxonomyClassification struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ID        *string    `json:"id,omitempty"`
	Name      string     `json:"name"`
	Parents   []string   `json:"parents,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (t TaxonomyClassification) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaxonomyClassification) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaxonomyClassification) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TaxonomyClassification) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TaxonomyClassification) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *TaxonomyClassification) GetParents() []string {
	if o == nil {
		return nil
	}
	return o.Parents
}

func (o *TaxonomyClassification) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}