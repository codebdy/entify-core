package graph

import (
	"github.com/codebdy/entify-core/model/meta"
	"github.com/codebdy/entify-core/model/table"
)

type Association struct {
	Name           string
	Relation       *Relation
	OwnerClassUuid string
}

func NewAssociation(name string, r *Relation, ownerUuid string) *Association {
	return &Association{
		Name:           name,
		Relation:       r,
		OwnerClassUuid: ownerUuid,
	}
}

// func (a *Association) Name() string {
// 	if a.IsSource() {
// 		return a.Relation.RoleOfTarget
// 	} else {
// 		return a.Relation.RoleOfSource
// 	}
// }

func (a *Association) Owner() *Entity {
	if a.IsSource() {
		return a.Relation.SourceEntity
	} else {
		return a.Relation.TargetEntity
	}
}

func (a *Association) TypeEntity() *Entity {
	if !a.IsSource() {
		return a.Relation.SourceEntity
	} else {
		return a.Relation.TargetEntity
	}
}

func (a *Association) Description() string {
	if a.IsSource() {
		return a.Relation.DescriptionOnTarget
	} else {
		return a.Relation.DescriptionOnSource
	}
}

func (a *Association) IsArray() bool {
	if a.IsSource() {
		return a.Relation.TargetMultiplicity == meta.ZERO_MANY
	} else {
		return a.Relation.SourceMutiplicity == meta.ZERO_MANY
	}
}

// func (a *Association) is1To1() bool {
// 	return a.Relation.SourceMutiplicity == meta.ZERO_ONE && a.Relation.TargetMultiplicity == meta.ZERO_ONE
// }

// func (a *Association) is1ToN() bool {
// 	if a.IsSource() {
// 		return a.Relation.SourceMutiplicity == meta.ZERO_ONE && a.Relation.TargetMultiplicity == meta.ZERO_MANY
// 	} else {
// 		return a.Relation.SourceMutiplicity == meta.ZERO_MANY && a.Relation.TargetMultiplicity == meta.ZERO_ONE
// 	}
// }

// func (a *Association) isNTo1() bool {
// 	if !a.IsSource() {
// 		return a.Relation.SourceMutiplicity == meta.ZERO_ONE && a.Relation.TargetMultiplicity == meta.ZERO_MANY
// 	} else {
// 		return a.Relation.SourceMutiplicity == meta.ZERO_MANY && a.Relation.TargetMultiplicity == meta.ZERO_ONE
// 	}
// }

// func (a *Association) isNToN() bool {
// 	return a.Relation.SourceMutiplicity == meta.ZERO_MANY && a.Relation.TargetMultiplicity == meta.ZERO_MANY
// }

// // 单向关联
// func (a *Association) isOneWay() bool {
// 	return a.Relation.RelationType != meta.ONE_WAY_ASSOCIATION
// }

//关系存本方
// func (a *Association) IsColumn() bool {
// 	if a.is1To1() { //单向双向是一样的
// 		if a.IsSource() {
// 			return true
// 		} else {
// 			return false
// 		}
// 	} else if a.is1ToN() { //存对方或中间表
// 		return false
// 	} else if a.isNTo1() {
// 		if a.isOneWay() && !a.IsSource() { //单向，被指向，存中间表
// 			return false
// 		} else {
// 			return true
// 		}
// 	}

// 	return false
// }

//关系存对方
// func (a *Association) IsTargetColumn() bool {
// 	if a.is1To1() { //单向双向是一样的
// 		if a.IsSource() {
// 			return false
// 		} else {
// 			return true
// 		}
// 	} else if a.is1ToN() { //存对方或中间表
// 		if a.isOneWay() && a.IsSource() { //单向，指向对方，存中间表
// 			return false
// 		} else {
// 			return true
// 		}
// 	} else if a.isNTo1() { //存本方或中间表
// 		return false
// 	}

// 	return false
// }

//关系存中间表
// func (a *Association) IsPovitTable() bool {
// 	if a.isNToN() {
// 		return true
// 	}

// 	if a.is1ToN() && a.IsSource() && a.isOneWay() {
// 		return true
// 	}

//		if a.isNTo1() && !a.IsSource() && a.isOneWay() {
//			return true
//		}
//		return false
//	}
func (a *Association) IsCombination() bool {
	return a.IsSource() &&
		(a.Relation.RelationType == meta.TWO_WAY_COMBINATION ||
			a.Relation.RelationType == meta.ONE_WAY_COMBINATION)
}

//UML的Role Name存在关系的对方，而不是本方
func (a *Association) IsSource() bool {
	return a.Relation.RoleOfTarget == a.Name
}

func (a *Association) GetName() string {
	return a.Name
}

func (a *Association) Path() string {
	return a.Owner().Domain.Name + "." + a.Name
}

func (a *Association) Table() *table.Table {
	return a.Relation.Table
}

func (r *Association) SourceColumn() *table.Column {
	for i := range r.Relation.Table.Columns {
		column := r.Relation.Table.Columns[i]
		if column.Name == r.Relation.SourceColumnName() {
			return column
		}
	}
	return nil
}

func (r *Association) TargetColumn() *table.Column {
	for i := range r.Relation.Table.Columns {
		column := r.Relation.Table.Columns[i]
		if column.Name == r.Relation.TargetColumnName() {
			return column
		}
	}
	return nil
}

func (r *Association) OwnerColumn() *table.Column {
	if r.IsSource() {
		return r.SourceColumn()
	} else {
		return r.TargetColumn()
	}
}

func (r *Association) TypeColumn() *table.Column {
	if !r.IsSource() {
		return r.SourceColumn()
	} else {
		return r.TargetColumn()
	}
}
