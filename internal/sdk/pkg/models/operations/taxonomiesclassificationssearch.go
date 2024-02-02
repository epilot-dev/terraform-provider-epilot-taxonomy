// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-purpose/internal/sdk/pkg/models/shared"
	"net/http"
)

type TaxonomiesClassificationsSearchRequestBody struct {
	ClassificationIds []string `json:"classificationIds,omitempty"`
}

func (o *TaxonomiesClassificationsSearchRequestBody) GetClassificationIds() []string {
	if o == nil {
		return nil
	}
	return o.ClassificationIds
}

type TaxonomiesClassificationsSearchRequest struct {
	RequestBody *TaxonomiesClassificationsSearchRequestBody `request:"mediaType=application/json"`
	// Taxonomy slug
	TaxonomySlug string `queryParam:"style=form,explode=true,name=taxonomySlug"`
}

func (o *TaxonomiesClassificationsSearchRequest) GetRequestBody() *TaxonomiesClassificationsSearchRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *TaxonomiesClassificationsSearchRequest) GetTaxonomySlug() string {
	if o == nil {
		return ""
	}
	return o.TaxonomySlug
}

// TaxonomiesClassificationsSearchResponseBody - Returns list of taxonomy classifications
type TaxonomiesClassificationsSearchResponseBody struct {
	Results []shared.TaxonomyClassification `json:"results,omitempty"`
}

func (o *TaxonomiesClassificationsSearchResponseBody) GetResults() []shared.TaxonomyClassification {
	if o == nil {
		return nil
	}
	return o.Results
}

type TaxonomiesClassificationsSearchResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns list of taxonomy classifications
	Object *TaxonomiesClassificationsSearchResponseBody
}

func (o *TaxonomiesClassificationsSearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TaxonomiesClassificationsSearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TaxonomiesClassificationsSearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TaxonomiesClassificationsSearchResponse) GetObject() *TaxonomiesClassificationsSearchResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
