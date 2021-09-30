package schema

import (
	"github.com/facebook/ent/dialect"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// UserInfo holds the schema definition for the UserInfo entity.
type UserInfo struct {
	ent.Schema
}

// Mixin of the UserInfo.
func (UserInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

func (UserInfo) Config() ent.Config {
	return ent.Config{
		Table: "user_info",
	}
}

// Fields of the UserInfo.
func (UserInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment(`主键`),
		field.String("user_identity").Optional().Comment(`标识`),
		field.Int32("channel").Default(0).Comment(`渠道：XXX`),
		field.Int32("form").Default(0).Comment(`来源：XXX`),
		field.Time("created_at").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Comment(`创建时间`),
		field.Time("updated_at").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Comment(`更新时间`),
		field.Time("deleted_at").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Comment(`删除时间`),
	}
}

// Edges of the UserInfo.
func (UserInfo) Edges() []ent.Edge {
	return nil
}
