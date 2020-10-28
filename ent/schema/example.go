package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

type Example struct {
	ent.Schema
}

func (Example) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("id").Default(uuid.New),
	}
}

func (Example) Hooks() []ent.Hook {
	return []ent.Hook{
		func(m ent.Mutator) ent.Mutator {
			return m
		},
	}
}
