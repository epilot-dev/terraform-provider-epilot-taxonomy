// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/models/shared"
	"net/http"
)

type DeleteSchemaGroupRequest struct {
	// Schema Slug and the Group ID
	CompositeID string `pathParam:"style=simple,explode=false,name=composite_id"`
}

func (o *DeleteSchemaGroupRequest) GetCompositeID() string {
	if o == nil {
		return ""
	}
	return o.CompositeID
}

type DeleteSchemaGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	EntitySchemaGroupWithCompositeID *shared.EntitySchemaGroupWithCompositeID
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteSchemaGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteSchemaGroupResponse) GetEntitySchemaGroupWithCompositeID() *shared.EntitySchemaGroupWithCompositeID {
	if o == nil {
		return nil
	}
	return o.EntitySchemaGroupWithCompositeID
}

func (o *DeleteSchemaGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteSchemaGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}