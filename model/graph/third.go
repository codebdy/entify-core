package graph

import (
	"github.com/codebdy/entify/entify/model/domain"
)

type ThirdParty struct {
	Class
}

func NewThirdParty(c *domain.Class) *ThirdParty {
	return &ThirdParty{
		Class: *NewClass(c),
	}
}

func (t *ThirdParty) Attributes() []*Attribute {
	return t.attributes
}
