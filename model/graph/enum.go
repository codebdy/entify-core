package graph

import (
	"github.com/codebdy/entify/entify/model/domain"
)

type Enum struct {
	domain.Enum
}

func NewEnum(e *domain.Enum) *Enum {
	return &Enum{Enum: *e}
}
