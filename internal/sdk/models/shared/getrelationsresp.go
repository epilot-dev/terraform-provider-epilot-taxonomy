// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"errors"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/internal/utils"
)

type GetRelationsRespType string

const (
	GetRelationsRespTypeRelationItem   GetRelationsRespType = "RelationItem"
	GetRelationsRespTypeRelationEntity GetRelationsRespType = "RelationEntity"
)

type GetRelationsResp struct {
	RelationItem   *RelationItem
	RelationEntity *RelationEntity

	Type GetRelationsRespType
}

func CreateGetRelationsRespRelationItem(relationItem RelationItem) GetRelationsResp {
	typ := GetRelationsRespTypeRelationItem

	return GetRelationsResp{
		RelationItem: &relationItem,
		Type:         typ,
	}
}

func CreateGetRelationsRespRelationEntity(relationEntity RelationEntity) GetRelationsResp {
	typ := GetRelationsRespTypeRelationEntity

	return GetRelationsResp{
		RelationEntity: &relationEntity,
		Type:           typ,
	}
}

func (u *GetRelationsResp) UnmarshalJSON(data []byte) error {

	var relationItem RelationItem = RelationItem{}
	if err := utils.UnmarshalJSON(data, &relationItem, "", true, true); err == nil {
		u.RelationItem = &relationItem
		u.Type = GetRelationsRespTypeRelationItem
		return nil
	}

	var relationEntity RelationEntity = RelationEntity{}
	if err := utils.UnmarshalJSON(data, &relationEntity, "", true, true); err == nil {
		u.RelationEntity = &relationEntity
		u.Type = GetRelationsRespTypeRelationEntity
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GetRelationsResp", string(data))
}

func (u GetRelationsResp) MarshalJSON() ([]byte, error) {
	if u.RelationItem != nil {
		return utils.MarshalJSON(u.RelationItem, "", true)
	}

	if u.RelationEntity != nil {
		return utils.MarshalJSON(u.RelationEntity, "", true)
	}

	return nil, errors.New("could not marshal union type GetRelationsResp: all fields are null")
}
