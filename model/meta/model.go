package meta

type Model struct {
	Classes   []*ClassMeta
	Relations []*RelationMeta
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
