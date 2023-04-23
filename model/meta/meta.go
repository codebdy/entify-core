package meta

type UMLMeta struct {
	Classes   []ClassMeta    `json:"classes"`
	Relations []RelationMeta `json:"relations"`

	ScriptLogics []MethodMeta `json:"scriptLogics"`
	GraphLogics  []MethodMeta `json:"graphLogics"`

	APIs []MethodMeta `json:"apis"`
}
