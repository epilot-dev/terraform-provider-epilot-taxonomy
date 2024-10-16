// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/models/shared"
	"net/http"
)

type TaxonomyAutocompleteRequest struct {
	// Input to autocomplete
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// Minimum number of results to return
	Size *float64 `queryParam:"style=form,explode=true,name=size"`
	// Taxonomy slug
	TaxonomySlug string `pathParam:"style=simple,explode=false,name=taxonomySlug"`
}

func (o *TaxonomyAutocompleteRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *TaxonomyAutocompleteRequest) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *TaxonomyAutocompleteRequest) GetTaxonomySlug() string {
	if o == nil {
		return ""
	}
	return o.TaxonomySlug
}

// TaxonomyAutocompleteResponseBody - Taxonomy classifications
type TaxonomyAutocompleteResponseBody struct {
	Results []shared.TaxonomyClassification `json:"results,omitempty"`
}

func (o *TaxonomyAutocompleteResponseBody) GetResults() []shared.TaxonomyClassification {
	if o == nil {
		return nil
	}
	return o.Results
}

type TaxonomyAutocompleteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Taxonomy classifications
	Object *TaxonomyAutocompleteResponseBody
}

func (o *TaxonomyAutocompleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TaxonomyAutocompleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TaxonomyAutocompleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TaxonomyAutocompleteResponse) GetObject() *TaxonomyAutocompleteResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
