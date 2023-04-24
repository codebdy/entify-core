package meta

type Model struct {
	Classes      []*ClassMeta
	Relations    []*RelationMeta
	ScriptLogics []*MethodMeta
	GraphLogics  []*MethodMeta
	APIs         []*MethodMeta
}

func New(m *UMLMeta, metaId uint64) *Model {
	model := Model{
		Classes:   make([]*ClassMeta, len(m.Classes)),
		Relations: make([]*RelationMeta, len(m.Relations)),
	}

	for i := range m.Classes {
		model.Classes[i] = &m.Classes[i]
		if model.Classes[i].MetaId == 0 {
			model.Classes[i].MetaId = metaId
		}
	}

	for i := range m.Relations {
		model.Relations[i] = &m.Relations[i]
		if model.Relations[i].AppId == 0 {
			model.Relations[i].AppId = metaId
		}
	}

	for i := range m.ScriptLogics {
		model.ScriptLogics[i] = &m.ScriptLogics[i]
	}

	for i := range m.GraphLogics {
		model.GraphLogics[i] = &m.GraphLogics[i]
	}

	for i := range m.APIs {
		model.APIs[i] = &m.APIs[i]
	}

	return &model
}

func (m *Model) GetClassByUuid(uuid string) *ClassMeta {
	for i := range m.Classes {
		if m.Classes[i].Uuid == uuid {
			return m.Classes[i]
		}
	}

	return nil
}

func (m *Model) GetAbstractClassByUuid(uuid string) *ClassMeta {
	for i := range m.Classes {
		if m.Classes[i].Uuid == uuid && m.Classes[i].StereoType == CLASSS_ABSTRACT {
			return m.Classes[i]
		}
	}
	return nil
}
