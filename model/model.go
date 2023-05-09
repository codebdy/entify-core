package model

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/codebdy/entify/model/domain"
	"github.com/codebdy/entify/model/graph"
	"github.com/codebdy/entify/model/meta"
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

func ReadContentFromJson(fileName string) meta.UMLMeta {
	data, err := ioutil.ReadFile(fileName)
	content := meta.UMLMeta{}
	if nil != err {
		log.Panic(err.Error())
	} else {
		err = json.Unmarshal(data, &content)
	}

	return content
}
