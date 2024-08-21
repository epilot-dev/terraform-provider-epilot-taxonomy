// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/internal/utils"
)

type EntitySchemaGroupWithCompositeIDInfoTooltipTitle struct {
	// Default string for info tooltip
	Default *string `json:"default,omitempty"`
	// Translation key for info tooltip
	Key *string `json:"key,omitempty"`
}

func (o *EntitySchemaGroupWithCompositeIDInfoTooltipTitle) GetDefault() *string {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *EntitySchemaGroupWithCompositeIDInfoTooltipTitle) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

// EntitySchemaGroupWithCompositeID - a readonly computed ID for the group including schema slug and the group ID
type EntitySchemaGroupWithCompositeID struct {
	// Only render group when one of the purposes is enabled
	Purpose     []string `json:"_purpose,omitempty"`
	CompositeID *string  `json:"composite_id,omitempty"`
	// Expanded by default
	Expanded *bool `default:"false" json:"expanded"`
	// This group should only be active when the feature flag is enabled
	FeatureFlag      *string                                           `json:"feature_flag,omitempty"`
	ID               string                                            `json:"id"`
	InfoTooltipTitle *EntitySchemaGroupWithCompositeIDInfoTooltipTitle `json:"info_tooltip_title,omitempty"`
	Label            string                                            `json:"label"`
	// Render order of the group
	Order *int64 `default:"0" json:"order"`
	// Only render group when render_condition resolves to true
	RenderCondition *string `json:"render_condition,omitempty"`
	// This group should only be active when all the settings have the correct value
	SettingsFlag []SettingFlag `json:"settings_flag,omitempty"`
}

func (e EntitySchemaGroupWithCompositeID) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EntitySchemaGroupWithCompositeID) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *EntitySchemaGroupWithCompositeID) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *EntitySchemaGroupWithCompositeID) GetCompositeID() *string {
	if o == nil {
		return nil
	}
	return o.CompositeID
}

func (o *EntitySchemaGroupWithCompositeID) GetExpanded() *bool {
	if o == nil {
		return nil
	}
	return o.Expanded
}

func (o *EntitySchemaGroupWithCompositeID) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *EntitySchemaGroupWithCompositeID) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *EntitySchemaGroupWithCompositeID) GetInfoTooltipTitle() *EntitySchemaGroupWithCompositeIDInfoTooltipTitle {
	if o == nil {
		return nil
	}
	return o.InfoTooltipTitle
}

func (o *EntitySchemaGroupWithCompositeID) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *EntitySchemaGroupWithCompositeID) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *EntitySchemaGroupWithCompositeID) GetRenderCondition() *string {
	if o == nil {
		return nil
	}
	return o.RenderCondition
}

func (o *EntitySchemaGroupWithCompositeID) GetSettingsFlag() []SettingFlag {
	if o == nil {
		return nil
	}
	return o.SettingsFlag
}

// EntitySchemaGroupWithCompositeIDInput - a readonly computed ID for the group including schema slug and the group ID
type EntitySchemaGroupWithCompositeIDInput struct {
	// Only render group when one of the purposes is enabled
	Purpose []string `json:"_purpose,omitempty"`
	// Expanded by default
	Expanded *bool `default:"false" json:"expanded"`
	// This group should only be active when the feature flag is enabled
	FeatureFlag      *string                                           `json:"feature_flag,omitempty"`
	ID               string                                            `json:"id"`
	InfoTooltipTitle *EntitySchemaGroupWithCompositeIDInfoTooltipTitle `json:"info_tooltip_title,omitempty"`
	Label            string                                            `json:"label"`
	// Render order of the group
	Order *int64 `default:"0" json:"order"`
	// Only render group when render_condition resolves to true
	RenderCondition *string `json:"render_condition,omitempty"`
	// This group should only be active when all the settings have the correct value
	SettingsFlag []SettingFlag `json:"settings_flag,omitempty"`
}

func (e EntitySchemaGroupWithCompositeIDInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EntitySchemaGroupWithCompositeIDInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *EntitySchemaGroupWithCompositeIDInput) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *EntitySchemaGroupWithCompositeIDInput) GetExpanded() *bool {
	if o == nil {
		return nil
	}
	return o.Expanded
}

func (o *EntitySchemaGroupWithCompositeIDInput) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *EntitySchemaGroupWithCompositeIDInput) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *EntitySchemaGroupWithCompositeIDInput) GetInfoTooltipTitle() *EntitySchemaGroupWithCompositeIDInfoTooltipTitle {
	if o == nil {
		return nil
	}
	return o.InfoTooltipTitle
}

func (o *EntitySchemaGroupWithCompositeIDInput) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *EntitySchemaGroupWithCompositeIDInput) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *EntitySchemaGroupWithCompositeIDInput) GetRenderCondition() *string {
	if o == nil {
		return nil
	}
	return o.RenderCondition
}

func (o *EntitySchemaGroupWithCompositeIDInput) GetSettingsFlag() []SettingFlag {
	if o == nil {
		return nil
	}
	return o.SettingsFlag
}