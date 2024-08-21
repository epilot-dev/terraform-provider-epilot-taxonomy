// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type GetRelationsRespWithPagination struct {
	Hits      *float64           `json:"hits,omitempty"`
	Relations []GetRelationsResp `json:"relations,omitempty"`
}

func (o *GetRelationsRespWithPagination) GetHits() *float64 {
	if o == nil {
		return nil
	}
	return o.Hits
}

func (o *GetRelationsRespWithPagination) GetRelations() []GetRelationsResp {
	if o == nil {
		return nil
	}
	return o.Relations
}