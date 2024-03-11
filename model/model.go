package model

import (
	"github.com/codebdy/entify-core/model/domain"
	"github.com/codebdy/entify-core/model/graph"
	"github.com/codebdy/entify-core/model/meta"
)

type Model struct {
	Meta   *meta.Model
	Domain *domain.Model
	Graph  *graph.Model
}

func New(c *meta.UMLMeta, metaId uint64) *Model {
	metaModel := meta.New(c, metaId)
	domainModel := domain.New(metaModel)
	grahpModel := graph.New(domainModel)
	model := Model{
		Meta:   metaModel,
		Domain: domainModel,
		Graph:  grahpModel,
	}
	return &model
}

var SystemModel *Model
