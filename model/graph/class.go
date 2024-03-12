package graph

import (
	"fmt"

	"github.com/codebdy/entify-core/model/domain"
	"github.com/codebdy/entify-core/shared"
)

type Class struct {
	attributes   []*Attribute
	associations []*Association
	Domain       *domain.Class
}

func NewClass(c *domain.Class) *Class {
	cls := Class{
		Domain:     c,
		attributes: make([]*Attribute, len(c.Attributes)),
	}

	for i := range c.Attributes {
		cls.attributes[i] = NewAttribute(c.Attributes[i], &cls)
	}

	return &cls
}

func (c *Class) Uuid() string {
	return c.Domain.Uuid
}

func (c *Class) InnerId() uint64 {
	return c.Domain.InnerId
}

func (c *Class) Name() string {
	return c.Domain.Name
}

func (c *Class) Description() string {
	return c.Domain.Description
}

func (c *Class) AddAssociation(a *Association) {
	c.associations = append(c.associations, a)
}

func (c *Class) TableName() string {
	name := shared.SnakeString(c.Domain.Name)
	if c.Domain.AppId == 0 {
		return name
	}
	return fmt.Sprintf("%s%d_%s", shared.TABLE_PREFIX, c.Domain.AppId, name)
}

func (c *Class) IsSoftDelete() bool {
	return c.Domain.SoftDelete
}

func (c *Class) QueryListName() string {
	return shared.FirstLower(c.Name() + "List")
}

func (c *Class) QueryOneName() string {
	return shared.ONE + shared.FirstUpper(c.Name())
}

func (c *Class) QueryAggregateName() string {
	return shared.FirstLower(c.Name()) + shared.FirstUpper(shared.AGGREGATE)
}

func (c *Class) DeleteName() string {
	return shared.DELETE + shared.FirstUpper(c.Name())
}

func (c *Class) DeleteByIdName() string {
	return shared.DELETE + shared.FirstUpper(c.Name()) + shared.BY_ID
}

func (c *Class) SetName() string {
	return shared.SET + shared.FirstUpper(c.Name())
}

func (c *Class) UpsertName() string {
	return shared.UPSERT + shared.FirstUpper(c.Name())
}

func (c *Class) UpsertOneName() string {
	return shared.UPSERT_ONE + shared.FirstUpper(c.Name())
}

func (c *Class) AggregateName() string {
	return c.Name() + shared.FirstUpper(shared.AGGREGATE)
}

func (c *Class) ListName() string {
	return c.Name() + shared.FirstUpper(shared.LIST)
}
