package schema

import (
	"github.com/facebook/ent/dialect"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// MiniProgramAccount holds the schema definition for the MiniProgramAccount entity.
type MiniProgramAccount struct {
	ent.Schema
}

// Mixin of the MiniProgramAccount.
func (MiniProgramAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

func (MiniProgramAccount) Config() ent.Config {
	return ent.Config{
		Table: "mini_program_account",
	}
}

// Fields of the MiniProgramAccount.
func (MiniProgramAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment(`主键`),
		field.String("user_identity").Optional().Comment(`用户标识`),
		field.String("open_id").Optional().Comment(`open_id`),
		field.String("nick_name").Default("").Comment(`用户昵称`),
		field.String("avatar_url").Default("").Comment(`用户头像`),
		field.Int32("gender").Default(0).Comment(`用户性别：0-未知，1-男性，2-女性，默认-0`),
		field.String("country").Default("").Comment(`用户所在国家`),
		field.String("province").Default("").Comment(`用户所在省份`),
		field.String("city").Default("").Comment(`用户所在城市`),
		field.String("language").Default("").Comment(`所用语言：en-英语，zh_CN-简体中文，zh_TW-繁体中文`),
		field.Time("created_at").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Comment(`创建时间`),
		field.Time("updated_at").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Comment(`更新时间`),
		field.Time("deleted_at").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Comment(`删除时间`),
	}
}

// Edges of the MiniProgramAccount.
func (MiniProgramAccount) Edges() []ent.Edge {
	return nil
}
