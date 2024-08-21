// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteSavedViewRequest struct {
	// View id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteSavedViewRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteSavedViewResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteSavedViewResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteSavedViewResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteSavedViewResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}