// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/models/shared"
	"net/http"
)

type DeleteSchemaCapabilityRequest struct {
	// Schema Slug and the Attribute ID
	CompositeID string `pathParam:"style=simple,explode=false,name=composite_id"`
}

func (o *DeleteSchemaCapabilityRequest) GetCompositeID() string {
	if o == nil {
		return ""
	}
	return o.CompositeID
}

type DeleteSchemaCapabilityResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	EntityCapabilityWithCompositeID *shared.EntityCapabilityWithCompositeID
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteSchemaCapabilityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteSchemaCapabilityResponse) GetEntityCapabilityWithCompositeID() *shared.EntityCapabilityWithCompositeID {
	if o == nil {
		return nil
	}
	return o.EntityCapabilityWithCompositeID
}

func (o *DeleteSchemaCapabilityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteSchemaCapabilityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
