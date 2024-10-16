// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/internal/utils"
)

// InternalUserAttributeConstraints - A set of constraints applicable to the attribute.
// These constraints should and will be enforced by the attribute renderer.
type InternalUserAttributeConstraints struct {
}

// InternalUserAttributeInfoHelpers - A set of configurations meant to document and assist the user in filling the attribute.
type InternalUserAttributeInfoHelpers struct {
	// The name of the custom component to be used as the hint helper.
	// The component should be registered in the `@epilot360/entity-ui` on the index of the components directory.
	// When specified it overrides the `hint_text` or `hint_text_key` configuration.
	//
	HintCustomComponent *string `json:"hint_custom_component,omitempty"`
	// The text to be displayed in the attribute hint helper.
	// When specified it overrides the `hint_text_key` configuration.
	//
	HintText *string `json:"hint_text,omitempty"`
	// The key of the hint text to be displayed in the attribute hint helper.
	// The key should be a valid i18n key.
	//
	HintTextKey *string `json:"hint_text_key,omitempty"`
	// The placement of the hint tooltip.
	// The value should be a valid `@mui/core` tooltip placement.
	//
	HintTooltipPlacement *string `json:"hint_tooltip_placement,omitempty"`
}

func (o *InternalUserAttributeInfoHelpers) GetHintCustomComponent() *string {
	if o == nil {
		return nil
	}
	return o.HintCustomComponent
}

func (o *InternalUserAttributeInfoHelpers) GetHintText() *string {
	if o == nil {
		return nil
	}
	return o.HintText
}

func (o *InternalUserAttributeInfoHelpers) GetHintTextKey() *string {
	if o == nil {
		return nil
	}
	return o.HintTextKey
}

func (o *InternalUserAttributeInfoHelpers) GetHintTooltipPlacement() *string {
	if o == nil {
		return nil
	}
	return o.HintTooltipPlacement
}

type InternalUserAttributeType string

const (
	InternalUserAttributeTypeInternalUser InternalUserAttributeType = "internal_user"
)

func (e InternalUserAttributeType) ToPointer() *InternalUserAttributeType {
	return &e
}
func (e *InternalUserAttributeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "internal_user":
		*e = InternalUserAttributeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InternalUserAttributeType: %v", v)
	}
}

// InternalUserAttribute - Epilot internal user info
type InternalUserAttribute struct {
	// Manifest ID used to create/update the schema attribute
	Manifest []string `json:"_manifest,omitempty"`
	Purpose  []string `json:"_purpose,omitempty"`
	// A set of constraints applicable to the attribute.
	// These constraints should and will be enforced by the attribute renderer.
	//
	Constraints  *InternalUserAttributeConstraints `json:"constraints,omitempty"`
	DefaultValue any                               `json:"default_value,omitempty"`
	Deprecated   *bool                             `default:"false" json:"deprecated"`
	// Setting to `true` disables editing the attribute on the entity builder UI
	EntityBuilderDisableEdit *bool `default:"false" json:"entity_builder_disable_edit"`
	// This attribute should only be active when the feature flag is enabled
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// Which group the attribute should appear in. Accepts group ID or group name
	Group *string `json:"group,omitempty"`
	// Do not render attribute in entity views
	Hidden *bool `default:"false" json:"hidden"`
	// When set to true, will hide the label of the field.
	HideLabel *bool `json:"hide_label,omitempty"`
	// Code name of the icon to used to represent this attribute.
	// The value must be a valid @epilot/base-elements Icon name
	//
	Icon *string `json:"icon,omitempty"`
	// ID for the entity attribute
	ID *string `json:"id,omitempty"`
	// A set of configurations meant to document and assist the user in filling the attribute.
	InfoHelpers *InternalUserAttributeInfoHelpers `json:"info_helpers,omitempty"`
	Label       string                            `json:"label"`
	Layout      *string                           `json:"layout,omitempty"`
	Name        string                            `json:"name"`
	// Attribute sort order (ascending) in group
	Order                 *int64  `json:"order,omitempty"`
	Placeholder           *string `json:"placeholder,omitempty"`
	PreviewValueFormatter *string `json:"preview_value_formatter,omitempty"`
	// Setting to `true` prevents the attribute from being modified / deleted
	Protected *bool `json:"protected,omitempty"`
	Readonly  *bool `default:"false" json:"readonly"`
	// Defines the conditional rendering expression for showing this field.
	// When a valid expression is parsed, their evaluation defines the visibility of this attribute.
	// Note: Empty or invalid expression have no effect on the field visibility.
	//
	RenderCondition *string `json:"render_condition,omitempty"`
	Required        *bool   `default:"false" json:"required"`
	// This attribute should only be active when one of the provided settings have the correct value
	SettingsFlag []SettingFlag `json:"settings_flag,omitempty"`
	// Render as a column in table views. When defined, overrides `hidden`
	ShowInTable *bool `json:"show_in_table,omitempty"`
	// Allow sorting by this attribute in table views if `show_in_table` is true
	Sortable       *bool                      `default:"true" json:"sortable"`
	Type           *InternalUserAttributeType `json:"type,omitempty"`
	ValueFormatter *string                    `json:"value_formatter,omitempty"`
}

func (i InternalUserAttribute) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InternalUserAttribute) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *InternalUserAttribute) GetManifest() []string {
	if o == nil {
		return nil
	}
	return o.Manifest
}

func (o *InternalUserAttribute) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *InternalUserAttribute) GetConstraints() *InternalUserAttributeConstraints {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *InternalUserAttribute) GetDefaultValue() any {
	if o == nil {
		return nil
	}
	return o.DefaultValue
}

func (o *InternalUserAttribute) GetDeprecated() *bool {
	if o == nil {
		return nil
	}
	return o.Deprecated
}

func (o *InternalUserAttribute) GetEntityBuilderDisableEdit() *bool {
	if o == nil {
		return nil
	}
	return o.EntityBuilderDisableEdit
}

func (o *InternalUserAttribute) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *InternalUserAttribute) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *InternalUserAttribute) GetHidden() *bool {
	if o == nil {
		return nil
	}
	return o.Hidden
}

func (o *InternalUserAttribute) GetHideLabel() *bool {
	if o == nil {
		return nil
	}
	return o.HideLabel
}

func (o *InternalUserAttribute) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *InternalUserAttribute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InternalUserAttribute) GetInfoHelpers() *InternalUserAttributeInfoHelpers {
	if o == nil {
		return nil
	}
	return o.InfoHelpers
}

func (o *InternalUserAttribute) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *InternalUserAttribute) GetLayout() *string {
	if o == nil {
		return nil
	}
	return o.Layout
}

func (o *InternalUserAttribute) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InternalUserAttribute) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *InternalUserAttribute) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *InternalUserAttribute) GetPreviewValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.PreviewValueFormatter
}

func (o *InternalUserAttribute) GetProtected() *bool {
	if o == nil {
		return nil
	}
	return o.Protected
}

func (o *InternalUserAttribute) GetReadonly() *bool {
	if o == nil {
		return nil
	}
	return o.Readonly
}

func (o *InternalUserAttribute) GetRenderCondition() *string {
	if o == nil {
		return nil
	}
	return o.RenderCondition
}

func (o *InternalUserAttribute) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *InternalUserAttribute) GetSettingsFlag() []SettingFlag {
	if o == nil {
		return nil
	}
	return o.SettingsFlag
}

func (o *InternalUserAttribute) GetShowInTable() *bool {
	if o == nil {
		return nil
	}
	return o.ShowInTable
}

func (o *InternalUserAttribute) GetSortable() *bool {
	if o == nil {
		return nil
	}
	return o.Sortable
}

func (o *InternalUserAttribute) GetType() *InternalUserAttributeType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *InternalUserAttribute) GetValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.ValueFormatter
}
