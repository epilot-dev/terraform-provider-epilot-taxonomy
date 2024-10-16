// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/internal/utils"
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/models/shared"
	"net/http"
)

type ListSchemasRequest struct {
	// Return unpublished draft schemas
	Unpublished *bool `default:"false" queryParam:"style=form,explode=true,name=unpublished"`
}

func (l ListSchemasRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListSchemasRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListSchemasRequest) GetUnpublished() *bool {
	if o == nil {
		return nil
	}
	return o.Unpublished
}

// ListSchemasResponseBody - Success
type ListSchemasResponseBody struct {
	Results []shared.EntitySchemaItem `json:"results,omitempty"`
}

func (o *ListSchemasResponseBody) GetResults() []shared.EntitySchemaItem {
	if o == nil {
		return nil
	}
	return o.Results
}

type ListSchemasResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Object *ListSchemasResponseBody
}

func (o *ListSchemasResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSchemasResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSchemasResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListSchemasResponse) GetObject() *ListSchemasResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
