// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Exists - Returns documents that have a value in the specified field.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-exists-query.html
type Exists struct {
	Field string `json:"field"`
}

func (o *Exists) GetField() string {
	if o == nil {
		return ""
	}
	return o.Field
}

// Ids - Returns documents based on their IDs.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-ids-query.html
type Ids struct {
	Values []string `json:"values,omitempty"`
}

func (o *Ids) GetValues() []string {
	if o == nil {
		return nil
	}
	return o.Values
}

// Relation - Indicates how the range query matches values for range fields.
type Relation string

const (
	RelationIntersects Relation = "INTERSECTS"
	RelationContains   Relation = "CONTAINS"
	RelationWithin     Relation = "WITHIN"
)

func (e Relation) ToPointer() *Relation {
	return &e
}
func (e *Relation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INTERSECTS":
		fallthrough
	case "CONTAINS":
		fallthrough
	case "WITHIN":
		*e = Relation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Relation: %v", v)
	}
}

type Range struct {
	// The date format used to parse date values.
	Format *string `json:"format,omitempty"`
	// A filter field value.
	Gt *SearchFilterValue `json:"gt,omitempty"`
	// A filter field value.
	Gte *SearchFilterValue `json:"gte,omitempty"`
	// A filter field value.
	Lt *SearchFilterValue `json:"lt,omitempty"`
	// A filter field value.
	Lte *SearchFilterValue `json:"lte,omitempty"`
	// Indicates how the range query matches values for range fields.
	Relation *Relation `json:"relation,omitempty"`
	// Coordinated Universal Time (UTC) offset or IANA time zone used to convert date values in the query to UTC.
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o *Range) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *Range) GetGt() *SearchFilterValue {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *Range) GetGte() *SearchFilterValue {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *Range) GetLt() *SearchFilterValue {
	if o == nil {
		return nil
	}
	return o.Lt
}

func (o *Range) GetLte() *SearchFilterValue {
	if o == nil {
		return nil
	}
	return o.Lte
}

func (o *Range) GetRelation() *Relation {
	if o == nil {
		return nil
	}
	return o.Relation
}

func (o *Range) GetTimeZone() *string {
	if o == nil {
		return nil
	}
	return o.TimeZone
}

type SearchFilter struct {
	// A subset of simplified Elasticsearch query clauses. The default operator is a logical AND. Use nested $and, $or, $not to combine filters using different logical operators.
	DollarAnd []SearchFilter `json:"$and,omitempty"`
	// A subset of simplified Elasticsearch query clauses. The default operator is a logical AND. Use nested $and, $or, $not to combine filters using different logical operators.
	DollarNot []SearchFilter `json:"$not,omitempty"`
	// A subset of simplified Elasticsearch query clauses. The default operator is a logical AND. Use nested $and, $or, $not to combine filters using different logical operators.
	DollarOr []SearchFilter `json:"$or,omitempty"`
	// Returns documents that have a value in the specified field.
	Exists *Exists `json:"exists,omitempty"`
	// Returns documents based on their IDs.
	Ids *Ids `json:"ids,omitempty"`
	// Returns documents with fields that have terms within a certain range.
	Range map[string]Range `json:"range,omitempty"`
	// Returns documents that contain an exact term in a provided field.
	//
	// To return a document, the query term must exactly match the queried field's value, including whitespace and capitalization.
	//
	// You likely DO NOT want to use this filter on text fields and want to target its .keyword instead.
	//
	Term map[string]*SearchFilterValue `json:"term,omitempty"`
	// Returns documents that contain one of the exact terms in a provided field. See term filter for more info.
	Terms map[string][]*SearchFilterValue `json:"terms,omitempty"`
}

func (o *SearchFilter) GetDollarAnd() []SearchFilter {
	if o == nil {
		return nil
	}
	return o.DollarAnd
}

func (o *SearchFilter) GetDollarNot() []SearchFilter {
	if o == nil {
		return nil
	}
	return o.DollarNot
}

func (o *SearchFilter) GetDollarOr() []SearchFilter {
	if o == nil {
		return nil
	}
	return o.DollarOr
}

func (o *SearchFilter) GetExists() *Exists {
	if o == nil {
		return nil
	}
	return o.Exists
}

func (o *SearchFilter) GetIds() *Ids {
	if o == nil {
		return nil
	}
	return o.Ids
}

func (o *SearchFilter) GetRange() map[string]Range {
	if o == nil {
		return nil
	}
	return o.Range
}

func (o *SearchFilter) GetTerm() map[string]*SearchFilterValue {
	if o == nil {
		return nil
	}
	return o.Term
}

func (o *SearchFilter) GetTerms() map[string][]*SearchFilterValue {
	if o == nil {
		return nil
	}
	return o.Terms
}
