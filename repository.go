// 一个UML meta 对应一个 Repository，维护自己的数据库配置
package entify

import (
	"github.com/codebdy/entify/db"
	"github.com/codebdy/entify/model"
	"github.com/codebdy/entify/model/meta"
	"github.com/codebdy/entify/orm"
	"github.com/codebdy/entify/shared"
	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	DbConfig db.DbConfig
	Model    *model.Model
	MetaId   shared.ID
}

func New(config db.DbConfig) *Repository {
	return &Repository{DbConfig: config}
}

func (r *Repository) Init(umlMeta meta.UMLMeta, metaId shared.ID) {
	r.MetaId = metaId
	r.Model = model.New(&umlMeta, metaId)
}

func (r *Repository) PublishMeta(published, next *meta.UMLMeta, metaId shared.ID) {
	publishedModel := model.New(published, metaId)
	nextModel := model.New(next, metaId)
	diff := model.CreateDiff(publishedModel, nextModel)
	orm.Migrage(diff, r.DbConfig)
}

func (r *Repository) OpenSession() (*orm.Session, error) {
	return orm.Open(r.DbConfig, r.Model)
}

func (r *Repository) IsEntityExists(name string) bool {
	return orm.IsEntityExists(name, r.DbConfig)
}
