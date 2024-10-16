// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/internal/utils"
)

type GroupHeadlineWithCompositeIDDivider string

const (
	GroupHeadlineWithCompositeIDDividerTopDivider    GroupHeadlineWithCompositeIDDivider = "top_divider"
	GroupHeadlineWithCompositeIDDividerBottomDivider GroupHeadlineWithCompositeIDDivider = "bottom_divider"
)

func (e GroupHeadlineWithCompositeIDDivider) ToPointer() *GroupHeadlineWithCompositeIDDivider {
	return &e
}
func (e *GroupHeadlineWithCompositeIDDivider) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "top_divider":
		fallthrough
	case "bottom_divider":
		*e = GroupHeadlineWithCompositeIDDivider(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GroupHeadlineWithCompositeIDDivider: %v", v)
	}
}

type GroupHeadlineWithCompositeIDType string

const (
	GroupHeadlineWithCompositeIDTypeHeadline GroupHeadlineWithCompositeIDType = "headline"
)

func (e GroupHeadlineWithCompositeIDType) ToPointer() *GroupHeadlineWithCompositeIDType {
	return &e
}
func (e *GroupHeadlineWithCompositeIDType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "headline":
		*e = GroupHeadlineWithCompositeIDType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GroupHeadlineWithCompositeIDType: %v", v)
	}
}

// GroupHeadlineWithCompositeID - a readonly computed ID for the entity group headline including schema slug and the headline ID
type GroupHeadlineWithCompositeID struct {
	// Manifest ID used to create/update the schema group headline
	Manifest      []string                             `json:"_manifest,omitempty"`
	CompositeID   *string                              `json:"composite_id,omitempty"`
	Divider       *GroupHeadlineWithCompositeIDDivider `json:"divider,omitempty"`
	EnableDivider *bool                                `default:"false" json:"enable_divider"`
	// The group of headline attribute
	Group  string  `json:"group"`
	ID     *string `json:"id,omitempty"`
	Label  string  `json:"label"`
	Layout *string `json:"layout,omitempty"`
	Name   string  `json:"name"`
	// The order of headline attribute
	Order *int64 `json:"order,omitempty"`
	// Schema slug the capability belongs to
	Schema *string                          `json:"schema,omitempty"`
	Type   GroupHeadlineWithCompositeIDType `json:"type"`
}

func (g GroupHeadlineWithCompositeID) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GroupHeadlineWithCompositeID) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GroupHeadlineWithCompositeID) GetManifest() []string {
	if o == nil {
		return nil
	}
	return o.Manifest
}

func (o *GroupHeadlineWithCompositeID) GetCompositeID() *string {
	if o == nil {
		return nil
	}
	return o.CompositeID
}

func (o *GroupHeadlineWithCompositeID) GetDivider() *GroupHeadlineWithCompositeIDDivider {
	if o == nil {
		return nil
	}
	return o.Divider
}

func (o *GroupHeadlineWithCompositeID) GetEnableDivider() *bool {
	if o == nil {
		return nil
	}
	return o.EnableDivider
}

func (o *GroupHeadlineWithCompositeID) GetGroup() string {
	if o == nil {
		return ""
	}
	return o.Group
}

func (o *GroupHeadlineWithCompositeID) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GroupHeadlineWithCompositeID) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *GroupHeadlineWithCompositeID) GetLayout() *string {
	if o == nil {
		return nil
	}
	return o.Layout
}

func (o *GroupHeadlineWithCompositeID) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GroupHeadlineWithCompositeID) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GroupHeadlineWithCompositeID) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *GroupHeadlineWithCompositeID) GetType() GroupHeadlineWithCompositeIDType {
	if o == nil {
		return GroupHeadlineWithCompositeIDType("")
	}
	return o.Type
}

// GroupHeadlineWithCompositeIDInput - a readonly computed ID for the entity group headline including schema slug and the headline ID
type GroupHeadlineWithCompositeIDInput struct {
	// Manifest ID used to create/update the schema group headline
	Manifest      []string                             `json:"_manifest,omitempty"`
	Divider       *GroupHeadlineWithCompositeIDDivider `json:"divider,omitempty"`
	EnableDivider *bool                                `default:"false" json:"enable_divider"`
	// The group of headline attribute
	Group  string  `json:"group"`
	ID     *string `json:"id,omitempty"`
	Label  string  `json:"label"`
	Layout *string `json:"layout,omitempty"`
	Name   string  `json:"name"`
	// The order of headline attribute
	Order *int64 `json:"order,omitempty"`
	// Schema slug the capability belongs to
	Schema *string                          `json:"schema,omitempty"`
	Type   GroupHeadlineWithCompositeIDType `json:"type"`
}

func (g GroupHeadlineWithCompositeIDInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GroupHeadlineWithCompositeIDInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GroupHeadlineWithCompositeIDInput) GetManifest() []string {
	if o == nil {
		return nil
	}
	return o.Manifest
}

func (o *GroupHeadlineWithCompositeIDInput) GetDivider() *GroupHeadlineWithCompositeIDDivider {
	if o == nil {
		return nil
	}
	return o.Divider
}

func (o *GroupHeadlineWithCompositeIDInput) GetEnableDivider() *bool {
	if o == nil {
		return nil
	}
	return o.EnableDivider
}

func (o *GroupHeadlineWithCompositeIDInput) GetGroup() string {
	if o == nil {
		return ""
	}
	return o.Group
}

func (o *GroupHeadlineWithCompositeIDInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GroupHeadlineWithCompositeIDInput) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *GroupHeadlineWithCompositeIDInput) GetLayout() *string {
	if o == nil {
		return nil
	}
	return o.Layout
}

func (o *GroupHeadlineWithCompositeIDInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GroupHeadlineWithCompositeIDInput) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GroupHeadlineWithCompositeIDInput) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *GroupHeadlineWithCompositeIDInput) GetType() GroupHeadlineWithCompositeIDType {
	if o == nil {
		return GroupHeadlineWithCompositeIDType("")
	}
	return o.Type
}