package domain

import "github.com/codebdy/entify-core/model/meta"

type Enum struct {
	Uuid   string
	Name   string
	Values []string
	AppId  uint64
}

func NewEnum(c *meta.ClassMeta) *Enum {
	enum := Enum{
		Uuid:   c.Uuid,
		Name:   c.Name,
		Values: make([]string, len(c.Attributes)),
		AppId:  c.MetaId,
	}

	for i := range c.Attributes {
		enum.Values[i] = c.Attributes[i].Name
	}

	return &enum
}
