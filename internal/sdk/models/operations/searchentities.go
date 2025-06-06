// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/models/shared"
	"net/http"
)

type SearchEntitiesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	EntitySearchResults *shared.EntitySearchResults
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Res *string
}

func (o *SearchEntitiesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SearchEntitiesResponse) GetEntitySearchResults() *shared.EntitySearchResults {
	if o == nil {
		return nil
	}
	return o.EntitySearchResults
}

func (o *SearchEntitiesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SearchEntitiesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SearchEntitiesResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
