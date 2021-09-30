package schema

import (
	"github.com/facebook/ent/dialect"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// UserAccount holds the schema definition for the UserAccount entity.
type UserAccount struct {
	ent.Schema
}

// Mixin of the UserAccount.
func (UserAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

func (UserAccount) Config() ent.Config {
	return ent.Config{
		Table: "user_account",
	}
}

// Fields of the UserAccount.
func (UserAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment(`主键`),
		field.String("user_identity").Optional().Comment(`用户标识`),
		field.Int64("account").Optional().Comment(`用户账号`),
		field.String("password").Default("").Comment(`用户登录密码`),
		field.String("salt").Default("").Comment(`密码盐`),
		field.Time("created_at").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Comment(`创建时间`),
		field.Time("updated_at").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Comment(`更新时间`),
		field.Time("deleted_at").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Comment(`删除时间`),
	}
}

// Edges of the UserAccount.
func (UserAccount) Edges() []ent.Edge {
	return nil
}
