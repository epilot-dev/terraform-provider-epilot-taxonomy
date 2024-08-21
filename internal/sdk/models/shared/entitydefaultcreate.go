// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ViewType string

const (
	ViewTypeDefault ViewType = "default"
)

func (e ViewType) ToPointer() *ViewType {
	return &e
}
func (e *ViewType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "default":
		*e = ViewType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ViewType: %v", v)
	}
}

type EntityDefaultCreate struct {
	SearchParams map[string]string `json:"search_params,omitempty"`
	ViewType     *ViewType         `json:"view_type,omitempty"`
}

func (o *EntityDefaultCreate) GetSearchParams() map[string]string {
	if o == nil {
		return nil
	}
	return o.SearchParams
}

func (o *EntityDefaultCreate) GetViewType() *ViewType {
	if o == nil {
		return nil
	}
	return o.ViewType
}