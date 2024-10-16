// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/internal/utils"
	"time"
)

// RelationAttributeActionType - The action type. Currently supported actions:
//
// | action | description |
// |--------|-------------|
// | add_existing | Enables the user to pick an existing entity to link as relation |
// | create_new | Enables the user to create a new entity using the first/main `allowed_schemas` schema
// | create_from_existing | Enables the user to pick an existing entity to clone from, while creating a blank new entity to link as relation |
type RelationAttributeActionType string

const (
	RelationAttributeActionTypeAddExisting        RelationAttributeActionType = "add_existing"
	RelationAttributeActionTypeCreateNew          RelationAttributeActionType = "create_new"
	RelationAttributeActionTypeCreateFromExisting RelationAttributeActionType = "create_from_existing"
)

func (e RelationAttributeActionType) ToPointer() *RelationAttributeActionType {
	return &e
}
func (e *RelationAttributeActionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "add_existing":
		fallthrough
	case "create_new":
		fallthrough
	case "create_from_existing":
		*e = RelationAttributeActionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RelationAttributeActionType: %v", v)
	}
}

type RelationAttributeNewEntityItem struct {
	AdditionalProperties any `additionalProperties:"true" json:"-"`
	// Access control list (ACL) for an entity. Defines sharing access to external orgs or users.
	ACL       *EntityACL `json:"_acl,omitempty"`
	CreatedAt *time.Time `json:"_created_at"`
	ID        string     `json:"_id"`
	// Manifest ID used to create/update the entity
	Manifest []string `json:"_manifest,omitempty"`
	// Organization Id the entity belongs to
	Org     string        `json:"_org"`
	Owners  []EntityOwner `json:"_owners,omitempty"`
	Purpose []string      `json:"_purpose,omitempty"`
	// URL-friendly identifier for the entity schema
	Schema string   `json:"_schema"`
	Tags   []string `json:"_tags,omitempty"`
	// Title of entity
	Title     *string    `json:"_title"`
	UpdatedAt *time.Time `json:"_updated_at"`
}

func (r RelationAttributeNewEntityItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RelationAttributeNewEntityItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RelationAttributeNewEntityItem) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *RelationAttributeNewEntityItem) GetACL() *EntityACL {
	if o == nil {
		return nil
	}
	return o.ACL
}

func (o *RelationAttributeNewEntityItem) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RelationAttributeNewEntityItem) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RelationAttributeNewEntityItem) GetManifest() []string {
	if o == nil {
		return nil
	}
	return o.Manifest
}

func (o *RelationAttributeNewEntityItem) GetOrg() string {
	if o == nil {
		return ""
	}
	return o.Org
}

func (o *RelationAttributeNewEntityItem) GetOwners() []EntityOwner {
	if o == nil {
		return nil
	}
	return o.Owners
}

func (o *RelationAttributeNewEntityItem) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *RelationAttributeNewEntityItem) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *RelationAttributeNewEntityItem) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *RelationAttributeNewEntityItem) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *RelationAttributeNewEntityItem) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

type RelationAttributeActions struct {
	// The action type. Currently supported actions:
	//
	// | action | description |
	// |--------|-------------|
	// | add_existing | Enables the user to pick an existing entity to link as relation |
	// | create_new | Enables the user to create a new entity using the first/main `allowed_schemas` schema
	// | create_from_existing | Enables the user to pick an existing entity to clone from, while creating a blank new entity to link as relation |
	//
	ActionType *RelationAttributeActionType `json:"action_type,omitempty"`
	// Sets the action as the default action, visible as the main action button.
	Default *bool `json:"default,omitempty"`
	// Name of the feature flag that enables this action
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// The action label or action translation key (i18n)
	Label         *string                         `json:"label,omitempty"`
	NewEntityItem *RelationAttributeNewEntityItem `json:"new_entity_item,omitempty"`
	// This action should only be active when all the settings have the correct value
	SettingsFlag []SettingFlag `json:"settings_flag,omitempty"`
}

func (o *RelationAttributeActions) GetActionType() *RelationAttributeActionType {
	if o == nil {
		return nil
	}
	return o.ActionType
}

func (o *RelationAttributeActions) GetDefault() *bool {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *RelationAttributeActions) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *RelationAttributeActions) GetLabel() *string {
	if o == nil {
		return nil
	}
	return o.Label
}

func (o *RelationAttributeActions) GetNewEntityItem() *RelationAttributeNewEntityItem {
	if o == nil {
		return nil
	}
	return o.NewEntityItem
}

func (o *RelationAttributeActions) GetSettingsFlag() []SettingFlag {
	if o == nil {
		return nil
	}
	return o.SettingsFlag
}

// RelationAttributeConstraints - A set of constraints applicable to the attribute.
// These constraints should and will be enforced by the attribute renderer.
type RelationAttributeConstraints struct {
}

type RelationAttributeDrawerSize string

const (
	RelationAttributeDrawerSizeSmall  RelationAttributeDrawerSize = "small"
	RelationAttributeDrawerSizeMedium RelationAttributeDrawerSize = "medium"
	RelationAttributeDrawerSizeLarge  RelationAttributeDrawerSize = "large"
)

func (e RelationAttributeDrawerSize) ToPointer() *RelationAttributeDrawerSize {
	return &e
}
func (e *RelationAttributeDrawerSize) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "small":
		fallthrough
	case "medium":
		fallthrough
	case "large":
		*e = RelationAttributeDrawerSize(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RelationAttributeDrawerSize: %v", v)
	}
}

type RelationAttributeEditMode string

const (
	RelationAttributeEditModeListView RelationAttributeEditMode = "list-view"
)

func (e RelationAttributeEditMode) ToPointer() *RelationAttributeEditMode {
	return &e
}
func (e *RelationAttributeEditMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "list-view":
		*e = RelationAttributeEditMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RelationAttributeEditMode: %v", v)
	}
}

// RelationAttributeInfoHelpers - A set of configurations meant to document and assist the user in filling the attribute.
type RelationAttributeInfoHelpers struct {
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

func (o *RelationAttributeInfoHelpers) GetHintCustomComponent() *string {
	if o == nil {
		return nil
	}
	return o.HintCustomComponent
}

func (o *RelationAttributeInfoHelpers) GetHintText() *string {
	if o == nil {
		return nil
	}
	return o.HintText
}

func (o *RelationAttributeInfoHelpers) GetHintTextKey() *string {
	if o == nil {
		return nil
	}
	return o.HintTextKey
}

func (o *RelationAttributeInfoHelpers) GetHintTooltipPlacement() *string {
	if o == nil {
		return nil
	}
	return o.HintTooltipPlacement
}

// RelationAttributeRelationAffinityMode - Weak relation attributes are kept when duplicating an entity. Strong relation attributes are discarded when duplicating an entity.
type RelationAttributeRelationAffinityMode string

const (
	RelationAttributeRelationAffinityModeWeak   RelationAttributeRelationAffinityMode = "weak"
	RelationAttributeRelationAffinityModeStrong RelationAttributeRelationAffinityMode = "strong"
)

func (e RelationAttributeRelationAffinityMode) ToPointer() *RelationAttributeRelationAffinityMode {
	return &e
}
func (e *RelationAttributeRelationAffinityMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "weak":
		fallthrough
	case "strong":
		*e = RelationAttributeRelationAffinityMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RelationAttributeRelationAffinityMode: %v", v)
	}
}

type RelationAttributeRelationType string

const (
	RelationAttributeRelationTypeHasMany RelationAttributeRelationType = "has_many"
	RelationAttributeRelationTypeHasOne  RelationAttributeRelationType = "has_one"
)

func (e RelationAttributeRelationType) ToPointer() *RelationAttributeRelationType {
	return &e
}
func (e *RelationAttributeRelationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "has_many":
		fallthrough
	case "has_one":
		*e = RelationAttributeRelationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RelationAttributeRelationType: %v", v)
	}
}

type RelationAttributeSummaryFieldsType string

const (
	RelationAttributeSummaryFieldsTypeStr          RelationAttributeSummaryFieldsType = "str"
	RelationAttributeSummaryFieldsTypeSummaryField RelationAttributeSummaryFieldsType = "SummaryField"
)

type RelationAttributeSummaryFields struct {
	Str          *string
	SummaryField *SummaryField

	Type RelationAttributeSummaryFieldsType
}

func CreateRelationAttributeSummaryFieldsStr(str string) RelationAttributeSummaryFields {
	typ := RelationAttributeSummaryFieldsTypeStr

	return RelationAttributeSummaryFields{
		Str:  &str,
		Type: typ,
	}
}

func CreateRelationAttributeSummaryFieldsSummaryField(summaryField SummaryField) RelationAttributeSummaryFields {
	typ := RelationAttributeSummaryFieldsTypeSummaryField

	return RelationAttributeSummaryFields{
		SummaryField: &summaryField,
		Type:         typ,
	}
}

func (u *RelationAttributeSummaryFields) UnmarshalJSON(data []byte) error {

	var summaryField SummaryField = SummaryField{}
	if err := utils.UnmarshalJSON(data, &summaryField, "", true, true); err == nil {
		u.SummaryField = &summaryField
		u.Type = RelationAttributeSummaryFieldsTypeSummaryField
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = RelationAttributeSummaryFieldsTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for RelationAttributeSummaryFields", string(data))
}

func (u RelationAttributeSummaryFields) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.SummaryField != nil {
		return utils.MarshalJSON(u.SummaryField, "", true)
	}

	return nil, errors.New("could not marshal union type RelationAttributeSummaryFields: all fields are null")
}

type RelationAttributeType string

const (
	RelationAttributeTypeRelation RelationAttributeType = "relation"
)

func (e RelationAttributeType) ToPointer() *RelationAttributeType {
	return &e
}
func (e *RelationAttributeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "relation":
		*e = RelationAttributeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RelationAttributeType: %v", v)
	}
}

// RelationAttribute - Entity Relationship
type RelationAttribute struct {
	// Manifest ID used to create/update the schema attribute
	Manifest []string                   `json:"_manifest,omitempty"`
	Purpose  []string                   `json:"_purpose,omitempty"`
	Actions  []RelationAttributeActions `json:"actions,omitempty"`
	// Optional label for the add button. The translated value for add_button_lable is used, if found else the string is used as is.
	AddButtonLabel *string  `json:"add_button_label,omitempty"`
	AllowedSchemas []string `json:"allowedSchemas,omitempty"`
	// A set of constraints applicable to the attribute.
	// These constraints should and will be enforced by the attribute renderer.
	//
	Constraints  *RelationAttributeConstraints `json:"constraints,omitempty"`
	DefaultValue any                           `json:"default_value,omitempty"`
	Deprecated   *bool                         `default:"false" json:"deprecated"`
	// Enables the preview, edition, and creation of relation items on a Master-Details view mode.
	DetailsViewModeEnabled *bool                        `default:"false" json:"details_view_mode_enabled"`
	DrawerSize             *RelationAttributeDrawerSize `json:"drawer_size,omitempty"`
	EditMode               *RelationAttributeEditMode   `json:"edit_mode,omitempty"`
	// When enable_relation_picker is set to true the user will be able to pick existing relations as values. Otherwise, the user will need to create new relation to link.
	EnableRelationPicker *bool `default:"true" json:"enable_relation_picker"`
	// When enable_relation_tags is set to true the user will be able to set tags(labels) in each relation item.
	EnableRelationTags *bool `default:"true" json:"enable_relation_tags"`
	// Setting to `true` disables editing the attribute on the entity builder UI
	EntityBuilderDisableEdit *bool `default:"false" json:"entity_builder_disable_edit"`
	// This attribute should only be active when the feature flag is enabled
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// Which group the attribute should appear in. Accepts group ID or group name
	Group      *string `json:"group,omitempty"`
	HasPrimary *bool   `json:"has_primary,omitempty"`
	// Do not render attribute in entity views
	Hidden *bool `default:"false" json:"hidden"`
	// When set to true, will hide the label of the field.
	HideLabel *bool   `json:"hide_label,omitempty"`
	Icon      *string `json:"icon,omitempty"`
	// ID for the entity attribute
	ID *string `json:"id,omitempty"`
	// A set of configurations meant to document and assist the user in filling the attribute.
	InfoHelpers *RelationAttributeInfoHelpers `json:"info_helpers,omitempty"`
	Label       string                        `json:"label"`
	Layout      *string                       `json:"layout,omitempty"`
	Name        string                        `json:"name"`
	// Attribute sort order (ascending) in group
	Order                 *int64  `json:"order,omitempty"`
	Placeholder           *string `json:"placeholder,omitempty"`
	PreviewValueFormatter *string `json:"preview_value_formatter,omitempty"`
	// Setting to `true` prevents the attribute from being modified / deleted
	Protected *bool `json:"protected,omitempty"`
	Readonly  *bool `default:"false" json:"readonly"`
	// Weak relation attributes are kept when duplicating an entity. Strong relation attributes are discarded when duplicating an entity.
	RelationAffinityMode *RelationAttributeRelationAffinityMode `json:"relation_affinity_mode,omitempty"`
	RelationType         *RelationAttributeRelationType         `json:"relation_type,omitempty"`
	// Defines the conditional rendering expression for showing this field.
	// When a valid expression is parsed, their evaluation defines the visibility of this attribute.
	// Note: Empty or invalid expression have no effect on the field visibility.
	//
	RenderCondition *string `json:"render_condition,omitempty"`
	Required        *bool   `default:"false" json:"required"`
	// Map of schema slug to target relation attribute
	ReverseAttributes map[string]string `json:"reverse_attributes,omitempty"`
	// Optional placeholder text for the relation search input. The translated value for search_placeholder is used, if found else the string is used as is.
	SearchPlaceholder *string `json:"search_placeholder,omitempty"`
	// This attribute should only be active when one of the provided settings have the correct value
	SettingsFlag []SettingFlag `json:"settings_flag,omitempty"`
	// Render as a column in table views. When defined, overrides `hidden`
	ShowInTable *bool `json:"show_in_table,omitempty"`
	// Allow sorting by this attribute in table views if `show_in_table` is true
	Sortable       *bool                            `default:"true" json:"sortable"`
	SummaryFields  []RelationAttributeSummaryFields `json:"summary_fields,omitempty"`
	Type           *RelationAttributeType           `json:"type,omitempty"`
	ValueFormatter *string                          `json:"value_formatter,omitempty"`
}

func (r RelationAttribute) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RelationAttribute) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *RelationAttribute) GetManifest() []string {
	if o == nil {
		return nil
	}
	return o.Manifest
}

func (o *RelationAttribute) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *RelationAttribute) GetActions() []RelationAttributeActions {
	if o == nil {
		return nil
	}
	return o.Actions
}

func (o *RelationAttribute) GetAddButtonLabel() *string {
	if o == nil {
		return nil
	}
	return o.AddButtonLabel
}

func (o *RelationAttribute) GetAllowedSchemas() []string {
	if o == nil {
		return nil
	}
	return o.AllowedSchemas
}

func (o *RelationAttribute) GetConstraints() *RelationAttributeConstraints {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *RelationAttribute) GetDefaultValue() any {
	if o == nil {
		return nil
	}
	return o.DefaultValue
}

func (o *RelationAttribute) GetDeprecated() *bool {
	if o == nil {
		return nil
	}
	return o.Deprecated
}

func (o *RelationAttribute) GetDetailsViewModeEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.DetailsViewModeEnabled
}

func (o *RelationAttribute) GetDrawerSize() *RelationAttributeDrawerSize {
	if o == nil {
		return nil
	}
	return o.DrawerSize
}

func (o *RelationAttribute) GetEditMode() *RelationAttributeEditMode {
	if o == nil {
		return nil
	}
	return o.EditMode
}

func (o *RelationAttribute) GetEnableRelationPicker() *bool {
	if o == nil {
		return nil
	}
	return o.EnableRelationPicker
}

func (o *RelationAttribute) GetEnableRelationTags() *bool {
	if o == nil {
		return nil
	}
	return o.EnableRelationTags
}

func (o *RelationAttribute) GetEntityBuilderDisableEdit() *bool {
	if o == nil {
		return nil
	}
	return o.EntityBuilderDisableEdit
}

func (o *RelationAttribute) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *RelationAttribute) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *RelationAttribute) GetHasPrimary() *bool {
	if o == nil {
		return nil
	}
	return o.HasPrimary
}

func (o *RelationAttribute) GetHidden() *bool {
	if o == nil {
		return nil
	}
	return o.Hidden
}

func (o *RelationAttribute) GetHideLabel() *bool {
	if o == nil {
		return nil
	}
	return o.HideLabel
}

func (o *RelationAttribute) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *RelationAttribute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RelationAttribute) GetInfoHelpers() *RelationAttributeInfoHelpers {
	if o == nil {
		return nil
	}
	return o.InfoHelpers
}

func (o *RelationAttribute) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *RelationAttribute) GetLayout() *string {
	if o == nil {
		return nil
	}
	return o.Layout
}

func (o *RelationAttribute) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RelationAttribute) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *RelationAttribute) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *RelationAttribute) GetPreviewValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.PreviewValueFormatter
}

func (o *RelationAttribute) GetProtected() *bool {
	if o == nil {
		return nil
	}
	return o.Protected
}

func (o *RelationAttribute) GetReadonly() *bool {
	if o == nil {
		return nil
	}
	return o.Readonly
}

func (o *RelationAttribute) GetRelationAffinityMode() *RelationAttributeRelationAffinityMode {
	if o == nil {
		return nil
	}
	return o.RelationAffinityMode
}

func (o *RelationAttribute) GetRelationType() *RelationAttributeRelationType {
	if o == nil {
		return nil
	}
	return o.RelationType
}

func (o *RelationAttribute) GetRenderCondition() *string {
	if o == nil {
		return nil
	}
	return o.RenderCondition
}

func (o *RelationAttribute) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *RelationAttribute) GetReverseAttributes() map[string]string {
	if o == nil {
		return nil
	}
	return o.ReverseAttributes
}

func (o *RelationAttribute) GetSearchPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.SearchPlaceholder
}

func (o *RelationAttribute) GetSettingsFlag() []SettingFlag {
	if o == nil {
		return nil
	}
	return o.SettingsFlag
}

func (o *RelationAttribute) GetShowInTable() *bool {
	if o == nil {
		return nil
	}
	return o.ShowInTable
}

func (o *RelationAttribute) GetSortable() *bool {
	if o == nil {
		return nil
	}
	return o.Sortable
}

func (o *RelationAttribute) GetSummaryFields() []RelationAttributeSummaryFields {
	if o == nil {
		return nil
	}
	return o.SummaryFields
}

func (o *RelationAttribute) GetType() *RelationAttributeType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *RelationAttribute) GetValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.ValueFormatter
}

type RelationAttributeNewEntityItemInput struct {
	AdditionalProperties any `additionalProperties:"true" json:"-"`
	// Access control list (ACL) for an entity. Defines sharing access to external orgs or users.
	ACL *EntityACL `json:"_acl,omitempty"`
	ID  string     `json:"_id"`
	// Manifest ID used to create/update the entity
	Manifest []string `json:"_manifest,omitempty"`
	Purpose  []string `json:"_purpose,omitempty"`
	// URL-friendly identifier for the entity schema
	Schema string   `json:"_schema"`
	Tags   []string `json:"_tags,omitempty"`
	// Title of entity
	Title *string `json:"_title"`
}

func (r RelationAttributeNewEntityItemInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RelationAttributeNewEntityItemInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RelationAttributeNewEntityItemInput) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *RelationAttributeNewEntityItemInput) GetACL() *EntityACL {
	if o == nil {
		return nil
	}
	return o.ACL
}

func (o *RelationAttributeNewEntityItemInput) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RelationAttributeNewEntityItemInput) GetManifest() []string {
	if o == nil {
		return nil
	}
	return o.Manifest
}

func (o *RelationAttributeNewEntityItemInput) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *RelationAttributeNewEntityItemInput) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *RelationAttributeNewEntityItemInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *RelationAttributeNewEntityItemInput) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

type RelationAttributeActionsInput struct {
	// The action type. Currently supported actions:
	//
	// | action | description |
	// |--------|-------------|
	// | add_existing | Enables the user to pick an existing entity to link as relation |
	// | create_new | Enables the user to create a new entity using the first/main `allowed_schemas` schema
	// | create_from_existing | Enables the user to pick an existing entity to clone from, while creating a blank new entity to link as relation |
	//
	ActionType *RelationAttributeActionType `json:"action_type,omitempty"`
	// Sets the action as the default action, visible as the main action button.
	Default *bool `json:"default,omitempty"`
	// Name of the feature flag that enables this action
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// The action label or action translation key (i18n)
	Label         *string                              `json:"label,omitempty"`
	NewEntityItem *RelationAttributeNewEntityItemInput `json:"new_entity_item,omitempty"`
	// This action should only be active when all the settings have the correct value
	SettingsFlag []SettingFlag `json:"settings_flag,omitempty"`
}

func (o *RelationAttributeActionsInput) GetActionType() *RelationAttributeActionType {
	if o == nil {
		return nil
	}
	return o.ActionType
}

func (o *RelationAttributeActionsInput) GetDefault() *bool {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *RelationAttributeActionsInput) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *RelationAttributeActionsInput) GetLabel() *string {
	if o == nil {
		return nil
	}
	return o.Label
}

func (o *RelationAttributeActionsInput) GetNewEntityItem() *RelationAttributeNewEntityItemInput {
	if o == nil {
		return nil
	}
	return o.NewEntityItem
}

func (o *RelationAttributeActionsInput) GetSettingsFlag() []SettingFlag {
	if o == nil {
		return nil
	}
	return o.SettingsFlag
}

// RelationAttributeInput - Entity Relationship
type RelationAttributeInput struct {
	// Manifest ID used to create/update the schema attribute
	Manifest []string                        `json:"_manifest,omitempty"`
	Purpose  []string                        `json:"_purpose,omitempty"`
	Actions  []RelationAttributeActionsInput `json:"actions,omitempty"`
	// Optional label for the add button. The translated value for add_button_lable is used, if found else the string is used as is.
	AddButtonLabel *string  `json:"add_button_label,omitempty"`
	AllowedSchemas []string `json:"allowedSchemas,omitempty"`
	// A set of constraints applicable to the attribute.
	// These constraints should and will be enforced by the attribute renderer.
	//
	Constraints  *RelationAttributeConstraints `json:"constraints,omitempty"`
	DefaultValue any                           `json:"default_value,omitempty"`
	Deprecated   *bool                         `default:"false" json:"deprecated"`
	// Enables the preview, edition, and creation of relation items on a Master-Details view mode.
	DetailsViewModeEnabled *bool                        `default:"false" json:"details_view_mode_enabled"`
	DrawerSize             *RelationAttributeDrawerSize `json:"drawer_size,omitempty"`
	EditMode               *RelationAttributeEditMode   `json:"edit_mode,omitempty"`
	// When enable_relation_picker is set to true the user will be able to pick existing relations as values. Otherwise, the user will need to create new relation to link.
	EnableRelationPicker *bool `default:"true" json:"enable_relation_picker"`
	// When enable_relation_tags is set to true the user will be able to set tags(labels) in each relation item.
	EnableRelationTags *bool `default:"true" json:"enable_relation_tags"`
	// Setting to `true` disables editing the attribute on the entity builder UI
	EntityBuilderDisableEdit *bool `default:"false" json:"entity_builder_disable_edit"`
	// This attribute should only be active when the feature flag is enabled
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// Which group the attribute should appear in. Accepts group ID or group name
	Group      *string `json:"group,omitempty"`
	HasPrimary *bool   `json:"has_primary,omitempty"`
	// Do not render attribute in entity views
	Hidden *bool `default:"false" json:"hidden"`
	// When set to true, will hide the label of the field.
	HideLabel *bool   `json:"hide_label,omitempty"`
	Icon      *string `json:"icon,omitempty"`
	// ID for the entity attribute
	ID *string `json:"id,omitempty"`
	// A set of configurations meant to document and assist the user in filling the attribute.
	InfoHelpers *RelationAttributeInfoHelpers `json:"info_helpers,omitempty"`
	Label       string                        `json:"label"`
	Layout      *string                       `json:"layout,omitempty"`
	Name        string                        `json:"name"`
	// Attribute sort order (ascending) in group
	Order                 *int64  `json:"order,omitempty"`
	Placeholder           *string `json:"placeholder,omitempty"`
	PreviewValueFormatter *string `json:"preview_value_formatter,omitempty"`
	// Setting to `true` prevents the attribute from being modified / deleted
	Protected *bool `json:"protected,omitempty"`
	Readonly  *bool `default:"false" json:"readonly"`
	// Weak relation attributes are kept when duplicating an entity. Strong relation attributes are discarded when duplicating an entity.
	RelationAffinityMode *RelationAttributeRelationAffinityMode `json:"relation_affinity_mode,omitempty"`
	RelationType         *RelationAttributeRelationType         `json:"relation_type,omitempty"`
	// Defines the conditional rendering expression for showing this field.
	// When a valid expression is parsed, their evaluation defines the visibility of this attribute.
	// Note: Empty or invalid expression have no effect on the field visibility.
	//
	RenderCondition *string `json:"render_condition,omitempty"`
	Required        *bool   `default:"false" json:"required"`
	// Map of schema slug to target relation attribute
	ReverseAttributes map[string]string `json:"reverse_attributes,omitempty"`
	// Optional placeholder text for the relation search input. The translated value for search_placeholder is used, if found else the string is used as is.
	SearchPlaceholder *string `json:"search_placeholder,omitempty"`
	// This attribute should only be active when one of the provided settings have the correct value
	SettingsFlag []SettingFlag `json:"settings_flag,omitempty"`
	// Render as a column in table views. When defined, overrides `hidden`
	ShowInTable *bool `json:"show_in_table,omitempty"`
	// Allow sorting by this attribute in table views if `show_in_table` is true
	Sortable       *bool                            `default:"true" json:"sortable"`
	SummaryFields  []RelationAttributeSummaryFields `json:"summary_fields,omitempty"`
	Type           *RelationAttributeType           `json:"type,omitempty"`
	ValueFormatter *string                          `json:"value_formatter,omitempty"`
}

func (r RelationAttributeInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RelationAttributeInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *RelationAttributeInput) GetManifest() []string {
	if o == nil {
		return nil
	}
	return o.Manifest
}

func (o *RelationAttributeInput) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *RelationAttributeInput) GetActions() []RelationAttributeActionsInput {
	if o == nil {
		return nil
	}
	return o.Actions
}

func (o *RelationAttributeInput) GetAddButtonLabel() *string {
	if o == nil {
		return nil
	}
	return o.AddButtonLabel
}

func (o *RelationAttributeInput) GetAllowedSchemas() []string {
	if o == nil {
		return nil
	}
	return o.AllowedSchemas
}

func (o *RelationAttributeInput) GetConstraints() *RelationAttributeConstraints {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *RelationAttributeInput) GetDefaultValue() any {
	if o == nil {
		return nil
	}
	return o.DefaultValue
}

func (o *RelationAttributeInput) GetDeprecated() *bool {
	if o == nil {
		return nil
	}
	return o.Deprecated
}

func (o *RelationAttributeInput) GetDetailsViewModeEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.DetailsViewModeEnabled
}

func (o *RelationAttributeInput) GetDrawerSize() *RelationAttributeDrawerSize {
	if o == nil {
		return nil
	}
	return o.DrawerSize
}

func (o *RelationAttributeInput) GetEditMode() *RelationAttributeEditMode {
	if o == nil {
		return nil
	}
	return o.EditMode
}

func (o *RelationAttributeInput) GetEnableRelationPicker() *bool {
	if o == nil {
		return nil
	}
	return o.EnableRelationPicker
}

func (o *RelationAttributeInput) GetEnableRelationTags() *bool {
	if o == nil {
		return nil
	}
	return o.EnableRelationTags
}

func (o *RelationAttributeInput) GetEntityBuilderDisableEdit() *bool {
	if o == nil {
		return nil
	}
	return o.EntityBuilderDisableEdit
}

func (o *RelationAttributeInput) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *RelationAttributeInput) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *RelationAttributeInput) GetHasPrimary() *bool {
	if o == nil {
		return nil
	}
	return o.HasPrimary
}

func (o *RelationAttributeInput) GetHidden() *bool {
	if o == nil {
		return nil
	}
	return o.Hidden
}

func (o *RelationAttributeInput) GetHideLabel() *bool {
	if o == nil {
		return nil
	}
	return o.HideLabel
}

func (o *RelationAttributeInput) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *RelationAttributeInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RelationAttributeInput) GetInfoHelpers() *RelationAttributeInfoHelpers {
	if o == nil {
		return nil
	}
	return o.InfoHelpers
}

func (o *RelationAttributeInput) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *RelationAttributeInput) GetLayout() *string {
	if o == nil {
		return nil
	}
	return o.Layout
}

func (o *RelationAttributeInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RelationAttributeInput) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *RelationAttributeInput) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *RelationAttributeInput) GetPreviewValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.PreviewValueFormatter
}

func (o *RelationAttributeInput) GetProtected() *bool {
	if o == nil {
		return nil
	}
	return o.Protected
}

func (o *RelationAttributeInput) GetReadonly() *bool {
	if o == nil {
		return nil
	}
	return o.Readonly
}

func (o *RelationAttributeInput) GetRelationAffinityMode() *RelationAttributeRelationAffinityMode {
	if o == nil {
		return nil
	}
	return o.RelationAffinityMode
}

func (o *RelationAttributeInput) GetRelationType() *RelationAttributeRelationType {
	if o == nil {
		return nil
	}
	return o.RelationType
}

func (o *RelationAttributeInput) GetRenderCondition() *string {
	if o == nil {
		return nil
	}
	return o.RenderCondition
}

func (o *RelationAttributeInput) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *RelationAttributeInput) GetReverseAttributes() map[string]string {
	if o == nil {
		return nil
	}
	return o.ReverseAttributes
}

func (o *RelationAttributeInput) GetSearchPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.SearchPlaceholder
}

func (o *RelationAttributeInput) GetSettingsFlag() []SettingFlag {
	if o == nil {
		return nil
	}
	return o.SettingsFlag
}

func (o *RelationAttributeInput) GetShowInTable() *bool {
	if o == nil {
		return nil
	}
	return o.ShowInTable
}

func (o *RelationAttributeInput) GetSortable() *bool {
	if o == nil {
		return nil
	}
	return o.Sortable
}

func (o *RelationAttributeInput) GetSummaryFields() []RelationAttributeSummaryFields {
	if o == nil {
		return nil
	}
	return o.SummaryFields
}

func (o *RelationAttributeInput) GetType() *RelationAttributeType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *RelationAttributeInput) GetValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.ValueFormatter
}
