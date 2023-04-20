// 一个UML meta 对应一个 Repository，维护自己的数据库配置
package entify

import (
	"github.com/codebdy/entify/db"
	"github.com/codebdy/entify/model"
	"github.com/codebdy/entify/model/meta"
	"github.com/codebdy/entify/shared"
)

type Repository struct {
	Config db.DbConfig
	Model  *model.Model
	MetaId shared.ID
}

func New(config db.DbConfig) *Repository {
	return &Repository{Config: config}
}

func (r *Repository) Init(umlMeta meta.UMLMeta, metaId shared.ID) {
	r.MetaId = metaId
}
