// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/models/shared"
	"net/http"
)

type GetRelatedEntitiesCountRequest struct {
	// Filter results to exclude schemas
	ExcludeSchemas []string `queryParam:"style=form,explode=false,name=exclude_schemas"`
	// Entity id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Entity Type
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
}

func (o *GetRelatedEntitiesCountRequest) GetExcludeSchemas() []string {
	if o == nil {
		return nil
	}
	return o.ExcludeSchemas
}

func (o *GetRelatedEntitiesCountRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetRelatedEntitiesCountRequest) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

type GetRelatedEntitiesCountResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	GetRelatedEntitiesCount *shared.GetRelatedEntitiesCount
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetRelatedEntitiesCountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRelatedEntitiesCountResponse) GetGetRelatedEntitiesCount() *shared.GetRelatedEntitiesCount {
	if o == nil {
		return nil
	}
	return o.GetRelatedEntitiesCount
}

func (o *GetRelatedEntitiesCountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRelatedEntitiesCountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
