package meta

type UMLMeta struct {
	Classes   []ClassMeta    `json:"classes"`
	Relations []RelationMeta `json:"relations"`
}
