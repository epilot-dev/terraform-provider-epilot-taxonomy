// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/internal/utils"
)

// MultiSelectAttributeConstraints - A set of constraints applicable to the attribute.
// These constraints should and will be enforced by the attribute renderer.
type MultiSelectAttributeConstraints struct {
}

// MultiSelectAttributeInfoHelpers - A set of configurations meant to document and assist the user in filling the attribute.
type MultiSelectAttributeInfoHelpers struct {
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

func (o *MultiSelectAttributeInfoHelpers) GetHintCustomComponent() *string {
	if o == nil {
		return nil
	}
	return o.HintCustomComponent
}

func (o *MultiSelectAttributeInfoHelpers) GetHintText() *string {
	if o == nil {
		return nil
	}
	return o.HintText
}

func (o *MultiSelectAttributeInfoHelpers) GetHintTextKey() *string {
	if o == nil {
		return nil
	}
	return o.HintTextKey
}

func (o *MultiSelectAttributeInfoHelpers) GetHintTooltipPlacement() *string {
	if o == nil {
		return nil
	}
	return o.HintTooltipPlacement
}

type MultiSelectAttribute2 struct {
	Title *string `json:"title,omitempty"`
	Value string  `json:"value"`
}

func (o *MultiSelectAttribute2) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *MultiSelectAttribute2) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type MultiSelectAttributeOptionsType string

const (
	MultiSelectAttributeOptionsTypeStr                   MultiSelectAttributeOptionsType = "str"
	MultiSelectAttributeOptionsTypeMultiSelectAttribute2 MultiSelectAttributeOptionsType = "MultiSelectAttribute_2"
)

type MultiSelectAttributeOptions struct {
	Str                   *string
	MultiSelectAttribute2 *MultiSelectAttribute2

	Type MultiSelectAttributeOptionsType
}

func CreateMultiSelectAttributeOptionsStr(str string) MultiSelectAttributeOptions {
	typ := MultiSelectAttributeOptionsTypeStr

	return MultiSelectAttributeOptions{
		Str:  &str,
		Type: typ,
	}
}

func CreateMultiSelectAttributeOptionsMultiSelectAttribute2(multiSelectAttribute2 MultiSelectAttribute2) MultiSelectAttributeOptions {
	typ := MultiSelectAttributeOptionsTypeMultiSelectAttribute2

	return MultiSelectAttributeOptions{
		MultiSelectAttribute2: &multiSelectAttribute2,
		Type:                  typ,
	}
}

func (u *MultiSelectAttributeOptions) UnmarshalJSON(data []byte) error {

	var multiSelectAttribute2 MultiSelectAttribute2 = MultiSelectAttribute2{}
	if err := utils.UnmarshalJSON(data, &multiSelectAttribute2, "", true, true); err == nil {
		u.MultiSelectAttribute2 = &multiSelectAttribute2
		u.Type = MultiSelectAttributeOptionsTypeMultiSelectAttribute2
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = MultiSelectAttributeOptionsTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for MultiSelectAttributeOptions", string(data))
}

func (u MultiSelectAttributeOptions) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.MultiSelectAttribute2 != nil {
		return utils.MarshalJSON(u.MultiSelectAttribute2, "", true)
	}

	return nil, errors.New("could not marshal union type MultiSelectAttributeOptions: all fields are null")
}

type MultiSelectAttributeType string

const (
	MultiSelectAttributeTypeMultiselect MultiSelectAttributeType = "multiselect"
	MultiSelectAttributeTypeCheckbox    MultiSelectAttributeType = "checkbox"
)

func (e MultiSelectAttributeType) ToPointer() *MultiSelectAttributeType {
	return &e
}
func (e *MultiSelectAttributeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "multiselect":
		fallthrough
	case "checkbox":
		*e = MultiSelectAttributeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MultiSelectAttributeType: %v", v)
	}
}

// MultiSelectAttribute - Multi Choice Selection
type MultiSelectAttribute struct {
	Purpose []string `json:"_purpose,omitempty"`
	// Allow arbitrary input values in addition to provided options
	AllowAny *bool `json:"allow_any,omitempty"`
	// controls if the 360 ui will allow the user to enter a value which is not defined by the options
	AllowExtraOptions *bool `json:"allow_extra_options,omitempty"`
	// A set of constraints applicable to the attribute.
	// These constraints should and will be enforced by the attribute renderer.
	//
	Constraints  *MultiSelectAttributeConstraints `json:"constraints,omitempty"`
	DefaultValue any                              `json:"default_value,omitempty"`
	Deprecated   *bool                            `default:"false" json:"deprecated"`
	// controls if the matching of values against the options is case sensitive or not
	DisableCaseSensitive *bool `json:"disable_case_sensitive,omitempty"`
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
	InfoHelpers *MultiSelectAttributeInfoHelpers `json:"info_helpers,omitempty"`
	Label       string                           `json:"label"`
	Layout      *string                          `json:"layout,omitempty"`
	Name        string                           `json:"name"`
	Options     []MultiSelectAttributeOptions    `json:"options,omitempty"`
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
	// This attribute should only be active when all the settings have the correct value
	SettingsFlag []SettingFlag `json:"settings_flag,omitempty"`
	// Render as a column in table views. When defined, overrides `hidden`
	ShowInTable *bool `json:"show_in_table,omitempty"`
	// Allow sorting by this attribute in table views if `show_in_table` is true
	Sortable       *bool                     `default:"true" json:"sortable"`
	Type           *MultiSelectAttributeType `json:"type,omitempty"`
	ValueFormatter *string                   `json:"value_formatter,omitempty"`
}

func (m MultiSelectAttribute) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MultiSelectAttribute) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *MultiSelectAttribute) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *MultiSelectAttribute) GetAllowAny() *bool {
	if o == nil {
		return nil
	}
	return o.AllowAny
}

func (o *MultiSelectAttribute) GetAllowExtraOptions() *bool {
	if o == nil {
		return nil
	}
	return o.AllowExtraOptions
}

func (o *MultiSelectAttribute) GetConstraints() *MultiSelectAttributeConstraints {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *MultiSelectAttribute) GetDefaultValue() any {
	if o == nil {
		return nil
	}
	return o.DefaultValue
}

func (o *MultiSelectAttribute) GetDeprecated() *bool {
	if o == nil {
		return nil
	}
	return o.Deprecated
}

func (o *MultiSelectAttribute) GetDisableCaseSensitive() *bool {
	if o == nil {
		return nil
	}
	return o.DisableCaseSensitive
}

func (o *MultiSelectAttribute) GetEntityBuilderDisableEdit() *bool {
	if o == nil {
		return nil
	}
	return o.EntityBuilderDisableEdit
}

func (o *MultiSelectAttribute) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *MultiSelectAttribute) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *MultiSelectAttribute) GetHidden() *bool {
	if o == nil {
		return nil
	}
	return o.Hidden
}

func (o *MultiSelectAttribute) GetHideLabel() *bool {
	if o == nil {
		return nil
	}
	return o.HideLabel
}

func (o *MultiSelectAttribute) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *MultiSelectAttribute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *MultiSelectAttribute) GetInfoHelpers() *MultiSelectAttributeInfoHelpers {
	if o == nil {
		return nil
	}
	return o.InfoHelpers
}

func (o *MultiSelectAttribute) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *MultiSelectAttribute) GetLayout() *string {
	if o == nil {
		return nil
	}
	return o.Layout
}

func (o *MultiSelectAttribute) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MultiSelectAttribute) GetOptions() []MultiSelectAttributeOptions {
	if o == nil {
		return nil
	}
	return o.Options
}

func (o *MultiSelectAttribute) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *MultiSelectAttribute) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *MultiSelectAttribute) GetPreviewValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.PreviewValueFormatter
}

func (o *MultiSelectAttribute) GetProtected() *bool {
	if o == nil {
		return nil
	}
	return o.Protected
}

func (o *MultiSelectAttribute) GetReadonly() *bool {
	if o == nil {
		return nil
	}
	return o.Readonly
}

func (o *MultiSelectAttribute) GetRenderCondition() *string {
	if o == nil {
		return nil
	}
	return o.RenderCondition
}

func (o *MultiSelectAttribute) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *MultiSelectAttribute) GetSettingsFlag() []SettingFlag {
	if o == nil {
		return nil
	}
	return o.SettingsFlag
}

func (o *MultiSelectAttribute) GetShowInTable() *bool {
	if o == nil {
		return nil
	}
	return o.ShowInTable
}

func (o *MultiSelectAttribute) GetSortable() *bool {
	if o == nil {
		return nil
	}
	return o.Sortable
}

func (o *MultiSelectAttribute) GetType() *MultiSelectAttributeType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *MultiSelectAttribute) GetValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.ValueFormatter
}