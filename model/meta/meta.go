package meta

import "github.com/codebdy/entify/shared"

type UMLMeta struct {
	//导出、导入文件时使用
	Id           shared.ID      `json:"id"`
	Classes      []ClassMeta    `json:"classes"`
	Relations    []RelationMeta `json:"relations"`
	Codes        []CodeMeta     `json:"codes"`
	ScriptLogics []MethodMeta   `json:"scriptLogics"` //脚本逻辑编排
	GraphLogics  []MethodMeta   `json:"graphLogics"`  //图形化逻辑编排
	APIs         []MethodMeta   `json:"apis"`         //用于生成服务接口
}
