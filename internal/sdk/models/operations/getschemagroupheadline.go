// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/models/shared"
	"net/http"
)

type GetSchemaGroupHeadlineRequest struct {
	// Schema Slug and the Schema Group ID
	CompositeID string `pathParam:"style=simple,explode=false,name=composite_id"`
}

func (o *GetSchemaGroupHeadlineRequest) GetCompositeID() string {
	if o == nil {
		return ""
	}
	return o.CompositeID
}

type GetSchemaGroupHeadlineResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	GroupHeadlineWithCompositeID *shared.GroupHeadlineWithCompositeID
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetSchemaGroupHeadlineResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSchemaGroupHeadlineResponse) GetGroupHeadlineWithCompositeID() *shared.GroupHeadlineWithCompositeID {
	if o == nil {
		return nil
	}
	return o.GroupHeadlineWithCompositeID
}

func (o *GetSchemaGroupHeadlineResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSchemaGroupHeadlineResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
