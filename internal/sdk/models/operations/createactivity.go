// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/models/shared"
	"net/http"
)

type CreateActivityRequest struct {
	Activity *shared.Activity `request:"mediaType=application/json"`
	// Comma-separated list of entities which the activity primarily concerns
	Entities []string `queryParam:"style=form,explode=false,name=entities"`
}

func (o *CreateActivityRequest) GetActivity() *shared.Activity {
	if o == nil {
		return nil
	}
	return o.Activity
}

func (o *CreateActivityRequest) GetEntities() []string {
	if o == nil {
		return nil
	}
	return o.Entities
}

type CreateActivityResponse struct {
	// Success
	BaseActivityItem *shared.BaseActivityItem
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateActivityResponse) GetBaseActivityItem() *shared.BaseActivityItem {
	if o == nil {
		return nil
	}
	return o.BaseActivityItem
}

func (o *CreateActivityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateActivityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateActivityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
