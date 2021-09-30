// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// MiniProgramAccountColumns holds the columns for the "mini_program_account" table.
	MiniProgramAccountColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_identity", Type: field.TypeString, Nullable: true},
		{Name: "open_id", Type: field.TypeString, Nullable: true},
		{Name: "nick_name", Type: field.TypeString, Default: ""},
		{Name: "avatar_url", Type: field.TypeString, Default: ""},
		{Name: "gender", Type: field.TypeInt32},
		{Name: "country", Type: field.TypeString, Default: ""},
		{Name: "province", Type: field.TypeString, Default: ""},
		{Name: "city", Type: field.TypeString, Default: ""},
		{Name: "language", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// MiniProgramAccountTable holds the schema information for the "mini_program_account" table.
	MiniProgramAccountTable = &schema.Table{
		Name:        "mini_program_account",
		Columns:     MiniProgramAccountColumns,
		PrimaryKey:  []*schema.Column{MiniProgramAccountColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PhoneAccountColumns holds the columns for the "phone_account" table.
	PhoneAccountColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_identity", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// PhoneAccountTable holds the schema information for the "phone_account" table.
	PhoneAccountTable = &schema.Table{
		Name:        "phone_account",
		Columns:     PhoneAccountColumns,
		PrimaryKey:  []*schema.Column{PhoneAccountColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "identity", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeInt32},
		{Name: "is_disable", Type: field.TypeInt32},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:        "user",
		Columns:     UserColumns,
		PrimaryKey:  []*schema.Column{UserColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UserAccountColumns holds the columns for the "user_account" table.
	UserAccountColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_identity", Type: field.TypeString, Nullable: true},
		{Name: "account", Type: field.TypeInt64, Nullable: true},
		{Name: "password", Type: field.TypeString, Default: ""},
		{Name: "salt", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// UserAccountTable holds the schema information for the "user_account" table.
	UserAccountTable = &schema.Table{
		Name:        "user_account",
		Columns:     UserAccountColumns,
		PrimaryKey:  []*schema.Column{UserAccountColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UserExtendColumns holds the columns for the "user_extend" table.
	UserExtendColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_identity", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// UserExtendTable holds the schema information for the "user_extend" table.
	UserExtendTable = &schema.Table{
		Name:        "user_extend",
		Columns:     UserExtendColumns,
		PrimaryKey:  []*schema.Column{UserExtendColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UserInfoColumns holds the columns for the "user_info" table.
	UserInfoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_identity", Type: field.TypeString, Nullable: true},
		{Name: "channel", Type: field.TypeInt32},
		{Name: "form", Type: field.TypeInt32},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// UserInfoTable holds the schema information for the "user_info" table.
	UserInfoTable = &schema.Table{
		Name:        "user_info",
		Columns:     UserInfoColumns,
		PrimaryKey:  []*schema.Column{UserInfoColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MiniProgramAccountTable,
		PhoneAccountTable,
		UserTable,
		UserAccountTable,
		UserExtendTable,
		UserInfoTable,
	}
)

func init() {
}
