package graph

import (
	"github.com/codebdy/entify-core/model/domain"
	"github.com/codebdy/entify-core/model/meta"
	"github.com/codebdy/entify-core/model/table"
	"github.com/codebdy/entify-core/shared"
)

type Relation struct {
	AppId                  uint64
	Uuid                   string
	InnerId                uint64
	RelationType           string
	SourceEntity           *Entity
	TargetEntity           *Entity
	RoleOfTarget           string
	RoleOfSource           string
	DescriptionOnSource    string
	DescriptionOnTarget    string
	SourceMutiplicity      string
	TargetMultiplicity     string
	EnableAssociaitonClass bool
	AssociationClass       meta.AssociationClass
	Table                  *table.Table
}

func NewRelation(
	r *domain.Relation,
	sourceEntity *Entity,
	targetEntity *Entity,
) *Relation {
	roleOfTarget := r.RoleOfTarget
	roleOfSource := r.RoleOfSource

	if sourceEntity.Uuid() != r.Source.Uuid {
		roleOfSource = roleOfSource + "Of" + shared.FirstUpper(sourceEntity.Name())
	}

	if targetEntity.Uuid() != r.Target.Uuid {
		roleOfTarget = roleOfTarget + "Of" + shared.FirstUpper(targetEntity.Name())
	}

	relation := &Relation{
		Uuid:                   r.Uuid,
		InnerId:                r.InnerId,
		RelationType:           r.RelationType,
		SourceEntity:           sourceEntity,
		TargetEntity:           targetEntity,
		RoleOfTarget:           roleOfTarget,
		RoleOfSource:           roleOfSource,
		DescriptionOnSource:    r.DescriptionOnSource,
		DescriptionOnTarget:    r.DescriptionOnTarget,
		SourceMutiplicity:      r.SourceMutiplicity,
		TargetMultiplicity:     r.TargetMultiplicity,
		EnableAssociaitonClass: r.EnableAssociaitonClass,
		AssociationClass:       r.AssociationClass,
		AppId:                  r.AppId,
	}

	return relation
}

func (r *Relation) SourceColumnName() string {
	return r.RoleOfTarget + "_id"
}

func (r *Relation) TargetColumnName() string {
	return r.RoleOfSource + "_id"
}