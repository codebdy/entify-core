package model

import (
	"github.com/codebdy/entify/entify/model/domain"
	"github.com/codebdy/entify/entify/model/graph"
	"github.com/codebdy/entify/entify/model/meta"
)

type Model struct {
	Meta   *meta.Model
	Domain *domain.Model
	Graph  *graph.Model
}

func New(c *meta.MetaContent, appid uint64) *Model {
	metaModel := meta.New(c, appid)
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
