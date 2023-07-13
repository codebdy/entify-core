package meta

import "github.com/codebdy/minions-go/dsl"

const (
	QUERY    string = "query"
	MUTATION string = "mutation"
)

type ArgMeta struct {
	Uuid      string `json:"uuid"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	TypeUuid  string `json:"typeUuid"`
	TypeLabel string `json:"typeLabel"`
}

type MethodMeta struct {
	Uuid        string            `json:"uuid"`
	Name        string            `json:"name"`
	Type        string            `json:"type"`
	TypeUuid    string            `json:"typeUuid"`
	TypeLabel   string            `json:"typeLabel"`
	Args        []ArgMeta         `json:"args"`
	OperateType string            `json:"operateType"` //mutation or query
	Description string            `json:"description"`
	LogicMetas  dsl.LogicFlowMeta `json:"logicMetas"`
	LogicScript string            `json:"logicScript"`
}
