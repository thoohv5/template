package schema

import (
	"github.com/facebook/ent/dialect"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// PhoneAccount holds the schema definition for the PhoneAccount entity.
type PhoneAccount struct {
	ent.Schema
}

// Mixin of the PhoneAccount.
func (PhoneAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

func (PhoneAccount) Config() ent.Config {
	return ent.Config{
		Table: "phone_account",
	}
}

// Fields of the PhoneAccount.
func (PhoneAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment(`主键`),
		field.String("user_identity").Optional().Comment(`用户标识`),
		field.String("phone").Optional().Comment(`手机号码`),
		field.Time("created_at").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Comment(`创建时间`),
		field.Time("updated_at").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Comment(`更新时间`),
		field.Time("deleted_at").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Comment(`删除时间`),
	}
}

// Edges of the PhoneAccount.
func (PhoneAccount) Edges() []ent.Edge {
	return nil
}
