package schema

import (
	"github.com/facebook/ent/dialect"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// UserExtend holds the schema definition for the UserExtend entity.
type UserExtend struct {
	ent.Schema
}

// Mixin of the UserExtend.
func (UserExtend) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

func (UserExtend) Config() ent.Config {
	return ent.Config{
		Table: "user_extend",
	}
}

// Fields of the UserExtend.
func (UserExtend) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment(`主键`),
		field.String("user_identity").Optional().Comment(`标识`),
		field.Time("created_at").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Comment(`创建时间`),
		field.Time("updated_at").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Comment(`更新时间`),
		field.Time("deleted_at").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Comment(`删除时间`),
	}
}

// Edges of the UserExtend.
func (UserExtend) Edges() []ent.Edge {
	return nil
}
