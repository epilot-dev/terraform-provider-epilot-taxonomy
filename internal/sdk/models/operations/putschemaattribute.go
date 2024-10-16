// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/models/shared"
	"net/http"
)

type PutSchemaAttributeRequest struct {
	AttributeWithCompositeID *shared.AttributeWithCompositeIDInput `request:"mediaType=application/json"`
	// Schema Slug and the Attribute ID
	CompositeID string `pathParam:"style=simple,explode=false,name=composite_id"`
}

func (o *PutSchemaAttributeRequest) GetAttributeWithCompositeID() *shared.AttributeWithCompositeIDInput {
	if o == nil {
		return nil
	}
	return o.AttributeWithCompositeID
}

func (o *PutSchemaAttributeRequest) GetCompositeID() string {
	if o == nil {
		return ""
	}
	return o.CompositeID
}

type PutSchemaAttributeResponse struct {
	// Success
	AttributeWithCompositeID *shared.AttributeWithCompositeID
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSchemaAttributeResponse) GetAttributeWithCompositeID() *shared.AttributeWithCompositeID {
	if o == nil {
		return nil
	}
	return o.AttributeWithCompositeID
}

func (o *PutSchemaAttributeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSchemaAttributeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSchemaAttributeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
