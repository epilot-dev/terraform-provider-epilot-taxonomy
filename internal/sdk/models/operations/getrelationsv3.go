// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/internal/utils"
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/models/shared"
	"net/http"
)

type GetRelationsV3Request struct {
	// Filter results to exclude schemas
	ExcludeSchemas []string `queryParam:"style=form,explode=false,name=exclude_schemas"`
	// Starting page number
	From *int64 `default:"0" queryParam:"style=form,explode=true,name=from"`
	// When true, enables entity hydration to resolve nested $relation & $relation_ref references in-place.
	Hydrate *bool `default:"false" queryParam:"style=form,explode=true,name=hydrate"`
	// Entity id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// When true, includes reverse relations in response (other entities pointing to this entity)
	// *It gets overriden by mode query parameter.*
	//
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	IncludeReverse *bool `default:"false" queryParam:"style=form,explode=true,name=include_reverse"`
	// Filter results to only include schemas
	IncludeSchemas []string `queryParam:"style=form,explode=false,name=include_schemas"`
	// Options to determine how relations will be included in the result.
	// *It overrides the include_reverse query param.*
	// Explanation of possible options:
	// - direct: include relations to which the searched entity refers
	// - reverse: include relations that refer to the entity you are looking for
	// - both: both direct and reverse relations
	//
	Mode *shared.EntityRelationsModeQueryParam `queryParam:"style=form,explode=true,name=mode"`
	// Number of results to return per page
	Size *int64 `default:"100" queryParam:"style=form,explode=true,name=size"`
	// Entity Type
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
}

func (g GetRelationsV3Request) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRelationsV3Request) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRelationsV3Request) GetExcludeSchemas() []string {
	if o == nil {
		return nil
	}
	return o.ExcludeSchemas
}

func (o *GetRelationsV3Request) GetFrom() *int64 {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *GetRelationsV3Request) GetHydrate() *bool {
	if o == nil {
		return nil
	}
	return o.Hydrate
}

func (o *GetRelationsV3Request) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetRelationsV3Request) GetIncludeReverse() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeReverse
}

func (o *GetRelationsV3Request) GetIncludeSchemas() []string {
	if o == nil {
		return nil
	}
	return o.IncludeSchemas
}

func (o *GetRelationsV3Request) GetMode() *shared.EntityRelationsModeQueryParam {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *GetRelationsV3Request) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetRelationsV3Request) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

type GetRelationsV3Response struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	GetRelationsRespWithPagination *shared.GetRelationsRespWithPagination
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetRelationsV3Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRelationsV3Response) GetGetRelationsRespWithPagination() *shared.GetRelationsRespWithPagination {
	if o == nil {
		return nil
	}
	return o.GetRelationsRespWithPagination
}

func (o *GetRelationsV3Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRelationsV3Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}